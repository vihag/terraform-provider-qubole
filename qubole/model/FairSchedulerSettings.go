package model

import (
	_ "fmt"
)

type FairSchedulerSettings struct {
	Default_pool             string `json:",default_pool,omitempty"`
	Fairscheduler_config_xml string `json:",fairscheduler_config_xml,omitempty"`
}
