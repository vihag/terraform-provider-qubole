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
