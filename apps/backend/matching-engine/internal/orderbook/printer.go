package orderbook

import (
	"fmt"
	"matching-engine/internal/models"
)

func PrintOrderBook(book *models.OrderBook) {

	fmt.Println("Asks")
	fmt.Println("------------------------------")
	for i := 0; i < len(book.Asks); i++ {
		fmt.Println()
		fmt.Println(book.Asks[i].OrderID, " SELL ", book.Asks[i].Quantity)
	}

	fmt.Println("Bids")
	fmt.Println("------------------------------")
	for i := 0; i < len(book.Bids); i++ {
		fmt.Println(book.Bids[i].OrderID, " BUY ", book.Bids[i].Quantity)
	}
}
