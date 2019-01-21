package entity

import (
	_ "fmt"
)

type HadoopSettings struct {
	Use_hadoop2                 bool					`json:"use_hadoop2,omitempty"`
	Use_spark                   bool					`json:"use_spark,omitempty"`
	Custom_config               string					`json:"custom_config,omitempty"`
	Fairscheduler_settings      FairSchedulerSettings	`json:"fairscheduler_settings,omitempty"`
	Use_qubole_placement_policy bool					`json:"use_qubole_placement_policy,omitempty"`
}
