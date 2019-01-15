package entity

import (
	_ "fmt"
)

type NodeConfiguration struct {
	Master_instance_type          string
	Slave_instance_type           string
	Heterogeneous_instance_config HeterogeneousInstanceConfig
	Initial_nodes                 int
	Max_nodes                     int
	Slave_request_type            string
	Spot_instance_settings        SpotInstanceSettings
	Stable_spot_instance_settings StableSpotInstanceSettings
	Spot_block_settings           SpotBlockSettings
	Fallback_to_ondemand          bool
	Ebs_volume_type               string
	Ebs_volume_size               int
	Ebs_volume_count              int
	Ebs_upscaling_config          EbsUpscalingConfig
	Custom_ec2_tags               map[string]string
	Idle_cluster_timeout_in_secs  int
	Node_base_cooldown_period     int
	Node_spot_cooldown_period     int
}
