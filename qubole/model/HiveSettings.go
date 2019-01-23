package model

import (
	_ "fmt"
)

type HiveSettings struct {
	Is_hs2                     bool   `json:"is_hs2,omitempty"`
	Hive_version               string `json:"hive_version,omitempty"`
	Hive_qubole_metadata_cache bool   `json:"hive_qubole_metadata_cache,omitempty"`
	Hs2_thrift_port            int    `json:"hs2_thrift_port,omitempty"`
	Overrides                  string `json:"overrides,omitempty"`
}
