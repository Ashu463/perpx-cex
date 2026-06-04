package events

import (
	"matching-engine/internal/models"

	"github.com/redis/go-redis/v9"
)

type Publisher struct {

	// Redis client

	// Stream names

	// PubSub channel names
}

func (p *Publisher) PublishTradeResult(client *redis.Client, trade *models.Trade) {

	// Write individual event
	// trade_results stream

	// consume this trade in nest app and add this to postgres
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
