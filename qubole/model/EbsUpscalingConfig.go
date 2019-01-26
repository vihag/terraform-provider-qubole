package model

import (
	_ "fmt"
)

type EbsUpscalingConfig struct {
	Max_ebs_volume_count          int     `json:"max_ebs_volume_count,omitempty"`
	Percent_free_space_threshold  float32 `json:"percent_free_space_threshold,omitempty"`
	Absolute_free_space_threshold float32 `json:"absolute_free_space_threshold,omitempty"`
	Sampling_interval             int     `json:"sampling_interval,omitempty"`
	Sampling_window               int     `json:"sampling_window,omitempty"`
}

/*
function to flatten EBS Upscaling Config
*/
func FlattenEbsUpscalingConfig(ia *EbsUpscalingConfig) []map[string]interface{} {
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
