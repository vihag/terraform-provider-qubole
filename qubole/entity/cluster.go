package entity

import (
	_ "fmt"
)

type Cluster struct {
	Dd                           string
	Label                        []string
	Presto_version               string
	Spark_version                string
	Zeppelin_interpreter_mode    string
	Ec2_settings                 Ec2Settings
	Node_configuration           NodeConfiguration
	Hadoop_settings              HadoopSettings
	Security_settings            SecuritySettings
	Presto_settings              PrestoSettings
	Spark_settings               SparkSettings
	Datadog_settings             DatadogSettings
	Disallow_cluster_termination bool
	Enable_ganglia_monitoring    bool
	Node_bootstrap_file          string
	Idle_cluster_timeout         int
}
