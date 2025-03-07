# Регистрация и Авторизация + Аутентификация

Данный проект реализует функционал регистрации, авторизации и аутентификации пользователей с использованием Golang, Gin и Keycloak.

## Стек технологий

- **Golang** - язык программирования
- **Gin** - веб-фреймворк для Go
- **Keycloak** - система управления идентификацией и доступом
- **Docker/Docker-compose** - для контейнеризации приложения

## Установка

Для запуска проекта вам потребуется Docker и Docker Compose. Убедитесь, что они установлены на вашем компьютере.

### Клонирование репозитория -> docker-compose up --build

Это создаст и запустит все необходимые контейнеры.

## Настройка Keycloak

1. Откройте Keycloak в браузере (лучше использовать локальный ip(настройте под себя docker-compose.yaml))
2. Создайте новый Realm или останьтесь в master.
3. Создайте нового клиента, указав URL вашего приложения.
4. Настройте роли и пользователей по мере необходимости.

## Использование

После успешного запуска приложения вы сможете получить доступ к следующим эндпоинтам:

- /register - для регистрации нового пользователя
- /login - для авторизации пользователя
