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

/*
function to flatten airflow settings to the schema that is defined
*/
func FlattenAirflowSettings(ia *AirflowSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Dbtap_id != nil {
		attrs["dbtap_id"] = ia.Dbtap_id
	}
	if &ia.Fernet_key != nil {
		attrs["fernet_key"] = ia.Fernet_key
	}
	if &ia.Overrides != nil {
		attrs["overrides"] = ia.Overrides
	}
	if &ia.Version != nil {
		attrs["version"] = ia.Version
	}
	if &ia.Airflow_python_version != nil {
		attrs["airflow_python_version"] = ia.Airflow_python_version
	}

	result = append(result, attrs)

	return result
}

func ReadAirflowSettingsFromTf(airflow_settings *AirflowSettings, airflowSettings []interface{}) bool {

	if len(airflowSettings) > 0 {
		configs := airflowSettings[0].(map[string]interface{})
		if v, ok := configs["dbtap_id"]; ok {
			airflow_settings.Dbtap_id = v.(int)
		}
		if v, ok := configs["fernet_key"]; ok {
			airflow_settings.Fernet_key = v.(string)
		}
		if v, ok := configs["overrides"]; ok {
			airflow_settings.Overrides = v.(string)
		}
		if v, ok := configs["version"]; ok {
			airflow_settings.Version = v.(string)
		}
		if v, ok := configs["airflow_python_version"]; ok {
			airflow_settings.Airflow_python_version = v.(string)
		}
	}

	return true
}
