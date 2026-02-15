package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "MORSE: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	logger.Println("Сервер успешно запущен")

	<-quit
	logger.Println("Получен сигнал остановки сервера")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Ошибка при остановке сервера: %v\n", err)
	}

	logger.Println("Сервер остановлен")
}
