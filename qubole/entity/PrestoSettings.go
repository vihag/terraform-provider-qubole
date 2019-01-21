package entity

import (
	_ "fmt"
)

type PrestoSettings struct {
	Enable_presto bool			`json:"enable_presto,omitempty"`
	Custom_config string		`json:"custom_config,omitempty"`
}
