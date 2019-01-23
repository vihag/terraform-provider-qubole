package model

import (
	_ "fmt"
)

type StorageConfig struct {
	Storage_access_key                       string              `json:"storage_access_key,omitempty"`
	Storage_account_name                     string              `json:"storage_account_name,omitempty"`
	Disk_storage_account_name                string              `json:"disk_storage_account_name,omitempty"`
	Disk_storage_account_resource_group_name string              `json:"disk_storage_account_resource_group_name,omitempty"`
	Managed_disk_account_type                string              `json:"managed_disk_account_type,omitempty"`
	Data_disk_count                          int                 `json:"data_disk_count,omitempty"`
	Data_disk_size                           int                 `json:"data_disk_size,omitempty"`
	Disk_upscaling_config                    DiskUpscalingConfig `json:"disk_upscaling_config,omitempty"`
}
