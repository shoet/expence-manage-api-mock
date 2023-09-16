package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/shoet/expense-manage-api-mock/handler"
)

type Server struct {
	srv http.Server
	l   net.Listener
}

func NewServer(port int) (*Server, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed listen port: %d", port)
	}
	mux := handler.NewMux()
	s := &Server{
		srv: http.Server{
			Handler: mux,
		},
		l: l,
	}
	return s, nil
}

func (s *Server) Run(ctx context.Context) error {

	if err := s.srv.Serve(s.l); err != nil {
		return fmt.Errorf("failed run server: %v", err)
	}

	return nil
}
