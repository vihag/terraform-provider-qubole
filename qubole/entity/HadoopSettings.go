package entity

import (
	"fmt"
)

type HadoopSettings struct {
	use_hadoop2                 bool
	use_spark                   bool
	custom_config               string
	fairscheduler_settings      FairSchedulerSettings
	use_qubole_placement_policy bool
}
