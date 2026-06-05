package internal

import (
	"fmt"
	"matching-engine/internal/models"
)

type Engine struct {
	OrderBooks map[string]*models.OrderBook
}

func NewEngine() *Engine {
	return &Engine{
		OrderBooks: make(
			map[string]*models.OrderBook,
		),
	}
}

func (e *Engine) ProcessOrder(
	order *models.Order,
) {

	book, exists := e.OrderBooks[order.MarketID]

	if !exists {

		book = models.NewOrderBook()

		fmt.Println(
			"CREATING ORDERBOOK:",
			order.MarketID,
		)

		e.OrderBooks[order.MarketID] = book
	}

	Match(
		book,
		order,
	)
	book.Print(order.MarketID)
}
