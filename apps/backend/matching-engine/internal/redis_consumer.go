package internal

import (
	"context"
	"matching-engine/internal/models"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

func StartConsumer(
	ctx context.Context,
	client *redis.Client,
	eng *Engine,
) {

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

					Quantity: qty,
				}

				eng.ProcessOrder(
					order,
				)

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
