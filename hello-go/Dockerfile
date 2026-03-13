# Estagio de build
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY main.go .
RUN go build -o hello-go ./main.go

# Estagio de runtime
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/hello-go ./hello-go

CMD ["./hello-go"]
