package models

import "github.com/shopspring/decimal"

type Position struct {
	UserID string

	MarketID string

	Quantity decimal.Decimal

	EntryPrice decimal.Decimal

	MarkPrice decimal.Decimal

	LiquidationPrice decimal.Decimal

	InitialMargin decimal.Decimal

	MaintenanceMargin decimal.Decimal

	UnrealizedPnL decimal.Decimal

	RealizedPnL decimal.Decimal
}
