# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /src

# copiar módulos e baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# copiar restante do código e compilar apenas o pacote producer
COPY . .
WORKDIR /src/producer
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/producer

# Final stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/producer /bin/producer

# Variáveis de ambiente padrão (sobrescritas por docker-compose)
ENV REDIS_HOST=redis \
    REDIS_PORT=6379 \
    REDIS_PASSWORD=local

ENTRYPOINT ["/bin/producer"]
