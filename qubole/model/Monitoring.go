package model

import (
	_ "fmt"
)

type Monitoring struct {
	Ganglia bool    `json:"ganglia,omitempty"`
	Datadog Datadog `json:"datadog,omitempty"`
}
