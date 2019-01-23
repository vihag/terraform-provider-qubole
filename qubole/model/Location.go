package model

import (
	_ "fmt"
)

type Location struct {
	Aws_region            string `json:"aws_region,omitempty"`
	Aws_availability_zone string `json:"aws_availability_zone,omitempty"`
	//Azure elements
	Location string `json:"location,omitempty"`
}
