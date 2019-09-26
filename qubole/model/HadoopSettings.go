package model

import (
	_ "fmt"
)

type HadoopSettings struct {
	Use_qubole_placement_policy bool                  `json:"use_qubole_placement_policy,omitempty"`
	Is_ha                       bool                  `json:"is_ha,omitempty"`
	Custom_hadoop_config        string                `json:"custom_hadoop_config,omitempty"`
	Fairscheduler_settings      FairSchedulerSettings `json:"fairscheduler_settings,omitempty"`
}

/*
function to flatten Hadoop Settings
*/
func FlattenHadoopSettings(ia *HadoopSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Use_qubole_placement_policy != nil {
		attrs["use_qubole_placement_policy"] = ia.Use_qubole_placement_policy
	}
	if &ia.Is_ha != nil {
		attrs["is_ha"] = ia.Is_ha
	}
	if &ia.Custom_hadoop_config != nil {
		attrs["custom_hadoop_config"] = ia.Custom_hadoop_config
	}
	if &ia.Fairscheduler_settings != nil {
		attrs["fairscheduler_settings"] = FlattenFairSchedulerSettings(&ia.Fairscheduler_settings)
	}

	result = append(result, attrs)

	return result
}

func ReadHadoopSettingsSettingsFromTf(hadoop_settings *HadoopSettings, hadoopSettings []interface{}) bool {

	if len(hadoopSettings) > 0 {
		configs := hadoopSettings[0].(map[string]interface{})
		if v, ok := configs["use_qubole_placement_policy"]; ok {
			hadoop_settings.Use_qubole_placement_policy = v.(bool)
		}
		if v, ok := configs["is_ha"]; ok {
			hadoop_settings.Is_ha = v.(bool)
		}
		if v, ok := configs["custom_hadoop_config"]; ok {
			hadoop_settings.Custom_hadoop_config = v.(string)
		}
		//Read Fair Scheduler Settings
		var fairscheduler_settings FairSchedulerSettings
		if v, ok := configs["fairscheduler_settings"]; ok {
			fairSchedulerSettings := v.([]interface{})
			ReadFairSchedulerSettingsFromTf(&fairscheduler_settings, fairSchedulerSettings)
			/*if len(fairSchedulerSettings) > 0 {
				configs := fairSchedulerSettings[0].(map[string]interface{})
				if v, ok := configs["default_pool"]; ok {
					fairscheduler_settings.Default_pool = v.(string)
				}
				if v, ok := configs["fairscheduler_config_xml"]; ok {
					fairscheduler_settings.Fairscheduler_config_xml = v.(string)
				}
				hadoop_settings.Fairscheduler_settings = fairscheduler_settings
			}*/
			hadoop_settings.Fairscheduler_settings = fairscheduler_settings
		}
	}

	return true
}
