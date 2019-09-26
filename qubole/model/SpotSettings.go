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

func ReadSpotSettingsFromTf(spot_settings *SpotSettings, spotSettings []interface{}) bool {

	if len(spotSettings) > 0 {
		configs := spotSettings[0].(map[string]interface{})

		//Read spot instance settings
		var spot_instance_settings SpotInstanceSettings
		if v, ok := configs["spot_instance_settings"]; ok {
			spotInstanceSettings := v.([]interface{})
			ReadSpotInstanceSettingsFromTf(&spot_instance_settings, spotInstanceSettings)
			spot_settings.Spot_instance_settings = spot_instance_settings
		}

		//Read stable spot instance settings
		var stable_spot_instance_settings StableSpotInstanceSettings
		if v, ok := configs["stable_spot_instance_settings"]; ok {
			stableSpotInstanceSettings := v.([]interface{})
			ReadStableSpotInstanceSettingsFromTf(&stable_spot_instance_settings, stableSpotInstanceSettings)
			spot_settings.Stable_spot_instance_settings = stable_spot_instance_settings
		}

		//Read spot block settings
		var spot_block_settings SpotBlockSettings
		if v, ok := configs["spot_block_settings"]; ok {
			spotBlockSettings := v.([]interface{})
			ReadSpotBlockSettingsFromTf(&spot_block_settings, spotBlockSettings)
			spot_settings.Spot_block_settings = spot_block_settings
		}
	}

	return true
}
