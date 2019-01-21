package entity

import (
	_ "fmt"
)

type EngineConfig struct {
	Dbtap_id      int				`json:"dbtap_id,omitempty"`
	Fernet_key    string			`json:"fernet_key,omitempty"`
	Engine_type   string			`json:"engine_type,omitempty"`
	Version       string			`json:"version,omitempty"`
	Overrides     bool				`json:"overrides,omitempty"`
	Hive_settings HiveSettings		`json:"hive_settings,omitempty"`
}
