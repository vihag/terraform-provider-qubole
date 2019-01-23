package model

import (
	_ "fmt"
)

type SparkSettings struct {
	Custom_spark_config string `json:"custom_spark_config,omitempty"`
	Spark_version       string `json:"spark_version,omitempty"`
	Enable_rubix        string `json:"enable_rubix,omitempty"`
}
