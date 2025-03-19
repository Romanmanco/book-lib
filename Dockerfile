####################      BUILD Stage      ####################

FROM golang:1.23-alpine AS build-env

# зависимости, но без лишних пакетов
RUN apk add --no-cache git bash gcc musl-dev

# рабочая директорию
WORKDIR /app

# Файлы go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Исходный код
COPY . .

# Соборка приложения
RUN go build -o library ./app/main.go

####################      RUN Stage      ####################

FROM alpine:latest

# Установка необходимых пакетов
RUN apk add --no-cache ca-certificates bash curl

# Установка рабочей директории
WORKDIR /app

# Только нужные файлы из этапа сборки
COPY --from=build-env /app/library /app/library
COPY config/.env /app/.env
RUN cat /app/.env

# Открыть порт (если требуется)
EXPOSE 8080

# Запуск приложения
CMD ["/app/library"]
