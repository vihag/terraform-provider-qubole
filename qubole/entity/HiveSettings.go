package entity

import (
	_ "fmt"
)

type HiveSettings struct {
	Is_hs2                     bool
	Hive_version               string
	Hive_qubole_metadata_cache bool
	Hs2_thrift_port            int
	Overrides                  bool
}
