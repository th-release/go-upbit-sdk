# go-upbit-sdk

Go SDK for the [Upbit](https://upbit.com) cryptocurrency exchange API.

## Installation

```bash
go get github.com/cth-release/go-upbit-sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    upbit "github.com/cth-release/go-upbit-sdk"
)

func main() {
    // Public API (no authentication required)
    client := upbit.NewClient("", "")

    // Get all markets
    markets, _ := client.GetMarkets(false)
    for _, m := range markets {
        fmt.Printf("%s: %s\n", m.Market, m.EnglishName)
    }

    // Get ticker
    tickers, _ := client.GetTicker([]string{"KRW-BTC"})
    fmt.Printf("BTC Price: %.0f KRW\n", tickers[0].TradePrice)
}
```

## Authentication

For private APIs (trading, account info), you need API keys from [Upbit](https://upbit.com):

```go
client := upbit.NewClient("your-access-key", "your-secret-key")

// Get account balances
accounts, _ := client.GetAccounts()
```

## API Reference

### Public APIs (Quotation)

| Method | Description |
|--------|-------------|
| `GetMarkets(isDetails)` | Get all available markets |
| `GetTicker(markets)` | Get current price for markets |
| `GetAllTickers(quoteCurrencies)` | Get tickers for all markets |
| `GetOrderbook(markets, level)` | Get orderbook |
| `GetTrades(market, to, count, cursor, daysAgo)` | Get recent trades |
| `GetMinuteCandles(market, unit, to, count)` | Get minute candles |
| `GetDayCandles(market, to, count, convertingPriceUnit)` | Get daily candles |
| `GetWeekCandles(market, to, count)` | Get weekly candles |
| `GetMonthCandles(market, to, count)` | Get monthly candles |

### Private APIs (Exchange)

| Method | Description |
|--------|-------------|
| `GetAccounts()` | Get account balances |
| `GetOrderChance(market)` | Get order constraints |
| `GetOrder(uuid)` | Get order details |
| `GetOrders(request)` | Get order list |
| `GetClosedOrders(...)` | Get completed orders |
| `PlaceOrder(request)` | Place a new order |
| `CancelOrder(uuid)` | Cancel an order |
| `GetWithdraws(...)` | Get withdrawal list |
| `GetWithdraw(uuid)` | Get withdrawal details |
| `GetWithdrawChance(currency, netType)` | Get withdrawal constraints |
| `WithdrawCoin(request)` | Withdraw cryptocurrency |
| `WithdrawKRW(amount, twoFactorType)` | Withdraw KRW |
| `GetDeposits(...)` | Get deposit list |
| `GetDeposit(uuid)` | Get deposit details |
| `GetDepositAddresses()` | Get all deposit addresses |
| `GetDepositAddress(currency, netType)` | Get specific deposit address |
| `GenerateDepositAddress(currency, netType)` | Generate new deposit address |
| `GetWalletStatus()` | Get wallet status |
| `GetAPIKeys()` | Get API key info |

## Examples

### Get Ticker

```go
tickers, err := client.GetTicker([]string{"KRW-BTC", "KRW-ETH"})
if err != nil {
    log.Fatal(err)
}
for _, t := range tickers {
    fmt.Printf("%s: %.0f KRW\n", t.Market, t.TradePrice)
}
```

### Get Orderbook

```go
orderbooks, err := client.GetOrderbook([]string{"KRW-BTC"}, 0)
if err != nil {
    log.Fatal(err)
}
for _, ob := range orderbooks {
    fmt.Printf("Best Bid: %.0f, Best Ask: %.0f\n",
        ob.OrderbookUnits[0].BidPrice,
        ob.OrderbookUnits[0].AskPrice)
}
```

### Get Candles

```go
// Minute candles
candles, _ := client.GetMinuteCandles("KRW-BTC", upbit.CandleUnit1, "", 200)

// Daily candles
candles, _ := client.GetDayCandles("KRW-BTC", "", 200, "")

// Weekly candles
candles, _ := client.GetWeekCandles("KRW-BTC", "", 52)
```

### Place Order

```go
// Limit buy order
order, err := client.PlaceOrder(&upbit.PlaceOrderRequest{
    Market:  "KRW-BTC",
    Side:    upbit.OrderSideBid,
    Volume:  "0.0001",
    Price:   "50000000",
    OrdType: upbit.OrderTypeLimit,
})

// Market buy order (by price)
order, err := client.PlaceOrder(&upbit.PlaceOrderRequest{
    Market:  "KRW-BTC",
    Side:    upbit.OrderSideBid,
    Price:   "10000",  // Buy 10,000 KRW worth
    OrdType: upbit.OrderTypePrice,
})

// Market sell order
order, err := client.PlaceOrder(&upbit.PlaceOrderRequest{
    Market:  "KRW-BTC",
    Side:    upbit.OrderSideAsk,
    Volume:  "0.0001",
    OrdType: upbit.OrderTypeMarket,
})
```

### Cancel Order

```go
order, err := client.CancelOrder("order-uuid-here")
```

### Get Pending Orders

```go
orders, err := client.GetOrders(&upbit.GetOrdersRequest{
    State: upbit.OrderStateWait,
    Limit: 100,
})
```

## Error Handling

```go
tickers, err := client.GetTicker([]string{"KRW-BTC"})
if err != nil {
    if apiErr, ok := err.(*upbit.APIError); ok {
        fmt.Printf("API Error: %s - %s\n", apiErr.Error.Name, apiErr.Error.Message)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## License

MIT License
