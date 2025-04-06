# Описание проекта
Микросервис для управления задачами на Go с аутентификацией через JWT и асинхронной обработкой через Kafka. 

## Стек технологий  
- **Backend**: Go (chi), PostgreSQL.  
- **Инфраструктура**: Docker и Docker Compose для контейнерезации, Kafka, Nginx как обратный прокси  
- **Инструменты**: Git для контроля версий, Jira для отслеживания прогресса

## Функционал  
- Регистрация и аутентификация пользователей.  
- Создание и управление задачами 

## Запуск
1. Склонируйте репозиторий:
```bash  
git clone https://github.com/MafiaLogiki/ToDoList-on-Go.git
cd ToDoList-on-Go
```
2. Запустите все сервисы локально через Docker Compose:
```bash
docker compose up --build
```
Важно: Порты самих микросервисов (Auth-service, Register-service, Task-service) не пробрасываются наружу напрямую. Доступ к API возможен только через Nginx, который настроен на порт 8080. Например, чтобы проверить API, используйте URL вроде http://localhost:8080/tasks или http://localhost:8080/auth. Файл конфигурации Nginx находится в nginx/nginx.conf. Проверьте статус контейнеров:
```bash
docker ps
```

## API Endpoints
Все запросы должны поступать к Nginx на порт 8080. Ниже таблица основных endpoint'ов:

|Сервис|Метод|Endpoint|Описание|Параметры|Пример ответа|
|------|-----|--------|--------|---------|-------------|
|Auth-service|POST|/login|Авторизация пользователя|username, password(в виде хэша)|{"token", "jwt_token_here"} |



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
