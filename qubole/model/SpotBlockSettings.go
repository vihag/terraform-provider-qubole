package model

import (
	_ "fmt"
)

type SpotBlockSettings struct {
	Duration int `json:"duration,omitempty"`
}

/*
function to flatten Spot Block Settings
*/
func FlattenSpotBlockSettings(ia *SpotBlockSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Duration != nil {
		attrs["duration"] = ia.Duration
	}

	result = append(result, attrs)

	return result
}

func ReadSpotBlockSettingsFromTf(spot_block_settings *SpotBlockSettings, spotBlockSettings []interface{}) bool {

	if len(spotBlockSettings) > 0 {
		configs := spotBlockSettings[0].(map[string]interface{})
		if v, ok := configs["duration"]; ok {
			spot_block_settings.Duration = v.(int)
		}
	}

	return true
}
