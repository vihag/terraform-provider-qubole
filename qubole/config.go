package qubole

import (
	"fmt"
)

type Config struct {
	AuthToken        string
	ApiEndpoint      string
	ConnectionString string
}

// Validate the provider config and return the same struct with formulated Connection String
func (c *Config) getValidatedConfig() (interface{}, error) {

	errEndpoint := c.ValidateEndpoint()
	if errEndpoint != nil {
		return nil, errEndpoint
	}

	fmt.Println("Using api endpoint version: %s for endpoint %s", "v1.3", c.ApiEndpoint)

	//Ok, so we have a correct version and endpoint, all is well
	c.ConnectionString = "https://" + c.ApiEndpoint + "/api/" + "v1.3" + "/clusters/"

	return c, nil

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
	case "oraclecloud.qubole.com":
		return nil
	default:
		return fmt.Errorf("Not a valid api endpoint: %s", c.ApiEndpoint)
	}
	
	return nil

}

