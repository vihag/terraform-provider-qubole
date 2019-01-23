package model

import (
	_ "fmt"
)

type PrestoSettings struct {
	Presto_version       string `json:"presto_version,omitempty"`
	Custom_presto_config string `json:"custom_presto_config,omitempty"`
	Enable_rubix         bool   `json:"enable_rubix,omitempty"`
}
