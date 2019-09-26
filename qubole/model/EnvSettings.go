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

func ReadEnvSettingsFromTf(env_settings *EnvSettings, envSettings []interface{}) bool {

	if len(envSettings) > 0 {
		configs := envSettings[0].(map[string]interface{})
		if v, ok := configs["python_version"]; ok {
			env_settings.Python_version = v.(string)
		}
		if v, ok := configs["r_version"]; ok {
			env_settings.R_version = v.(string)
		}
		if v, ok := configs["name"]; ok {
			env_settings.Name = v.(string)
		}
	}

	return true
}
