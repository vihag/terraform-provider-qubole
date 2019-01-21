package entity

import (
	_ "fmt"
)

type StableSpotInstanceSettings struct {
	Maximum_bid_price_percentage     int	`json:"maximum_bid_price_percentage,omitempty"`
	Timeout_for_request              int	`json:"timeout_for_request,omitempty"`
}

