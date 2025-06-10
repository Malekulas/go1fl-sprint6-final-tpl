package main

import (
	"go1fl-sprint6-final-tpl/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	srv := server.CreateServer(logger)

	logger.Println("Сервер запускается на порту :8080...")
	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("ошибка при запуске сервера: ", err)
	}
}
