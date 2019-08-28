package model

import (
	_ "fmt"
	"github.com/hashicorp/terraform/helper/schema"
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

/*
Read Engine Config from Terraform File
*/
func ReadEngineConfigFromTf(d *schema.ResourceData) (EngineConfig, bool) {

	var engine_config EngineConfig
	if v, ok := d.GetOk("engine_config"); ok {
		engineConfig := v.([]interface{})
		if len(engineConfig) > 0 {
			configs := engineConfig[0].(map[string]interface{})
			//Read engine type spark/presto/airflow/hadoop2
			if v, ok := configs["flavour"]; ok {
				engine_config.Flavour = v.(string)
			}

			//Read hadoop settings
			var hadoop_settings HadoopSettings
			if v, ok := configs["hadoop_settings"]; ok {
				hadoopSettings := v.([]interface{})
				ReadHadoopSettingsSettingsFromTf(&hadoop_settings, hadoopSettings)
				engine_config.Hadoop_settings = hadoop_settings
			}
			//Read presto settings
			var presto_settings PrestoSettings
			if v, ok := configs["presto_settings"]; ok {
				prestoSettings := v.([]interface{})
				ReadPrestoSettingsFromTf(&presto_settings, prestoSettings)
				engine_config.Presto_settings = presto_settings
			}
			//Read spark settings
			var spark_settings SparkSettings
			if v, ok := configs["spark_settings"]; ok {
				sparkSettings := v.([]interface{})
				ReadSparkSettingsFromTf(&spark_settings, sparkSettings)
				engine_config.Spark_settings = spark_settings
			}
			//Read hive settings
			var hive_settings HiveSettings
			if v, ok := configs["hive_settings"]; ok {
				hiveSettings := v.([]interface{})
				ReadHiveSettingsFromTf(&hive_settings, hiveSettings)
				engine_config.Hive_settings = hive_settings
			}
			//Read airflow settings
			var airflow_settings AirflowSettings
			if v, ok := configs["airflow_settings"]; ok {
				airflowSettings := v.([]interface{})
				ReadAirflowSettingsFromTf(&airflow_settings, airflowSettings)
				engine_config.Airflow_settings = airflow_settings
			}
			return engine_config, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return engine_config, false
}
