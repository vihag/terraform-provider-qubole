package model

import (
	_ "bytes"
	_ "encoding/gob"
	_ "encoding/json"
	_ "fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	_ "strconv"
)

type Cluster struct {
	Id            int          `json:"id,omitempty"`
	State         string       `json:"state,omitempty"`
	Cloud_config  CloudConfig  `json:"cloud_config,omitempty"`
	Cluster_info  ClusterInfo  `json:"cluster_info,omitempty"`
	Engine_config EngineConfig `json:"engine_config,omitempty"`
	Monitoring    Monitoring   `json:"monitoring,omitempty"`
	Internal      Internal     `json:"internal,omitempty"`
}

/*
Read cluster information from terraform file
*/
func ReadClusterFromTf(d *schema.ResourceData) (Cluster, bool) {

	//Create the representative json object here
	var cluster Cluster

	//create nested datas structures
	//1. Cloud Config
	if cloud_config, ok := ReadCloudConfigFromTf(d); ok {
		cluster.Cloud_config = cloud_config
	} else {
		log.Printf("[WARN] No cloud_config seen.")
	}

	//2. Cluster Info
	if cluster_info, ok := ReadClusterInfoFromTf(d); ok {
		cluster.Cluster_info = cluster_info
	} else {
		log.Printf("[WARN] No cluster_info seen.")
	}

	//3. Engine Config
	if engine_config, ok := ReadEngineConfigFromTf(d); ok {
		cluster.Engine_config = engine_config
	} else {
		log.Printf("[WARN] No engine_config seen.")
	}

	//4. Monitoring
	if monitoring, ok := ReadMonitoringFromTf(d); ok {
		cluster.Monitoring = monitoring
	} else {
		log.Printf("[WARN] No monitoring seen.")
	}

	//5. Internal
	if internal, ok := ReadInternalFromTf(d); ok {
		cluster.Internal = internal
	} else {
		log.Printf("[WARN] No internal seen.")
	}

	//Finally, the cluster
	return cluster, true

}
