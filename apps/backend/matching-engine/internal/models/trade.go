package models

import "github.com/shopspring/decimal"

type Trade struct {
	TradeID string

	BuyOrderID  string
	SellOrderID string

	BuyerID  string
	SellerID string

	MarketID string

	Price decimal.Decimal

	Quantity decimal.Decimal

	Timestamp int64
}
