package main

import (
	"context"

	"github.com/redis/go-redis/v9"

	"matching-engine/internal/consumer"
	"matching-engine/internal/engine"
)

// for starting multiple parallel go routines
func main() {

	ctx := context.Background()

	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	eng := engine.NewEngine()
	// #TODO: do this by go routines
	consumer.StartConsumer(
		ctx,
		client,
		eng,
	)
}
