package upbit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("access", "secret")
	if client.accessKey != "access" {
		t.Errorf("Expected accessKey 'access', got '%s'", client.accessKey)
	}
	if client.secretKey != "secret" {
		t.Errorf("Expected secretKey 'secret', got '%s'", client.secretKey)
	}
	if client.baseURL != BaseURL {
		t.Errorf("Expected baseURL '%s', got '%s'", BaseURL, client.baseURL)
	}
}

func TestGetMarkets(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/market/all" {
			t.Errorf("Expected path '/v1/market/all', got '%s'", r.URL.Path)
		}

		markets := []Market{
			{Market: "KRW-BTC", KoreanName: "비트코인", EnglishName: "Bitcoin"},
			{Market: "KRW-ETH", KoreanName: "이더리움", EnglishName: "Ethereum"},
		}
		json.NewEncoder(w).Encode(markets)
	}))
	defer server.Close()

	client := NewClient("", "")
	client.SetBaseURL(server.URL + "/v1")

	markets, err := client.GetMarkets(false)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(markets) != 2 {
		t.Errorf("Expected 2 markets, got %d", len(markets))
	}

	if markets[0].Market != "KRW-BTC" {
		t.Errorf("Expected market 'KRW-BTC', got '%s'", markets[0].Market)
	}
}

func TestGetTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/ticker" {
			t.Errorf("Expected path '/v1/ticker', got '%s'", r.URL.Path)
		}

		markets := r.URL.Query().Get("markets")
		if markets != "KRW-BTC" {
			t.Errorf("Expected markets 'KRW-BTC', got '%s'", markets)
		}

		tickers := []Ticker{
			{
				Market:     "KRW-BTC",
				TradePrice: 50000000,
				Change:     "RISE",
			},
		}
		json.NewEncoder(w).Encode(tickers)
	}))
	defer server.Close()

	client := NewClient("", "")
	client.SetBaseURL(server.URL + "/v1")

	tickers, err := client.GetTicker([]string{"KRW-BTC"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tickers) != 1 {
		t.Errorf("Expected 1 ticker, got %d", len(tickers))
	}

	if tickers[0].TradePrice != 50000000 {
		t.Errorf("Expected trade price 50000000, got %f", tickers[0].TradePrice)
	}
}

func TestGetOrderbook(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/orderbook" {
			t.Errorf("Expected path '/v1/orderbook', got '%s'", r.URL.Path)
		}

		orderbooks := []Orderbook{
			{
				Market:       "KRW-BTC",
				TotalAskSize: 10.5,
				TotalBidSize: 8.3,
				OrderbookUnits: []OrderbookUnit{
					{AskPrice: 50100000, BidPrice: 50000000, AskSize: 1.5, BidSize: 2.0},
				},
			},
		}
		json.NewEncoder(w).Encode(orderbooks)
	}))
	defer server.Close()

	client := NewClient("", "")
	client.SetBaseURL(server.URL + "/v1")

	orderbooks, err := client.GetOrderbook([]string{"KRW-BTC"}, 0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(orderbooks) != 1 {
		t.Errorf("Expected 1 orderbook, got %d", len(orderbooks))
	}

	if orderbooks[0].TotalAskSize != 10.5 {
		t.Errorf("Expected total ask size 10.5, got %f", orderbooks[0].TotalAskSize)
	}
}

func TestAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":{"name":"invalid_parameter","message":"Invalid market"}}`))
	}))
	defer server.Close()

	client := NewClient("", "")
	client.SetBaseURL(server.URL + "/v1")

	_, err := client.GetTicker([]string{"INVALID"})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("Expected APIError, got %T", err)
	}

	if apiErr.Err.Name != "invalid_parameter" {
		t.Errorf("Expected error name 'invalid_parameter', got '%s'", apiErr.Err.Name)
	}
}

func TestGenerateToken(t *testing.T) {
	client := NewClient("test-access-key", "test-secret-key")

	token, err := client.generateToken(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Error("Expected non-empty token")
	}
}

func TestGetAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/accounts" {
			t.Errorf("Expected path '/v1/accounts', got '%s'", r.URL.Path)
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			t.Error("Expected Authorization header")
		}

		accounts := []Account{
			{Currency: "KRW", Balance: "1000000", Locked: "0"},
			{Currency: "BTC", Balance: "0.1", Locked: "0"},
		}
		json.NewEncoder(w).Encode(accounts)
	}))
	defer server.Close()

	client := NewClient("access", "secret")
	client.SetBaseURL(server.URL + "/v1")

	accounts, err := client.GetAccounts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(accounts) != 2 {
		t.Errorf("Expected 2 accounts, got %d", len(accounts))
	}
}

func TestPlaceOrder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/v1/orders" {
			t.Errorf("Expected path '/v1/orders', got '%s'", r.URL.Path)
		}

		order := Order{
			UUID:    "test-uuid",
			Market:  "KRW-BTC",
			Side:    "bid",
			OrdType: "limit",
			State:   "wait",
			Volume:  "0.0001",
			Price:   "50000000",
		}
		json.NewEncoder(w).Encode(order)
	}))
	defer server.Close()

	client := NewClient("access", "secret")
	client.SetBaseURL(server.URL + "/v1")

	order, err := client.PlaceOrder(&PlaceOrderRequest{
		Market:  "KRW-BTC",
		Side:    OrderSideBid,
		Volume:  "0.0001",
		Price:   "50000000",
		OrdType: OrderTypeLimit,
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if order.UUID != "test-uuid" {
		t.Errorf("Expected UUID 'test-uuid', got '%s'", order.UUID)
	}
}

func TestCancelOrder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		order := Order{
			UUID:  "test-uuid",
			State: "cancel",
		}
		json.NewEncoder(w).Encode(order)
	}))
	defer server.Close()

	client := NewClient("access", "secret")
	client.SetBaseURL(server.URL + "/v1")

	order, err := client.CancelOrder("test-uuid")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if order.State != "cancel" {
		t.Errorf("Expected state 'cancel', got '%s'", order.State)
	}
}

func TestGetCandles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		candles := []Candle{
			{
				Market:       "KRW-BTC",
				OpeningPrice: 50000000,
				HighPrice:    51000000,
				LowPrice:     49000000,
				TradePrice:   50500000,
			},
		}
		json.NewEncoder(w).Encode(candles)
	}))
	defer server.Close()

	client := NewClient("", "")
	client.SetBaseURL(server.URL + "/v1")

	// Test minute candles
	candles, err := client.GetMinuteCandles("KRW-BTC", CandleUnit1, "", 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(candles) != 1 {
		t.Errorf("Expected 1 candle, got %d", len(candles))
	}

	// Test day candles
	candles, err = client.GetDayCandles("KRW-BTC", "", 1, "")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test week candles
	candles, err = client.GetWeekCandles("KRW-BTC", "", 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test month candles
	candles, err = client.GetMonthCandles("KRW-BTC", "", 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
