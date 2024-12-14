package http

import (
	"orderAPI/service/internal/usecase"
	"net/http"
	"log"
	"context"
)

type Server struct {
	httpServer	*http.Server 
}

func NewServer(ucOrder usecase.Order) *Server {
	handler := NewHandler(ucOrder)
	router := handler.InitRouter()
	server :=  &http.Server{
		Handler:	router,
		Addr:		":8080",
	}
	return &Server{
		httpServer: server,
	}
}


func (s *Server) StartServer() {
	log.Println("start server on port:", s.httpServer.Addr)
	s.httpServer.ListenAndServe()
}

func (s *Server)Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}