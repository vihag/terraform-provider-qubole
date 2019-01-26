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
				if len(hadoopSettings) > 0 {
					configs := hadoopSettings[0].(map[string]interface{})
					if v, ok := configs["use_qubole_placement_policy"]; ok {
						hadoop_settings.Use_qubole_placement_policy = v.(bool)
					}
					if v, ok := configs["is_ha"]; ok {
						hadoop_settings.Is_ha = v.(bool)
					}
					if v, ok := configs["custom_hadoop_config"]; ok {
						hadoop_settings.Custom_hadoop_config = v.(string)
					}
					//Read Fair Scheduler Settings
					var fairscheduler_settings FairSchedulerSettings
					if v, ok := configs["fairscheduler_settings"]; ok {
						fairSchedulerSettings := v.([]interface{})
						if len(fairSchedulerSettings) > 0 {
							configs := fairSchedulerSettings[0].(map[string]interface{})
							if v, ok := configs["default_pool"]; ok {
								fairscheduler_settings.Default_pool = v.(string)
							}
							if v, ok := configs["fairscheduler_config_xml"]; ok {
								fairscheduler_settings.Fairscheduler_config_xml = v.(string)
							}
							hadoop_settings.Fairscheduler_settings = fairscheduler_settings
						}
					}
					engine_config.Hadoop_settings = hadoop_settings
				}
			}
			//Read presto settings
			var presto_settings PrestoSettings
			if v, ok := configs["presto_settings"]; ok {
				prestoSettings := v.([]interface{})
				if len(prestoSettings) > 0 {
					configs := prestoSettings[0].(map[string]interface{})
					if v, ok := configs["presto_version"]; ok {
						presto_settings.Presto_version = v.(string)
					}
					if v, ok := configs["custom_presto_config"]; ok {
						presto_settings.Custom_presto_config = v.(string)
					}
					if v, ok := configs["enable_rubix"]; ok {
						presto_settings.Enable_rubix = v.(bool)
					}
					engine_config.Presto_settings = presto_settings
				}
			}
			//Read spark settings
			var spark_settings SparkSettings
			if v, ok := configs["spark_settings"]; ok {
				sparkSettings := v.([]interface{})
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
					engine_config.Spark_settings = spark_settings
				}
			}
			//Read hive settings
			var hive_settings HiveSettings
			if v, ok := configs["hive_settings"]; ok {
				hiveSettings := v.([]interface{})
				if len(hiveSettings) > 0 {
					configs := hiveSettings[0].(map[string]interface{})
					if v, ok := configs["is_hs2"]; ok {
						hive_settings.Is_hs2 = v.(bool)
					}
					if v, ok := configs["hive_version"]; ok {
						hive_settings.Hive_version = v.(string)
					}
					if v, ok := configs["is_metadata_cache_enabled"]; ok {
						hive_settings.Is_metadata_cache_enabled = v.(bool)
					}
					if v, ok := configs["hs2_thrift_port"]; ok {
						hive_settings.Hs2_thrift_port = v.(int)
					}
					if v, ok := configs["overrides"]; ok {
						hive_settings.Overrides = v.(string)
					}
					if v, ok := configs["execution_engine"]; ok {
						hive_settings.Execution_engine = v.(string)
					}
					engine_config.Hive_settings = hive_settings
				}
			}
			//Read airflow settings
			var airflow_settings AirflowSettings
			if v, ok := configs["airflow_settings"]; ok {
				airflowSettings := v.([]interface{})
				if len(airflowSettings) > 0 {
					configs := airflowSettings[0].(map[string]interface{})
					if v, ok := configs["dbtap_id"]; ok {
						airflow_settings.Dbtap_id = v.(int)
					}
					if v, ok := configs["fernet_key"]; ok {
						airflow_settings.Fernet_key = v.(string)
					}
					if v, ok := configs["overrides"]; ok {
						airflow_settings.Overrides = v.(string)
					}
					if v, ok := configs["version"]; ok {
						airflow_settings.Version = v.(string)
					}
					if v, ok := configs["airflow_python_version"]; ok {
						airflow_settings.Airflow_python_version = v.(string)
					}
					engine_config.Airflow_settings = airflow_settings
				}
			}
			return engine_config, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return engine_config, false
}
