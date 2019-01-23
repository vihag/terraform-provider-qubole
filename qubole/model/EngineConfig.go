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
