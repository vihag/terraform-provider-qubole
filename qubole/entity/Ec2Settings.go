package entity

import (
	"fmt"
)

type Ec2Settings struct {
	compute_access_key              string
	compute_secret_key              string
	compute_access_key              string
	aws_region                      string
	aws_preferred_availability_zone string
	vpc_id                          string
	subnet_id                       string
	master_elastic_ip               string
	bastion_node_public_dns         string
	bastion_node_port               int
	bastion_node_user               string
	role_instance_profile           string
	use_account_compute_creds       bool
}
