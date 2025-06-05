package engine

import (
	"fmt"
	"math"
	"time"
	"trading-system/common"
	"trading-system/models"
	"trading-system/store"
	"trading-system/utils"
)

type MatchingEngine struct {
	Store          *store.MemoryStore
	ExpiryDuration time.Duration
}

func NewMatchingEngine(store *store.MemoryStore) *MatchingEngine {
	return &MatchingEngine{
		Store:          store,
		ExpiryDuration: 30 * time.Second,
	}
}

func (me *MatchingEngine) MatchOrders(symbol string) {
	ob := me.Store.GetOrderBook(symbol)

	ob.Mutex.Lock()
	defer ob.Mutex.Unlock()

	now := time.Now()

	// call function to expire the older orders
	me.expireOldOrders(ob, now)

	for {
		// check for matching orders
		buy := ob.BuyOrders.Peek()
		sell := ob.SellOrders.Peek()

		if buy == nil || sell == nil || buy.Price < sell.Price {
			// return if no valid trade is found
			break
		}

		tradeQty := int(math.Min(float64(buy.Quantity), float64(sell.Quantity)))
		tradePrice := sell.Price

		trade := &models.Trade{
			ID:            generateID(),
			Type:          models.Buy,
			BuyerOrderID:  buy.ID,
			SellerOrderID: sell.ID,
			Symbol:        symbol,
			Quantity:      tradeQty,
			Price:         tradePrice,
			Timestamp:     now,
		}

		// remove lock to avoid deadlock and add to trade model and relock
		ob.Mutex.Unlock()
		me.Store.AddTrade(trade)
		ob.Mutex.Lock()

		// update quantity and status in order book
		buy.Quantity -= tradeQty
		sell.Quantity -= tradeQty

		if buy.Quantity == 0 {
			ob.BuyOrders.Remove()
			me.Store.MarkOrderCompleted(buy.ID)
		}
		if sell.Quantity == 0 {
			ob.SellOrders.Remove()
			me.Store.MarkOrderCompleted(sell.ID)
		}
	}
}

func (me *MatchingEngine) expireOldOrders(ob *common.OrderBook, now time.Time) {
	expire := func(queue *utils.OrderPriorityQueue, orderType string) {
		for {
			top := queue.Peek()
			if top == nil {
				break
			}

			// check if the order timestamp is past the expiry duration
			if now.Sub(top.Timestamp) > me.ExpiryDuration {
				expired := queue.Remove()
				me.Store.MarkOrderExpired(expired.ID)
				fmt.Printf("Expired Order: ID=%s, Type=%s, Symbol=%s, Quantity=%d, Price=%.2f, Timestamp=%s\n",
					expired.ID, orderType, expired.Symbol, expired.Quantity, expired.Price, expired.Timestamp.Format("2006-01-02 15:04:05"))
			} else {
				break
			}
		}
	}
	expire(ob.BuyOrders, "Buy")
	expire(ob.SellOrders, "Sell")
}

func generateID() string {
	return time.Now().Format("20060102150405.000000")
}
