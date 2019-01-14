package entity

import (
	"fmt"
)

type NodeConfiguration struct {
	master_instance_type          string
	slave_instance_type           string
	heterogeneous_instance_config HeterogeneousInstanceConfig
	initial_nodes                 int
	max_nodes                     int
	slave_request_type            string
	spot_instance_settings        SpotInstanceSettings
	stable_spot_instance_settings StableSpotInstanceSettings
	spot_block_settings           SpotBlockSettings
	fallback_to_ondemand          bool
	ebs_volume_type               string
	ebs_volume_size               string
	ebs_volume_count              int
	ebs_upscaling_config          EbsUpscalingConfig
	custom_ec2_tags               map[string]string
	idle_cluster_timeout_in_secs  int
	node_base_cooldown_period     int
	node_spot_cooldown_period     int
}
