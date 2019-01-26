package model

import (
	_ "fmt"
)

type SpotSettings struct {
	Spot_instance_settings        SpotInstanceSettings       `json:"spot_instance_settings,omitempty"`
	Spot_block_settings           SpotBlockSettings          `json:"spot_block_settings,omitempty"`
	Stable_spot_instance_settings StableSpotInstanceSettings `json:"stable_spot_instance_settings,omitempty"`
}

/*
function to flatten Spot Settings
*/
func FlattenSpotSettings(ia *SpotSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Spot_instance_settings != nil {
		attrs["spot_instance_settings"] = FlattenSpotInstanceSettings(&ia.Spot_instance_settings)
	}
	if &ia.Spot_block_settings != nil {
		attrs["spot_block_settings"] = FlattenSpotBlockSettings(&ia.Spot_block_settings)
	}
	if &ia.Stable_spot_instance_settings != nil {
		attrs["stable_spot_instance_settings"] = FlattenStableSpotInstanceSettings(&ia.Stable_spot_instance_settings)
	}

	result = append(result, attrs)

	return result
}
