package model

import (
	_ "fmt"
)

type Ec2Settings struct {
	Compute_access_key              string		`json:"compute_access_key,omitempty"`
	Compute_secret_key              string		`json:"compute_secret_key,omitempty"`
	Aws_region                      string		`json:"aws_region,omitempty"`
	Aws_preferred_availability_zone string		`json:"aws_preferred_availability_zone,omitempty"`
	Vpc_id                          string		`json:"vpc_id,omitempty"`
	Subnet_id                       string		`json:"subnet_id,omitempty"`
	Master_elastic_ip               string		`json:"master_elastic_ip,omitempty"`
	Bastion_node_public_dns         string		`json:"bastion_node_public_dns,omitempty"`
	Bastion_node_port               int			`json:"bastion_node_port,omitempty"`
	Bastion_node_user               string		`json:"bastion_node_user,omitempty"`
	Role_instance_profile           string		`json:"role_instance_profile,omitempty"`
	Use_account_compute_creds       bool		`json:"use_account_compute_creds,omitempty"`
	Compute_validated				bool		`json:"compute_validated,omitempty"`
	Instance_tenancy				string		`json:"instance_tenancy,omitempty"`
	Compute_role_arn				string		`json:"compute_role_arn,omitempty"`
	Compute_external_id				string		`json:"compute_external_id,omitempty"`
}
