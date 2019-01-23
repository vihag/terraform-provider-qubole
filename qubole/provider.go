package qubole

/*
The *schema.Provider type describes the provider's properties including:

1. the configuration keys it accepts
2. the resources it supports
3. any callbacks to configure
*/

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		//Provider Schema Map
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["auth_token"],
			},

			"api_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["api_endpoint"],
			},
		},

		//Resources Map
		ResourcesMap: map[string]*schema.Resource{
			"qubole_cluster": resourceQuboleCluster(),
		},
		
		//Provider Configurer: https://www.terraform.io/docs/plugins/provider.html#configurefunc
		ConfigureFunc: providerConfigure,
	}
}

//Schema Descriptions
var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_token": "The auth_token for the account. You can retreive this\n" +
			"from the My Accounts section of the QDS Control Panel. Press Show in the token column",

		"api_endpoint": "The API endpoint in which your account is hosted.\n" +
			"Possible options are api.qubole.com, us.qubole.com, in.qubole.com, eu-central-1.qubole.com, wellness.qubole.com, azure.qubole.com, oraclecloud.qubole.com",

		"api_version": "The API version to use.",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AuthToken:   d.Get("auth_token").(string),
		ApiEndpoint: d.Get("api_endpoint").(string),
	}
	return config.getValidatedConfig();
}
