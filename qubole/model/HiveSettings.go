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
