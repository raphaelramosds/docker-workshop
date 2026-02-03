package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Producer Service")

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

	// Publicar mensagem
	err := rdb.Set(ctx, "msg_producer_2", "Mensagem publicada por Producer 2", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Mensagem publicada no Redis com sucesso!")
}
