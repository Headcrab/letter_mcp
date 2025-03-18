package main

import (
	"flag"
	"letter_mcp/app"
	"log/slog"
	"os"
)

func main() {
	var transport string
	var testMode bool
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.StringVar(&transport, "transport", "stdio", "Transport type (stdio or sse)")
	flag.BoolVar(&testMode, "test", false, "Run in test mode")
	flag.Parse()

	// Настраиваем текстовый логгер с указанием времени
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	// Конфигурация сервера
	config := app.ServerConfig{
		Transport: transport,
		TestMode:  testMode,
	}

	// Создаем и запускаем сервер
	server := app.NewServer(config)
	if err := server.Start(); err != nil {
		slog.Error("Ошибка сервера", "err", err)
		os.Exit(1)
	}
}
