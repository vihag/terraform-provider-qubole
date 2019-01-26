package model

import (
	_ "fmt"
)

type Datadog struct {
	Datadog_api_token string `json:"datadog_api_token,omitempty"`
	Datadog_app_token string `json:"datadog_app_token,omitempty"`
}

/*
function to flatten Datadog Settings
*/
func FlattenDatadog(ia *Datadog) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Datadog_api_token != nil {
		attrs["datadog_api_token"] = ia.Datadog_api_token
	}

	if &ia.Datadog_app_token != nil {
		attrs["datadog_app_token"] = ia.Datadog_app_token
	}

	result = append(result, attrs)

	return result
}

func ReadDatadogFromTf(datadog *Datadog, datadogSettings []interface{}) bool {

	if len(datadogSettings) > 0 {
		configs := datadogSettings[0].(map[string]interface{})
		if v, ok := configs["datadog_api_token"]; ok {
			datadog.Datadog_api_token = v.(string)
		}
		if v, ok := configs["datadog_app_token"]; ok {
			datadog.Datadog_app_token = v.(string)
		}
	}

	return true
}
