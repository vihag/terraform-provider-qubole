package model

import (
	_ "fmt"
	"log"
)

type CloudConfig struct {
	Provider       string        `json:"provider,omitempty"`
	Compute_config ComputeConfig `json:"compute_config,omitempty"`
	Location       Location      `json:"location,omitempty"`
	Network_config NetworkConfig `json:"network_config,omitempty"`
	//Azure elements
	Storage_config StorageConfig `json:"storage_config,omitempty"`
}

/*
function to flatten Cloud Config
*/
func FlattenCloudConfig(ia *CloudConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Provider != nil {
		attrs["provider"] = ia.Provider
	}

	if &ia.Compute_config != nil {
		attrs["compute_config"] = FlattenComputeConfig(&ia.Compute_config)
	}

	if &ia.Location != nil {
		attrs["location"] = FlattenLocation(&ia.Location)
	}

	if &ia.Network_config != nil {
		attrs["network_config"] = FlattenNetworkConfig(&ia.Network_config)
	}

	if &ia.Storage_config != nil {
		attrs["storage_config"] = FlattenStorageConfig(&ia.Storage_config)
	}

	result = append(result, attrs)

	log.Print(result)
	return result
}
