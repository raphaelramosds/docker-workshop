# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
COPY producer/main.go ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/producer

# Final stage
FROM alpine:3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/producer /bin/producer

ENTRYPOINT ["/bin/producer"]

