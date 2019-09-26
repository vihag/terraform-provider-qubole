package model

import (
	_ "fmt"
)

type UpscalingConfig struct {
	Max_ebs_volume_count          int     `json:"max_ebs_volume_count,omitempty"`
	Percent_free_space_threshold  float32 `json:"percent_free_space_threshold,omitempty"`
	Absolute_free_space_threshold float32 `json:"absolute_free_space_threshold,omitempty"`
	Sampling_interval             int     `json:"sampling_interval,omitempty"`
	Sampling_window               int     `json:"sampling_window,omitempty"`
}

/*
function to flatten EBS Upscaling Config
*/
func FlattenUpscalingConfig(ia *UpscalingConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Max_ebs_volume_count != nil {
		attrs["max_ebs_volume_count"] = ia.Max_ebs_volume_count
	}
	if &ia.Percent_free_space_threshold != nil {
		attrs["percent_free_space_threshold"] = ia.Percent_free_space_threshold
	}
	if &ia.Absolute_free_space_threshold != nil {
		attrs["absolute_free_space_threshold"] = ia.Absolute_free_space_threshold
	}
	if &ia.Sampling_interval != nil {
		attrs["sampling_interval"] = ia.Sampling_interval
	}
	if &ia.Sampling_window != nil {
		attrs["sampling_window"] = ia.Sampling_window
	}

	result = append(result, attrs)

	return result
}

func ReadUpscalingConfigFromTf(upscaling_config *UpscalingConfig, ebsUpscalingConfigs []interface{}) bool {

	if len(ebsUpscalingConfigs) > 0 {
		configs := ebsUpscalingConfigs[0].(map[string]interface{})
		if v, ok := configs["max_ebs_volume_count"]; ok {
			upscaling_config.Max_ebs_volume_count = v.(int)
		}
		if v, ok := configs["percent_free_space_threshold"]; ok {
			upscaling_config.Percent_free_space_threshold = float32(v.(int))
		}
		if v, ok := configs["absolute_free_space_threshold"]; ok {
			upscaling_config.Absolute_free_space_threshold = float32(v.(int))
		}
		if v, ok := configs["sampling_interval"]; ok {
			upscaling_config.Sampling_interval = v.(int)
		}
		if v, ok := configs["sampling_window"]; ok {
			upscaling_config.Sampling_window = v.(int)
		}
	}

	return true
}
