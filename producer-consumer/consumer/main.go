package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Consumer Service")

	// Instanciar Redis Client
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "local",
			DB:       0,
			Protocol: 2,
		})

	// Criar contexto
	ctx := context.Background()

	// Ler mensagem
	val, err := rdb.Get(ctx, "msg_producer_2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mensagem lida do Redis:", val)
}
