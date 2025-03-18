# Сервер подсчета букв MCP

[![Go Version](https://img.shields.io/github/go-mod/go-version/Headcrab/letter_mcp)](https://go.dev)
[![License](https://img.shields.io/github/license/Headcrab/letter_mcp)](LICENSE)
[![Coverage](https://codecov.io/gh/Headcrab/letter_mcp/graph/badge.svg?token=WSRWMHXMTA)](https://codecov.io/gh/Headcrab/letter_mcp)

Этот проект представляет собой пример MCP-совместимого сервера реализующего  подсчет букв в словах.

## Возможности

- Подсчет заданных букв в словах с разделением по регистру
- Поддержка русских букв
- Запуск в режиме тестирования для демонстрации работы
- Поддержка различных транспортов (stdio и SSE)

## Структура проекта

Проект имеет четкое разделение ответственности по принципам SOLID:

```tree
letter_mcp/
├── app/            # Основная логика приложения
│   └── server.go   # Настройка и запуск сервера
├── letters/        # Пакет для работы с буквами
│   ├── counter.go  # Счетчик букв
│   └── formatter.go # Форматирование результатов
├── mcp/            # Работа с протоколом MCP
│   └── tools.go    # Инструменты MCP
└── main.go         # Точка входа
```

## Использование

### Сборка

```bash
go build -o letter_mcp
```

### Запуск

Запуск через stdio (по умолчанию):

```bash
./letter_mcp
```

Запуск через SSE:

```bash
./letter_mcp -t sse
```

Запуск в тестовом режиме:

```bash
./letter_mcp -test
```

Запуск в Docker

```bash
docker build -t letter_mcp .
docker run -d -p 8080:8080 --name letter_mcp letter_mcp:latest
```

Запуск с Docker Compose

```bash
# Запуск со стандартным портом 8080
docker-compose up -d

# Запуск с пользовательским портом
PORT=9090 docker-compose up -d
```

## Формат запросов и ответов

### Запрос

```json
{
  "jsonrpc": "2.0",
  "id": "test",
  "method": "mcp.call",
  "params": {
    "tool": "count_letters",
    "arguments": {
      "word": "ПриВЕТ",
      "letters": "пр"
    }
  }
}
```

### Ответ

```json
{
  "jsonrpc": "2.0",
  "id": "test",
  "result": {
    "type": "text",
    "text": "Результаты подсчёта в слове 'ПриВЕТ':\n'п' (строчная): 0\n'П' (заглавная): 1\n'п' (всего): 1\n\n'р' (строчная): 1\n'Р' (заглавная): 0\n'р' (всего): 1\n\n"
  }
}
```

## Настройка MCP клиента

```json
{
  "mcpServers": {
    "count_letters": {
      "command": "/your/path/to/letter_mcp.exe",
      "args": [],
      "disabled": false,
      "alwaysAllow": []
    }
  }
}
```

## Настройка MCP клиента c sse

```json
{
  "mcpServers": {
    "count_letters": {
      "url": "http://localhost:8080/sse",
      "env": {
        "API_KEY": ""
      }
    }
  }
}
```

## Лицензия

MIT License. См. файл [LICENSE](LICENSE) для подробностей.

## Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для ваших изменений
3. Внесите изменения и создайте pull request

## Контакты

Создайте issue в репозитории для сообщения о проблемах или предложений по улучшению.

## Спасибо

- [@Headcrab](https://github.com/Headcrab)
