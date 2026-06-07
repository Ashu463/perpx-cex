package internal

import (
	"context"
	"fmt"
	"matching-engine/internal/models"

	"github.com/redis/go-redis/v9"
)

type Engine struct {
	// map<marketId, orderbook>
	OrderBooks map[string]*models.OrderBook

	// map<userId, Balance table>
	Balances map[string]*models.Balance

	// map<marketId, map<userId, Positions table>>
	Positions map[string]map[string]*models.Position
}

func NewEngine() *Engine {
	return &Engine{
		OrderBooks: make(
			map[string]*models.OrderBook,
		),
		Balances:  make(map[string]*models.Balance),
		Positions: make(map[string]map[string]*models.Position),
	}
}

type Publisher struct {
	client *redis.Client
	ctx    context.Context
	engine *Engine

	// Redis client
	// Stream names
	// PubSub channel names
}

func NewPublisher(
	client *redis.Client,
	ctx context.Context,
	engine *Engine,
) *Publisher {

	return &Publisher{
		client: client,
		ctx:    ctx,
		engine: engine,
	}
}

func (e *Engine) ProcessOrder(
	order *models.Order,
	p *Publisher,
) {
	book, exists := e.OrderBooks[order.MarketID]
	// fmt.Println("Process order called")

	if !exists {

		book = models.NewOrderBook()

		fmt.Println(
			"CREATING ORDERBOOK:",
			order.MarketID,
		)

		e.OrderBooks[order.MarketID] = book
	}
	fmt.Println("match called from engine")

	// udpate in memory balances here, this is the common step of all three cases,
	SyncBalances(*order, e)
	Match(
		book,
		order,
		e,
		p,
	)
	book.Print(order.MarketID)
}
