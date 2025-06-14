package server

import (
	"go1fl-sprint6-final-tpl/internal/handlers"
	"log"
	"net/http"
	"time"
)

type MyServer struct {
	Logger *log.Logger
	Server http.Server
}

func CreateServer(l *log.Logger) *MyServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.PageHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	return &MyServer{
		Server: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     l,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		Logger: l,
	}
}