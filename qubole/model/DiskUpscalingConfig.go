package model

import (
	_ "fmt"
)

type DiskUpscalingConfig struct {
	Percent_free_space_threshold  float32 `json:"percent_free_space_threshold,omitempty"`
	Max_data_disk_count           int     `json:"max_data_disk_count,omitempty"`
	Absolute_free_space_threshold float32 `json:"absolute_free_space_threshold,omitempty"`
}

/*
function to Disk Upscaling Settings
*/
func FlattenDiskUpscalingConfig(ia *DiskUpscalingConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Percent_free_space_threshold != nil {
		attrs["percent_free_space_threshold"] = ia.Percent_free_space_threshold
	}
	if &ia.Max_data_disk_count != nil {
		attrs["max_data_disk_count"] = ia.Max_data_disk_count
	}
	if &ia.Absolute_free_space_threshold != nil {
		attrs["absolute_free_space_threshold"] = ia.Absolute_free_space_threshold
	}

	result = append(result, attrs)

	return result
}

func ReadDiskUpscalingConfigFromTf(disk_upscaling_config *DiskUpscalingConfig, diskUpscalingConfig []interface{}) bool {

	if len(diskUpscalingConfig) > 0 {
		configs := diskUpscalingConfig[0].(map[string]interface{})
		if v, ok := configs["percent_free_space_threshold"]; ok {
			disk_upscaling_config.Percent_free_space_threshold = float32(v.(int))
		}
		if v, ok := configs["max_data_disk_count"]; ok {
			disk_upscaling_config.Max_data_disk_count = v.(int)
		}
		if v, ok := configs["absolute_free_space_threshold"]; ok {
			disk_upscaling_config.Absolute_free_space_threshold = float32(v.(int))
		}
	}
	return true
}
