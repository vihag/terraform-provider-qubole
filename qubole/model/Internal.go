package model

import (
	_ "fmt"
	"github.com/hashicorp/terraform/helper/schema"
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

/*
Read Internal Information from terraform file
*/
func ReadInternalFromTf(d *schema.ResourceData) (Internal, bool) {

	var internal Internal
	if v, ok := d.GetOk("internal"); ok {
		internalConfig := v.([]interface{})
		if len(internalConfig) > 0 {
			configs := internalConfig[0].(map[string]interface{})

			if v, ok := configs["zeppelin_interpreter_mode"]; ok {
				internal.Zeppelin_interpreter_mode = v.(string)
			}
			if v, ok := configs["spark_s3_package_name"]; ok {
				internal.Spark_s3_package_name = v.(string)
			}
			if v, ok := configs["zeppelin_s3_package_name"]; ok {
				internal.Zeppelin_s3_package_name = v.(string)
			}

			return internal, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return internal, false
}
