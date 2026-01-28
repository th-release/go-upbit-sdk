package upbit

// Account represents a user's account balance information.
type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

// Market represents a trading pair.
type Market struct {
	Market        string `json:"market"`
	KoreanName    string `json:"korean_name"`
	EnglishName   string `json:"english_name"`
	MarketWarning string `json:"market_warning,omitempty"`
	MarketEvent   *struct {
		Warning bool `json:"warning"`
		Caution *struct {
			PriceFluctuations            bool `json:"PRICE_FLUCTUATIONS"`
			TradingVolumeSoaring         bool `json:"TRADING_VOLUME_SOARING"`
			DepositAmountSoaring         bool `json:"DEPOSIT_AMOUNT_SOARING"`
			GlobalPriceDifferences       bool `json:"GLOBAL_PRICE_DIFFERENCES"`
			ConcentrationOfSmallAccounts bool `json:"CONCENTRATION_OF_SMALL_ACCOUNTS"`
		} `json:"caution,omitempty"`
	} `json:"market_event,omitempty"`
}

// Ticker represents the current price information for a market.
type Ticker struct {
	Market             string  `json:"market"`
	TradeDate          string  `json:"trade_date"`
	TradeTime          string  `json:"trade_time"`
	TradeDateKst       string  `json:"trade_date_kst"`
	TradeTimeKst       string  `json:"trade_time_kst"`
	TradeTimestamp     int64   `json:"trade_timestamp"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	PrevClosingPrice   float64 `json:"prev_closing_price"`
	Change             string  `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          int64   `json:"timestamp"`
}

// Orderbook represents the order book for a market.
type Orderbook struct {
	Market         string           `json:"market"`
	Timestamp      int64            `json:"timestamp"`
	TotalAskSize   float64          `json:"total_ask_size"`
	TotalBidSize   float64          `json:"total_bid_size"`
	OrderbookUnits []OrderbookUnit  `json:"orderbook_units"`
	Level          int              `json:"level,omitempty"`
}

// OrderbookUnit represents a single price level in the order book.
type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"`
	BidPrice float64 `json:"bid_price"`
	AskSize  float64 `json:"ask_size"`
	BidSize  float64 `json:"bid_size"`
}

// Trade represents a recent trade.
type Trade struct {
	Market           string  `json:"market"`
	TradeDateUtc     string  `json:"trade_date_utc"`
	TradeTimeUtc     string  `json:"trade_time_utc"`
	Timestamp        int64   `json:"timestamp"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	ChangePrice      float64 `json:"change_price"`
	AskBid           string  `json:"ask_bid"`
	SequentialID     int64   `json:"sequential_id"`
}

// Candle represents OHLCV candlestick data.
type Candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUtc    string  `json:"candle_date_time_utc"`
	CandleDateTimeKst    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int64   `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	Unit                 int     `json:"unit,omitempty"`
	FirstDayOfPeriod     string  `json:"first_day_of_period,omitempty"`
}

// Order represents an order.
type Order struct {
	UUID            string  `json:"uuid"`
	Side            string  `json:"side"`
	OrdType         string  `json:"ord_type"`
	Price           string  `json:"price,omitempty"`
	State           string  `json:"state"`
	Market          string  `json:"market"`
	CreatedAt       string  `json:"created_at"`
	Volume          string  `json:"volume"`
	RemainingVolume string  `json:"remaining_volume"`
	ReservedFee     string  `json:"reserved_fee"`
	RemainingFee    string  `json:"remaining_fee"`
	PaidFee         string  `json:"paid_fee"`
	Locked          string  `json:"locked"`
	ExecutedVolume  string  `json:"executed_volume"`
	TradesCount     int     `json:"trades_count"`
	TimeInForce     string  `json:"time_in_force,omitempty"`
}

// OrderDetail represents detailed order information including trades.
type OrderDetail struct {
	Order
	Trades []OrderTrade `json:"trades,omitempty"`
}

// OrderTrade represents a trade that is part of an order.
type OrderTrade struct {
	Market    string `json:"market"`
	UUID      string `json:"uuid"`
	Price     string `json:"price"`
	Volume    string `json:"volume"`
	Funds     string `json:"funds"`
	Side      string `json:"side"`
	CreatedAt string `json:"created_at"`
}

// OrderChance represents the order constraints for a market.
type OrderChance struct {
	BidFee     string `json:"bid_fee"`
	AskFee     string `json:"ask_fee"`
	Market     *struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		OrderTypes []string `json:"order_types"`
		AskTypes   []string `json:"ask_types"`
		BidTypes   []string `json:"bid_types"`
		OrderSides []string `json:"order_sides"`
		Bid        *struct {
			Currency  string `json:"currency"`
			MinTotal  string `json:"min_total"`
		} `json:"bid"`
		Ask        *struct {
			Currency  string `json:"currency"`
			MinTotal  string `json:"min_total"`
		} `json:"ask"`
		MaxTotal   string   `json:"max_total"`
		State      string   `json:"state"`
	} `json:"market"`
	BidAccount *Account `json:"bid_account"`
	AskAccount *Account `json:"ask_account"`
}

// Withdraw represents a withdrawal record.
type Withdraw struct {
	Type            string `json:"type"`
	UUID            string `json:"uuid"`
	Currency        string `json:"currency"`
	NetType         string `json:"net_type,omitempty"`
	TxID            string `json:"txid,omitempty"`
	State           string `json:"state"`
	CreatedAt       string `json:"created_at"`
	DoneAt          string `json:"done_at,omitempty"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	TransactionType string `json:"transaction_type"`
}

// WithdrawChance represents withdrawal constraints.
type WithdrawChance struct {
	MemberLevel *struct {
		SecurityLevel        int  `json:"security_level"`
		FeeLevel             int  `json:"fee_level"`
		EmailVerified        bool `json:"email_verified"`
		IdentityAuthVerified bool `json:"identity_auth_verified"`
		BankAccountVerified  bool `json:"bank_account_verified"`
		KakaoPayAuthVerified bool `json:"kakao_pay_auth_verified"`
		Locked               bool `json:"locked"`
		WalletLocked         bool `json:"wallet_locked"`
	} `json:"member_level"`
	Currency *struct {
		Code          string   `json:"code"`
		WithdrawFee   string   `json:"withdraw_fee"`
		IsCoin        bool     `json:"is_coin"`
		WalletState   string   `json:"wallet_state"`
		WalletSupport []string `json:"wallet_support"`
	} `json:"currency"`
	Account *struct {
		Currency            string `json:"currency"`
		Balance             string `json:"balance"`
		Locked              string `json:"locked"`
		AvgBuyPrice         string `json:"avg_buy_price"`
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
		UnitCurrency        string `json:"unit_currency"`
	} `json:"account"`
	WithdrawLimit *struct {
		Currency                   string `json:"currency"`
		Minimum                    string `json:"minimum"`
		Onetime                    string `json:"onetime"`
		Daily                      string `json:"daily"`
		RemainingDaily             string `json:"remaining_daily"`
		RemainingDailyKrw          string `json:"remaining_daily_krw"`
		Fixed                      int    `json:"fixed"`
		CanWithdraw                bool   `json:"can_withdraw"`
	} `json:"withdraw_limit"`
}

// Deposit represents a deposit record.
type Deposit struct {
	Type            string `json:"type"`
	UUID            string `json:"uuid"`
	Currency        string `json:"currency"`
	NetType         string `json:"net_type,omitempty"`
	TxID            string `json:"txid,omitempty"`
	State           string `json:"state"`
	CreatedAt       string `json:"created_at"`
	DoneAt          string `json:"done_at,omitempty"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	TransactionType string `json:"transaction_type"`
}

// DepositAddress represents a deposit address.
type DepositAddress struct {
	Currency       string `json:"currency"`
	NetType        string `json:"net_type"`
	DepositAddress string `json:"deposit_address"`
	SecondaryAddress string `json:"secondary_address,omitempty"`
}

// CoinAddress represents a coin address for withdrawal.
type CoinAddress struct {
	Currency       string `json:"currency"`
	NetType        string `json:"net_type"`
	Network        string `json:"network_name"`
	DepositAddress string `json:"deposit_address"`
	SecondaryAddress string `json:"secondary_address,omitempty"`
}

// APIKey represents API key information.
type APIKey struct {
	AccessKey string `json:"access_key"`
	ExpireAt  string `json:"expire_at"`
}

// WalletStatus represents the wallet status for a currency.
type WalletStatus struct {
	Currency     string `json:"currency"`
	WalletState  string `json:"wallet_state"`
	BlockState   string `json:"block_state"`
	BlockHeight  int64  `json:"block_height"`
	BlockUpdated string `json:"block_updated_at"`
	NetType      string `json:"net_type"`
}
