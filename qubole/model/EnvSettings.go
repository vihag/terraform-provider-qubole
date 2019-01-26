package model

import (
	_ "fmt"
)

type EnvSettings struct {
	Python_version string `json:"python_version,omitempty"`
	R_version      string `json:"r_version,omitempty"`
	Name           string `json:"name,omitempty"`
}

/*
function to flatten Env Settings
*/
func FlattenEnvSettings(ia *EnvSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Python_version != nil {
		attrs["python_version"] = ia.Python_version
	}
	if &ia.R_version != nil {
		attrs["r_version"] = ia.R_version
	}
	if &ia.Name != nil {
		attrs["name"] = ia.Name
	}

	result = append(result, attrs)

	return result
}


