package models

type OrderBook struct {

	// Buy orders
	Bids []*Order

	// Sell orders
	Asks []*Order
}
