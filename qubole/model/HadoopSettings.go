package model

import (
	_ "fmt"
)

type HadoopSettings struct {
	Use_qubole_placement_policy bool                  `json:"use_qubole_placement_policy,omitempty"`
	Is_ha                       bool                  `json:"is_ha,omitempty"`
	Custom_hadoop_config        string                `json:"custom_hadoop_config,omitempty"`
	Fairscheduler_settings      FairSchedulerSettings `json:"fairscheduler_settings,omitempty"`
}
