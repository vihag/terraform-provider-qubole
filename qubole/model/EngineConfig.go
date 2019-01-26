package model

import (
	_ "fmt"
)

type EngineConfig struct {
	Flavour          string          `json:"flavour,omitempty"`
	Hadoop_settings  HadoopSettings  `json:"hadoop_settings,omitempty"`
	Presto_settings  PrestoSettings  `json:"presto_settings,omitempty"`
	Spark_settings   SparkSettings   `json:"spark_settings,omitempty"`
	Hive_settings    HiveSettings    `json:"hive_settings,omitempty"`
	Airflow_settings AirflowSettings `json:"airflow_settings,omitempty"`
}

/*
function to flatten Engine Config Settings
*/
func FlattenEngineConfig(ia *EngineConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Flavour != nil {
		attrs["flavour"] = ia.Flavour
	}
	if &ia.Hadoop_settings != nil {
		attrs["hadoop_settings"] = FlattenHadoopSettings(&ia.Hadoop_settings)
	}
	if &ia.Presto_settings != nil {
		attrs["presto_settings"] = FlattenPrestoSettings(&ia.Presto_settings)
	}
	if &ia.Spark_settings != nil {
		attrs["spark_settings"] = FlattenSparkSettings(&ia.Spark_settings)
	}
	if &ia.Hive_settings != nil {
		attrs["hive_settings"] = FlattenHiveSettings(&ia.Hive_settings)
	}
	if &ia.Airflow_settings != nil {
		attrs["airflow_settings"] = FlattenAirflowSettings(&ia.Airflow_settings)
	}

	result = append(result, attrs)

	return result
}
