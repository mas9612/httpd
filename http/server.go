package http

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// Server is the HTTP server implementation.
type Server struct {
	Port         int
	DocumentRoot string
	listener     net.Listener
}

// Serve starts hTTP request handling.
func (s *Server) Serve() error {
	if _, err := os.Stat(s.DocumentRoot); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}
	s.listener = listener

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return nil
		}
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		req, err := parseRequestMessage(reader)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("handleRequest: %+v\n", err)
			writeResponse(conn, nil, err)
			return
		}
		fmt.Printf("req: %+v\n", req)

		res, err := buildResponseFromRequest(req, s.DocumentRoot)
		fmt.Printf("res: %+v\n", res)
		writeResponse(conn, res, err)

		// For now, we close connection if some error occured while processing request.
		if req == nil {
			return
		}

		// If connection option is close, we will close this connection.
		// Note: Connection header value is case-insensitive.
		connection := strings.ToLower(req.Headers.Get("Connection"))
		if connection == "close" {
			return
		}
	}
}

func writeResponse(w io.Writer, res *Response, err error) {
	if err == nil {
		fmt.Fprintf(w, "%s %d %s\r\n", res.Version, res.StatusCode, res.ReasonPhrase)
		for k, v := range res.Headers {
			fmt.Fprintf(w, "%s: %s\r\n", k, strings.Join(v, ", "))
		}
		fmt.Fprintf(w, "\r\n")

		w.Write(res.Body)
	} else {
		var statusCode int
		var reasonPhrase string

		switch err {
		case ErrInvalidRequest:
			statusCode = StatusBadRequest
			reasonPhrase = reasonPhrases[StatusBadRequest]
		case ErrNotFound:
			statusCode = StatusNotFound
			reasonPhrase = reasonPhrases[StatusNotFound]
		case ErrInternalServerError:
			statusCode = StatusInternalServerError
			reasonPhrase = reasonPhrases[StatusInternalServerError]
		case ErrMethodNotImplemented:
			statusCode = StatusMethodNotImplemented
			reasonPhrase = reasonPhrases[StatusMethodNotImplemented]
		}

		// TODO: don't use constant for HTTP version
		fmt.Fprintf(w, "%s %d %s\r\n", "HTTP/1.1", statusCode, reasonPhrase)
		fmt.Fprintf(w, "Content-Length: 0\r\n\r\n")
	}
}
