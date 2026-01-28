package main

import (
	"fmt"
	"log"
	"os"

	upbit "github.com/cth-release/go-upbit-sdk"
)

func main() {
	// For public APIs (market data), you don't need API keys
	publicClient := upbit.NewClient("", "")

	// For private APIs (trading), you need API keys
	accessKey := os.Getenv("UPBIT_ACCESS_KEY")
	secretKey := os.Getenv("UPBIT_SECRET_KEY")
	privateClient := upbit.NewClient(accessKey, secretKey)

	// Example 1: Get all markets
	fmt.Println("=== Markets ===")
	markets, err := publicClient.GetMarkets(false)
	if err != nil {
		log.Printf("Failed to get markets: %v", err)
	} else {
		for i, m := range markets {
			if i >= 5 {
				fmt.Printf("... and %d more markets\n", len(markets)-5)
				break
			}
			fmt.Printf("Market: %s (%s)\n", m.Market, m.EnglishName)
		}
	}

	// Example 2: Get ticker for BTC
	fmt.Println("\n=== Ticker ===")
	tickers, err := publicClient.GetTicker([]string{"KRW-BTC", "KRW-ETH"})
	if err != nil {
		log.Printf("Failed to get ticker: %v", err)
	} else {
		for _, t := range tickers {
			fmt.Printf("%s: %.0f KRW (Change: %.2f%%)\n",
				t.Market, t.TradePrice, t.SignedChangeRate*100)
		}
	}

	// Example 3: Get orderbook
	fmt.Println("\n=== Orderbook ===")
	orderbooks, err := publicClient.GetOrderbook([]string{"KRW-BTC"}, 0)
	if err != nil {
		log.Printf("Failed to get orderbook: %v", err)
	} else {
		for _, ob := range orderbooks {
			fmt.Printf("%s Orderbook:\n", ob.Market)
			fmt.Printf("  Total Ask Size: %.8f\n", ob.TotalAskSize)
			fmt.Printf("  Total Bid Size: %.8f\n", ob.TotalBidSize)
			if len(ob.OrderbookUnits) > 0 {
				unit := ob.OrderbookUnits[0]
				fmt.Printf("  Best Ask: %.0f (%.8f)\n", unit.AskPrice, unit.AskSize)
				fmt.Printf("  Best Bid: %.0f (%.8f)\n", unit.BidPrice, unit.BidSize)
			}
		}
	}

	// Example 4: Get candles
	fmt.Println("\n=== Candles (Day) ===")
	candles, err := publicClient.GetDayCandles("KRW-BTC", "", 5, "")
	if err != nil {
		log.Printf("Failed to get candles: %v", err)
	} else {
		fmt.Printf("Candles %d\n", len(candles))
		for _, c := range candles {
			fmt.Printf("%s: Open=%.0f High=%.0f Low=%.0f Close=%.0f\n",
				c.CandleDateTimeKst, c.OpeningPrice, c.HighPrice, c.LowPrice, c.TradePrice)
		}
	}

	// Example 5: Get recent trades
	fmt.Println("\n=== Recent Trades ===")
	trades, err := publicClient.GetTrades("KRW-BTC", "", 5, "", 0)
	if err != nil {
		log.Printf("Failed to get trades: %v", err)
	} else {
		for _, t := range trades {
			fmt.Printf("%s %s: %.0f KRW x %.8f\n",
				t.TradeTimeUtc, t.AskBid, t.TradePrice, t.TradeVolume)
		}
	}

	// Private API examples (requires valid API keys)
	if accessKey != "" && secretKey != "" {
		// Example 6: Get account balances
		fmt.Println("\n=== Account Balances ===")
		accounts, err := privateClient.GetAccounts()
		if err != nil {
			log.Printf("Failed to get accounts: %v", err)
		} else {
			for _, a := range accounts {
				fmt.Printf("%s: %s (Locked: %s)\n", a.Currency, a.Balance, a.Locked)
			}
		}

		// Example 7: Get order chance (market constraints)
		fmt.Println("\n=== Order Chance (KRW-BTC) ===")
		chance, err := privateClient.GetOrderChance("KRW-BTC")
		if err != nil {
			log.Printf("Failed to get order chance: %v", err)
		} else {
			fmt.Printf("Bid Fee: %s\n", chance.BidFee)
			fmt.Printf("Ask Fee: %s\n", chance.AskFee)
			if chance.Market != nil {
				fmt.Printf("Min Total: %s\n", chance.Market.Bid.MinTotal)
			}
		}

		// Example 8: Place a limit order (commented out for safety)
		/*
			fmt.Println("\n=== Place Order ===")
			order, err := privateClient.PlaceOrder(&upbit.PlaceOrderRequest{
				Market:  "KRW-BTC",
				Side:    upbit.OrderSideBid,
				Volume:  "0.0001",
				Price:   "50000000",
				OrdType: upbit.OrderTypeLimit,
			})
			if err != nil {
				log.Printf("Failed to place order: %v", err)
			} else {
				fmt.Printf("Order placed: %s (State: %s)\n", order.UUID, order.State)
			}
		*/

		// Example 9: Get pending orders
		fmt.Println("\n=== Pending Orders ===")
		orders, err := privateClient.GetOrders(&upbit.GetOrdersRequest{
			State: upbit.OrderStateWait,
			Limit: 10,
		})
		if err != nil {
			log.Printf("Failed to get orders: %v", err)
		} else {
			if len(orders) == 0 {
				fmt.Println("No pending orders")
			}
			for _, o := range orders {
				fmt.Printf("Order %s: %s %s %s @ %s\n",
					o.UUID, o.Side, o.Market, o.Volume, o.Price)
			}
		}

		// Example 10: Get wallet status
		fmt.Println("\n=== Wallet Status ===")
		statuses, err := privateClient.GetWalletStatus()
		if err != nil {
			log.Printf("Failed to get wallet status: %v", err)
		} else {
			for i, s := range statuses {
				if i >= 5 {
					fmt.Printf("... and %d more wallets\n", len(statuses)-5)
					break
				}
				fmt.Printf("%s: %s (Block: %s)\n", s.Currency, s.WalletState, s.BlockState)
			}
		}
	} else {
		fmt.Println("\n[Private API examples skipped - set UPBIT_ACCESS_KEY and UPBIT_SECRET_KEY environment variables]")
	}
}
