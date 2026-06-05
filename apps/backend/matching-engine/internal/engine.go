package internal

import "matching-engine/internal/models"

// "matching-engine/internal/models"

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

		book = &models.OrderBook{}

		e.OrderBooks[order.MarketID] = book
	}

	Match(
		book,
		order,
	)
}
