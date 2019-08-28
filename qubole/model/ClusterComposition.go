package model

import (
	_ "fmt"
	_ "github.com/hashicorp/terraform/helper/schema"
	"log"
)

type ClusterComposition struct {
	Master            Master           `json:"master,omitempty"`
	Min_nodes         MinNodes         `json:"min_nodes,omitempty"`
	Autoscaling_nodes AutoscalingNodes `json:"autoscaling_nodes,omitempty"`
}

/*
function to flatten Cluster Composition
*/
func FlattenClusterComposition(ia *ClusterComposition) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Master != nil {
		attrs["master"] = FlattenMaster(&ia.Master)
	}

	if &ia.Min_nodes != nil {
		attrs["min_nodes"] = FlattenMinNodes(&ia.Min_nodes)
	}

	if &ia.Autoscaling_nodes != nil {
		attrs["autoscaling_nodes"] = FlattenAutoscalingNodes(&ia.Autoscaling_nodes)
	}

	result = append(result, attrs)

	log.Print(result)
	return result
}

/*
Read Cluster Composition from terraform file
*/
func ReadClusterCompositionFromTf(cluster_composition *ClusterComposition, clusterComposition []interface{}) bool {

	if len(clusterComposition) > 0 {
		configs := clusterComposition[0].(map[string]interface{})

		//Read master config
		var master Master
		if v, ok := configs["master"]; ok {
			masterConfig := v.([]interface{})
			ReadMasterFromTf(&master, masterConfig)
			cluster_composition.Master = master
		}

		//Read min nodes config
		var min_nodes MinNodes
		if v, ok := configs["min_nodes"]; ok {
			minNodesConfig := v.([]interface{})
			ReadMinNodesFromTf(&min_nodes, minNodesConfig)
			cluster_composition.Min_nodes = min_nodes
		}
		
		//Read autoscaling nodes config
		var autoscaling_nodes AutoscalingNodes
		if v, ok := configs["autoscaling_nodes"]; ok {
			autoscalingNodesConfig := v.([]interface{})
			ReadAutoscalingNodesFromTf(&autoscaling_nodes, autoscalingNodesConfig)
			cluster_composition.Autoscaling_nodes = autoscaling_nodes
		}
	}

	return true
}
