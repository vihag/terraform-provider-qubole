package model

import (
	_ "fmt"
)

type Internal struct {
	Zeppelin_interpreter_mode string `json:"zeppelin_interpreter_mode,omitempty"`
	Spark_s3_package_name     string `json:"spark_s3_package_name,omitempty"`
	Zeppelin_s3_package_name  string `json:"zeppelin_s3_package_name,omitempty"`
}
