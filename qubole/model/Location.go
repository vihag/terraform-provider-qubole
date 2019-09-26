package model

import (
	_ "fmt"
)

type Location struct {
	Aws_region            string `json:"aws_region,omitempty"`
	Aws_availability_zone string `json:"aws_availability_zone,omitempty"`
	//Azure elements
	Location string `json:"location,omitempty"`
	//GCP elements
	Region string `json:"region,omitempty"`
	Zone   string `json:"zone,omitempty"`
}

/*
function to flatten Location Config
*/
func FlattenLocation(ia *Location) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Aws_region != nil {
		attrs["aws_region"] = ia.Aws_region
	}
	if &ia.Aws_availability_zone != nil {
		attrs["aws_availability_zone"] = ia.Aws_availability_zone
	}
	if &ia.Location != nil {
		attrs["location"] = ia.Location
	}

	//GCP Elements
	if &ia.Region != nil {
		attrs["region"] = ia.Region
	}
	if &ia.Zone != nil {
		attrs["zone"] = ia.Zone
	}

	result = append(result, attrs)

	return result
}

func ReadLocationFromTf(location *Location, locationConfig []interface{}) bool {

	if len(locationConfig) > 0 {
		configs := locationConfig[0].(map[string]interface{})
		if v, ok := configs["aws_region"]; ok {
			location.Aws_region = v.(string)
		}
		if v, ok := configs["aws_availability_zone"]; ok {
			location.Aws_availability_zone = v.(string)
		}
		if v, ok := configs["location"]; ok {
			location.Location = v.(string)
		}

		//GCP Elements
		if v, ok := configs["region"]; ok {
			location.Region = v.(string)
		}
		if v, ok := configs["zone"]; ok {
			location.Zone = v.(string)
		}

	}

	return true
}
