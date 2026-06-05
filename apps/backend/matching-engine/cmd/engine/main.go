package main

import (
	"context"
	"fmt"
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

	fmt.Println("Engine starts to listen through redis stream")
	eng := internal.NewEngine()
	// #TODO: do this by go routines
	internal.StartConsumer(
		ctx,
		client,
		eng,
	)
}

/* Test cases
XADD order_submissions * orderId buy1 userId u1 marketId BTCUSDT side BUY price 100000 quantity 1

XADD order_submissions * orderId buy2 userId u2 marketId BTCUSDT side BUY price 110000 quantity 1

XADD order_submissions * orderId buy3 userId u3 marketId BTCUSDT side BUY price 110000 quantity 1

XADD order_submissions * orderId sell1 userId u4 marketId BTCUSDT side SELL price 120000 quantity 1

XADD order_submissions * orderId sell2 userId u5 marketId BTCUSDT side SELL price 115000 quantity 1

XADD order_submissions * orderId ethbuy1 userId u6 marketId ETHUSDT side BUY price 3000 quantity 1

XADD order_submissions * orderId sell3 userId u7 marketId BTCUSDT side SELL price 100000 quantity 1
*/
