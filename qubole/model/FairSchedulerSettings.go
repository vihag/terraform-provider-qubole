package model

import (
	_ "fmt"
)

type FairSchedulerSettings struct {
	Default_pool             string `json:",default_pool,omitempty"`
	Fairscheduler_config_xml string `json:",fairscheduler_config_xml,omitempty"`
}

/*
function to flatten fair scheduler Settings
*/
func FlattenFairSchedulerSettings(ia *FairSchedulerSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Default_pool != nil {
		attrs["default_pool"] = ia.Default_pool
	}
	if &ia.Fairscheduler_config_xml != nil {
		attrs["fairscheduler_config_xml"] = ia.Fairscheduler_config_xml
	}

	result = append(result, attrs)

	return result
}
