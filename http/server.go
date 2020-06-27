package http

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// Server is the HTTP server implementation.
type Server struct {
	Port     int
	listener net.Listener
}

// Serve starts hTTP request handling.
func (s *Server) Serve() error {
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
		}
		fmt.Printf("req: %+v\n", req)

		// TODO: send resposne back to the client

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
