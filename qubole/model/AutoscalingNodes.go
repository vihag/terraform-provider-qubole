package model

import ()

type AutoscalingNodes struct {
	Preemptible bool    `json:"preemptible,omitempty"`
	Percentage  float32 `json:"percentage,omitempty"`
}

/*
function to flatten Autoscaling Nodes Settings
*/
func FlattenAutoscalingNodes(ia *AutoscalingNodes) []map[string]interface{} {
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

func ReadAutoscalingNodesFromTf(autoscalingNodes *AutoscalingNodes, autoscalingNodesSettings []interface{}) bool {

	if len(autoscalingNodesSettings) > 0 {
		configs := autoscalingNodesSettings[0].(map[string]interface{})
		if v, ok := configs["preemptible"]; ok {
			autoscalingNodes.Preemptible = v.(bool)
		}
		if v, ok := configs["percentage"]; ok {
			autoscalingNodes.Percentage = float32(v.(float64))
		}
	}

	return true
}
