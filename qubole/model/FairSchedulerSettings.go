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

func ReadFairSchedulerSettingsFromTf(fairscheduler_settings *FairSchedulerSettings, fairSchedulerSettings []interface{}) bool {

	if len(fairSchedulerSettings) > 0 {
		configs := fairSchedulerSettings[0].(map[string]interface{})
		if v, ok := configs["default_pool"]; ok {
			fairscheduler_settings.Default_pool = v.(string)
		}
		if v, ok := configs["fairscheduler_config_xml"]; ok {
			fairscheduler_settings.Fairscheduler_config_xml = v.(string)
		}
	}

	return true
}
