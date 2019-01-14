package qubole

import (
	"fmt"
)

type Config struct {
	AuthToken        string
	ApiEndpoint      string
	ApiVersion       string
	ConnectionString string
}

// Validate the provider config and return the same struct with formulated Connection String
func (c *Config) getValidatedConfig() (interface{}, error) {

	errEndpoint := c.ValidateEndpoint()
	if errEndpoint != nil {
		return nil, errEndpoint
	}

	errEndpointVersion := c.ValidateEndpointVersion()
	if errEndpointVersion != nil {
		return nil, errEndpointVersion
	}

	fmt.Println("Using api endpoint version: %s for endpoint %s", c.ApiVersion, c.ApiEndpoint)

	//Ok, so we have a correct version and endpoint, all is well
	c.connectionString = "https://" + c.ApiEndpoint + "/api/" + c.ApiVersion + "/"

	return &c, nil

}

// ValidateEndpoint returns an error if the configured endpoint is not a
// valid qubole endpoint and nil otherwise.
func (c *Config) ValidateEndpoint() error {

	//Validate API Endpoints
	switch endpoint := c.ApiEndpoint; endpoint {
	case "api.qubole.com":
	case "us.qubole.com":
	case "in.qubole.com":
	case "eu-central-1.qubole.com":
	case "wellness.qubole.com":
	case "azure.qubole.com":
	case "oraclecloud.qubole.com":
		return nil
	default:
		return fmt.Errorf("Not a valid api endpoint: %s", c.ApiEndpoint)
	}

}

// ValidateEndpointVersion returns an error if the configured endpoint version is not a
// valid qubole endpoint version and nil otherwise.
func (c *Config) ValidateEndpointVersion() error {

	//Validate API Endpoint Versions
	switch endpoint := c.ApiEndpoint; endpoint {
	case "api.qubole.com":
		if c.ApiVersion == "v1.2" || c.ApiVersion == "v1.3" || c.ApiVersion == "v2" {
			return nil
		}
	case "us.qubole.com":
		if c.ApiVersion == "v1.2" || c.ApiVersion == "v1.3" || c.ApiVersion == "v2" {
			return nil
		}
	case "in.qubole.com":
		if c.ApiVersion == "v1.2" || c.ApiVersion == "v1.3" || c.ApiVersion == "v2" {
			return nil
		}
	case "eu-central-1.qubole.com":
		if c.ApiVersion == "v1.2" || c.ApiVersion == "v1.3" || c.ApiVersion == "v2" {
			return nil
		}
	case "wellness.qubole.com":
		if c.ApiVersion == "v1.2" || c.ApiVersion == "v1.3" || c.ApiVersion == "v2" {
			return nil
		}
	case "azure.qubole.com":
		if c.ApiVersion == "v2" {
			return nil
		}
	case "oraclecloud.qubole.com":
		if c.ApiVersion == "v2" {
			return nil
		}
	default:
		return fmt.Errorf("Not a valid api endpoint: %s", c.ApiEndpoint)
	}

	return fmt.Errorf("Not a valid api endpoint version for api.qubole.com: %s", c.ApiVersion)

}
