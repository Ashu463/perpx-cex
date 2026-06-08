package models

import "github.com/shopspring/decimal"

type Trade struct {
	TradeID string

	BuyOrderID  string
	SellOrderID string

	BuyerID  string
	SellerID string

	MarketID   string
	SellerSide PositionSide

	Price decimal.Decimal

	Quantity decimal.Decimal

	Timestamp int64
}

// XADD order_submissions * orderId ask1 userId seller1 marketId BTCUSDT side SELL price 100 quantity 5
// XADD order_submissions * orderId buy1 userId buyer1 marketId BTCUSDT side BUY price 100 quantity 5
