package store

import (
	"fmt"
	"sync"
	"trading-system/common"
	"trading-system/models"
)

type MemoryStore struct {
	Users      map[string]*models.User
	Orders     map[string]*models.Order
	Trades     map[string]*models.Trade
	OrderBooks map[string]*common.OrderBook
	Mutex      sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Users:      make(map[string]*models.User),
		Orders:     make(map[string]*models.Order),
		Trades:     make(map[string]*models.Trade),
		OrderBooks: make(map[string]*common.OrderBook),
	}
}

func (ms *MemoryStore) AddUser(user *models.User) {
	ms.Mutex.Lock()
	defer ms.Mutex.Unlock()
	ms.Users[user.ID] = user
}

func (ms *MemoryStore) AddOrder(order *models.Order) {
	ms.Mutex.Lock()

	// Check if OrderBook exists
	ob, exists := ms.OrderBooks[order.Symbol]
	if !exists {
		ob = common.NewOrderBook(order.Symbol)
		ms.OrderBooks[order.Symbol] = ob
	}

	// Store the order
	ms.Orders[order.ID] = order

	ms.Mutex.Unlock()
	ob.AddOrder(order)
}

func (ms *MemoryStore) ModifyOrder(orderID string, updatedOrder *models.Order) error {
	ms.Mutex.Lock()
	defer ms.Mutex.Unlock()

	// Check if the order exists
	existingOrder, exists := ms.Orders[orderID]
	if !exists {
		return fmt.Errorf("order with ID %s not found", orderID)
	}

	// Get the order book
	ob, exists := ms.OrderBooks[existingOrder.Symbol]
	if exists {
		// Remove the existing order from the appropriate queue
		if existingOrder.Type == models.Buy {
			ob.BuyOrders.RemoveById(orderID)
		} else {
			ob.SellOrders.RemoveById(orderID)
		}

		// Update the order attributes
		existingOrder.Quantity = updatedOrder.Quantity
		existingOrder.Price = updatedOrder.Price
		existingOrder.Timestamp = updatedOrder.Timestamp
		existingOrder.Status = updatedOrder.Status

		// Add the updated order back to the book
		ob.AddOrder(existingOrder)
	}

	return nil
}

func (ms *MemoryStore) GetOrderStatus(orderID string) models.OrderStatus {
	ms.Mutex.RLock()
	defer ms.Mutex.RUnlock()
	if order, ok := ms.Orders[orderID]; ok {
		return order.Status
	}
	return models.Rejected
}

func (ms *MemoryStore) GetOrderBook(symbol string) *common.OrderBook {
	ms.Mutex.RLock()
	ob := ms.OrderBooks[symbol]
	ms.Mutex.RUnlock()
	return ob
}

func (ms *MemoryStore) AddTrade(trade *models.Trade) {
	ms.Mutex.Lock()
	defer ms.Mutex.Unlock()
	ms.Trades[trade.ID] = trade
}

func (ms *MemoryStore) MarkOrderCompleted(orderID string) {
	if order, ok := ms.Orders[orderID]; ok {
		order.Status = models.Completed
	}
}

func (ms *MemoryStore) MarkOrderExpired(orderID string) {
	if order, ok := ms.Orders[orderID]; ok {
		order.Status = models.Rejected
	}
}
