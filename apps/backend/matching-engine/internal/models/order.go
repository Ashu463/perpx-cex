package models

import (
	"github.com/shopspring/decimal"
)

type OrderType string

const (
	Market OrderType = "MARKET"
	Limit  OrderType = "LIMIT"
)

type OrderSide string

const (
	Buy  OrderSide = "BUY"
	Sell OrderSide = "SELL"
)

type PositionSide string

const (
	Long  PositionSide = "LONG"
	Short PositionSide = "SHORT"
)

type Status string

const (
	Pending          Status = "PENDING"
	Partially_filled Status = "PARTIALLY_FILLED"
	Cancelled        Status = "CANCELLED"
	Filled           Status = "FILLED"
)

type Order struct {
	OrderID      string
	UserID       string
	MarketID     string
	OrderType    OrderType
	Side         OrderSide
	PositionSide PositionSide

	Price             decimal.Decimal
	Quantity          decimal.Decimal
	RemainingQuantity decimal.Decimal
	Status            string
	Leverage          int64

	Margin    decimal.Decimal
	CreatedAt int64
}
