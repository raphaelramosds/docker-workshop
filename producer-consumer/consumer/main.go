package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Consumer Service")

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

	for {
		// Tenta adquirir o lock (SetNx - Set if Not exists)
		ok, err := rdb.SetNX(ctx, lockKey, "consumer", 5*time.Second).Result()
		if err != nil || !ok {
			// Caso nao consiga, dorme por 100ms e tenta novamente
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Remove e Consome os itens do buffer
		for i := 0; i < N; i++ {
			val, err := rdb.LPop(ctx, key).Result()
			if err != nil {
				panic(err)
			}
			fmt.Printf("CONSUMER: Consumido: %s\n", val)
		}

		// Libera o lock
		rdb.Del(ctx, lockKey)
		fmt.Printf("CONSUMER: Lock liberado.\n")

		// Dorme por 1s antes de consumir novamente
		time.Sleep(1 * time.Second)
	}
}