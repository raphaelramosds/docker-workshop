package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
			Protocol: 2,
		})
	ctx := context.Background()
	key := "buffer"
	lockKey := "buffer_lock"
	N := 5

	for i := 0; ; i++ {
		// Tenta adquirir o lock (SetNx - Set if Not exists)
		ok, err := rdb.SetNX(ctx, lockKey, "producer", 5*time.Second).Result()
		if err != nil || !ok {
			// Caso nao consiga, dorme por 100ms e tenta novamente
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Remove e Insere os itens no buffer
		rdb.Del(ctx, key)
		for j := 0; j < N; j++ {
			rdb.RPush(ctx, key, i*N+j)
			fmt.Printf("PRODUCER: Produzido: %d\n", i*N+j)
		}

		// Libera o lock
		rdb.Del(ctx, lockKey)
		fmt.Printf("PRODUCER: Lock liberado.\n")

		// Dorme por 1s antes de produzir novamente
		time.Sleep(1 * time.Second)
	}
}
