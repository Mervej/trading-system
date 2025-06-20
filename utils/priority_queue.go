package utils

import (
	"container/heap"
	"trading-system/models"
)

type OrderPriorityQueue struct {
	orders  []*models.Order
	isMaxPQ bool
}

func NewOrderPriorityQueue(isMax bool) *OrderPriorityQueue {
	return &OrderPriorityQueue{
		orders:  []*models.Order{},
		isMaxPQ: isMax,
	}
}

func (pq OrderPriorityQueue) Len() int { return len(pq.orders) }

func (pq OrderPriorityQueue) Less(i, j int) bool {
	a, b := pq.orders[i], pq.orders[j]

	// if price is same return the first order
	if a.Price == b.Price {
		return a.Timestamp.Before(b.Timestamp)
	}

	// now return the response basis on sell or buy order
	if pq.isMaxPQ {
		return a.Price > b.Price
	} else {
		return a.Price < b.Price
	}
}

func (pq OrderPriorityQueue) Swap(i, j int) {
	pq.orders[i], pq.orders[j] = pq.orders[j], pq.orders[i]
}

func (pq *OrderPriorityQueue) Push(x interface{}) {
	item := x.(*models.Order)
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
	return pq.orders[0]
}

func (pq *OrderPriorityQueue) Add(order *models.Order) {
	heap.Push(pq, order)
}

func (pq *OrderPriorityQueue) Remove() *models.Order {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(pq).(*models.Order)
}

func (pq *OrderPriorityQueue) RemoveById(orderId string) {
	if pq.Len() == 0 {
		return
	}

	// Find the index of the order with the given ID
	removeIdx := -1
	for i, order := range pq.orders {
		if order.ID == orderId {
			removeIdx = i
			break
		}
	}

	// If order is found remove it and sort the heap
	if removeIdx != -1 {
		// Move the last element to curr position
		lastIdx := len(pq.orders) - 1
		pq.orders[removeIdx] = pq.orders[lastIdx]
		pq.orders = pq.orders[:lastIdx]

		// Reheapify if not removing the last element
		if removeIdx != lastIdx {
			heap.Fix(pq, removeIdx)
		}
	}
}

func (pq *OrderPriorityQueue) Init() {
	heap.Init(pq)
}
