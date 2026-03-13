FROM golang:1.22-alpine

WORKDIR /app

COPY main.go .
RUN go build -o hello-go ./main.go

CMD ["./hello-go"]