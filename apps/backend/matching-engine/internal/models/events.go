package models

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
	Side         string
	PositionSide string
	Type         string
	Price        string
	Quantity     string
	Leverage     string
	Margin       string
}
