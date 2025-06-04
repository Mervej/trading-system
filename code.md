Order Placement:

Users place orders via the MemoryStore, which adds them to the appropriate OrderBook.
Order Matching:

The MatchingEngine retrieves the OrderBook for a stock symbol and matches buy and sell orders based on price and timestamp.
Trade Execution:

When a match is found, a Trade is created and stored in the MemoryStore.
The quantities of the matched orders are updated, and completed orders are removed.
Concurrency:

sync.RWMutex ensures thread-safe access to shared resources like MemoryStore and OrderBook.
Trade Expiry:

The MatchingEngine periodically checks for expired orders and removes them from the OrderBook.


The code references pointers (e.g., *models.User, *models.Order, *models.Trade) instead of the objects directly for the following reasons:
1. Memory Efficiency - Using pointers avoids copying the entire object when passing
2. Mutability - Pointers allow the modification of the original object directly
3. Consistency Across References
5. Concurrency - In a multi-threaded environment, pointers simplify synchronization.