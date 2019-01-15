package entity

import (
	_ "fmt"
)

type HadoopSettings struct {
	Use_hadoop2                 bool
	Use_spark                   bool
	Custom_config               string
	Fairscheduler_settings      FairSchedulerSettings
	Use_qubole_placement_policy bool
}
