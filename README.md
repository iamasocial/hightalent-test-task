# hightalent-test-task

Простой Q&A сервис написанный на Golang с использованием PostgreSQL. Предоставляет HTTP API для работы с вопросами и ответами

## Возможности
- CRUD операции для вопросов и ответов
- HTTP API
- Graceful Shutdown
- Миграция базы данных через Goose
- Docker контейнеризация
- Unit-тесты для сервисов и HTTP обработчиков

## Структура проекта
- `cmd/app/main.go` - точка входа в приложение
- `config/config.yaml` - конфигурация приложения 
- `internal/` 
  - `config/` - загрузка конфигурации приложения
  - `db/` - подключение к PostgreSQL и модели данных
  - `entities/` - бизнес-сущности
  - `logger/` - интерфейс и реализация интерфейса логгера
  - `migrations/` - выполнение миграций базы данных
  - `repository/` - работа с базой данных
  - `server/` - создание сервера
  - `service/` - бизнес-логика приложения
  - `transport/` - DTO и HTTP обработчики
- `migrations/` - файлы миграций
- `.env` - файл с переменными окружения
- `docker-compose.yaml` - Оркестрация Docker-контейнеров приложения и базы данных PostgreSQL
- `dockerfile` - Сборка Docker-образа приложения Go с включением бинарника и конфигурации

## Запуск
### Клонирование репозитория
```bash
git clone https://github.com/iamasocial/hightalent-test-task
```
Перейдите в директорию приложения
```bash
cd hightalent-test-task
```
Создайте `.env` файл с переменными окружения со следующим содеражинием
```.env
DB_USER=user
DB_PASSWORD=password
DB_NAME=name
```
Запустите Docker Compose
```bash
docker compose up --build
```

## API
Приложение слушает порт `8080`

## Эндпоинты
### Работа с вопросами
- `GET /questions/` - возвращает список всех вопросов
```bash
curl -X GET http://localhost:8080/questions/
[{"id":2,"text":"whats up?","created_at":"2025-11-25T12:00:44.877523Z"},{"id":3,"text":"how long?","created_at":"2025-11-25T12:01:48.724884Z"}]
```
- `POST /questions/` - создает новый вопрос
```bash
curl -X POST http://localhost:8080/questions/ -H "Content-Type: application/json" -d '{"text":"whats up?"}'

{"id":2,"text":"whats up?","created_at":"2025-11-25T12:00:44.877523801Z"}
```
- `GET /questions/{id}` - возвращает конкретный вопрос со всеми ответами на него
```bash
curl -X GET http://localhost:8080/questions/3

{"id":3,"text":"how long?","created_at":"2025-11-25T12:01:48.724884Z","answers":[{"id":1,"question_id":3,"user_id":"1234_uuid","text":"3 min","created_at":"2025-11-25T12:04:33.219354Z"},{"id":2,"question_id":3,"user_id":"1234_uuid","text":"4 min","created_at":"2025-11-25T12:04:41.678202Z"}]}

```
- `DELETE /questions/{id}` - удаляет конкретный вопрос и все ответы на него
```bash
curl -X DELETE http://localhost:8080/questions/3
```
### Работа с ответами
- `POST /questions/{id}/answers/` - создает новый ответ на конкретный вопрос
```bash
curl -X POST http://localhost:8080/questions/3/answers/ -H "Content-Type: application/json" -d '{"user_id":"1234_uuid","text":"4 min"}'

{"id":2,"question_id":3,"user_id":"1234_uuid","text":"4 min","created_at":"2025-11-25T12:04:41.678202571Z"}

```
- `GET /answers/{id}` - возвращает конкретный ответ
```bash
curl -X GET http://localhost:8080/answers/2

{"id":2,"question_id":3,"user_id":"1234_uuid","text":"4 min","created_at":"2025-11-25T12:04:41.678202Z"}

```
- `DELTE /answers/{id}` - удаляет конкретный ответ
```bash
curl -X DELETE http://localhost:8080/answers/2
```
