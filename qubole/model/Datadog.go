package model

import (
	_ "fmt"
)

type Datadog struct {
	Datadog_api_token string `json:"datadog_api_token,omitempty"`
	Datadog_app_token string `json:"datadog_app_token,omitempty"`
}
