package model

import ()

type MinNodes struct {
	Preemptible bool    `json:"preemptible,omitempty"`
	Percentage  float32 `json:"percentage,omitempty"`
}

/*
function to flatten Min Nodes Settings
*/
func FlattenMinNodes(ia *MinNodes) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Preemptible != nil {
		attrs["preemptible"] = ia.Preemptible
	}

	if &ia.Percentage != nil {
		attrs["percentage"] = ia.Percentage
	}

	result = append(result, attrs)

	return result
}

func ReadMinNodesFromTf(minNodes *MinNodes, minNodesSettings []interface{}) bool {

	if len(minNodesSettings) > 0 {
		configs := minNodesSettings[0].(map[string]interface{})
		if v, ok := configs["preemptible"]; ok {
			minNodes.Preemptible = v.(bool)
		}
		if v, ok := configs["percentage"]; ok {
			minNodes.Percentage = float32(v.(float64))
		}
	}

	return true
}
