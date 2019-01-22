package entity

import (
	_ "fmt"
)

type Cluster struct {
	Id                           int               `json:"id,omitempty"`
	State                        string            `json:"state,omitempty"`
	Force_tunnel                 bool              `json:"force_tunnel,omitempty"`
	Tunnel_server_ip             string            `json:"tunnel_server_ip,omitempty"`
	Label                        []*string         `json:"label,omitempty"`
	Presto_version               string            `json:"presto_version,omitempty"`
	Spark_version                string            `json:"spark_version,omitempty"`
	Zeppelin_interpreter_mode    string            `json:"zeppelin_interpreter_mode,omitempty"`
	Ec2_settings                 Ec2Settings       `json:"ec2_settings,omitempty"`
	Node_configuration           NodeConfiguration `json:"node_configuration,omitempty"`
	Hadoop_settings              HadoopSettings    `json:"hadoop_settings,omitempty"`
	Security_settings            SecuritySettings  `json:"security_settings,omitempty"`
	Presto_settings              PrestoSettings    `json:"presto_settings,omitempty"`
	Spark_settings               SparkSettings     `json:"spark_settings,omitempty"`
	Datadog_settings             DatadogSettings   `json:"datadog_settings,omitempty"`
	Disallow_cluster_termination bool              `json:"disallow_cluster_termination,omitempty"`
	Enable_ganglia_monitoring    bool              `json:"enable_ganglia_monitoring,omitempty"`
	Node_bootstrap_file          string            `json:"node_bootstrap_file,omitempty"`
	Idle_cluster_timeout         int               `json:"idle_cluster_timeout,omitempty"`
	Spark_s3_package_name        string            `json:"spark_s3_package_name,omitempty"`
	Zeppelin_s3_package_name     string            `json:"zeppelin_s3_package_name,omitempty"`
	Engine_config                EngineConfig      `json:"engine_config,omitempty"`
}
