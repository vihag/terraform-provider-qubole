package model

import (
	_ "encoding/json"
	"fmt"
	"log"
	"github.com/hashicorp/terraform/helper/schema"
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

/*
Read Cluster Info From Terraform File
*/
func ReadClusterInfoFromTf(d *schema.ResourceData) (ClusterInfo, bool) {

	var cluster_info ClusterInfo
	if v, ok := d.GetOk("cluster_info"); ok {
		clusterInfo := v.([]interface{})
		if len(clusterInfo) > 0 {
			configs := clusterInfo[0].(map[string]interface{})

			if v, ok := configs["master_instance_type"]; ok {
				cluster_info.Master_instance_type = v.(string)
			}

			if v, ok := configs["slave_instance_type"]; ok {
				cluster_info.Slave_instance_type = v.(string)
			}

			if v, ok := configs["node_base_cooldown_period"]; ok {
				cluster_info.Node_base_cooldown_period = v.(int)
			}

			if v, ok := configs["label"]; ok {
				labelSet := v.(*schema.Set)

				labels := make([]string, labelSet.Len())
				for i, label := range labelSet.List() {
					labels[i] = label.(string)
				}

				cluster_info.Label = labels
			}

			if v, ok := configs["min_nodes"]; ok {
				cluster_info.Min_nodes = v.(int)
			}

			if v, ok := configs["max_nodes"]; ok {
				cluster_info.Max_nodes = v.(int)
			}

			if v, ok := configs["idle_cluster_timeout_in_secs"]; ok {
				cluster_info.Idle_cluster_timeout_in_secs = v.(int)
			}

			if v, ok := configs["cluster_name"]; ok {
				cluster_info.Cluster_name = v.(string)
			}

			if v, ok := configs["node_bootstrap"]; ok {
				cluster_info.Node_bootstrap = v.(string)
			}

			if v, ok := configs["disallow_cluster_termination"]; ok {
				cluster_info.Disallow_cluster_termination = v.(bool)
			}

			if v, ok := configs["force_tunnel"]; ok {
				cluster_info.Force_tunnel = v.(bool)
			}

			if v, ok := configs["customer_ssh_key"]; ok {
				cluster_info.Customer_ssh_key = v.(string)
			}

			if v, ok := configs["child_hs2_cluster_id"]; ok {
				cluster_info.Child_hs2_cluster_id = v.(int)
			}

			if v, ok := configs["parent_cluster_id"]; ok {
				cluster_info.Parent_cluster_id = v.(int)
			}
			//Read Env Settings
			var env_settings EnvSettings
			if v, ok := configs["env_settings"]; ok {
				envSettings := v.([]interface{})
				if len(envSettings) > 0 {
					configs := envSettings[0].(map[string]interface{})
					if v, ok := configs["python_version"]; ok {
						env_settings.Python_version = v.(string)
					}
					if v, ok := configs["r_version"]; ok {
						env_settings.R_version = v.(string)
					}
					if v, ok := configs["name"]; ok {
						env_settings.Name = v.(string)
					}
					cluster_info.Env_settings = env_settings
				}
			}
			//Read datadisk
			var datadisk Datadisk
			if v, ok := configs["datadisk"]; ok {
				datadiskConfig := v.([]interface{})
				if len(datadiskConfig) > 0 {
					configs := datadiskConfig[0].(map[string]interface{})
					if v, ok := configs["count"]; ok {
						datadisk.Count = v.(int)
					}
					if v, ok := configs["type"]; ok {
						datadisk.Disktype = v.(string)
					}
					if v, ok := configs["encryption"]; ok {
						datadisk.Encryption = v.(bool)
					}
					if v, ok := configs["size"]; ok {
						datadisk.Size = v.(int)
					}
					//Read disk upscaling config
					var upscaling_config UpscalingConfig
					if v, ok := configs["upscaling_config"]; ok {
						ebsUpscalingConfigs := v.([]interface{})
						if len(ebsUpscalingConfigs) > 0 {
							configs := ebsUpscalingConfigs[0].(map[string]interface{})
							if v, ok := configs["max_ebs_volume_count"]; ok {
								upscaling_config.Max_ebs_volume_count = v.(int)
							}
							if v, ok := configs["percent_free_space_threshold"]; ok {
								upscaling_config.Percent_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["absolute_free_space_threshold"]; ok {
								upscaling_config.Absolute_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["sampling_interval"]; ok {
								upscaling_config.Sampling_interval = v.(int)
							}
							if v, ok := configs["sampling_window"]; ok {
								upscaling_config.Sampling_window = v.(int)
							}
							datadisk.Upscaling_config = upscaling_config
						}
					}
					cluster_info.Datadisk = datadisk
				}
			}
			//Read Heterogeneous Config
			var heterogeneous_config HeterogeneousConfig
			if v, ok := configs["heterogeneous_config"]; ok {
				heterogeneousConfigs := v.([]interface{})
				if len(heterogeneousConfigs) > 0 {
					configs := heterogeneousConfigs[0].(map[string]interface{})
					//level 1 of hetro config is the memory sub-object . there will be only one of these

					if v, ok := configs["memory"]; ok {
						inst_type_array := v.([]interface{})
						if len(inst_type_array) > 0 {
							insts := make([]map[string]interface{}, len(inst_type_array))
							for i, ins := range inst_type_array {
								datamap := ins.(map[string]interface{})
								log.Printf("[DEBUG] PRINTING DATAMAP %s", datamap)

								insts[i] = datamap

							}

							heterogeneous_config.Memory = insts
						}
					}
					cluster_info.Heterogeneous_config = heterogeneous_config
				}
			}

			if v, ok := configs["slave_request_type"]; ok {
				cluster_info.Slave_request_type = v.(string)
			}

			//Read spot settings
			var spot_settings SpotSettings
			if v, ok := configs["spot_settings"]; ok {
				spotSettings := v.([]interface{})
				if len(spotSettings) > 0 {
					configs := spotSettings[0].(map[string]interface{})

					//Read spot instance settings
					var spot_instance_settings SpotInstanceSettings
					if v, ok := configs["spot_instance_settings"]; ok {
						spotInstanceSettings := v.([]interface{})
						if len(spotInstanceSettings) > 0 {
							configs := spotInstanceSettings[0].(map[string]interface{})
							if v, ok := configs["maximum_bid_price_percentage"]; ok {
								spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
							}
							if v, ok := configs["timeout_for_request"]; ok {
								spot_instance_settings.Timeout_for_request = v.(int)
							}
							if v, ok := configs["maximum_spot_instance_percentage"]; ok {
								spot_instance_settings.Maximum_spot_instance_percentage = float32(v.(int))
							}
							spot_settings.Spot_instance_settings = spot_instance_settings
						}
					}

					//Read stable spot instance settings
					var stable_spot_instance_settings StableSpotInstanceSettings
					if v, ok := configs["stable_spot_instance_settings"]; ok {
						stableSpotInstanceSettings := v.([]interface{})
						if len(stableSpotInstanceSettings) > 0 {
							configs := stableSpotInstanceSettings[0].(map[string]interface{})
							if v, ok := configs["maximum_bid_price_percentage"]; ok {
								stable_spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
							}
							if v, ok := configs["timeout_for_request"]; ok {
								stable_spot_instance_settings.Timeout_for_request = v.(int)
							}
							spot_settings.Stable_spot_instance_settings = stable_spot_instance_settings
						}
					}

					//Read spot block settings
					var spot_block_settings SpotBlockSettings
					if v, ok := configs["spot_block_settings"]; ok {
						spotBlockSettings := v.([]interface{})
						if len(spotBlockSettings) > 0 {
							configs := spotBlockSettings[0].(map[string]interface{})
							if v, ok := configs["duration"]; ok {
								spot_block_settings.Duration = v.(int)
							}
							spot_settings.Spot_block_settings = spot_block_settings
						}
					}

					cluster_info.Spot_settings = spot_settings
				}
			}

			if v, ok := configs["custom_tags"]; ok {
				billing_tags := v.(map[string]interface{})
				custom_tags := make(map[string]string)
				for key, value := range billing_tags {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)

					custom_tags[strKey] = strValue
				}
				cluster_info.Custom_tags = custom_tags
			}

			if v, ok := configs["fallback_to_ondemand"]; ok {
				cluster_info.Fallback_to_ondemand = v.(bool)
			}

			return cluster_info, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return cluster_info, false

}
