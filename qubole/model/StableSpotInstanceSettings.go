package model

import (
	_ "fmt"
)

type StableSpotInstanceSettings struct {
	Maximum_bid_price_percentage float32 `json:"maximum_bid_price_percentage,omitempty"`
	Timeout_for_request          int     `json:"timeout_for_request,omitempty"`
}

/*
function to flatten Stable Spot Instance Settings
*/
func FlattenStableSpotInstanceSettings(ia *StableSpotInstanceSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Maximum_bid_price_percentage != nil {
		attrs["maximum_bid_price_percentage"] = ia.Maximum_bid_price_percentage
	}
	if &ia.Timeout_for_request != nil {
		attrs["timeout_for_request"] = ia.Timeout_for_request
	}

	result = append(result, attrs)

	return result
}
