package orderbook

import "matching-engine/internal/models"

func AddBid(
	book *models.OrderBook,
	order *models.Order,
) {
	// #TODO max heap to be implemented in bid.go
	book.Bids = append(book.Bids, order)
}

func AddAsk(
	book *models.OrderBook,
	order *models.Order,
) {
	// #todo min heap to be implmented in ask.go
	book.Asks = append(book.Asks, order)
}
