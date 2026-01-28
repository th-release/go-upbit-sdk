package upbit

import (
	"encoding/json"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

// === Market API ===

// GetMarkets retrieves all available markets.
func (c *Client) GetMarkets(isDetails bool) ([]Market, error) {
	params := url.Values{}
	if isDetails {
		params.Set("is_details", "true")
	}

	body, err := c.get("/market/all", params, false)
	if err != nil {
		return nil, err
	}

	var markets []Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}
	return markets, nil
}

// === Candle API ===

// CandleUnit represents the unit for minute candles.
type CandleUnit int

const (
	CandleUnit1   CandleUnit = 1
	CandleUnit3   CandleUnit = 3
	CandleUnit5   CandleUnit = 5
	CandleUnit10  CandleUnit = 10
	CandleUnit15  CandleUnit = 15
	CandleUnit30  CandleUnit = 30
	CandleUnit60  CandleUnit = 60
	CandleUnit240 CandleUnit = 240
)

// GetMinuteCandles retrieves minute candles for a market.
func (c *Client) GetMinuteCandles(market string, unit CandleUnit, to string, count int) ([]Candle, error) {
	params := url.Values{}
	params.Set("market", market)
	if to != "" {
		params.Set("to", to)
	}
	if count > 0 {
		params.Set("count", strconv.Itoa(count))
	}

	endpoint := "/candles/minutes/" + strconv.Itoa(int(unit))
	body, err := c.get(endpoint, params, false)
	if err != nil {
		return nil, err
	}

	var candles []Candle
	if err := json.Unmarshal(body, &candles); err != nil {
		return nil, err
	}
	slices.Reverse(candles)
	return candles, nil
}

// GetDayCandles retrieves daily candles for a market.
func (c *Client) GetDayCandles(market string, to string, count int, convertingPriceUnit string) ([]Candle, error) {
	params := url.Values{}
	params.Set("market", market)
	if to != "" {
		params.Set("to", to)
	}
	if count > 0 {
		params.Set("count", strconv.Itoa(count))
	}
	if convertingPriceUnit != "" {
		params.Set("converting_price_unit", convertingPriceUnit)
	}

	body, err := c.get("/candles/days", params, false)
	if err != nil {
		return nil, err
	}

	var candles []Candle
	if err := json.Unmarshal(body, &candles); err != nil {
		return nil, err
	}
	slices.Reverse(candles)
	return candles, nil
}

// GetWeekCandles retrieves weekly candles for a market.
func (c *Client) GetWeekCandles(market string, to string, count int) ([]Candle, error) {
	params := url.Values{}
	params.Set("market", market)
	if to != "" {
		params.Set("to", to)
	}
	if count > 0 {
		params.Set("count", strconv.Itoa(count))
	}

	body, err := c.get("/candles/weeks", params, false)
	if err != nil {
		return nil, err
	}

	var candles []Candle
	if err := json.Unmarshal(body, &candles); err != nil {
		return nil, err
	}
	slices.Reverse(candles)
	return candles, nil
}

// GetMonthCandles retrieves monthly candles for a market.
func (c *Client) GetMonthCandles(market string, to string, count int) ([]Candle, error) {
	params := url.Values{}
	params.Set("market", market)
	if to != "" {
		params.Set("to", to)
	}
	if count > 0 {
		params.Set("count", strconv.Itoa(count))
	}

	body, err := c.get("/candles/months", params, false)
	if err != nil {
		return nil, err
	}

	var candles []Candle
	if err := json.Unmarshal(body, &candles); err != nil {
		return nil, err
	}
	slices.Reverse(candles)
	return candles, nil
}

// === Ticker API ===

// GetTicker retrieves the current ticker for specified markets.
func (c *Client) GetTicker(markets []string) ([]Ticker, error) {
	params := url.Values{}
	params.Set("markets", strings.Join(markets, ","))

	body, err := c.get("/ticker", params, false)
	if err != nil {
		return nil, err
	}

	var tickers []Ticker
	if err := json.Unmarshal(body, &tickers); err != nil {
		return nil, err
	}
	return tickers, nil
}

// GetAllTickers retrieves tickers for all markets of specified quote currencies.
func (c *Client) GetAllTickers(quoteCurrencies []string) ([]Ticker, error) {
	params := url.Values{}
	if len(quoteCurrencies) > 0 {
		params.Set("quote_currencies", strings.Join(quoteCurrencies, ","))
	}

	body, err := c.get("/ticker/all", params, false)
	if err != nil {
		return nil, err
	}

	var tickers []Ticker
	if err := json.Unmarshal(body, &tickers); err != nil {
		return nil, err
	}
	return tickers, nil
}

// === Orderbook API ===

// GetOrderbook retrieves the orderbook for specified markets.
func (c *Client) GetOrderbook(markets []string, level int) ([]Orderbook, error) {
	params := url.Values{}
	params.Set("markets", strings.Join(markets, ","))
	if level > 0 {
		params.Set("level", strconv.Itoa(level))
	}

	body, err := c.get("/orderbook", params, false)
	if err != nil {
		return nil, err
	}

	var orderbooks []Orderbook
	if err := json.Unmarshal(body, &orderbooks); err != nil {
		return nil, err
	}
	return orderbooks, nil
}

// === Trade API ===

// GetTrades retrieves recent trades for a market.
func (c *Client) GetTrades(market string, to string, count int, cursor string, daysAgo int) ([]Trade, error) {
	params := url.Values{}
	params.Set("market", market)
	if to != "" {
		params.Set("to", to)
	}
	if count > 0 {
		params.Set("count", strconv.Itoa(count))
	}
	if cursor != "" {
		params.Set("cursor", cursor)
	}
	if daysAgo > 0 {
		params.Set("days_ago", strconv.Itoa(daysAgo))
	}

	body, err := c.get("/trades/ticks", params, false)
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err := json.Unmarshal(body, &trades); err != nil {
		return nil, err
	}
	return trades, nil
}
