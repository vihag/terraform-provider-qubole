package model

import (
	_ "fmt"
)

type CloudConfig struct {
	Provider       string        `json:"provider,omitempty"`
	Compute_config ComputeConfig `json:"compute_config,omitempty"`
	Location       Location      `json:"location,omitempty"`
	Network_config NetworkConfig `json:"network_config,omitempty"`
	//Azure elements
	Storage_config StorageConfig `json:"storage_config,omitempty"`
}
