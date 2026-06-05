package models

import (
	"fmt"

	"github.com/google/btree"
)

type BidTree struct {
	Tree *btree.BTree
}

type AskTree struct {
	Tree *btree.BTree
}

type OrderBook struct {
	Bids *BidTree
	Asks *AskTree
}

func NewOrderBook() *OrderBook {

	return &OrderBook{
		Bids: NewBidTree(),
		Asks: NewAskTree(),
	}
}

func NewBidTree() *BidTree {

	return &BidTree{
		Tree: btree.New(32),
	}
}

func (b *BidTree) Insert(
	order *Order,
) {
	fmt.Printf(
		"[BID INSERT] price=%s order=%s\n",
		order.Price.String(),
		order.OrderID,
	)
	level := &PriceLevel{
		Price: order.Price,
	}

	existing := b.Tree.Get(level)

	if existing != nil {

		pl := existing.(*PriceLevel)

		pl.Orders = append(
			pl.Orders,
			order,
		)

		return
	}

	level.Orders = []*Order{
		order,
	}

	b.Tree.ReplaceOrInsert(
		level,
	)
}

func (b *BidTree) BestBid() *PriceLevel {

	item := b.Tree.Max()

	if item == nil {
		return nil
	}

	return item.(*PriceLevel)
}
func (b *BidTree) PopBestBidOrder() *Order {

	best := b.BestBid()

	if best == nil {
		return nil
	}

	order := best.Orders[0]

	best.Orders = best.Orders[1:]

	if len(best.Orders) == 0 {

		b.Tree.Delete(best)
	}

	return order
}

func NewAskTree() *AskTree {

	return &AskTree{
		Tree: btree.New(32),
	}
}
func (a *AskTree) Insert(
	order *Order,
) {
	fmt.Printf(
		"[ASK INSERT] price=%s order=%s\n",
		order.Price.String(),
		order.OrderID,
	)
	level := &PriceLevel{
		Price: order.Price,
	}

	existing := a.Tree.Get(level)

	if existing != nil {

		pl := existing.(*PriceLevel)

		pl.Orders = append(
			pl.Orders,
			order,
		)

		return
	}

	level.Orders = []*Order{
		order,
	}

	a.Tree.ReplaceOrInsert(
		level,
	)
}
func (a *AskTree) BestAsk() *PriceLevel {

	item := a.Tree.Min()

	if item == nil {
		return nil
	}

	return item.(*PriceLevel)
}
func (a *AskTree) PopBestAskOrder() *Order {

	best := a.BestAsk()

	if best == nil {
		return nil
	}

	order := best.Orders[0]

	best.Orders = best.Orders[1:]

	if len(best.Orders) == 0 {

		a.Tree.Delete(best)
	}

	return order
}
func (o *OrderBook) Print(MarketID string) {

	fmt.Println("\n======================")
	fmt.Println("ORDERBOOK - ", MarketID)
	fmt.Println("======================")

	fmt.Println("\nASKS")

	o.Asks.Tree.Ascend(
		func(item btree.Item) bool {

			level := item.(*PriceLevel)

			fmt.Printf(
				"Price=%s Orders=%d\n",
				level.Price.String(),
				len(level.Orders),
			)

			for _, order := range level.Orders {

				fmt.Printf(
					"   %s Qty=%s\n",
					order.OrderID,
					order.Quantity.String(),
				)
			}

			return true
		},
	)

	fmt.Println("\nBIDS")

	o.Bids.Tree.Descend(
		func(item btree.Item) bool {

			level := item.(*PriceLevel)

			fmt.Printf(
				"Price=%s Orders=%d\n",
				level.Price.String(),
				len(level.Orders),
			)

			for _, order := range level.Orders {

				fmt.Printf(
					"   %s Qty=%s\n",
					order.OrderID,
					order.Quantity.String(),
				)
			}

			return true
		},
	)

	fmt.Println("======================")
}
