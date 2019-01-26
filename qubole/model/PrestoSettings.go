package model

import (
	_ "fmt"
)

type PrestoSettings struct {
	Presto_version       string `json:"presto_version,omitempty"`
	Custom_presto_config string `json:"custom_presto_config,omitempty"`
	Enable_rubix         bool   `json:"enable_rubix,omitempty"`
}

/*
function to flatten Presto Settings
*/
func FlattenPrestoSettings(ia *PrestoSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Presto_version != nil {
		attrs["presto_version"] = ia.Presto_version
	}
	if &ia.Custom_presto_config != nil {
		attrs["custom_presto_config"] = ia.Custom_presto_config
	}
	if &ia.Enable_rubix != nil {
		attrs["enable_rubix"] = ia.Enable_rubix
	}

	result = append(result, attrs)

	return result
}
