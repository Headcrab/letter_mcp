package app

import (
	"fmt"
	"letter_mcp/letters"
	"letter_mcp/mcp"
	"log/slog"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

// ServerConfig содержит конфигурацию сервера
type ServerConfig struct {
	Transport string
	TestMode  bool
}

// Server инкапсулирует логику запуска и настройки MCP сервера
type Server struct {
	config    ServerConfig
	mcpServer *server.MCPServer
	tools     mcp.ToolHandler
}

// NewServer создает новый экземпляр сервера
func NewServer(config ServerConfig) *Server {
	// Настраиваем логгер
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	// Создаем счетчик букв и форматтер
	counter := letters.NewCounter()
	formatter := letters.NewTextFormatter()

	// Создаем обработчик инструментов
	tools := mcp.NewToolHandler(counter, formatter)

	// Создаем MCP сервер
	mcpServer := server.NewMCPServer(
		"letter-counter",     // имя сервера
		"1.0.0",              // версия
		server.WithLogging(), // включаем логирование
	)

	// Регистрируем инструменты
	mcp.RegisterTools(mcpServer, tools)

	slog.Info("Инструмент count_letters добавлен")

	return &Server{
		config:    config,
		mcpServer: mcpServer,
		tools:     tools,
	}
}

// RunTests запускает тестовые примеры
func (s *Server) RunTests() {
	slog.Info("Запуск тестовых примеров")

	// Тестируем подсчет букв
	testWords := []struct {
		word    string
		letters string
	}{
		{"ПриВЕТ", "пр"},
		{"ПрограммироВАНИЕ", "рае"},
	}

	// Создаем счетчик и форматтер для тестов
	counter := letters.NewCounter()
	formatter := letters.NewTextFormatter()

	for _, test := range testWords {
		counts := counter.CountLetters(test.word, test.letters)
		result := formatter.Format(counts)
		fmt.Println("=== Тестирование ===")
		fmt.Println(result)
	}

	// Информация по использованию через MCP клиент
	fmt.Println("\n=== Тестирование через MCP сервер ===")
	fmt.Println("Для тестирования через MCP клиент можно использовать запрос:")
	fmt.Println(`{"jsonrpc":"2.0","id":"test","method":"mcp.call","params":{"tool":"count_letters","arguments":{"word":"ПриВЕТ","letters":"пр"}}}`)
	fmt.Println("\nЗапустите сервер без флага -test и отправьте запрос через клиент MCP")
}

// Start запускает сервер
func (s *Server) Start() error {
	if s.config.TestMode {
		s.RunTests()
		return nil
	}

	if s.config.Transport == "sse" {
		sseServer := server.NewSSEServer(s.mcpServer, server.WithBaseURL("http://localhost:8080"))
		slog.Info("SSE server listening on :8080")
		if err := sseServer.Start(":8080"); err != nil {
			return fmt.Errorf("ошибка запуска SSE сервера: %w", err)
		}
	} else {
		slog.Info("Запуск сервера подсчёта букв через stdio")
		if err := server.ServeStdio(s.mcpServer); err != nil {
			return fmt.Errorf("ошибка запуска stdio сервера: %w", err)
		}
	}

	return nil
}
