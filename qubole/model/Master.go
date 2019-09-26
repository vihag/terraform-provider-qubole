package model

import (
	_ "fmt"
)

type Master struct {
	Preemptible bool `json:"preemptible,omitempty"`
}

/*
function to flatten Master Settings
*/
func FlattenMaster(ia *Master) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Preemptible != nil {
		attrs["preemptible"] = ia.Preemptible
	}

	result = append(result, attrs)

	return result
}

func ReadMasterFromTf(master *Master, masterSettings []interface{}) bool {

	if len(masterSettings) > 0 {
		configs := masterSettings[0].(map[string]interface{})
		if v, ok := configs["preemptible"]; ok {
			master.Preemptible = v.(bool)
		}
	}

	return true
}
