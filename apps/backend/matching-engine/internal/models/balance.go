package models

import "github.com/shopspring/decimal"

type Balance struct {
	UserID string

	AvailableBalance decimal.Decimal

	LockedBalance decimal.Decimal

	Equity decimal.Decimal
}
