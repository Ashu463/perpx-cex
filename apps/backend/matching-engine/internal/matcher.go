package internal

import (
	"matching-engine/internal/models"
)

func Match(
	book *models.OrderBook,
	order *models.Order,
) []models.Trade {

	var trades []models.Trade

	// BUY ORDER
	if order.Side == models.Buy {

		bestAsk := book.Asks.BestAsk()

		if bestAsk == nil {

			book.Bids.Insert(order)

			return trades
		}

		if order.Price.LessThan(bestAsk.Price) {

			book.Bids.Insert(order)

			return trades
		}

		maker := bestAsk.Orders[0]

		trade := models.Trade{
			BuyerID:  order.UserID,
			SellerID: maker.UserID,

			BuyOrderID:  order.OrderID,
			SellOrderID: maker.OrderID,

			MarketID: order.MarketID,

			Price: bestAsk.Price,

			Quantity: order.Quantity,
		}

		trades = append(
			trades,
			trade,
		)

		book.Asks.PopBestAskOrder()

		return trades
	}

	// SELL ORDER

	bestBid := book.Bids.BestBid()

	if bestBid == nil {

		book.Asks.Insert(order)

		return trades
	}

	if order.Price.GreaterThan(bestBid.Price) {

		book.Asks.Insert(order)

		return trades
	}

	maker := bestBid.Orders[0]

	trade := models.Trade{
		BuyerID:  maker.UserID,
		SellerID: order.UserID,

		BuyOrderID:  maker.OrderID,
		SellOrderID: order.OrderID,

		MarketID: order.MarketID,

		Price: bestBid.Price,

		Quantity: order.Quantity,
	}

	trades = append(
		trades,
		trade,
	)

	book.Bids.PopBestBidOrder()

	return trades
}
