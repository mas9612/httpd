package http

import (
	"bufio"
	"net/url"
	"strconv"
	"strings"
)

const (
	bufSize = 4096
)

var (
	supportedMethods = []string{
		"HEAD",
		"GET",
		"POST",
	}
)

// Headers represents HTTP headers.
type Headers map[string][]string

// Add adds given key-value pair to Headers.
func (h Headers) Add(key, value string) {
	h[key] = append(h[key], value)
}

// Get returns the first element of given key. If given key doesn't exist in Headers, then empty string will be returned.
func (h Headers) Get(key string) string {
	if _, ok := h[key]; !ok {
		return ""
	}
	return h[key][0]
}

// Request represents the HTTP request.
type Request struct {
	Method  string
	Target  string
	Version string

	Headers Headers

	Body []byte
}

func parseRequestMessage(reader *bufio.Reader) (*Request, error) {
	var req Request
	req.Headers = make(Headers)

	if err := parseRequestLine(reader, &req); err != nil {
		return nil, err
	}

	if err := parseHeaders(reader, &req); err != nil {
		return nil, err
	}

	if lenStr := req.Headers.Get("Content-Length"); lenStr != "" {
		length, err := strconv.Atoi(lenStr)
		if err != nil {
			return nil, err
		}
		if err := readBody(reader, &req, length); err != nil {
			return nil, err
		}
	}
	if req.Headers.Get("Transfer-Encoding") != "" {
		// TODO: read body.
		// refer to RFC 7230 Section 4
	}

	return &req, nil
}

func parseRequestLine(reader *bufio.Reader, req *Request) error {
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	line = strings.TrimRight(line, "\r\n")

	tmp := strings.Split(line, " ")
	if len(tmp) != 3 {
		return ErrInvalidRequest
	}
	if !isValidMethod(tmp[0]) {
		return ErrMethodNotImplemented
	}

	// To accept also absolute-form.
	// Refer to RFC 7230 Section 5.3.2.
	url, err := url.Parse(tmp[1])
	if err != nil {
		return err
	}

	req.Method = tmp[0]
	req.Target = url.Path
	req.Version = tmp[2]

	return nil
}

func isValidMethod(method string) bool {
	for _, validMethod := range supportedMethods {
		if method == validMethod {
			return true
		}
	}
	return false
}

func parseHeaders(reader *bufio.Reader, req *Request) error {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimRight(line, "\r\n")

		// Empty line means the end of header section.
		// So we need to exit from this method.
		if line == "" {
			break
		}

		tmp := strings.Split(line, ":")

		if strings.Contains(tmp[0], " ") {
			return ErrInvalidRequest
		}

		req.Headers.Add(tmp[0], strings.Trim(tmp[1], " "))
	}

	if req.Headers.Get("Host") == "" {
		return ErrInvalidRequest
	}

	return nil
}

func readBody(reader *bufio.Reader, req *Request, length int) error {
	buf := make([]byte, bufSize)
	read := 0
	for read < length {
		n, err := reader.Read(buf)
		if err != nil {
			return err
		}
		read += n

		if read >= length {
			break
		}
	}
	req.Body = buf[:read]
	return nil
}

// Response represents the HTTP response.
type Response struct {
	Version      string
	StatusCode   int
	ReasonPhrase string
}
