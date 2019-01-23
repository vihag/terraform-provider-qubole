package model

import (
	_ "fmt"
)

type AirflowSettings struct {
	Dbtap_id               int    `json:"dbtap_id,omitempty"`
	Fernet_key             string `json:"fernet_key,omitempty"`
	Overrides              string `json:"overrides,omitempty"`
	Version                string `json:"version,omitempty"`
	Airflow_python_version string `json:"airflow_python_version,omitempty"`
}
