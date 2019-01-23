package model

import (
	_ "fmt"
)

type HadoopSettings struct {
	Use_hadoop2                 bool                  `json:"use_hadoop2,omitempty"`
	Use_spark                   bool                  `json:"use_spark,omitempty"`
	Use_hbase                   bool                  `json:"use_hbase,omitempty"`
	Use_qubole_placement_policy bool                  `json:"use_qubole_placement_policy,omitempty"`
	Is_ha                       bool                  `json:"is_ha,omitempty"`
	Enable_rubix                bool                  `json:"enable_rubix,omitempty"`
	Node_bootstrap_timeout      int                   `json:"node_bootstrap_timeout,omitempty"`
	Custom_config               string                `json:"custom_config,omitempty"`
	Fairscheduler_settings      FairSchedulerSettings `json:"fairscheduler_settings,omitempty"`
}
