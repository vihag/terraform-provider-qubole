package model

import (
	_ "fmt"
)

type Internal struct {
	Zeppelin_interpreter_mode string `json:"zeppelin_interpreter_mode,omitempty"`
	Spark_s3_package_name     string `json:"spark_s3_package_name,omitempty"`
	Zeppelin_s3_package_name  string `json:"zeppelin_s3_package_name,omitempty"`
}

/*
function to flatten Internal Config
*/
func FlattenInternal(ia *Internal) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Zeppelin_interpreter_mode != nil {
		attrs["zeppelin_interpreter_mode"] = ia.Zeppelin_interpreter_mode
	}
	if &ia.Spark_s3_package_name != nil {
		attrs["spark_s3_package_name"] = ia.Spark_s3_package_name
	}
	if &ia.Zeppelin_s3_package_name != nil {
		attrs["zeppelin_s3_package_name"] = ia.Zeppelin_s3_package_name
	}

	result = append(result, attrs)

	return result
}
