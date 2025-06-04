package utils

import (
	"container/heap"
	"trading-system/models"
)

type OrderItem struct {
	Order *models.Order
	index int
}

type OrderPriorityQueue struct {
	orders  []*OrderItem
	isMaxPQ bool
}

func NewOrderPriorityQueue(isMax bool) *OrderPriorityQueue {
	return &OrderPriorityQueue{
		orders:  []*OrderItem{},
		isMaxPQ: isMax,
	}
}

func (pq OrderPriorityQueue) Len() int { return len(pq.orders) }

func (pq OrderPriorityQueue) Less(i, j int) bool {
	a, b := pq.orders[i], pq.orders[j]
	if pq.isMaxPQ {
		if a.Order.Price == b.Order.Price {
			return a.Order.Timestamp.Before(b.Order.Timestamp)
		}
		return a.Order.Price > b.Order.Price
	} else {
		if a.Order.Price == b.Order.Price {
			return a.Order.Timestamp.Before(b.Order.Timestamp)
		}
		return a.Order.Price < b.Order.Price
	}
}

func (pq OrderPriorityQueue) Swap(i, j int) {
	pq.orders[i], pq.orders[j] = pq.orders[j], pq.orders[i]
	pq.orders[i].index = i
	pq.orders[j].index = j
}

func (pq *OrderPriorityQueue) Push(x interface{}) {
	n := len(pq.orders)
	item := x.(*OrderItem)
	item.index = n
	pq.orders = append(pq.orders, item)
}

func (pq *OrderPriorityQueue) Pop() interface{} {
	old := pq.orders
	n := len(old)
	item := old[n-1]
	pq.orders = old[0 : n-1]
	return item
}

func (pq *OrderPriorityQueue) Peek() *models.Order {
	if pq.Len() == 0 {
		return nil
	}
	return pq.orders[0].Order
}

func (pq *OrderPriorityQueue) Add(order *models.Order) {
	heap.Push(pq, &OrderItem{Order: order})
}

func (pq *OrderPriorityQueue) Remove() *models.Order {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(pq).(*OrderItem).Order
}

func (pq *OrderPriorityQueue) Init() {
	heap.Init(pq)
}
