package model

import (
	_ "fmt"
)

type HiveSettings struct {
	Is_hs2                    bool   `json:"is_hs2,omitempty"`
	Hive_version              string `json:"hive_version,omitempty"`
	Is_metadata_cache_enabled bool   `json:"is_metadata_cache_enabled,omitempty"`
	Hs2_thrift_port           int    `json:"hs2_thrift_port,omitempty"`
	Overrides                 string `json:"overrides,omitempty"`
	Execution_engine          string `json:"execution_engine,omitempty"`
}

/*
function to flatten Hive Settings
*/
func FlattenHiveSettings(ia *HiveSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Is_hs2 != nil {
		attrs["is_hs2"] = ia.Is_hs2
	}
	if &ia.Hive_version != nil {
		attrs["hive_version"] = ia.Hive_version
	}
	if &ia.Is_metadata_cache_enabled != nil {
		attrs["is_metadata_cache_enabled"] = ia.Is_metadata_cache_enabled
	}
	if &ia.Hs2_thrift_port != nil {
		attrs["hs2_thrift_port"] = ia.Hs2_thrift_port
	}
	if &ia.Overrides != nil {
		attrs["overrides"] = ia.Overrides
	}
	if &ia.Execution_engine != nil {
		attrs["execution_engine"] = ia.Execution_engine
	}

	result = append(result, attrs)

	return result
}
