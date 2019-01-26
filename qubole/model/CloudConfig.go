package model

import (
	_ "fmt"
	"log"
	"github.com/hashicorp/terraform/helper/schema"
)

type CloudConfig struct {
	Provider            string        `json:"provider,omitempty"`
	Resource_group_name string        `json:"resource_group_name,omitempty"`
	Compute_config      ComputeConfig `json:"compute_config,omitempty"`
	Location            Location      `json:"location,omitempty"`
	Network_config      NetworkConfig `json:"network_config,omitempty"`
	//Azure elements
	Storage_config StorageConfig `json:"storage_config,omitempty"`
}

/*
function to flatten Cloud Config
*/
func FlattenCloudConfig(ia *CloudConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Provider != nil {
		attrs["provider"] = ia.Provider
	}
	
	if &ia.Resource_group_name != nil {
		attrs["resource_group_name"] = ia.Resource_group_name
	}

	if &ia.Compute_config != nil {
		attrs["compute_config"] = FlattenComputeConfig(&ia.Compute_config)
	}

	if &ia.Location != nil {
		attrs["location"] = FlattenLocation(&ia.Location)
	}

	if &ia.Network_config != nil {
		attrs["network_config"] = FlattenNetworkConfig(&ia.Network_config)
	}

	if &ia.Storage_config != nil {
		attrs["storage_config"] = FlattenStorageConfig(&ia.Storage_config)
	}

	result = append(result, attrs)

	log.Print(result)
	return result
}

/*
Read Cloud Config from terraform file
*/
func ReadCloudConfigFromTf(d *schema.ResourceData) (CloudConfig, bool) {

	var cloud_config CloudConfig
	if v, ok := d.GetOk("cloud_config"); ok {
		cloudConfig := v.([]interface{})
		if len(cloudConfig) > 0 {
			configs := cloudConfig[0].(map[string]interface{})
			//Read cloud provider aws/azure/gcp
			if v, ok := configs["provider"]; ok {
				cloud_config.Provider = v.(string)
			}
			
			//Read cloud provider resource group for azure
			if v, ok := configs["resource_group_name"]; ok {
				cloud_config.Resource_group_name = v.(string)
			}
			//Read compute config
			var compute_config ComputeConfig
			if v, ok := configs["compute_config"]; ok {
				computeConfig := v.([]interface{})
				if len(computeConfig) > 0 {
					configs := computeConfig[0].(map[string]interface{})
					if v, ok := configs["compute_validated"]; ok {
						compute_config.Compute_validated = v.(bool)
					}
					if v, ok := configs["use_account_compute_creds"]; ok {
						compute_config.Use_account_compute_creds = v.(bool)
					}
					if v, ok := configs["instance_tenancy"]; ok {
						compute_config.Instance_tenancy = v.(string)
					}
					if v, ok := configs["role_instance_profile"]; ok {
						compute_config.Role_instance_profile = v.(string)
					}
					if v, ok := configs["compute_role_arn"]; ok {
						compute_config.Compute_role_arn = v.(string)
					}
					if v, ok := configs["compute_external_id"]; ok {
						compute_config.Compute_external_id = v.(string)
					}
					if v, ok := configs["compute_client_id"]; ok {
						compute_config.Compute_client_id = v.(string)
					}
					if v, ok := configs["compute_client_secret"]; ok {
						compute_config.Compute_client_secret = v.(string)
					}
					if v, ok := configs["compute_tenant_id"]; ok {
						compute_config.Compute_tenant_id = v.(string)
					}
					if v, ok := configs["compute_subscription_id"]; ok {
						compute_config.Compute_subscription_id = v.(string)
					}
					cloud_config.Compute_config = compute_config
				}
			}
			//Read location
			var location Location
			if v, ok := configs["location"]; ok {
				locationConfig := v.([]interface{})
				if len(locationConfig) > 0 {
					configs := locationConfig[0].(map[string]interface{})
					if v, ok := configs["aws_region"]; ok {
						location.Aws_region = v.(string)
					}
					if v, ok := configs["aws_availability_zone"]; ok {
						location.Aws_availability_zone = v.(string)
					}
					if v, ok := configs["location"]; ok {
						location.Location = v.(string)
					}
					cloud_config.Location = location
				}
			}
			//Read network config
			var network_config NetworkConfig
			if v, ok := configs["network_config"]; ok {
				networkConfig := v.([]interface{})
				if len(networkConfig) > 0 {
					configs := networkConfig[0].(map[string]interface{})
					if v, ok := configs["vpc_id"]; ok {
						network_config.Vpc_id = v.(string)
					}
					if v, ok := configs["subnet_id"]; ok {
						network_config.Subnet_id = v.(string)
					}
					if v, ok := configs["bastion_node_public_dns"]; ok {
						network_config.Bastion_node_public_dns = v.(string)
					}
					if v, ok := configs["bastion_node_port"]; ok {
						network_config.Bastion_node_port = v.(int)
					}
					if v, ok := configs["bastion_node_user"]; ok {
						network_config.Bastion_node_user = v.(string)
					}
					if v, ok := configs["master_elastic_ip"]; ok {
						network_config.Master_elastic_ip = v.(string)
					}
					if v, ok := configs["persistent_security_groups"]; ok {
						network_config.Persistent_security_groups = v.(string)
					}
					if v, ok := configs["persistent_security_group_resource_group_name"]; ok {
						network_config.Persistent_security_group_resource_group_name = v.(string)
					}
					if v, ok := configs["persistent_security_group_name"]; ok {
						network_config.Persistent_security_group_name = v.(string)
					}
					if v, ok := configs["vnet_name"]; ok {
						network_config.Vnet_name = v.(string)
					}
					if v, ok := configs["subnet_name"]; ok {
						network_config.Subnet_name = v.(string)
					}
					if v, ok := configs["vnet_resource_group_name"]; ok {
						network_config.Vnet_resource_group_name = v.(string)
					}
					if v, ok := configs["master_static_nic_name"]; ok {
						network_config.Master_static_nic_name = v.(string)
					}
					if v, ok := configs["master_static_public_ip_name"]; ok {
						network_config.Master_static_public_ip_name = v.(string)
					}
					cloud_config.Network_config = network_config
				}
			}
			//Read storage config
			var storage_config StorageConfig
			if v, ok := configs["storage_config"]; ok {
				storageConfig := v.([]interface{})
				if len(storageConfig) > 0 {
					configs := storageConfig[0].(map[string]interface{})
					if v, ok := configs["storage_access_key"]; ok {
						storage_config.Storage_access_key = v.(string)
					}
					if v, ok := configs["storage_account_name"]; ok {
						storage_config.Storage_account_name = v.(string)
					}
					if v, ok := configs["disk_storage_account_name"]; ok {
						storage_config.Disk_storage_account_name = v.(string)
					}
					if v, ok := configs["disk_storage_account_resource_group_name"]; ok {
						storage_config.Disk_storage_account_resource_group_name = v.(string)
					}
					if v, ok := configs["managed_disk_account_type"]; ok {
						storage_config.Managed_disk_account_type = v.(string)
					}
					if v, ok := configs["data_disk_count"]; ok {
						storage_config.Data_disk_count = v.(int)
					}
					if v, ok := configs["data_disk_size"]; ok {
						storage_config.Data_disk_size = v.(int)
					}
					//Read disk upscaling config
					var disk_upscaling_config DiskUpscalingConfig
					if v, ok := configs["disk_upscaling_config"]; ok {
						diskUpscalingConfig := v.([]interface{})
						if len(diskUpscalingConfig) > 0 {
							configs := diskUpscalingConfig[0].(map[string]interface{})
							if v, ok := configs["percent_free_space_threshold"]; ok {
								disk_upscaling_config.Percent_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["max_data_disk_count"]; ok {
								disk_upscaling_config.Max_data_disk_count = v.(int)
							}
							if v, ok := configs["absolute_free_space_threshold"]; ok {
								disk_upscaling_config.Absolute_free_space_threshold = float32(v.(int))
							}
							storage_config.Disk_upscaling_config = disk_upscaling_config
						}
					}
					cloud_config.Storage_config = storage_config
				}
			}
			return cloud_config, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return cloud_config, false
}
