package main

import (
	"context"
	"matching-engine/internal"

	"github.com/redis/go-redis/v9"
)

// for starting multiple parallel go routines
func main() {

	ctx := context.Background()

	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	eng := internal.NewEngine()
	// #TODO: do this by go routines
	internal.StartConsumer(
		ctx,
		client,
		eng,
	)
}
