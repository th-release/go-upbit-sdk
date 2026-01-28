package upbit

import "fmt"

// ErrorDetail contains the error details from the API.
type ErrorDetail struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// APIError represents an error response from the Upbit API.
type APIError struct {
	Err ErrorDetail `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("upbit API error: %s - %s", e.Err.Name, e.Err.Message)
}

// Common error names returned by Upbit API
const (
	ErrInvalidParameter       = "invalid_parameter"
	ErrUnauthorized           = "unauthorized"
	ErrInvalidQuery           = "invalid_query"
	ErrJWTVerificationFailed  = "jwt_verification_fail"
	ErrExpiredAccessKey       = "expired_access_key"
	ErrNonceUsed              = "nonce_used"
	ErrNoAuthorizationIP      = "no_authorization_ip"
	ErrOutOfScope             = "out_of_scope"
	ErrTooManyRequests        = "too_many_requests"
	ErrOrderNotFound          = "order_not_found"
	ErrInsufficientFunds      = "insufficient_funds"
	ErrUnderMinTotalBid       = "under_min_total_bid"
	ErrUnderMinTotalAsk       = "under_min_total_ask"
	ErrWidgetMakerOnlyOrder   = "widgetmaker_only_order"
	ErrMarketOrderDisabled    = "market_order_disabled"
	ErrInvalidVolume          = "invalid_volume"
	ErrInvalidPrice           = "invalid_price"
	ErrInvalidMarket          = "invalid_market"
	ErrOrderCancelled         = "order_cancelled"
	ErrOrderExecuted          = "order_executed"
	ErrServerError            = "server_error"
	ErrInternalServerError    = "internal_server_error"
	ErrUnknown                = "unknown"
)
