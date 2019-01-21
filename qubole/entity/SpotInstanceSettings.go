package entity

import (
	_ "fmt"
)

type SpotInstanceSettings struct {
	Maximum_bid_price_percentage     int	`json:"maximum_bid_price_percentage,omitempty"`
	Timeout_for_request              int	`json:"timeout_for_request,omitempty"`
	Maximum_spot_instance_percentage int	`json:"maximum_spot_instance_percentage,omitempty"`
}
