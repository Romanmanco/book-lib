# Мини REST api для работы с библиотекой книг и CRUD операции

## Переменные окружения, для настройки подключения бд, ссылки на внешний апи и хост: [.env](config%2F.env)

## Загрузка библиотек для работы со сваггером
- go get -u github.com/swaggo/swag/cmd/swag
- go get -u github.com/swaggo/echo-swagger

## Генерация документации
- swag init -g ./app/main.go -o ./docs

## Проверка развернутого сваггера
- http://localhost:8080/swagger/index.html

## Пример подключения к базе данных: docs/database.md