package qbittorrent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	url        url.URL
	httpClient http.Client
	user       string
	password   string
	sessionID  string
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
	Method string `json:"method"`
	Params T      `json:"params,omitempty"`
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

type Response[T any] struct {
	JSONRPC string           `json:"jsonrpc"`
	Error   ResponseError[T] `json:"error"`
	ID      int              `json:"id"`
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

func post(ctx context.Context, client *Client, body any) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, client.url.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error building req: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(client.user, client.password)

	req2 := req.Clone(ctx)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		sessionID := resp.Header.Get(sessionIDHeader)
		if sessionID == "" {
			return fmt.Errorf("server returned no session ID")
		}
		client.sessionID = sessionID

		resp, err = client.httpClient.Do(req2)
		if err != nil {
			return fmt.Errorf("error making request: %w", err)
		}
		defer resp.Body.Close()
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	// if dst == nil {
	//	return nil
	// }
	//
	// if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
	//	return fmt.Errorf("error decoding response body: %w", err)
	// }

	return nil
}
