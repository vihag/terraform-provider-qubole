package model

import (
	_ "fmt"
)

type ComputeConfig struct {
	Compute_validated         bool   `json:"compute_validated,omitempty"`
	Use_account_compute_creds bool   `json:"use_account_compute_creds,omitempty"`
	Instance_tenancy          string `json:"instance_tenancy,omitempty"`
	Compute_role_arn          string `json:"compute_role_arn,omitempty"`
	Compute_external_id       string `json:"compute_external_id,omitempty"`
	Role_instance_profile     string `json:"role_instance_profile,omitempty"`
	//Azure elements
	Compute_client_id       string `json:"compute_client_id,omitempty"`
	Compute_client_secret   string `json:"compute_client_secret,omitempty"`
	Compute_tenant_id       string `json:"compute_tenant_id,omitempty"`
	Compute_subscription_id string `json:"compute_subscription_id,omitempty"`
}
