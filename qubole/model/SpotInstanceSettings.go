package model

import (
	_ "fmt"
)

type SpotInstanceSettings struct {
	Maximum_bid_price_percentage     float32	`json:"maximum_bid_price_percentage,omitempty"`
	Timeout_for_request              int	`json:"timeout_for_request,omitempty"`
	Maximum_spot_instance_percentage float32	`json:"maximum_spot_instance_percentage,omitempty"`
}
