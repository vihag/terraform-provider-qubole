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

