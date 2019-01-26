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
	Label                        []string            `json:"label,omitempty"`
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

/*
function to Cluster Info
*/
func FlattenClusterInfo(ia *ClusterInfo) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Master_instance_type != nil {
		attrs["master_instance_type"] = ia.Master_instance_type
	}

	if &ia.Slave_instance_type != nil {
		attrs["slave_instance_type"] = ia.Slave_instance_type
	}

	if &ia.Node_base_cooldown_period != nil {
		attrs["node_base_cooldown_period"] = ia.Node_base_cooldown_period
	}

	if &ia.Label != nil {
		attrs["label"] = ia.Label
	}

	if &ia.Min_nodes != nil {
		attrs["min_nodes"] = ia.Min_nodes
	}

	if &ia.Max_nodes != nil {
		attrs["max_nodes"] = ia.Max_nodes
	}

	if &ia.Idle_cluster_timeout_in_secs != nil {
		attrs["idle_cluster_timeout_in_secs"] = ia.Idle_cluster_timeout_in_secs
	}

	if &ia.Cluster_name != nil {
		attrs["cluster_name"] = ia.Cluster_name
	}

	if &ia.Node_bootstrap != nil {
		attrs["node_bootstrap"] = ia.Node_bootstrap
	}

	if &ia.Disallow_cluster_termination != nil {
		attrs["disallow_cluster_termination"] = ia.Disallow_cluster_termination
	}

	if &ia.Force_tunnel != nil {
		attrs["force_tunnel"] = ia.Force_tunnel
	}

	if &ia.Customer_ssh_key != nil {
		attrs["customer_ssh_key"] = ia.Customer_ssh_key
	}

	if &ia.Child_hs2_cluster_id != nil {
		attrs["child_hs2_cluster_id"] = ia.Child_hs2_cluster_id
	}

	if &ia.Parent_cluster_id != nil {
		attrs["parent_cluster_id"] = ia.Parent_cluster_id
	}

	if &ia.Env_settings != nil {
		attrs["env_settings"] = FlattenEnvSettings(&ia.Env_settings)
	}

	if &ia.Datadisk != nil {
		attrs["datadisk"] = FlattenDatadisk(&ia.Datadisk)
	}

	if &ia.Heterogeneous_config != nil {
		attrs["heterogeneous_config"] = FlattenHeterogeneousConfig(&ia.Heterogeneous_config)
	}

	if &ia.Slave_request_type != nil {
		attrs["slave_request_type"] = ia.Slave_request_type
	}

	if &ia.Custom_tags != nil {
		attrs["custom_tags"] = ia.Custom_tags
	}

	if &ia.Fallback_to_ondemand != nil {
		attrs["fallback_to_ondemand"] = ia.Fallback_to_ondemand
	}

	result = append(result, attrs)

	return result
}
