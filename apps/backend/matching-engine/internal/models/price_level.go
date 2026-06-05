package models

import (
	"github.com/google/btree"
	"github.com/shopspring/decimal"
)

type PriceLevel struct {
	Price  decimal.Decimal
	Orders []*Order
}

func (p *PriceLevel) Less(other btree.Item) bool {

	return p.Price.LessThan(
		other.(*PriceLevel).Price,
	)
}
