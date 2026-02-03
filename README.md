# hitalent

## Описание

Тестовое задание: API чатов и сообщений

## Запуск

Создает локальный .env из примера с переменными окружения.

```ash
cp .env.example .env
```

Запускает сервисы через Docker Compose.

```ash
docker compose up
```

Собирает утилиту миграций goose внутри контейнера.

```ash
docker compose exec -T go go build -o goose.exe ./cmd/goose
```

Применяет миграции базы данных.

```ash
docker compose exec -T go go run ./cmd/goose -dir ./database/migrations up
```

Запускает тесты контроллера чата.

```ash
docker compose exec -T go go test ./app/Http/Controllers/Chats -run TestChatController
```

## Примеры запросов

Создание чата.

```ash
postman request POST 'http://localhost:8080/chats' \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --body '{
    "title": "My chat"
}'
```

Получение чата.

```ash
postman request 'http://localhost:8080/chats/1'
```

Создание сообщения в чате.

```ash
postman request POST 'http://localhost:8080/chats/1/messages' \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --body '{
    "text": "My message"
}'
```

Удаление чата.

```ash
postman request DELETE 'http://localhost:8080/chats/1'
```
