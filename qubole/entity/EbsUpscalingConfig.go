package entity

import (
	_ "fmt"
)

type EbsUpscalingConfig struct {
	Max_ebs_volume_count          int
	Percent_free_space_threshold  int
	Absolute_free_space_threshold int
	Sampling_interval             int
	Sampling_window               int
}
