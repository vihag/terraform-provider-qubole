package entity

import (
	"fmt"
)

type Cluster struct {
	label                        []string
	presto_version               string
	spark_version                string
	zeppelin_interpreter_mode    boolean
	ec2_settings                 Ec2Settings
	node_configuration           NodeConfiguration
	hadoop_settings              HadoopSettings
	security_settings            SecuritySettings
	presto_settings              PrestoSettings
	spark_settings               SparkSettings
	datadog_settings             DatadogSettings
	disallow_cluster_termination bool
	enable_ganglia_monitoring    bool
	node_bootstrap_file          string
	idle_cluster_timeout         int
}
