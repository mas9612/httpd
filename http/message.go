package http

import (
	"bufio"
	"strings"
)

type Request struct {
	Method  string
	Target  string
	Version string
}

func parseRequestMessage(reader *bufio.Reader) (*Request, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimRight(line, " \r\n")

	tmp := strings.Split(line, " ")
	if len(tmp) != 3 {
		return nil, ErrInvalidRequest
	}
	if !isValidMethod(tmp[0]) {
		return nil, ErrMethodNotImplemented
	}

	return &Request{
		Method:  tmp[0],
		Target:  tmp[1],
		Version: tmp[2],
	}, nil
}

func isValidMethod(method string) bool {
	for _, validMethod := range supportedMethods {
		if method == validMethod {
			return true
		}
	}
	return false
}

type Response struct {
	Version      string
	StatusCode   int
	ReasonPhrase string
}

// RFC7230 3
// A sender MUST NOT send whitespace between the start-line and the
//    first header field.  A recipient that receives whitespace between the
//    start-line and the first header field MUST either reject the message
//    as invalid or consume each whitespace-preceded line without further
//    processing of it (i.e., ignore the entire line, along with any
//    subsequent lines preceded by whitespace, until a properly formed
//    header field is received or the header section is terminated).
