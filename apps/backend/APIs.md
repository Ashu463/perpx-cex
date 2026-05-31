# API Design

## Auth-service (/auth)

### POST /signup

Request

```json
{
  "username": "ashu",
  "password": "password123"
}
```

Response

```json
{
  "id": 1,
  "username": "ashu",
  "createdAt": "2026-06-01T10:00:00Z"
}
```

---

### POST /login

Request

```json
{
  "username": "ashu",
  "password": "password123"
}
```

Response

```json
{
  "accessToken": "jwt-token",
  "refreshToken": "refresh-token"
}
```

---

### POST /refresh-token

Request

```json
{
  "refreshToken": "refresh-token"
}
```

Response

```json
{
  "accessToken": "new-jwt-token"
}
```

---

### POST /logout

Request

```json
{}
```

Response

```json
{
  "message": "Logged out successfully"
}
```

---

### GET /me

Response

```json
{
  "id": 1,
  "username": "ashu"
}
```

---

## Wallet-service (/wallet)

### POST /create

Request

```json
{
  "userId": 1
}
```

Response

```json
{
  "walletId": 1,
  "availableBalance": 0,
  "lockedBalance": 0
}
```

---

### GET /balance

Response

```json
{
  "availableBalance": 1000,
  "lockedBalance": 200,
  "equity": 1050
}
```

Notes:

```text
equity =
availableBalance
+ unrealizedPnl
```

---

### POST /deposit

Request

```json
{
  "amount": 1000
}
```

Response

```json
{
  "message": "Deposit successful",
  "availableBalance": 1000
}
```

---

### POST /withdraw

Request

```json
{
  "amount": 250
}
```

Response

```json
{
  "message": "Withdrawal successful",
  "availableBalance": 750
}
```

---

### GET /transactions

Response

```json
[
  {
    "id": 1,
    "type": "DEPOSIT",
    "amount": 1000
  },
  {
    "id": 2,
    "type": "WITHDRAWAL",
    "amount": 250
  }
]
```

---

## Market-service (/market)

### POST /create

Request

```json
{
  "symbol": "SOL-PERP",
  "maxLeverage": 20,
  "tickSize": 0.01,
  "stepSize": 0.001
}
```

Response

```json
{
  "id": 1,
  "symbol": "SOL-PERP"
}
```

---

### GET /all

Response

```json
[
  {
    "id": 1,
    "symbol": "SOL-PERP"
  },
  {
    "id": 2,
    "symbol": "BTC-PERP"
  }
]
```

---

### GET /:marketId

Response

```json
{
  "id": 1,
  "symbol": "SOL-PERP",
  "maxLeverage": 20,
  "tickSize": 0.01,
  "stepSize": 0.001
}
```

---

## Order-service (/orders)

### POST /create

Request

```json
{
  "marketId": 1,
  "side": "BUY",
  "type": "LIMIT",
  "price": 82.50,
  "quantity": 10,
  "leverage": 5
}
```

Response

```json
{
  "orderId": 101,
  "status": "PENDING"
}
```

---

### GET /all

Response

```json
[
  {
    "orderId": 101,
    "market": "SOL-PERP",
    "side": "BUY",
    "type": "LIMIT",
    "price": 82.5,
    "quantity": 10,
    "filledQuantity": 0,
    "status": "PENDING"
  }
]
```

---

### GET /open

Response

```json
[
  {
    "orderId": 101,
    "status": "PENDING"
  }
]
```

---

### PATCH /update

Request

```json
{
  "price": 82.6,
  "quantity": 12
}
```

Response

```json
{
  "message": "Order updated"
}
```

---

### POST /cancel

Request

```json
{
  "orderId": 101
}
```

Response

```json
{
  "message": "Order cancelled"
}
```

---

## OrderBook (/orderbook)

TBD

Future Endpoints:

```http
GET /orderbook/:marketId
GET /orderbook/:marketId/depth
```

---

## Matching Engine (/matching-engine)

TBD

Internal service only.

Responsibilities:

```text
Match BUY and SELL orders
Generate trades
Update positions
Update balances
```

---

## Trades (/trade)

### GET /history

Response

```json
[
  {
    "tradeId": 1,
    "market": "SOL-PERP",
    "price": 82.64,
    "quantity": 10,
    "side": "BUY",
    "timestamp": "2026-06-01T12:00:00Z"
  }
]
```

---

### GET /market/:marketId

Response

```json
[
  {
    "tradeId": 1,
    "price": 82.64,
    "quantity": 10
  }
]
```

---

## Positions (/position)

### GET /open

Response

```json
[
  {
    "positionId": 1,
    "market": "SOL-PERP",
    "side": "LONG",
    "quantity": 10,
    "entryPrice": 82.5,
    "markPrice": 84,
    "unrealizedPnl": 15,
    "liquidationPrice": 70
  }
]
```

---

### GET /closed

Response

```json
[
  {
    "positionId": 2,
    "market": "BTC-PERP",
    "side": "SHORT",
    "realizedPnl": 120
  }
]
```

---

### GET /:positionId

Response

```json
{
  "positionId": 1,
  "market": "SOL-PERP",
  "side": "LONG",
  "quantity": 10,
  "entryPrice": 82.5,
  "markPrice": 84,
  "unrealizedPnl": 15,
  "liquidationPrice": 70
}
```

---

# Phase 1 Deliverables

* User Registration
* Login & JWT Authentication
* Wallet Creation
* Deposit / Withdraw
* Market Creation
* View Markets
* Create Order
* Modify Order
* Cancel Order
* View Orders
* View Trades
* View Open Positions
* View Closed Positions

# Not In Phase 1

* Oracle Price
* Mark Price Engine
* Funding Rates
* Liquidation Engine
* Insurance Fund
* ADL
* Matching Engine
* Order Book
