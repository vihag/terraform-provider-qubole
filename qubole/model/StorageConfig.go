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

/*
function to flatten Storage Config
*/
func FlattenStorageConfig(ia *StorageConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Storage_access_key != nil {
		attrs["storage_access_key"] = ia.Storage_access_key
	}
	if &ia.Storage_account_name != nil {
		attrs["storage_account_name"] = ia.Storage_account_name
	}
	if &ia.Disk_storage_account_name != nil {
		attrs["disk_storage_account_name"] = ia.Disk_storage_account_name
	}
	if &ia.Disk_storage_account_resource_group_name != nil {
		attrs["disk_storage_account_resource_group_name"] = ia.Disk_storage_account_resource_group_name
	}
	if &ia.Managed_disk_account_type != nil {
		attrs["managed_disk_account_type"] = ia.Managed_disk_account_type
	}
	if &ia.Data_disk_count != nil {
		attrs["data_disk_count"] = ia.Data_disk_count
	}
	if &ia.Data_disk_size != nil {
		attrs["data_disk_size"] = ia.Data_disk_size
	}
	if &ia.Disk_upscaling_config != nil {
		attrs["disk_upscaling_config"] = FlattenDiskUpscalingConfig(&ia.Disk_upscaling_config)
	}
	result = append(result, attrs)

	return result
}
