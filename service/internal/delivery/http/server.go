package http

import (
	"orderAPI/service/internal/usecase"
	"net/http"
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
	s.httpServer.ListenAndServe()
}
