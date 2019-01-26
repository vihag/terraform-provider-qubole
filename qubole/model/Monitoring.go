package model

import (
	_ "fmt"
	"github.com/hashicorp/terraform/helper/schema"
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

/*
Read Monitoring Informtion from Terraform file
*/
func ReadMonitoringFromTf(d *schema.ResourceData) (Monitoring, bool) {

	var monitoring Monitoring
	if v, ok := d.GetOk("monitoring"); ok {
		monitoringConfig := v.([]interface{})
		if len(monitoringConfig) > 0 {
			configs := monitoringConfig[0].(map[string]interface{})

			if v, ok := configs["ganglia"]; ok {
				monitoring.Ganglia = v.(bool)
			}

			//Read datadog settings
			var datadog Datadog
			if v, ok := configs["datadog"]; ok {
				datadogSettings := v.([]interface{})
				if len(datadogSettings) > 0 {
					configs := datadogSettings[0].(map[string]interface{})
					if v, ok := configs["datadog_api_token"]; ok {
						datadog.Datadog_api_token = v.(string)
					}
					if v, ok := configs["datadog_app_token"]; ok {
						datadog.Datadog_app_token = v.(string)
					}
					monitoring.Datadog = datadog
				}
			}

			return monitoring, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return monitoring, false
}
