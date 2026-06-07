package internal

import (
	"fmt"
	"github.com/shopspring/decimal"
	"matching-engine/internal/models"
)

// flow for matcher would be
/*
- run matching
- if matches then go for the trade
- if trade happens: could be in two cases in balances
	- fully, lockedBalances -= tradedAmt which is margin in this case(buyer side) while for seller
	lockedBalance -= tradedAmt and add that in avaBalance of seller(tradedAmt).
	- partially, update balances with same above logic while keeping qty as not fully satisfied.
	- Open and close position for buyer and seller respectively.

- if no match then list up on orderBook.


*/
func Match(
	book *models.OrderBook,
	order *models.Order,
	engine *Engine,
	publisher *Publisher,
) []models.Trade {
	// fmt.Println("matcher called")

	fmt.Println(
		"\nMATCHING:",
		order.OrderID,
	)
	var trades []models.Trade

	if order.Side == models.Buy {
		demandedQty := order.Quantity

		for demandedQty.IsPositive() {
			bestAsk := book.Asks.BestAsk()

			if bestAsk == nil || bestAsk.Price.GreaterThan(order.Price) {
				break
			}

			maker := bestAsk.Orders[0]
			fillQty := decimal.Min(demandedQty, maker.RemainingQuantity)

			trade := models.Trade{
				BuyerID:     order.UserID,
				SellerID:    maker.UserID,
				BuyOrderID:  order.OrderID,
				SellOrderID: maker.OrderID,
				MarketID:    order.MarketID,
				Price:       bestAsk.Price,
				Quantity:    fillQty,
			}
			// #TODO DONE - trade poora krde bhai - kr diya bhai
			trades = append(trades, trade)
			// update balance and positions just after appending into the trade
			// UpdateBalances(trade, engine)
			// UpdatePositions(trade, engine)
			fmt.Printf(
				"\nTRADE EXECUTED buyer=%s seller=%s price=%s qty=%s\n",
				trade.BuyerID,
				trade.SellerID,
				trade.Price.String(),
				trade.Quantity.String(),
			)

			// reduce quantities
			demandedQty = demandedQty.Sub(fillQty)
			maker.RemainingQuantity = maker.RemainingQuantity.Sub(fillQty)
			// maker.Quantity = maker.RemainingQuantity
			order.RemainingQuantity = demandedQty

			// maker fully filled → remove from orderbook
			if maker.RemainingQuantity.IsZero() {
				book.Asks.PopBestAskOrder()
				fmt.Printf(
					"ASK FILLED order=%s removed from book\n",
					maker.OrderID,
				)
			}
			// maker partially filled → stays in orderbook with reduced qty
			// demandedQty loop continues looking for next ask
		}

		// incoming order still has remaining qty → list on orderbook
		if demandedQty.IsPositive() {
			fmt.Println(demandedQty, " is the demanded qty")
			order.RemainingQuantity = demandedQty
			// order.Quantity = demandedQty
			book.Bids.Insert(order)
			fmt.Printf(
				"BID INSERTED order=%s remaining=%s\n",
				order.OrderID,
				order.RemainingQuantity.String(),
			)
			fmt.Println("below bid insert logger")
		}

	} else {
		// SELL ORDER

		for order.RemainingQuantity.IsPositive() {
			fmt.Println("inside sell loop")
			bestBid := book.Bids.BestBid()

			if bestBid == nil || bestBid.Price.LessThan(order.Price) {
				fmt.Println("inside bid nil")
				break
			}

			maker := bestBid.Orders[0]

			fillQty := decimal.Min(
				order.RemainingQuantity,
				maker.RemainingQuantity,
			)

			trade := models.Trade{
				BuyerID:  maker.UserID,
				SellerID: order.UserID,

				BuyOrderID:  maker.OrderID,
				SellOrderID: order.OrderID,

				MarketID: order.MarketID,

				Price:    bestBid.Price,
				Quantity: fillQty,
			}

			trades = append(trades, trade)
			// update balance and position first
			// UpdateBalances(trade, engine)
			// UpdatePositions(trade, engine)
			fmt.Printf(
				"\nTRADE EXECUTED buyer=%s seller=%s price=%s qty=%s\n",
				trade.BuyerID,
				trade.SellerID,
				trade.Price.String(),
				trade.Quantity.String(),
			)

			order.RemainingQuantity =
				order.RemainingQuantity.Sub(fillQty)

			// order.Quantity = order.RemainingQuantity

			maker.RemainingQuantity =
				maker.RemainingQuantity.Sub(fillQty)
			// maker.Quantity = maker.RemainingQuantity

			if maker.RemainingQuantity.IsZero() {

				book.Bids.PopBestBidOrder()

				fmt.Printf(
					"BID FILLED order=%s removed from book\n",
					maker.OrderID,
				)
			}
		}

		if order.RemainingQuantity.IsPositive() {

			book.Asks.Insert(order)

			fmt.Printf(
				"ASK INSERTED order=%s remaining=%s\n",
				order.OrderID,
				order.RemainingQuantity.String(),
			)
		}
	}
	fmt.Println("above loop below bid inserted body")
	for i := 0; i < len(trades); i++ {
		fmt.Println("inside update loop")
		UpdateBalances(trades[i], engine)
		UpdatePositions(trades[i], engine)
	}
	fmt.Println("calling handle trades now.")
	HandleTrades(trades, engine, publisher)
	return trades
}

// logics behind handling cases of matching

// BUY ORDER
// if order.Side == models.Buy {

// 	bestAsk := book.Asks.BestAsk()
// 	// if either the table is null of most optimised ask price is strictly greater than order price
// 	if bestAsk == nil || bestAsk.Price.GreaterThan(order.Price) {
// 		book.Bids.Insert(order)
// 		return trades
// 	}

// 	// demandedQty := order.Quantity
// 	// C-1: A = B :simply do the trade, push the incoming order into trade event and remove that askOrder
// 	// C-2: A < B :do the trade for incoming order side while decrease the quantity of askOrder

// 	// execute the trade for incoming order
// 	/*
// 		trade execution means that you have to update the
// 		balance(it might be done somewhere else, coz balance updation doesn't upon whether
// 		the user order get executed instantly or listed on orderbook, but yeah probably we
// 		have to update opposite side balances from locked to available), positions
// 		(we have to update position on the both side),
// 		orderbook(based upon available number of units) and
// 		trade(which needs to be given to the queue for database updation).
// 	*/

// 	// update the position, balance of the order holder(incoming trader)
// 	// if bestAsk.Orders[0].Quantity.GreaterThanOrEqual(demandedQty) {
// 	// 	book.Asks.PopBestAskOrder()

// 	// } else {
// 	// 	for demandedQty.GreaterThanOrEqual(0) || bestAsk == nil {

// 	// 		// remove the bestAsk guy from orderBook, do it's position, balance everything
// 	// 		// append that number of the units into the incoming order guy and yeah execute trade on that side.
// 	// 		demandedQty.Sub(bestAsk.Orders[0].Quantity)
// 	// 		book.Asks.PopBestAskOrder()

// 	// 		bestAsk = book.Asks.BestAsk()
// 	// 	}
// 	// 	if demandedQty != 0 {
// 	// 		// push this order with whatever qty left to the orderBook.
// 	// 	}
// 	// }

// 	// I have to run a loop till, either all the quantity gets exhausted or incoming order didn't
// 	// desired price.
// 	// incoming order's price must be compared with bestAsk's Price
// 	// if ok, the compare the no of units both parties have
// 	// three cases
// 	// order's quantity k/a 'A' while bestAsk quantity k/a 'B'
// 	// C-3: A > B :give B to the order(i.e. trade for this) while search for new askOrder till either
// 	// A becomes zero or get listed on orderBook
// }
