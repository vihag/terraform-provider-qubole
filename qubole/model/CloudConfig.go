package model

import (
	_ "fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

type CloudConfig struct {
	Provider            string        `json:"provider,omitempty"`
	Resource_group_name string        `json:"resource_group_name,omitempty"`
	Compute_config      ComputeConfig `json:"compute_config,omitempty"`
	Location            Location      `json:"location,omitempty"`
	Network_config      NetworkConfig `json:"network_config,omitempty"`
	//Azure elements
	Storage_config StorageConfig `json:"storage_config,omitempty"`
	//GCP elements
	Cluster_composition ClusterComposition `json:"cluster_composition,omitempty"`
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

	//GCP Elements
	if &ia.Cluster_composition != nil {
		attrs["cluster_composition"] = FlattenClusterComposition(&ia.Cluster_composition)
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
				ReadComputeConfigFromTf(&compute_config, computeConfig)
				cloud_config.Compute_config = compute_config
			}
			//Read location
			var location Location
			if v, ok := configs["location"]; ok {
				locationConfig := v.([]interface{})
				ReadLocationFromTf(&location, locationConfig)
				cloud_config.Location = location
			}
			//Read network config
			var network_config NetworkConfig
			if v, ok := configs["network_config"]; ok {
				networkConfig := v.([]interface{})
				ReadNetworkConfigFromTf(&network_config, networkConfig)
				cloud_config.Network_config = network_config
			}
			//Read storage config
			var storage_config StorageConfig
			if v, ok := configs["storage_config"]; ok {
				storageConfig := v.([]interface{})
				ReadStorageConfigFromTf(&storage_config, storageConfig)
				cloud_config.Storage_config = storage_config
			}
			//GCP Elements
			//Read cluster composition
			var cluster_composition ClusterComposition
			if v, ok := configs["cluster_composition"]; ok {
				clusterComposition := v.([]interface{})
				ReadClusterCompositionFromTf(&cluster_composition, clusterComposition)
				cloud_config.Cluster_composition = cluster_composition
			}
			return cloud_config, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return cloud_config, false
}
