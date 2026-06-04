// func Match(
// 	book *models.OrderBook,
// 	incomingOrder *models.Order,
// ) bool {

// 	// Matching logic goes here
// 	// we have to check for relevant makers or takers.

// 	// if orderSide === 'BUY' and book.asks.length > equal to zero {
// 	// if price available in asks register is less than or equal to buyer's price then assign that order to buyer, remove the seller and push the event to queue
// 	//	i.e. if price is matched with someone on the asks table then place an order push event to queue
// 	// else add them into the orderBook along with the timestamp for better compairson.
// 	// }

// 	// similarly for SELL one.

// 	if incomingOrder.Side == models.Buy {

// 		for i, ask := range book.Asks {

// 			if incomingOrder.Price.GreaterThanOrEqual(ask.Price) {

// 				trade := models.Trade{
// 					BuyerID:  incomingOrder.UserID,
// 					SellerID: ask.UserID,

// 					BuyOrderID:  incomingOrder.OrderID,
// 					SellOrderID: ask.OrderID,

// 					MarketID: incomingOrder.MarketID,

// 					Price: ask.Price,

// 					Quantity: incomingOrder.Quantity,
// 				}
// 				fmt.Println("Trade exectued: ", trade)
// 				fmt.Println("BUYER:", incomingOrder.UserID)
// 				fmt.Println("SELLER:", ask.UserID)

// 				book.Asks = slices.Delete(book.Bids, i, i+1)
// 				// events.Publisher
// 				return true
// 			}
// 		}

// 		book.Bids = append(
// 			book.Bids,
// 			incomingOrder,
// 		)

// 		return false
// 	}

// 	if incomingOrder.Side == models.Sell {

// 		for i, bid := range book.Bids {

// 			if incomingOrder.Price.GreaterThanOrEqual(bid.Price) {

// 				trade := models.Trade{
// 					BuyerID:  incomingOrder.UserID,
// 					SellerID: bid.UserID,

// 					BuyOrderID:  incomingOrder.OrderID,
// 					SellOrderID: bid.OrderID,

// 					MarketID: incomingOrder.MarketID,

// 					Price: bid.Price,

// 					Quantity: incomingOrder.Quantity,
// 				}
// 				fmt.Println("Trade exectued: ", trade)

// 				fmt.Println("BUYER:", incomingOrder.UserID)
// 				fmt.Println("SELLER:", bid.UserID)

// 				book.Bids = slices.Delete(book.Asks, i, i+1)
// 				//

// 				return true
// 			}
// 		}

// 		book.Asks = append(
// 			book.Bids,
// 			incomingOrder,
// 		)

// 		return false
// 	}

// 	return false

// }

package matching

import (
	"matching-engine/internal/models"
	"matching-engine/internal/orderbook"
)

func Match(
	book *models.OrderBook,
	order *models.Order,
) {

	if order.Side == models.Buy {

		for _, ask := range book.Asks {

			if order.Price.GreaterThanOrEqual(
				ask.Price,
			) {

				// trade occurs

				return
			}
		}

		orderbook.AddBid(
			book,
			order,
		)

		return
	}

	for _, bid := range book.Bids {

		if order.Price.LessThanOrEqual(
			bid.Price,
		) {

			// trade occurs

			return
		}
	}

	orderbook.AddAsk(
		book,
		order,
	)

}

// demo streams
// XADD order_submissions * orderId btc-buy userId u1 marketId BTCUSDT side BUY price 10000 quantity 1 margin 100

// XADD order_submissions * orderId eth-buy userId u1 marketId ETHUSDT side BUY price 1000 quantity 1 margin 10
