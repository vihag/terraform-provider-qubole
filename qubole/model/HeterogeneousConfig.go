package model

import (
	_ "fmt"
	"log"
)

type HeterogeneousConfig struct {
	Memory []map[string]interface{} `json:"memory,omitempty"`
}

/*
function to flatten Heterogeneous Configuration
*/
func FlattenHeterogeneousConfig(ia *HeterogeneousConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Memory != nil {
		attrs["memory"] = ia.Memory
	}

	result = append(result, attrs)

	return result
}

func ReadHeterogeneousConfigFromTf(heterogeneous_config *HeterogeneousConfig, heterogeneousConfigs []interface{}) bool {

	if len(heterogeneousConfigs) > 0 {
		configs := heterogeneousConfigs[0].(map[string]interface{})
		//level 1 of hetro config is the memory sub-object . there will be only one of these

		if v, ok := configs["memory"]; ok {
			inst_type_array := v.([]interface{})
			if len(inst_type_array) > 0 {
				insts := make([]map[string]interface{}, len(inst_type_array))
				for i, ins := range inst_type_array {
					datamap := ins.(map[string]interface{})
					log.Printf("[DEBUG] PRINTING DATAMAP %s", datamap)

					insts[i] = datamap

				}

				heterogeneous_config.Memory = insts
			}
		}
	}

	return true
}
