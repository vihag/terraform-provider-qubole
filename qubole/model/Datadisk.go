package model

import (
	_ "fmt"
)

type Datadisk struct {
	Count                int                `json:"count,omitempty"`
	Disktype             string             `json:"type,omitempty"`
	Encryption           bool               `json:"encryption,omitempty"`
	Size                 int                `json:"size,omitempty"`
	Ebs_upscaling_config EbsUpscalingConfig `json:"ebs_upscaling_config,omitempty"`
}
