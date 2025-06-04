package main

import (
	"fmt"
	"time"
	"trading-system/engine"
	"trading-system/models"
	"trading-system/store"
)

func main() {
	memStore := store.NewMemoryStore()
	matcher := engine.NewMatchingEngine(memStore)

	// Dummy users
	user := &models.User{ID: "1", Name: "Mervej", PhoneNumber: "9876565432", Email: "mervej@example.com"}

	// Add orders
	order1 := &models.Order{
		ID:        "1",
		UserID:    "1",
		Type:      models.Buy,
		Symbol:    "WIPRO",
		Quantity:  100,
		Price:     250.0,
		Timestamp: time.Now().Add(-10 * time.Second), // Older order
	}

	order2 := &models.Order{
		ID:        "2",
		UserID:    "1",
		Type:      models.Sell,
		Symbol:    "WIPRO",
		Quantity:  100,
		Price:     245.0,
		Timestamp: time.Now(),
	}

	order3 := &models.Order{
		ID:        "3",
		UserID:    "1",
		Type:      models.Buy,
		Symbol:    "WIPRO",
		Quantity:  50,
		Price:     255.0,
		Timestamp: time.Now().Add(-40 * time.Second), // Expired order
	}

	order4 := &models.Order{
		ID:        "4",
		UserID:    "1",
		Type:      models.Sell,
		Symbol:    "WIPRO",
		Quantity:  50,
		Price:     240.0,
		Timestamp: time.Now().Add(-35 * time.Second), // Expired order
	}

	memStore.AddUser(user)
	memStore.AddOrder(order1)
	memStore.AddOrder(order2)
	memStore.AddOrder(order3)
	memStore.AddOrder(order4)

	// Match orders
	matcher.MatchOrders("WIPRO")

	// Print trades
	for _, trade := range memStore.Trades {
		fmt.Printf("Trade ID: %s\n", trade.ID)
		fmt.Printf("  Type: %s\n", trade.Type)
		fmt.Printf("  Buyer Order ID: %s\n", trade.BuyerOrderID)
		fmt.Printf("  Seller Order ID: %s\n", trade.SellerOrderID)
		fmt.Printf("  Symbol: %s\n", trade.Symbol)
		fmt.Printf("  Quantity: %d\n", trade.Quantity)
		fmt.Printf("  Price: %.2f\n", trade.Price)
		fmt.Printf("  Timestamp: %s\n", trade.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Println("-----------------------")

	}
}
