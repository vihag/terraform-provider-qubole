package model

import (
	_ "fmt"
)

type DiskUpscalingConfig struct {
	Percent_free_space_threshold  float32 `json:"percent_free_space_threshold,omitempty"`
	Max_data_disk_count           int     `json:"max_data_disk_count,omitempty"`
	Absolute_free_space_threshold float32 `json:"absolute_free_space_threshold,omitempty"`
}
