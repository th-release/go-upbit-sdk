/*
Package upbit provides a Go SDK for the Upbit cryptocurrency exchange API.

# Overview

This SDK provides access to both public (quotation) and private (exchange) APIs.
Public APIs don't require authentication and can be used to get market data.
Private APIs require API keys and can be used for trading operations.

# Quick Start

Create a client:

	// For public APIs (no authentication needed)
	client := upbit.NewClient("", "")

	// For private APIs (authentication required)
	client := upbit.NewClient("your-access-key", "your-secret-key")

# Public APIs

Get market information:

	markets, _ := client.GetMarkets(false)
	tickers, _ := client.GetTicker([]string{"KRW-BTC"})
	orderbooks, _ := client.GetOrderbook([]string{"KRW-BTC"}, 0)
	trades, _ := client.GetTrades("KRW-BTC", "", 100, "", 0)
	candles, _ := client.GetDayCandles("KRW-BTC", "", 200, "")

# Private APIs

Account operations:

	accounts, _ := client.GetAccounts()

Order operations:

	// Place order
	order, _ := client.PlaceOrder(&upbit.PlaceOrderRequest{
		Market:  "KRW-BTC",
		Side:    upbit.OrderSideBid,
		Volume:  "0.0001",
		Price:   "50000000",
		OrdType: upbit.OrderTypeLimit,
	})

	// Cancel order
	cancelled, _ := client.CancelOrder(order.UUID)

	// Get orders
	orders, _ := client.GetOrders(&upbit.GetOrdersRequest{
		State: upbit.OrderStateWait,
	})

# Error Handling

All API errors are returned as *APIError:

	tickers, err := client.GetTicker([]string{"KRW-BTC"})
	if err != nil {
		if apiErr, ok := err.(*upbit.APIError); ok {
			fmt.Printf("API Error: %s - %s\n", apiErr.Err.Name, apiErr.Err.Message)
		}
	}
*/
package upbit
