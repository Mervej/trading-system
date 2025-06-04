package common

import (
	"sync"
	"trading-system/models"
	"trading-system/utils"
)

type OrderBook struct {
	Symbol     string
	BuyOrders  *utils.OrderPriorityQueue
	SellOrders *utils.OrderPriorityQueue
	Mutex      sync.RWMutex
}

func NewOrderBook(symbol string) *OrderBook {
	buyPQ := utils.NewOrderPriorityQueue(true)
	sellPQ := utils.NewOrderPriorityQueue(false)
	buyPQ.Init()
	sellPQ.Init()
	return &OrderBook{
		Symbol:     symbol,
		BuyOrders:  buyPQ,
		SellOrders: sellPQ,
	}
}

func (ob *OrderBook) AddOrder(order *models.Order) {

	// Add the new order to the order book
	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()
	if order.Type == models.Buy {
		ob.BuyOrders.Add(order)
	} else {
		ob.SellOrders.Add(order)
	}
}
