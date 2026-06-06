package internal

import (
	"fmt"
	"matching-engine/internal/models"

	"github.com/redis/go-redis/v9"
)

func (p *Publisher) PublishTradeResult(trade *models.Trade) error {

	// Write individual event
	// trade_results stream

	// consume this trade in nest app and add this to postgres
	result, err := p.client.XAdd(
		p.ctx,
		&redis.XAddArgs{
			Stream: "Trade_Results",
			Values: map[string]interface{}{
				"buyerId":     trade.BuyerID,
				"sellerId":    trade.SellerID,
				"buyOrderId":  trade.BuyOrderID,
				"sellOrderId": trade.SellOrderID,

				"marketId": trade.MarketID,

				"price": trade.Price.String(),

				"quantity": trade.Quantity.String(),
			},
		},
	).Result()

	if err != nil {
		return err
	}
	fmt.Println("Trade published, ", result)

	return nil

}

func (p *Publisher) PublishOrderbookUpdate() {

	// Publish global event
}

func (p *Publisher) PublishTradeExecuted() {

	// Publish trade broadcast
}

func (p *Publisher) PublishMarkPriceUpdate() {

	// Publish mark price update
}
