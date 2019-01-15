package entity

import (
	_ "fmt"
)

type Ec2Settings struct {
	Compute_access_key              string
	Compute_secret_key              string
	Aws_region                      string
	Aws_preferred_availability_zone string
	Vpc_id                          string
	Subnet_id                       string
	Master_elastic_ip               string
	Bastion_node_public_dns         string
	Bastion_node_port               int
	Bastion_node_user               string
	Role_instance_profile           string
	Use_account_compute_creds       bool
}
