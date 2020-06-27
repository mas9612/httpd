package http

import (
	"bytes"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	inputs := []struct {
		res *Response
		err error
	}{
		{
			res: &Response{
				Version:      "HTTP/1.1",
				StatusCode:   200,
				ReasonPhrase: "OK",
				Headers: Headers{
					"Content-Length": []string{"15"},
				},
				Body: []byte("Hello from test"),
			},
			err: nil,
		},
		{
			res: &Response{
				Version:      "HTTP/1.1",
				StatusCode:   200,
				ReasonPhrase: "OK",
				Headers: Headers{
					"Content-Length": []string{"15"},
				},
				Body: []byte("Hello from test"),
			},
			err: nil,
		},
		{
			res: nil,
			err: ErrNotFound,
		},
	}
	expects := []struct {
		b []byte
	}{
		{[]byte("HTTP/1.1 200 OK\r\nContent-Length: 15\r\n\r\nHello from test")},
		{[]byte("HTTP/1.1 200 OK\r\nContent-Length: 15\r\n\r\nHello from test")},
		{[]byte("HTTP/1.1 404 Not Found\r\nContent-Length: 0\r\n\r\n")},
	}

	for i, tt := range expects {
		var buf bytes.Buffer
		writeResponse(&buf, inputs[i].res, inputs[i].err)
		if !bytes.Equal(buf.Bytes(), tt.b) {
			t.Errorf("wants %s, but got %s\n", string(tt.b), string(buf.Bytes()))
		}
	}
}
