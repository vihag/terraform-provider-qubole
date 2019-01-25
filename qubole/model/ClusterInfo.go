package model

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"
)

type ClusterInfo struct {
	Master_instance_type         string              `json:"master_instance_type,omitempty"`
	Slave_instance_type          string              `json:"slave_instance_type,omitempty"`
	Node_base_cooldown_period    int                 `json:"node_base_cooldown_period,omitempty"`
	Label                        []string           `json:"label,omitempty"`
	Min_nodes                    int                 `json:"min_nodes,omitempty"`
	Max_nodes                    int                 `json:"max_nodes,omitempty"`
	Idle_cluster_timeout_in_secs int                 `json:"idle_cluster_timeout_in_secs,omitempty"`
	Cluster_name                 string              `json:"cluster_name,omitempty"`
	Node_bootstrap               string              `json:"node_bootstrap,omitempty"`
	Disallow_cluster_termination bool                `json:"disallow_cluster_termination,omitempty"`
	Force_tunnel                 bool                `json:"force_tunnel,omitempty"`
	Customer_ssh_key             string              `json:"customer_ssh_key,omitempty"`
	Child_hs2_cluster_id         int                 `json:"child_hs2_cluster_id,omitempty"`
	Parent_cluster_id            int                 `json:"parent_cluster_id,omitempty"`
	Env_settings                 EnvSettings         `json:"env_settings,omitempty"`
	Datadisk                     Datadisk            `json:"datadisk,omitempty"`
	Heterogeneous_config         HeterogeneousConfig `json:"heterogeneous_config,omitempty"`
	Slave_request_type           string              `json:"slave_request_type,omitempty"`
	Spot_settings                SpotSettings        `json:"spot_settings,omitempty"`
	Custom_tags                  map[string]string   `json:"custom_tags,omitempty"`
	Fallback_to_ondemand         bool                `json:"fallback_to_ondemand,omitempty"`
}

//Custom unmarshalling logic
/*
func (u *ClusterInfo) UnmarshalJSON(data []byte) error {
	log.Printf("[ERR]using custom unmarshaller for umarshalling cluster object: %s", "ClusterInfo")
	type Alias ClusterInfo
	aux := &struct {
		Datadisk Datadisk `json:"datadisk,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	log.Printf("[INFO]Pretty Printing Unmarshalled Response for the rest of cluster info %#v", aux)
	//Now concentrate on datadisk
	var datadisk *Datadisk
	err := json.Unmarshal(GetBytes(aux.Datadisk), &datadisk)
	if err != nil {

		log.Printf("verbose error info: %#v", err)
		return fmt.Errorf("verbose error info: %#v", err)
	}
	u.Datadisk = *datadisk
	return nil
}*/
