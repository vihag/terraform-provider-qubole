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
	//GCP elements
	Customer_project_id string `json:"customer_project_id,omitempty"`
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

	if &ia.Customer_project_id != nil {
		attrs["customer_project_id"] = ia.Customer_project_id
	}

	result = append(result, attrs)

	return result
}

func ReadComputeConfigFromTf(compute_config *ComputeConfig, computeConfig []interface{}) bool {
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
	}
	return true
}
