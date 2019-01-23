package model

import (
	_ "fmt"
)

type EnvSettings struct {
	Python_version string `json:"python_version,omitempty"`
	R_version      string `json:"r_version,omitempty"`
	Name           string `json:"name,omitempty"`
}
