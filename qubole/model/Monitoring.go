package model

import (
	_ "fmt"
)

type Monitoring struct {
	Ganglia bool    `json:"ganglia,omitempty"`
	Datadog Datadog `json:"datadog,omitempty"`
}

/*
function to flatten Monitoring Config
*/
func FlattenMonitoring(ia *Monitoring) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Ganglia != nil {
		attrs["ganglia"] = ia.Ganglia
	}
	if &ia.Datadog != nil {
		attrs["datadog"] = FlattenDatadog(&ia.Datadog)
	}

	result = append(result, attrs)

	return result
}
