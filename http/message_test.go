package http

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestParseRequestMessage(t *testing.T) {
	inputs := []struct {
		reader *bufio.Reader
	}{
		{bufio.NewReader(strings.NewReader("GET / HTTP/1.1\r\nHost: example.com\r\nAccept-Encoding: gzip, deflate\r\n\r\n"))},
		{bufio.NewReader(strings.NewReader("GET / HTTP/1.1\r\n Host: example.com\r\nAccept-Encoding: gzip, deflate\r\n\r\n"))},
		{bufio.NewReader(strings.NewReader("GET / HTTP/1.1 toomanyfields\r\nHost: example.com\r\n"))},
		{bufio.NewReader(strings.NewReader("GET /\r\nHost: example.com\r\n"))},
		{bufio.NewReader(strings.NewReader("INVALID / HTTP/1.1\r\nHost: example.com\r\n"))},
	}
	expects := []struct {
		Request *Request
		Err     error
	}{
		{
			Request: &Request{
				Method:  "GET",
				Target:  "/",
				Version: "HTTP/1.1",
				Headers: Headers{
					"Host":            []string{"example.com"},
					"Accept-Encoding": []string{"gzip, deflate"},
				},
			},
			Err: nil,
		},
		{
			Request: nil,
			Err:     ErrInvalidRequest,
		},
		{
			Request: nil,
			Err:     ErrInvalidRequest,
		},
		{
			Request: nil,
			Err:     ErrInvalidRequest,
		},
		{
			Request: nil,
			Err:     ErrMethodNotImplemented,
		},
	}

	for i, tt := range expects {
		output, err := parseRequestMessage(inputs[i].reader)
		if tt.Err == nil {
			if !reflect.DeepEqual(output, tt.Request) {
				t.Errorf("wants %+v, but got %+v\n", tt.Request, output)
			}
		} else {
			if err != tt.Err {
				t.Errorf("wants error %+v, but got %+v\n", tt.Err, err)
			}
		}
	}
}
