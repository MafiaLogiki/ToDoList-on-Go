# Описание проекта
Микросервис для управления задачами на Go с аутентификацией через JWT и асинхронной обработкой через Kafka. 

## Стек технологий  
- **Backend**: Go (chi), PostgreSQL.  
- **Инфраструктура**: Docker и Docker Compose для контейнерезации, Kafka в качестве брокера сообщений, Nginx как обратный прокси  
- **Инструменты**: Git для контроля версий, Jira для отслеживания прогресса

## Функционал  
- Регистрация и аутентификация пользователей.  
- Создание и управление задачами 

## Архитектура


## Запуск
1. Склонируйте репозиторий:
```bash  
git clone https://github.com/MafiaLogiki/ToDoList-on-Go.git
cd ToDoList-on-Go
```
2. Запустите все сервисы локально через Docker Compose:
```bash
docker-compose up --build
```
Важно: Порты самих микросервисов (Auth-service, Register-service, Task-service) не пробрасываются наружу напрямую. Доступ к API возможен только через Nginx, который настроен на порт 8080. Например, чтобы проверить API, используйте URL вроде http://localhost:8080/tasks или http://localhost:8080/auth. Файл конфигурации Nginx находится в nginx/nginx.conf. Проверьте статус контейнеров:
```bash
docker ps
```

## API Endpoints
Все запросы должны поступать к Nginx на порт 8080. Ниже таблица основных endpoint'ов:

|Сервис|Метод|Endpoint|Описание|Параметры|Пример ответа|Статусы ответа|
|------|-----|--------|--------|---------|-------------|--------------|
|Auth-service|POST|/api/login|Авторизация пользователя. При успехе устанавливает HTTP-only куки с JWT токеном|username, password(хэш)|{"status": "success"} ||
|Register-serivce|POST|/api/register|Регистрация нового пользователя. При успехе создает JWT токен и устанавливает его в HTTP-only куки |username, password (хэш)|{"status": "success"} ||
|Task-service|GET|/api/tasks|Список задач авторизованного пользователя по его id, взятого из JWT токена||{{"id":1,"title":"test title","descriprion":"test description","status":"in progress","userId":1}, ...}||
|Task-service|GET|/api/tasks/{id}|Одна конкретная задача. Задача должна принадлежать пользователю ||{"id":1,"title":"test title","descriprion":"test description","status":"in progress","userId":1}||
|Task-service|POST|/api/tasks/create|Создание задачи|title, description (по умолчанию статус у всех задач: todo)|{"status": "success"}||



## Структура проекта
```
ToDoList-on-Go/
├── common/             # Общие для всех микросервисов функции
│   ├── auth/               # Все, что связано с токенами
│   ├── domain/             # Все структуры
│   ├── logger/             # Конфигурация логера
│   └── middleware/         # Промежуточное ПО (аутентификация)
├── docker-compose.yml
├── go.work
├── go.work.sum
├── init.sql
├── nginx/              # Все, что связано с nginx
│   ├── Dockerfile      
│   ├── nginx.conf
│   └── static/             # Статические файлы для сервисов
├── README.md
└── services/
    ├── auth-service/   # Сервис авторизации 
    │   ├── cmd/        # Точка входа
    │   ├── internal/   # Не импортируемые модули (обработчики, база данных)
    │   ├── config.yml
    │   ├── go.mod
    │   ├── go.sum
    │   └── auth-service.Dockerfile
    ├── register-service
    │   ├── cmd/
    │   ├── internal/
    │   ├── config.yml
    │   ├── go.mod
    │   ├── go.sum
    │   └── register-service.Dockerfile
    └── task-service
        ├── cmd/ 
        ├── internal/
        ├── config.yml
        ├── go.mod
        ├── go.sum
        └── task-service.Dockerfile
```

## Тестирование
Тесты в разработке. Планируется запускать через Docker Compose

## Необходимо:

- Go 1.23+
- Docker (версия 20.10 и выше)
- Docker compose (Версия 2.0 и выше)

Все остальные сервисы (Kafka, Zookeeper, PostgreSQL, Nginx) автоматически устанавливаются и конфигурируются через `docker-compose.yml`. Важен доступ в интернет для скачивания образов