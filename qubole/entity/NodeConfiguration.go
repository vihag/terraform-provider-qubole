package entity

import (
	_ "fmt"
)

type NodeConfiguration struct {
	Master_instance_type          string                      `json:"master_instance_type,omitempty"`
	Slave_instance_type           string                      `json:"slave_instance_type,omitempty"`
	Heterogeneous_instance_config HeterogeneousInstanceConfig `json:"heterogeneous_instance_config,omitempty"`
	Initial_nodes                 int                         `json:"initial_nodes,omitempty"`
	Max_nodes                     int                         `json:"max_nodes,omitempty"`
	Slave_request_type            string                      `json:"slave_request_type,omitempty"`
	Spot_instance_settings        SpotInstanceSettings        `json:"spot_instance_settings,omitempty"`
	Stable_spot_instance_settings StableSpotInstanceSettings  `json:"stable_spot_instance_settings,omitempty"`
	Spot_block_settings           SpotBlockSettings           `json:"spot_block_settings,omitempty"`
	Fallback_to_ondemand          bool                        `json:"fallback_to_ondemand,omitempty"`
	Ebs_volume_type               string                      `json:"ebs_volume_type,omitempty"`
	Ebs_volume_size               int                         `json:"ebs_volume_size,omitempty"`
	Ebs_volume_count              int                         `json:"ebs_volume_count,omitempty"`
	Ebs_upscaling_config          EbsUpscalingConfig          `json:"ebs_upscaling_config,omitempty"`
	Custom_ec2_tags               map[string]string           `json:"custom_ec2_tags,omitempty"`
	Idle_cluster_timeout_in_secs  int                         `json:"idle_cluster_timeout_in_secs,omitempty"`
	Node_base_cooldown_period     int                         `json:"node_base_cooldown_period,omitempty"`
	Node_spot_cooldown_period     int                         `json:"node_spot_cooldown_period,omitempty"`
	Child_hs2_cluster_id          int                         `json:"child_hs2_cluster_id,omitempty"`
	Parent_cluster_id             int                         `json:"parent_cluster_id,omitempty"`
	Cluster_name                  string                      `json:"cluster_name,omitempty"`
}
