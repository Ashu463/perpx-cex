package internal

import (
	"fmt"

	"matching-engine/internal/models"

	"github.com/shopspring/decimal"
)

func HandleTrades(
	trades []models.Trade,
	engine *Engine,
	publisher *Publisher,
) {

	if len(trades) == 0 {
		return
	}

	for _, trade := range trades {
		ProcessTrade(
			trade,
			engine,
			publisher,
		)
	}
}

func ProcessTrade(
	trade models.Trade,
	engine *Engine,
	publisher *Publisher,
) {

	fmt.Println(
		"\n======================",
	)
	fmt.Println(
		"TRADE EVENT",
	)
	fmt.Println(
		"======================",
	)

	fmt.Println(
		"Buyer:",
		trade.BuyerID,
		"Seller:",
		trade.SellerID,
		"Price:",
		trade.Price,
		"Quantity:",
		trade.Quantity,
		"Market:",
		trade.MarketID,
	)

	// TODO:
	// UpdateBalances(trade, engine)

	// TODO:
	// UpdatePositions(trade, engine)

	// TODO: PublishTrade(trade)
	PublishTrade(&trade, publisher)

	// TODO:
	// PublishOrderUpdate(trade)

	// TODO:
	// PersistTrade(trade)
}

func PublishTrade(
	trade *models.Trade,
	publisher *Publisher,
) {

	err := publisher.PublishTradeResult(trade)

	if err != nil {

		fmt.Println(
			"Error while publishing trade:",
			err,
		)

		return
	}

	fmt.Println(
		"Trade published successfully",
	)
}
func SyncBalances(order models.Order, engine *Engine) {

	fmt.Println(order, "is the order model")
	user := engine.Balances[order.UserID]

	if user == nil {
		user = &models.Balance{
			UserID:           order.UserID,
			AvailableBalance: decimal.Zero,
			LockedBalance:    order.Margin,
			Equity:           decimal.Zero,
		}
		engine.Balances[order.UserID] = user
		return
	}
	user.LockedBalance.Add(order.Margin)
}

// rewrite this update balance without taking trade as param
func UpdateBalances(trade models.Trade, engine *Engine) {

	fmt.Println(trade, "is the trade model")

	tradedAmt := trade.Price.Mul(trade.Quantity)

	buyer := engine.Balances[trade.BuyerID]
	if buyer == nil {
		// now neither the buyer nor the seller could be nil, coz I synced every user before matching
		fmt.Println("Buyer is nil, ", buyer)
		return
	}

	seller := engine.Balances[trade.SellerID]
	if seller == nil {
		fmt.Println("Seller is nil, ", seller)
		return
	}

	fmt.Println(buyer, "is the buyer fetched from engine")

	buyer.LockedBalance = buyer.LockedBalance.Sub(tradedAmt)

	seller.AvailableBalance = seller.AvailableBalance.Add(tradedAmt)
	seller.LockedBalance = seller.LockedBalance.Sub(tradedAmt)

	fmt.Printf(
		"\nBUYER BALANCE: %+v\n",
		*buyer,
	)

	fmt.Printf(
		"\nSELLER BALANCE: %+v\n",
		*seller,
	)
}
func UpdatePositions(trade models.Trade, engine *Engine) {
	// add position for buyer and delete position for seller
	// #TODO: update MarkPrice, LiquidationPrice, InitialMargin, MaintainenceMargin, UnrealizedPnl, RealizedPnl

	fmt.Println("Positions, ", trade.MarketID)

	for userID, pos := range engine.Positions[trade.MarketID] {

		fmt.Println(
			"User= Qty= EntryPrice=",
			userID,
			pos.Quantity.String(),
			pos.EntryPrice.String(),
		)
	}

	fmt.Println("====================================")

	buyerPos := GetPosition(trade.MarketID, trade.BuyerID, engine)

	if engine.Positions[trade.MarketID] == nil {

		engine.Positions[trade.MarketID] = make(
			map[string]*models.Position,
		)
	}
	if buyerPos == nil {

		buyerPos = &models.Position{
			UserID:     trade.BuyerID,
			MarketID:   trade.MarketID,
			Quantity:   decimal.Zero,
			EntryPrice: decimal.Zero,
		}

		engine.Positions[trade.MarketID][trade.BuyerID] = buyerPos
	}

	// weighted average entry

	if buyerPos.Quantity.IsZero() {

		buyerPos.EntryPrice = trade.Price

	} else {

		totalCost := buyerPos.EntryPrice.Mul(buyerPos.Quantity).Add(trade.Price.Mul(trade.Quantity))

		newQty := buyerPos.Quantity.Add(trade.Quantity)

		buyerPos.EntryPrice = totalCost.Div(newQty)
	}

	buyerPos.Quantity = buyerPos.Quantity.Add(trade.Quantity)

	// seller side

	sellerPos := GetPosition(
		trade.MarketID,
		trade.SellerID,
		engine,
	)

	if sellerPos == nil {
		fmt.Println("ERROR: seller has no position to close", trade.SellerID)
		return
	}
	// intially seller pos bhi toh nil hi hogi ya fer order pdte hi yeh position and balance update krni hogi
	// toh fix that thing(update of balance and position just after the order creation) pehle then test
	sellerPos.Quantity = sellerPos.Quantity.Sub(trade.Quantity)

	if sellerPos.Quantity.IsZero() {
		delete(engine.Positions[trade.MarketID], trade.SellerID)
	}
}

func GetPosition(
	marketID string,
	userID string,
	engine *Engine,
) *models.Position {

	marketPositions, exists := engine.Positions[marketID]

	// market doesn't exsit
	if !exists {
		return nil
	}
	// that position doesn't exist
	position, exists := marketPositions[userID]

	if !exists {
		return nil
	}

	return position
}
