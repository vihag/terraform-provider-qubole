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
