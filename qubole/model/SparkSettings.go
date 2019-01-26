package model

import (
	_ "fmt"
)

type SparkSettings struct {
	Custom_spark_config string `json:"custom_spark_config,omitempty"`
	Spark_version       string `json:"spark_version,omitempty"`
	Enable_rubix        bool   `json:"enable_rubix,omitempty"`
}

/*
function to flatten Spark Settings
*/
func FlattenSparkSettings(ia *SparkSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Spark_version != nil {
		attrs["spark_version"] = ia.Spark_version
	}
	if &ia.Custom_spark_config != nil {
		attrs["custom_spark_config"] = ia.Custom_spark_config
	}
	if &ia.Enable_rubix != nil {
		attrs["enable_rubix"] = ia.Enable_rubix
	}

	result = append(result, attrs)

	return result
}

func ReadSparkSettingsFromTf(spark_settings *SparkSettings, sparkSettings []interface{}) bool {

	if len(sparkSettings) > 0 {
		configs := sparkSettings[0].(map[string]interface{})
		if v, ok := configs["spark_version"]; ok {
			spark_settings.Spark_version = v.(string)
		}
		if v, ok := configs["custom_spark_config"]; ok {
			spark_settings.Custom_spark_config = v.(string)
		}
		if v, ok := configs["enable_rubix"]; ok {
			spark_settings.Enable_rubix = v.(bool)
		}
	}

	return true
}
