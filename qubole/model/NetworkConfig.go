package model

import (
	_ "fmt"
)

type NetworkConfig struct {
	Vpc_id                     string `json:"vpc_id,omitempty"`
	Subnet_id                  string `json:"subnet_id,omitempty"`
	Bastion_node_public_dns    string `json:"bastion_node_public_dns,omitempty"`
	Bastion_node_port          int    `json:"bastion_node_port,omitempty"`
	Bastion_node_user          string `json:"bastion_node_user,omitempty"`
	Master_elastic_ip          string `json:"master_elastic_ip,omitempty"`
	Persistent_security_groups string `json:"persistent_security_groups,omitempty"`
	//Azure elements
	Persistent_security_group_resource_group_name string `json:"persistent_security_group_resource_group_name,omitempty"`
	Persistent_security_group_name                string `json:"persistent_security_group_name,omitempty"`
	Vnet_name                                     string `json:"vnet_name,omitempty"`
	Subnet_name                                   string `json:"subnet_name,omitempty"`
	Vnet_resource_group_name                      string `json:"vnet_resource_group_name,omitempty"`
	Master_static_nic_name                        string `json:"master_static_nic_name,omitempty"`
	Master_static_public_ip_name                  string `json:"master_static_public_ip_name,omitempty"`
	//GCP Elements
	Network          string `json:"network,omitempty"`
	Subnet           string `json:"subnet,omitempty"`
	Master_static_ip string `json:"master_static_ip,omitempty"`
}

/*
function to flatten Network Config
*/
func FlattenNetworkConfig(ia *NetworkConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Vpc_id != nil {
		attrs["vpc_id"] = ia.Vpc_id
	}
	if &ia.Subnet_id != nil {
		attrs["subnet_id"] = ia.Subnet_id
	}
	if &ia.Bastion_node_public_dns != nil {
		attrs["bastion_node_public_dns"] = ia.Bastion_node_public_dns
	}
	if &ia.Bastion_node_port != nil {
		attrs["bastion_node_port"] = ia.Bastion_node_port
	}
	if &ia.Bastion_node_user != nil {
		attrs["bastion_node_user"] = ia.Bastion_node_user
	}
	if &ia.Master_elastic_ip != nil {
		attrs["master_elastic_ip"] = ia.Master_elastic_ip
	}
	if &ia.Persistent_security_groups != nil {
		attrs["persistent_security_groups"] = ia.Persistent_security_groups
	}
	if &ia.Persistent_security_group_resource_group_name != nil {
		attrs["persistent_security_group_resource_group_name"] = ia.Persistent_security_group_resource_group_name
	}
	if &ia.Persistent_security_group_name != nil {
		attrs["persistent_security_group_name"] = ia.Persistent_security_group_name
	}
	if &ia.Vnet_name != nil {
		attrs["vnet_name"] = ia.Vnet_name
	}
	if &ia.Subnet_name != nil {
		attrs["subnet_name"] = ia.Subnet_name
	}
	if &ia.Vnet_resource_group_name != nil {
		attrs["vnet_resource_group_name"] = ia.Vnet_resource_group_name
	}
	if &ia.Master_static_nic_name != nil {
		attrs["master_static_nic_name"] = ia.Master_static_nic_name
	}
	if &ia.Master_static_public_ip_name != nil {
		attrs["master_static_public_ip_name"] = ia.Master_static_public_ip_name
	}

	//GCP Elements
	if &ia.Network != nil {
		attrs["network"] = ia.Network
	}
	if &ia.Subnet != nil {
		attrs["subnet"] = ia.Subnet
	}
	if &ia.Master_static_ip != nil {
		attrs["master_static_ip"] = ia.Master_static_ip
	}

	result = append(result, attrs)

	return result
}

func ReadNetworkConfigFromTf(network_config *NetworkConfig, networkConfig []interface{}) bool {

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

		//GCP Elements
		if v, ok := configs["network"]; ok {
			network_config.Network = v.(string)
		}
		if v, ok := configs["Subnet"]; ok {
			network_config.Subnet = v.(string)
		}
		if v, ok := configs["master_static_ip"]; ok {
			network_config.Master_static_ip = v.(string)
		}

	}

	return true
}
