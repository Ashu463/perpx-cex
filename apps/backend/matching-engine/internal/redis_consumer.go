package internal

import (
	"context"
	"fmt"
	"matching-engine/internal/models"
	// "time"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

func StartConsumer(
	ctx context.Context,
	client *redis.Client,
	engine *Engine,
	p *Publisher,
) {
	fmt.Println("start consumer called")

	for {

		result, err := client.XReadGroup(
			ctx,
			&redis.XReadGroupArgs{
				Group:    "matching-engine",
				Consumer: "engine-1",
				Streams: []string{
					"order_submissions",
					">",
				},
				// Block: time.Second * 5,
			},
		).Result()

		if err != nil {
			continue
		}

		for _, stream := range result {

			for _, msg := range stream.Messages {

				price, _ := decimal.NewFromString(
					msg.Values["price"].(string),
				)

				qty, _ := decimal.NewFromString(
					msg.Values["quantity"].(string),
				)

				order := &models.Order{
					OrderID: msg.Values["orderId"].(string),

					UserID: msg.Values["userId"].(string),

					MarketID: msg.Values["marketId"].(string),

					Side: models.OrderSide(
						msg.Values["side"].(string),
					),

					Price: price,

					Quantity:          qty,
					RemainingQuantity: qty,
				}

				fmt.Println("\n==========================")
				fmt.Println("ORDER RECEIVED")
				fmt.Println("OrderID :", order.OrderID)
				fmt.Println("UserID  :", order.UserID)
				fmt.Println("Market  :", order.MarketID)
				fmt.Println("Side    :", order.Side)
				fmt.Println("Price   :", order.Price)
				fmt.Println("Qty     :", order.Quantity)
				fmt.Println("==========================")

				engine.ProcessOrder(order, p)

				client.XAck(
					ctx,
					"order_submissions",
					"matching-engine",
					msg.ID,
				)
			}
		}
	}
}
