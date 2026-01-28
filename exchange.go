package upbit

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// === Account API ===

// GetAccounts retrieves all account balances.
func (c *Client) GetAccounts() ([]Account, error) {
	body, err := c.get("/accounts", nil, true)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	if err := json.Unmarshal(body, &accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

// === Order API ===

// OrderSide represents the side of an order.
type OrderSide string

const (
	OrderSideBid OrderSide = "bid" // Buy
	OrderSideAsk OrderSide = "ask" // Sell
)

// OrderType represents the type of an order.
type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"  // Limit order
	OrderTypePrice  OrderType = "price"  // Market buy (by price)
	OrderTypeMarket OrderType = "market" // Market sell
	OrderTypeBest   OrderType = "best"   // Best price order
)

// TimeInForce represents the time-in-force option.
type TimeInForce string

const (
	TimeInForceIOC TimeInForce = "ioc" // Immediate or Cancel
	TimeInForceFOK TimeInForce = "fok" // Fill or Kill
)

// OrderState represents the state of an order.
type OrderState string

const (
	OrderStateWait   OrderState = "wait"   // Pending
	OrderStateWatch  OrderState = "watch"  // Reserved
	OrderStateDone   OrderState = "done"   // Completed
	OrderStateCancel OrderState = "cancel" // Cancelled
)

// GetOrderChance retrieves the order constraints for a market.
func (c *Client) GetOrderChance(market string) (*OrderChance, error) {
	params := url.Values{}
	params.Set("market", market)

	body, err := c.get("/orders/chance", params, true)
	if err != nil {
		return nil, err
	}

	var chance OrderChance
	if err := json.Unmarshal(body, &chance); err != nil {
		return nil, err
	}
	return &chance, nil
}

// GetOrder retrieves a single order by UUID.
func (c *Client) GetOrder(uuid string) (*OrderDetail, error) {
	params := url.Values{}
	params.Set("uuid", uuid)

	body, err := c.get("/order", params, true)
	if err != nil {
		return nil, err
	}

	var order OrderDetail
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrderByIdentifier retrieves a single order by custom identifier.
func (c *Client) GetOrderByIdentifier(identifier string) (*OrderDetail, error) {
	params := url.Values{}
	params.Set("identifier", identifier)

	body, err := c.get("/order", params, true)
	if err != nil {
		return nil, err
	}

	var order OrderDetail
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrdersRequest represents the request parameters for GetOrders.
type GetOrdersRequest struct {
	Market      string       // Market code (e.g., "KRW-BTC")
	UUIDs       []string     // List of order UUIDs
	Identifiers []string     // List of custom identifiers
	State       OrderState   // Order state filter
	States      []OrderState // Multiple order states filter
	Page        int          // Page number (deprecated, use cursor)
	Limit       int          // Number of results per page
	OrderBy     string       // "asc" or "desc"
}

// GetOrders retrieves a list of orders.
func (c *Client) GetOrders(req *GetOrdersRequest) ([]Order, error) {
	params := url.Values{}

	if req != nil {
		if req.Market != "" {
			params.Set("market", req.Market)
		}
		if len(req.UUIDs) > 0 {
			for _, uuid := range req.UUIDs {
				params.Add("uuids[]", uuid)
			}
		}
		if len(req.Identifiers) > 0 {
			for _, id := range req.Identifiers {
				params.Add("identifiers[]", id)
			}
		}
		if req.State != "" {
			params.Set("state", string(req.State))
		}
		if len(req.States) > 0 {
			for _, state := range req.States {
				params.Add("states[]", string(state))
			}
		}
		if req.Page > 0 {
			params.Set("page", strconv.Itoa(req.Page))
		}
		if req.Limit > 0 {
			params.Set("limit", strconv.Itoa(req.Limit))
		}
		if req.OrderBy != "" {
			params.Set("order_by", req.OrderBy)
		}
	}

	body, err := c.get("/orders", params, true)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// GetClosedOrders retrieves completed/cancelled orders.
func (c *Client) GetClosedOrders(market string, states []OrderState, startTime, endTime string, limit int, orderBy string) ([]Order, error) {
	params := url.Values{}

	if market != "" {
		params.Set("market", market)
	}
	if len(states) > 0 {
		for _, state := range states {
			params.Add("states[]", string(state))
		}
	}
	if startTime != "" {
		params.Set("start_time", startTime)
	}
	if endTime != "" {
		params.Set("end_time", endTime)
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if orderBy != "" {
		params.Set("order_by", orderBy)
	}

	body, err := c.get("/orders/closed", params, true)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := json.Unmarshal(body, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// PlaceOrderRequest represents the request parameters for placing an order.
type PlaceOrderRequest struct {
	Market      string      // Market code (required)
	Side        OrderSide   // Order side: bid or ask (required)
	Volume      string      // Order volume (required for limit/market sell)
	Price       string      // Order price (required for limit/market buy)
	OrdType     OrderType   // Order type (required)
	Identifier  string      // Custom identifier (optional)
	TimeInForce TimeInForce // Time in force option (optional)
}

// PlaceOrder places a new order.
func (c *Client) PlaceOrder(req *PlaceOrderRequest) (*Order, error) {
	params := url.Values{}
	params.Set("market", req.Market)
	params.Set("side", string(req.Side))
	params.Set("ord_type", string(req.OrdType))

	if req.Volume != "" {
		params.Set("volume", req.Volume)
	}
	if req.Price != "" {
		params.Set("price", req.Price)
	}
	if req.Identifier != "" {
		params.Set("identifier", req.Identifier)
	}
	if req.TimeInForce != "" {
		params.Set("time_in_force", string(req.TimeInForce))
	}

	body, err := c.post("/orders", params, true)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

// CancelOrder cancels an order by UUID.
func (c *Client) CancelOrder(uuid string) (*Order, error) {
	params := url.Values{}
	params.Set("uuid", uuid)

	body, err := c.delete("/order", params, true)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

// CancelOrderByIdentifier cancels an order by custom identifier.
func (c *Client) CancelOrderByIdentifier(identifier string) (*Order, error) {
	params := url.Values{}
	params.Set("identifier", identifier)

	body, err := c.delete("/order", params, true)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := json.Unmarshal(body, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

// === Withdrawal API ===

// GetWithdraws retrieves a list of withdrawals.
func (c *Client) GetWithdraws(currency, state string, uuids, txids []string, limit int, page int, orderBy string) ([]Withdraw, error) {
	params := url.Values{}

	if currency != "" {
		params.Set("currency", currency)
	}
	if state != "" {
		params.Set("state", state)
	}
	if len(uuids) > 0 {
		params.Set("uuids", strings.Join(uuids, ","))
	}
	if len(txids) > 0 {
		params.Set("txids", strings.Join(txids, ","))
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if orderBy != "" {
		params.Set("order_by", orderBy)
	}

	body, err := c.get("/withdraws", params, true)
	if err != nil {
		return nil, err
	}

	var withdraws []Withdraw
	if err := json.Unmarshal(body, &withdraws); err != nil {
		return nil, err
	}
	return withdraws, nil
}

// GetWithdraw retrieves a single withdrawal by UUID.
func (c *Client) GetWithdraw(uuid string) (*Withdraw, error) {
	params := url.Values{}
	params.Set("uuid", uuid)

	body, err := c.get("/withdraw", params, true)
	if err != nil {
		return nil, err
	}

	var withdraw Withdraw
	if err := json.Unmarshal(body, &withdraw); err != nil {
		return nil, err
	}
	return &withdraw, nil
}

// GetWithdrawChance retrieves withdrawal constraints for a currency.
func (c *Client) GetWithdrawChance(currency, netType string) (*WithdrawChance, error) {
	params := url.Values{}
	params.Set("currency", currency)
	if netType != "" {
		params.Set("net_type", netType)
	}

	body, err := c.get("/withdraws/chance", params, true)
	if err != nil {
		return nil, err
	}

	var chance WithdrawChance
	if err := json.Unmarshal(body, &chance); err != nil {
		return nil, err
	}
	return &chance, nil
}

// WithdrawCoinRequest represents the request parameters for coin withdrawal.
type WithdrawCoinRequest struct {
	Currency         string // Currency code (required)
	NetType          string // Network type (required)
	Amount           string // Withdrawal amount (required)
	Address          string // Withdrawal address (required)
	SecondaryAddress string // Secondary address (optional, e.g., XRP tag)
	TransactionType  string // Transaction type (optional)
}

// WithdrawCoin withdraws cryptocurrency to an external address.
func (c *Client) WithdrawCoin(req *WithdrawCoinRequest) (*Withdraw, error) {
	params := url.Values{}
	params.Set("currency", req.Currency)
	params.Set("net_type", req.NetType)
	params.Set("amount", req.Amount)
	params.Set("address", req.Address)

	if req.SecondaryAddress != "" {
		params.Set("secondary_address", req.SecondaryAddress)
	}
	if req.TransactionType != "" {
		params.Set("transaction_type", req.TransactionType)
	}

	body, err := c.post("/withdraws/coin", params, true)
	if err != nil {
		return nil, err
	}

	var withdraw Withdraw
	if err := json.Unmarshal(body, &withdraw); err != nil {
		return nil, err
	}
	return &withdraw, nil
}

// WithdrawKRW withdraws KRW to a registered bank account.
func (c *Client) WithdrawKRW(amount string, twoFactorType string) (*Withdraw, error) {
	params := url.Values{}
	params.Set("amount", amount)
	if twoFactorType != "" {
		params.Set("two_factor_type", twoFactorType)
	}

	body, err := c.post("/withdraws/krw", params, true)
	if err != nil {
		return nil, err
	}

	var withdraw Withdraw
	if err := json.Unmarshal(body, &withdraw); err != nil {
		return nil, err
	}
	return &withdraw, nil
}

// === Deposit API ===

// GetDeposits retrieves a list of deposits.
func (c *Client) GetDeposits(currency, state string, uuids, txids []string, limit int, page int, orderBy string) ([]Deposit, error) {
	params := url.Values{}

	if currency != "" {
		params.Set("currency", currency)
	}
	if state != "" {
		params.Set("state", state)
	}
	if len(uuids) > 0 {
		params.Set("uuids", strings.Join(uuids, ","))
	}
	if len(txids) > 0 {
		params.Set("txids", strings.Join(txids, ","))
	}
	if limit > 0 {
		params.Set("limit", strconv.Itoa(limit))
	}
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if orderBy != "" {
		params.Set("order_by", orderBy)
	}

	body, err := c.get("/deposits", params, true)
	if err != nil {
		return nil, err
	}

	var deposits []Deposit
	if err := json.Unmarshal(body, &deposits); err != nil {
		return nil, err
	}
	return deposits, nil
}

// GetDeposit retrieves a single deposit by UUID.
func (c *Client) GetDeposit(uuid string) (*Deposit, error) {
	params := url.Values{}
	params.Set("uuid", uuid)

	body, err := c.get("/deposit", params, true)
	if err != nil {
		return nil, err
	}

	var deposit Deposit
	if err := json.Unmarshal(body, &deposit); err != nil {
		return nil, err
	}
	return &deposit, nil
}

// GenerateDepositAddress generates a new deposit address for a currency.
func (c *Client) GenerateDepositAddress(currency, netType string) (*DepositAddress, error) {
	params := url.Values{}
	params.Set("currency", currency)
	params.Set("net_type", netType)

	body, err := c.post("/deposits/generate_coin_address", params, true)
	if err != nil {
		return nil, err
	}

	var address DepositAddress
	if err := json.Unmarshal(body, &address); err != nil {
		return nil, err
	}
	return &address, nil
}

// GetDepositAddresses retrieves all deposit addresses.
func (c *Client) GetDepositAddresses() ([]DepositAddress, error) {
	body, err := c.get("/deposits/coin_addresses", nil, true)
	if err != nil {
		return nil, err
	}

	var addresses []DepositAddress
	if err := json.Unmarshal(body, &addresses); err != nil {
		return nil, err
	}
	return addresses, nil
}

// GetDepositAddress retrieves a specific deposit address.
func (c *Client) GetDepositAddress(currency, netType string) (*DepositAddress, error) {
	params := url.Values{}
	params.Set("currency", currency)
	params.Set("net_type", netType)

	body, err := c.get("/deposits/coin_address", params, true)
	if err != nil {
		return nil, err
	}

	var address DepositAddress
	if err := json.Unmarshal(body, &address); err != nil {
		return nil, err
	}
	return &address, nil
}

// === Status API ===

// GetWalletStatus retrieves the wallet status for all currencies.
func (c *Client) GetWalletStatus() ([]WalletStatus, error) {
	body, err := c.get("/status/wallet", nil, true)
	if err != nil {
		return nil, err
	}

	var statuses []WalletStatus
	if err := json.Unmarshal(body, &statuses); err != nil {
		return nil, err
	}
	return statuses, nil
}

// GetAPIKeys retrieves API key information.
func (c *Client) GetAPIKeys() ([]APIKey, error) {
	body, err := c.get("/api_keys", nil, true)
	if err != nil {
		return nil, err
	}

	var keys []APIKey
	if err := json.Unmarshal(body, &keys); err != nil {
		return nil, err
	}
	return keys, nil
}
