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

/*
function to flatten Compute Config
*/
func FlattenComputeConfig(ia *ComputeConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Compute_validated != nil {
		attrs["compute_validated"] = ia.Compute_validated
	}

	if &ia.Use_account_compute_creds != nil {
		attrs["use_account_compute_creds"] = ia.Use_account_compute_creds
	}

	if &ia.Instance_tenancy != nil {
		attrs["instance_tenancy"] = ia.Instance_tenancy
	}

	if &ia.Compute_role_arn != nil {
		attrs["compute_role_arn"] = ia.Compute_role_arn
	}

	if &ia.Compute_external_id != nil {
		attrs["compute_external_id"] = ia.Compute_external_id
	}

	if &ia.Role_instance_profile != nil {
		attrs["role_instance_profile"] = ia.Role_instance_profile
	}

	if &ia.Compute_client_id != nil {
		attrs["compute_client_id"] = ia.Compute_client_id
	}

	if &ia.Compute_client_secret != nil {
		attrs["compute_client_secret"] = ia.Compute_client_secret
	}

	if &ia.Compute_tenant_id != nil {
		attrs["compute_tenant_id"] = ia.Compute_tenant_id
	}

	if &ia.Compute_subscription_id != nil {
		attrs["compute_subscription_id"] = ia.Compute_subscription_id
	}

	result = append(result, attrs)

	return result
}

