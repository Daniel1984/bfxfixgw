package convert

import (
	"strconv"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/common"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/order"
	bfxv1 "github.com/bitfinexcom/bitfinex-api-go/v1"
)

// converts messages from FIX to bitfinex
// Bitfinex types.

// Int64OrZero tries to get an int64 from a generic interface or returns 0
func Int64OrZero(i interface{}) int64 {
	if r, ok := i.(int64); ok {
		return r
	}
	return 0
}

// Float64OrZero tries to get a float64 from a generic interface or returns 0
func Float64OrZero(i interface{}) float64 {
	if r, ok := i.(float64); ok {
		return r
	}
	return 0.0
}

// BoolOrFalse tries to get a bool from a generic interface or returns false
func BoolOrFalse(i interface{}) bool {
	if r, ok := i.(bool); ok {
		return r
	}
	return false
}

// StringOrEmpty tries to get a string from a generic interface or returns an empty string
func StringOrEmpty(i interface{}) string {
	if r, ok := i.(string); ok {
		return r
	}
	return ""
}

// OrderFromV1Order converts a bitfinex v1 type order to v2
func OrderFromV1Order(o bfxv1.Order) (*order.Order, error) {
	out := &order.Order{}

	out.ID = o.ID
	out.Symbol = o.Symbol
	out.Hidden = o.IsHidden

	ts, err := strconv.ParseFloat(o.Timestamp, 64)
	if err != nil {
		return nil, err
	}
	out.MTSCreated = int64(ts)
	out.MTSUpdated = int64(ts)

	p, err := strconv.ParseFloat(o.Price, 64)
	if err != nil {
		return nil, err
	}
	out.Price = p

	ap, err := strconv.ParseFloat(o.AvgExecutionPrice, 64)
	if err != nil {
		return nil, err
	}
	out.PriceAvg = ap

	switch {
	case o.IsCanceled:
		out.Status = common.OrderStatusCanceled
	case o.IsLive:
		out.Status = common.OrderStatusActive
	}

	mul := 1
	if o.Side == "sell" {
		mul = -1
	}
	oa, err := strconv.ParseFloat(o.OriginalAmount, 64)
	if err != nil {
		return nil, err
	}
	out.AmountOrig = oa
	or, err := strconv.ParseFloat(o.RemainingAmount, 64)
	if err != nil {
		return nil, err
	}
	out.Amount = or * float64(mul)

	switch o.Type {
	case "market":
		out.Type = common.OrderTypeMarket
	case "limit":
		out.Type = common.OrderTypeLimit
	case "exchange limit":
		out.Type = common.OrderTypeExchangeLimit
	case "stop":
		out.Type = common.OrderTypeStop
	case "trailing-stop":
		out.Type = common.OrderTypeTrailingStop
	}

	//out.PlacedID = o.
	return out, nil
}
