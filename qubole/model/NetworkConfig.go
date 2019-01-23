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
}
