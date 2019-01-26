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

func ReadStorageConfigFromTf(storage_config *StorageConfig, storageConfig []interface{}) bool {

	if len(storageConfig) > 0 {
		configs := storageConfig[0].(map[string]interface{})
		if v, ok := configs["storage_access_key"]; ok {
			storage_config.Storage_access_key = v.(string)
		}
		if v, ok := configs["storage_account_name"]; ok {
			storage_config.Storage_account_name = v.(string)
		}
		if v, ok := configs["disk_storage_account_name"]; ok {
			storage_config.Disk_storage_account_name = v.(string)
		}
		if v, ok := configs["disk_storage_account_resource_group_name"]; ok {
			storage_config.Disk_storage_account_resource_group_name = v.(string)
		}
		if v, ok := configs["managed_disk_account_type"]; ok {
			storage_config.Managed_disk_account_type = v.(string)
		}
		if v, ok := configs["data_disk_count"]; ok {
			storage_config.Data_disk_count = v.(int)
		}
		if v, ok := configs["data_disk_size"]; ok {
			storage_config.Data_disk_size = v.(int)
		}
		//Read disk upscaling config
		var disk_upscaling_config DiskUpscalingConfig
		if v, ok := configs["disk_upscaling_config"]; ok {
			diskUpscalingConfig := v.([]interface{})
			ReadDiskUpscalingConfigFromTf(&disk_upscaling_config, diskUpscalingConfig)
			storage_config.Disk_upscaling_config = disk_upscaling_config
		}
	}

	return true
}
