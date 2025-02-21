# Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG APP_VERSION=dev
RUN go build -ldflags="-X 'chappie_bot/config.APP_VERSION=${APP_VERSION}'" -o chappie_bot ./cmd/bot/main.go

# Runtime
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/chappie_bot .
COPY .env /app/.env
CMD ["./chappie_bot"]