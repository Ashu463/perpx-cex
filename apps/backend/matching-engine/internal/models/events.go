package models

import "github.com/shopspring/decimal"

type Event struct {
}
type OrderFilledEvent struct {

	// Event payload
}

type OrderPartiallyFilledEvent struct {

	// Event payload
}

type PositionUpdatedEvent struct {

	// Event payload
}

type MarginUnlockedEvent struct {

	// Event payload
}

type TradeExecutedEvent struct {

	// Event payload
}

type OrderbookUpdatedEvent struct {

	// Event payload
}

type MarkPriceUpdatedEvent struct {

	// Event payload
}

type OrderCreatedEvent struct {
	Event        string
	OrderID      string
	UserID       string
	MarketID     string
	Side         OrderSide
	PositionSide PositionSide
	Type         OrderType
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	Leverage     int
	Margin       decimal.Decimal
}
