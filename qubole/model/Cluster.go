package model

import (
	_ "fmt"
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
