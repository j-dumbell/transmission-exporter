package transmission

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type Client struct {
	url        url.URL
	httpClient http.Client
	user       string
	password   string
	sessionID  string
	mutex      sync.RWMutex
}

type ClientParams struct {
	Host     string
	User     string
	Password string
}

type Request struct {
	Method string `json:"method"`
}

type RequestWithParams[T any] struct {
	Method    string `json:"method"`
	Arguments T      `json:"arguments,omitempty"`
}

type ErrorData[T any] struct {
	ErrorString string `json:"errorString"`
	Result      T      `json:"result"`
}

type ResponseError[T any] struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    ErrorData[T] `json:"data"`
}

func (re ResponseError[T]) Error() string {
	return fmt.Sprintf("response code: %d, message: %s", re.Code, re.Message)
}

type Response[R any] struct {
	Arguments R      `json:"arguments"`
	Result    string `json:"result"`
}

func (r Response[R]) isSuccess() bool {
	return r.Result == "success"
}

func New(params ClientParams) (*Client, error) {
	hostPart, err := url.Parse(params.Host)
	if err != nil {
		return nil, fmt.Errorf("error parsing host URL: %w", err)
	}
	apiPart, _ := url.Parse("transmission/rpc")
	fullURL := hostPart.ResolveReference(apiPart)

	return &Client{
		url:        *fullURL,
		httpClient: http.Client{},
		user:       params.User,
		password:   params.Password,
	}, nil
}

const (
	sessionIDHeader = "X-Transmission-Session-Id"
)

func (c *Client) doRequest(ctx context.Context, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url.String(), body)
	if err != nil {
		return nil, fmt.Errorf("error building request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.user, c.password)
	if c.sessionID != "" {
		c.mutex.RLock()
		req.Header.Set(sessionIDHeader, c.sessionID)
		c.mutex.RUnlock()
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	return resp, nil
}

func (c *Client) post(ctx context.Context, body any, dst any) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling body: %w", err)
	}

	// fmt.Println("===> body", string(jsonBody))

	resp, err := c.doRequest(ctx, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		sessionID := resp.Header.Get(sessionIDHeader)
		if sessionID == "" {
			return fmt.Errorf("server returned no session ID")
		}

		c.mutex.Lock()
		c.sessionID = sessionID
		c.mutex.Unlock()

		resp, err = c.doRequest(ctx, bytes.NewBuffer(jsonBody))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("unexpected %d status returned: %s", resp.StatusCode, bytes)
	}

	// bytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(bytes))

	if dst == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}

	return nil
}

func post[R any](ctx context.Context, client *Client, method string) (*R, error) {
	var response Response[R]
	if err := client.post(ctx, Request{Method: method}, &response); err != nil {
		return nil, err
	}
	if !response.isSuccess() {
		return nil, errors.New(response.Result)
	}
	return &response.Arguments, nil
}

func postWithArgs[P any, R any](ctx context.Context, client *Client, method string, params P) (*R, error) {
	var response Response[R]
	if err := client.post(ctx, RequestWithParams[P]{Method: method, Arguments: params}, &response); err != nil {
		return nil, err
	}
	if !response.isSuccess() {
		return nil, errors.New(response.Result)
	}
	return &response.Arguments, nil
}
