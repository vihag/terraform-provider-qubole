package entity

import (
	_ "fmt"
)

type EngineConfig struct {
	Dbtap_id      int
	Fernet_key    string
	Engine_type   string
	Version       string
	Overrides     bool
	Hive_settings HiveSettings
}
