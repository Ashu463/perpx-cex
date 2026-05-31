# Perpetual Exchange Roadmap

## Goal

Build a Backpack/Hyperliquid-style perpetual futures exchange from scratch.

---

# Phase 1: Trading Core (MVP)

Goal: Users can place orders and open positions.

## Auth & User Management

* [ ] User
* [ ] Session
* [ ] Authentication
* [ ] Authorization
* [ ] API Keys

### Deliverables

* Register
* Login
* Refresh Token
* API Key Management

---

## Wallet & Balances

* [ ] Wallet
* [ ] Balance

### Deliverables

* Wallet Creation
* Available Balance
* Locked Balance
* Equity Calculation

---

## Market

* [ ] Market

### Fields

* marketId
* symbol
* status
* tickSize
* stepSize
* maxLeverage

### Deliverables

* Create Market
* List Markets
* Market Metadata

---

## Orders

* [ ] Order

### Fields

* orderId
* marketId
* userId
* side
* type
* quantity
* price
* status

### Deliverables

* Place Order
* Cancel Order
* Modify Order
* Query Orders

---

## Order Book

* [ ] OrderBook

### Deliverables

* Bid Side
* Ask Side
* Price Levels
* Aggregation
* Best Bid
* Best Ask

---

## Matching Engine

### Deliverables

* Match Buy/Sell Orders
* Partial Fills
* Full Fills
* Market Orders
* Limit Orders
* FIFO Matching

---

## Trades

* [ ] Trade

### Fields

* tradeId
* buyerId
* sellerId
* price
* quantity
* timestamp

### Deliverables

* Trade History
* Recent Trades Feed

---

## Positions

* [ ] Position

### Fields

* positionId
* marketId
* userId
* side
* size
* entryPrice
* leverage
* margin

### Deliverables

* Open Position
* Increase Position
* Reduce Position
* Close Position

---

# Phase 2: Risk Engine

Goal: Calculate fair valuation and trader profitability.

## Oracle Service

* [ ] Oracle Price

### Deliverables

* External Price Feed
* Price Aggregation
* Fair Price Calculation

---

## Mark Price Engine

* [ ] Mark Price

### Deliverables

* Mark Price Calculation
* Oracle Integration

---

## Margin Engine

* [ ] Margin

### Deliverables

* Initial Margin
* Maintenance Margin
* Margin Usage

---

## PnL Engine

* [ ] Unrealized PnL
* [ ] Realized PnL

### Deliverables

* Position Valuation
* Profit/Loss Tracking

---

## Risk Engine

### Deliverables

* Margin Monitoring
* Risk Scoring
* Liquidation Eligibility

---

# Phase 3: Liquidation Infrastructure

Goal: Protect exchange from bankrupt accounts.

## Liquidation Engine

* [ ] Liquidation

### Deliverables

* Liquidation Trigger
* Forced Close Logic
* Liquidation Queue

---

## Insurance Fund

* [ ] Insurance Fund

### Deliverables

* Fund Accounting
* Loss Coverage
* Bad Debt Protection

---

## Auto Deleveraging (ADL)

* [ ] ADL

### Deliverables

* ADL Ranking
* Position Reduction
* Emergency Recovery

---

# Phase 4: Perpetual-Specific Features

Goal: Keep perpetual markets aligned with spot markets.

## Funding Rate Engine

* [ ] Funding Rate

### Deliverables

* Funding Calculation
* Long/Short Imbalance Detection

---

## Funding Payments

* [ ] Funding Payment

### Deliverables

* Funding Settlement
* Funding History

---

## Market Analytics

### Deliverables

* Open Interest
* Long/Short Ratio
* Funding Dashboard

---

# Phase 5: Accounting & Exchange Operations

Goal: Production-grade exchange accounting.

## Ledger System

* [ ] Ledger

### Deliverables

* Double Entry Accounting
* Immutable Records

---

## Transactions

* [ ] Deposit
* [ ] Withdrawal

### Deliverables

* Deposit Processing
* Withdrawal Processing

---

## Fee Engine

* [ ] Trading Fees

### Deliverables

* Maker Fees
* Taker Fees
* Fee Settlement

---

## Balance Reconciliation

### Deliverables

* Balance Verification
* Ledger Auditing

---

## Reporting

### Deliverables

* User Statements
* PnL Reports
* Funding Reports
* Audit Logs

---

# Future Enhancements

## Trading Features

* [ ] Stop Loss
* [ ] Take Profit
* [ ] OCO Orders
* [ ] Trailing Stop

---

## Exchange Features

* [ ] Copy Trading
* [ ] Referral System
* [ ] Leaderboards
* [ ] Market Maker Program

---

## Scalability

* [ ] Redis
* [ ] Kafka
* [ ] WebSockets
* [ ] Horizontal Scaling

---

# Final Architecture

User
↓
Wallet
↓
Order
↓
Order Book
↓
Matching Engine
↓
Trade
↓
Position
↓
Oracle Price
↓
Mark Price
↓
PnL Engine
↓
Risk Engine
↓
Liquidation Engine
↓
Insurance Fund
↓
ADL
↓
Ledger
↓
Reporting
