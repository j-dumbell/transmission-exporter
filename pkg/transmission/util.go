package transmission

import (
	"encoding/base64"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func structJSONFields[T any]() []string {
	var zero T
	t := reflect.TypeOf(zero)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("structFields: T must be a struct or pointer to struct")
	}

	var fields []string
	collectJSONFields(t, &fields)

	return fields
}

func collectJSONFields(t reflect.Type, out *[]string) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if !f.IsExported() {
			continue
		}

		// Follow embedded structs (anonymous fields)
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			collectJSONFields(f.Type, out)
			continue
		}

		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}

		// Extract the field name from the JSON tag
		name := tag
		if idx := strings.IndexByte(tag, ','); idx != -1 {
			name = tag[:idx]
		}

		// If no JSON tag name is specified, use the Go field name
		if name == "" {
			name = f.Name
		}

		// Skip if name is still empty or "-"
		if name == "" || name == "-" {
			continue
		}

		*out = append(*out, name)
	}
}

func TorrentBase64(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading torrent file: %w", err)
	}
	return base64.StdEncoding.EncodeToString(contents), nil
}
