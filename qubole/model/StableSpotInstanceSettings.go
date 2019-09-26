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

func ReadStableSpotInstanceSettingsFromTf(stable_spot_instance_settings *StableSpotInstanceSettings, stableSpotInstanceSettings []interface{}) bool {

	if len(stableSpotInstanceSettings) > 0 {
		configs := stableSpotInstanceSettings[0].(map[string]interface{})
		if v, ok := configs["maximum_bid_price_percentage"]; ok {
			stable_spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
		}
		if v, ok := configs["timeout_for_request"]; ok {
			stable_spot_instance_settings.Timeout_for_request = v.(int)
		}
	}

	return true
}
