# Trading System

## Overview
This project is an in-memory trading system that simulates a stock exchange. It allows registered users to place, modify, cancel, and query orders, as well as execute trades based on matching buy and sell orders. The system demonstrates synchronization and concurrency in a multi-threaded environment.

## Project Structure
```
trading-system
├── cmd
│   └── main.go        # Entry point of the application
├── common
│   └── order_book.go  # OrderBook implementation
├── engine
│   └── matcher.go     # Matching engine for trade execution
├── models
│   └── models.go      # Data models for User, Order, and Trade
├── store
│   └── memory_store.go # In-memory data store for users, orders, and trades
├── utils
│   └── priority_queue.go # Priority queue implementation for order matching
├── Requirement.md     # Project requirements and expectations
├── go.mod             # Module definition and dependencies
└── README.md          # Project documentation
```

## Features
### Functional Requirements
- **Order Management**:
  - Place, modify, and cancel orders.
  - Query the status of an order.
- **Trade Execution**:
  - Match buy and sell orders based on price and timestamp.
  - Execute trades when buy price ≥ sell price.
  - Prioritize older orders if prices are equal.
- **Concurrency**:
  - Handle concurrent order placement, modification, cancellation, and execution.
- **Order Book**:
  - Maintain an order book per stock symbol with unexecuted orders.
- **Trade Expiry**:
  - Automatically cancel orders that exceed a specific expiry duration.

### Non-Functional Requirements
- Thread-safe operations using `sync.RWMutex`.
- Clean and refactored code with graceful exception handling.
- Abstraction for in-memory data structures to allow future integration with persistent stores.

## Getting Started

### Prerequisites
- Go 1.18 or later installed on your machine.

### Installation

1. Navigate to the project directory:
   ```
   cd trading-system
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Application
To run the application, execute the following command:
```
go run cmd/main.go
```

## Usage
### Example Workflow
1. Add dummy users to the system.
2. Place multiple buy and sell orders for a stock symbol.
3. Execute trades using the matching engine.
4. Query the status of orders and view executed trades.
5. Handle expired orders automatically.

### Sample Output
```
--- Executed Trades ---
Trade ID: 1
  Type: Buy
  Buyer Order ID: 1
  Seller Order ID: 2
  Symbol: WIPRO
  Quantity: 100
  Price: 245.00
  Timestamp: 2025-06-05 12:34:56
-----------------------
Expired Order: ID=3, Type=Buy, Symbol=WIPRO, Quantity=50, Price=255.00, Timestamp=2025-06-05 12:34:16
``