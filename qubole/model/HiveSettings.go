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

func ReadHiveSettingsFromTf(hive_settings *HiveSettings, hiveSettings []interface{}) bool {

	if len(hiveSettings) > 0 {
		configs := hiveSettings[0].(map[string]interface{})
		if v, ok := configs["is_hs2"]; ok {
			hive_settings.Is_hs2 = v.(bool)
		}
		if v, ok := configs["hive_version"]; ok {
			hive_settings.Hive_version = v.(string)
		}
		if v, ok := configs["is_metadata_cache_enabled"]; ok {
			hive_settings.Is_metadata_cache_enabled = v.(bool)
		}
		if v, ok := configs["hs2_thrift_port"]; ok {
			hive_settings.Hs2_thrift_port = v.(int)
		}
		if v, ok := configs["overrides"]; ok {
			hive_settings.Overrides = v.(string)
		}
		if v, ok := configs["execution_engine"]; ok {
			hive_settings.Execution_engine = v.(string)
		}
	}

	return true
}
