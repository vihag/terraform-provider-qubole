package model

import (
	_ "fmt"
)

type HeterogeneousConfig struct {
	Memory []map[string]interface{} `json:"memory,omitempty"`
}

/*
function to flatten Heterogeneous Configuration
*/
func FlattenHeterogeneousConfig(ia *HeterogeneousConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Memory != nil {
		attrs["memory"] = ia.Memory
	}

	result = append(result, attrs)

	return result
}
