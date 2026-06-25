package server

import (
	"subscriptionmanager/internal/infrastructure/config"
	"time"
)

type Server struct {
	httpTimeout time.Duration
}

func New(cfg *config.ServerConfig) *Server {
	return &Server{}
}
