package model

import (
	_ "fmt"
)

type SpotSettings struct {
	Spot_instance_settings        SpotInstanceSettings       `json:"spot_instance_settings,omitempty"`
	Spot_block_settings           SpotBlockSettings          `json:"spot_block_settings,omitempty"`
	Stable_spot_instance_settings StableSpotInstanceSettings `json:"stable_spot_instance_settings,omitempty"`
}
