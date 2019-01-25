package qubole

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	model "github.com/terraform-providers/terraform-provider-qubole/qubole/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func resourceQuboleCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceQuboleClusterCreate,
		Read:   resourceQuboleClusterRead,
		Update: resourceQuboleClusterUpdate,
		Delete: resourceQuboleClusterDelete,

		Schema: map[string]*schema.Schema{
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Required: true,
						},
						"compute_config": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"compute_validated": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"use_account_compute_creds": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"instance_tenancy": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"role_instance_profile": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compute_role_arn": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compute_external_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compute_client_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compute_client_secret": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compute_tenant_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compute_subscription_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"location": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aws_region": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"aws_availability_zone": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"location": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"network_config": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"bastion_node_public_dns": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"bastion_node_port": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"bastion_node_user": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"master_elastic_ip": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"persistent_security_groups": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"persistent_security_group_resource_group_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"persistent_security_group_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vnet_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"subnet_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vnet_resource_group_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"master_static_nic_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"master_static_public_ip_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"storage_config": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"storage_access_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"storage_account_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"disk_storage_account_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"disk_storage_account_resource_group_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"managed_disk_account_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"data_disk_count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"data_disk_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"disk_upscaling_config": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"percent_free_space_threshold": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"max_data_disk_count": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"absolute_free_space_threshold": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"cluster_info": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_instance_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"slave_instance_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"node_base_cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"label": &schema.Schema{
							Type:     schema.TypeSet,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
						"min_nodes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_nodes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"idle_cluster_timeout_in_secs": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cluster_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"node_bootstrap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disallow_cluster_termination": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"force_tunnel": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"customer_ssh_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"child_hs2_cluster_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"parent_cluster_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"env_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"python_version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"r_version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"datadisk": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"encryption": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"size": { //Be careful, the GET will return this as an array
										Type:     schema.TypeInt,
										Optional: true,
									},
									"ebs_upscaling_config": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_ebs_volume_count": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"percent_free_space_threshold": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"absolute_free_space_threshold": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"sampling_interval": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"sampling_window": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"heterogeneous_config": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									/*
														Sample JSON input:
														{
											             [
										                   {"instance_type": "m4.4xlarge", "weight": 1.0},
										                   {"instance_type": "m4.2xlarge", "weight": 0.5}
										                  ]
										                 }
									*/
									//TODO better schema representation of this JSON instead of String
									"memory": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"slave_request_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"spot_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spot_instance_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"maximum_bid_price_percentage": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"timeout_for_request": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"maximum_spot_instance_percentage": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"stable_spot_instance_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"maximum_bid_price_percentage": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"timeout_for_request": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"spot_block_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"duration": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"custom_tags": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"fallback_to_ondemand": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"engine_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flavour": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"hadoop_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"use_qubole_placement_policy": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"is_ha": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"custom_hadoop_config": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"fairscheduler_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fairscheduler_config_xml": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"default_pool": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"presto_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"presto_version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"custom_presto_config": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"enable_rubix": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"spark_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_spark_config": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"spark_version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"enable_rubix": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"hive_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_hs2": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"hive_version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"is_metadata_cache_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"hs2_thrift_port": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"overrides": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"execution_engine": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"airflow_settings": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dbtap_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"fernet_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"overrides": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"airflow_python_version": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"monitoring": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ganglia": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"datadog": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"datadog_api_token": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"datadog_app_token": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"internal": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zeppelin_interpreter_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"spark_s3_package_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"zeppelin_s3_package_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

/*
Parser methods
*/
//TODO move to util classes
func readCloudConfigFromTf(d *schema.ResourceData) (model.CloudConfig, bool) {

	var cloud_config model.CloudConfig
	if v, ok := d.GetOk("cloud_config"); ok {
		cloudConfig := v.([]interface{})
		if len(cloudConfig) > 0 {
			configs := cloudConfig[0].(map[string]interface{})
			//Read cloud provider aws/azure/gcp
			if v, ok := configs["provider"]; ok {
				cloud_config.Provider = v.(string)
			}
			//Read compute config
			var compute_config model.ComputeConfig
			if v, ok := configs["compute_config"]; ok {
				computeConfig := v.([]interface{})
				if len(computeConfig) > 0 {
					configs := computeConfig[0].(map[string]interface{})
					if v, ok := configs["compute_validated"]; ok {
						compute_config.Compute_validated = v.(bool)
					}
					if v, ok := configs["use_account_compute_creds"]; ok {
						compute_config.Use_account_compute_creds = v.(bool)
					}
					if v, ok := configs["instance_tenancy"]; ok {
						compute_config.Instance_tenancy = v.(string)
					}
					if v, ok := configs["role_instance_profile"]; ok {
						compute_config.Role_instance_profile = v.(string)
					}
					if v, ok := configs["compute_role_arn"]; ok {
						compute_config.Compute_role_arn = v.(string)
					}
					if v, ok := configs["compute_external_id"]; ok {
						compute_config.Compute_external_id = v.(string)
					}
					if v, ok := configs["compute_client_id"]; ok {
						compute_config.Compute_client_id = v.(string)
					}
					if v, ok := configs["compute_client_secret"]; ok {
						compute_config.Compute_client_secret = v.(string)
					}
					if v, ok := configs["compute_tenant_id"]; ok {
						compute_config.Compute_tenant_id = v.(string)
					}
					if v, ok := configs["compute_subscription_id"]; ok {
						compute_config.Compute_subscription_id = v.(string)
					}
					cloud_config.Compute_config = compute_config
				}
			}
			//Read location
			var location model.Location
			if v, ok := configs["location"]; ok {
				locationConfig := v.([]interface{})
				if len(locationConfig) > 0 {
					configs := locationConfig[0].(map[string]interface{})
					if v, ok := configs["aws_region"]; ok {
						location.Aws_region = v.(string)
					}
					if v, ok := configs["aws_availability_zone"]; ok {
						location.Aws_availability_zone = v.(string)
					}
					if v, ok := configs["location"]; ok {
						location.Location = v.(string)
					}
					cloud_config.Location = location
				}
			}
			//Read network config
			var network_config model.NetworkConfig
			if v, ok := configs["network_config"]; ok {
				networkConfig := v.([]interface{})
				if len(networkConfig) > 0 {
					configs := networkConfig[0].(map[string]interface{})
					if v, ok := configs["vpc_id"]; ok {
						network_config.Vpc_id = v.(string)
					}
					if v, ok := configs["subnet_id"]; ok {
						network_config.Subnet_id = v.(string)
					}
					if v, ok := configs["bastion_node_public_dns"]; ok {
						network_config.Bastion_node_public_dns = v.(string)
					}
					if v, ok := configs["bastion_node_port"]; ok {
						network_config.Bastion_node_port = v.(int)
					}
					if v, ok := configs["bastion_node_user"]; ok {
						network_config.Bastion_node_user = v.(string)
					}
					if v, ok := configs["master_elastic_ip"]; ok {
						network_config.Master_elastic_ip = v.(string)
					}
					if v, ok := configs["persistent_security_groups"]; ok {
						network_config.Persistent_security_groups = v.(string)
					}
					if v, ok := configs["persistent_security_group_resource_group_name"]; ok {
						network_config.Persistent_security_group_resource_group_name = v.(string)
					}
					if v, ok := configs["persistent_security_group_name"]; ok {
						network_config.Persistent_security_group_name = v.(string)
					}
					if v, ok := configs["vnet_name"]; ok {
						network_config.Vnet_name = v.(string)
					}
					if v, ok := configs["subnet_name"]; ok {
						network_config.Subnet_name = v.(string)
					}
					if v, ok := configs["vnet_resource_group_name"]; ok {
						network_config.Vnet_resource_group_name = v.(string)
					}
					if v, ok := configs["master_static_nic_name"]; ok {
						network_config.Master_static_nic_name = v.(string)
					}
					if v, ok := configs["master_static_public_ip_name"]; ok {
						network_config.Master_static_public_ip_name = v.(string)
					}
					cloud_config.Network_config = network_config
				}
			}
			//Read storage config
			var storage_config model.StorageConfig
			if v, ok := configs["storage_config"]; ok {
				storageConfig := v.([]interface{})
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
					var disk_upscaling_config model.DiskUpscalingConfig
					if v, ok := configs["disk_upscaling_config"]; ok {
						diskUpscalingConfig := v.([]interface{})
						if len(diskUpscalingConfig) > 0 {
							configs := diskUpscalingConfig[0].(map[string]interface{})
							if v, ok := configs["percent_free_space_threshold"]; ok {
								disk_upscaling_config.Percent_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["max_data_disk_count"]; ok {
								disk_upscaling_config.Max_data_disk_count = v.(int)
							}
							if v, ok := configs["absolute_free_space_threshold"]; ok {
								disk_upscaling_config.Absolute_free_space_threshold = float32(v.(int))
							}
							storage_config.Disk_upscaling_config = disk_upscaling_config
						}
					}
					cloud_config.Storage_config = storage_config
				}
			}
			return cloud_config, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return cloud_config, false
}

func readClusterInfoFromTf(d *schema.ResourceData) (model.ClusterInfo, bool) {

	var cluster_info model.ClusterInfo
	if v, ok := d.GetOk("cluster_info"); ok {
		clusterInfo := v.([]interface{})
		if len(clusterInfo) > 0 {
			configs := clusterInfo[0].(map[string]interface{})

			if v, ok := configs["master_instance_type"]; ok {
				cluster_info.Master_instance_type = v.(string)
			}

			if v, ok := configs["slave_instance_type"]; ok {
				cluster_info.Slave_instance_type = v.(string)
			}

			if v, ok := configs["node_base_cooldown_period"]; ok {
				cluster_info.Node_base_cooldown_period = v.(int)
			}

			if v, ok := configs["label"]; ok {
				labelSet := v.(*schema.Set)

				labels := make([]string, labelSet.Len())
				for i, label := range labelSet.List() {
					labels[i] = label.(string)
				}

				cluster_info.Label = labels
			}

			if v, ok := configs["min_nodes"]; ok {
				cluster_info.Min_nodes = v.(int)
			}

			if v, ok := configs["max_nodes"]; ok {
				cluster_info.Max_nodes = v.(int)
			}

			if v, ok := configs["idle_cluster_timeout_in_secs"]; ok {
				cluster_info.Idle_cluster_timeout_in_secs = v.(int)
			}

			if v, ok := configs["cluster_name"]; ok {
				cluster_info.Cluster_name = v.(string)
			}

			if v, ok := configs["node_bootstrap"]; ok {
				cluster_info.Node_bootstrap = v.(string)
			}

			if v, ok := configs["disallow_cluster_termination"]; ok {
				cluster_info.Disallow_cluster_termination = v.(bool)
			}

			if v, ok := configs["force_tunnel"]; ok {
				cluster_info.Force_tunnel = v.(bool)
			}

			if v, ok := configs["customer_ssh_key"]; ok {
				cluster_info.Customer_ssh_key = v.(string)
			}

			if v, ok := configs["child_hs2_cluster_id"]; ok {
				cluster_info.Child_hs2_cluster_id = v.(int)
			}

			if v, ok := configs["parent_cluster_id"]; ok {
				cluster_info.Parent_cluster_id = v.(int)
			}
			//Read Env Settings
			var env_settings model.EnvSettings
			if v, ok := configs["env_settings"]; ok {
				envSettings := v.([]interface{})
				if len(envSettings) > 0 {
					configs := envSettings[0].(map[string]interface{})
					if v, ok := configs["python_version"]; ok {
						env_settings.Python_version = v.(string)
					}
					if v, ok := configs["r_version"]; ok {
						env_settings.R_version = v.(string)
					}
					if v, ok := configs["name"]; ok {
						env_settings.Name = v.(string)
					}
					cluster_info.Env_settings = env_settings
				}
			}
			//Read datadisk
			var datadisk model.Datadisk
			if v, ok := configs["datadisk"]; ok {
				datadiskConfig := v.([]interface{})
				if len(datadiskConfig) > 0 {
					configs := datadiskConfig[0].(map[string]interface{})
					if v, ok := configs["count"]; ok {
						datadisk.Count = v.(int)
					}
					if v, ok := configs["type"]; ok {
						datadisk.Disktype = v.(string)
					}
					if v, ok := configs["encryption"]; ok {
						datadisk.Encryption = v.(bool)
					}
					if v, ok := configs["size"]; ok {
						datadisk.Size = v.(int)
					}
					//Read disk upscaling config
					var ebs_upscaling_config model.EbsUpscalingConfig
					if v, ok := configs["ebs_upscaling_config"]; ok {
						ebsUpscalingConfigs := v.([]interface{})
						if len(ebsUpscalingConfigs) > 0 {
							configs := ebsUpscalingConfigs[0].(map[string]interface{})
							if v, ok := configs["max_ebs_volume_count"]; ok {
								ebs_upscaling_config.Max_ebs_volume_count = v.(int)
							}
							if v, ok := configs["percent_free_space_threshold"]; ok {
								ebs_upscaling_config.Percent_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["absolute_free_space_threshold"]; ok {
								ebs_upscaling_config.Absolute_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["sampling_interval"]; ok {
								ebs_upscaling_config.Sampling_interval = v.(int)
							}
							if v, ok := configs["sampling_window"]; ok {
								ebs_upscaling_config.Sampling_window = v.(int)
							}
							datadisk.Ebs_upscaling_config = ebs_upscaling_config
						}
					}
					cluster_info.Datadisk = datadisk
				}
			}
			//Read Heterogeneous Config
			var heterogeneous_config model.HeterogeneousConfig
			if v, ok := configs["heterogeneous_config"]; ok {
				heterogeneousConfigs := v.([]interface{})
				if len(heterogeneousConfigs) > 0 {
					configs := heterogeneousConfigs[0].(map[string]interface{})
					if v, ok := configs["memory"]; ok {
						heterogeneous_config.Memory = v.(string)
					}
					cluster_info.Heterogeneous_config = heterogeneous_config
				}
			}

			if v, ok := configs["slave_request_type"]; ok {
				cluster_info.Slave_request_type = v.(string)
			}

			//Read spot settings
			var spot_settings model.SpotSettings
			if v, ok := configs["spot_settings"]; ok {
				spotSettings := v.([]interface{})
				if len(spotSettings) > 0 {
					configs := spotSettings[0].(map[string]interface{})

					//Read spot instance settings
					var spot_instance_settings model.SpotInstanceSettings
					if v, ok := configs["spot_instance_settings"]; ok {
						spotInstanceSettings := v.([]interface{})
						if len(spotInstanceSettings) > 0 {
							configs := spotInstanceSettings[0].(map[string]interface{})
							if v, ok := configs["maximum_bid_price_percentage"]; ok {
								spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
							}
							if v, ok := configs["timeout_for_request"]; ok {
								spot_instance_settings.Timeout_for_request = v.(int)
							}
							if v, ok := configs["maximum_spot_instance_percentage"]; ok {
								spot_instance_settings.Maximum_spot_instance_percentage = float32(v.(int))
							}
							spot_settings.Spot_instance_settings = spot_instance_settings
						}
					}

					//Read stable spot instance settings
					var stable_spot_instance_settings model.StableSpotInstanceSettings
					if v, ok := configs["stable_spot_instance_settings"]; ok {
						stableSpotInstanceSettings := v.([]interface{})
						if len(stableSpotInstanceSettings) > 0 {
							configs := stableSpotInstanceSettings[0].(map[string]interface{})
							if v, ok := configs["maximum_bid_price_percentage"]; ok {
								stable_spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
							}
							if v, ok := configs["timeout_for_request"]; ok {
								stable_spot_instance_settings.Timeout_for_request = v.(int)
							}
							spot_settings.Stable_spot_instance_settings = stable_spot_instance_settings
						}
					}

					//Read spot block settings
					var spot_block_settings model.SpotBlockSettings
					if v, ok := configs["spot_block_settings"]; ok {
						spotBlockSettings := v.([]interface{})
						if len(spotBlockSettings) > 0 {
							configs := spotBlockSettings[0].(map[string]interface{})
							if v, ok := configs["duration"]; ok {
								spot_block_settings.Duration = v.(int)
							}
							spot_settings.Spot_block_settings = spot_block_settings
						}
					}

					cluster_info.Spot_settings = spot_settings
				}
			}

			if v, ok := configs["custom_tags"]; ok {
				billing_tags := v.(map[string]interface{})
				custom_tags := make(map[string]string)
				for key, value := range billing_tags {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)

					custom_tags[strKey] = strValue
				}
				cluster_info.Custom_tags = custom_tags
			}

			if v, ok := configs["fallback_to_ondemand"]; ok {
				cluster_info.Fallback_to_ondemand = v.(bool)
			}

			return cluster_info, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return cluster_info, false

}

func readEngineConfigFromTf(d *schema.ResourceData) (model.EngineConfig, bool) {

	var engine_config model.EngineConfig
	if v, ok := d.GetOk("engine_config"); ok {
		engineConfig := v.([]interface{})
		if len(engineConfig) > 0 {
			configs := engineConfig[0].(map[string]interface{})
			//Read engine type spark/presto/airflow/hadoop2
			if v, ok := configs["flavour"]; ok {
				engine_config.Flavour = v.(string)
			}

			//Read hadoop settings
			var hadoop_settings model.HadoopSettings
			if v, ok := configs["hadoop_settings"]; ok {
				hadoopSettings := v.([]interface{})
				if len(hadoopSettings) > 0 {
					configs := hadoopSettings[0].(map[string]interface{})
					if v, ok := configs["use_qubole_placement_policy"]; ok {
						hadoop_settings.Use_qubole_placement_policy = v.(bool)
					}
					if v, ok := configs["is_ha"]; ok {
						hadoop_settings.Is_ha = v.(bool)
					}
					if v, ok := configs["custom_hadoop_config"]; ok {
						hadoop_settings.Custom_hadoop_config = v.(string)
					}
					//Read Fair Scheduler Settings
					var fairscheduler_settings model.FairSchedulerSettings
					if v, ok := configs["fairscheduler_settings"]; ok {
						fairSchedulerSettings := v.([]interface{})
						if len(fairSchedulerSettings) > 0 {
							configs := fairSchedulerSettings[0].(map[string]interface{})
							if v, ok := configs["default_pool"]; ok {
								fairscheduler_settings.Default_pool = v.(string)
							}
							if v, ok := configs["fairscheduler_config_xml"]; ok {
								fairscheduler_settings.Fairscheduler_config_xml = v.(string)
							}
							hadoop_settings.Fairscheduler_settings = fairscheduler_settings
						}
					}
					engine_config.Hadoop_settings = hadoop_settings
				}
			}
			//Read presto settings
			var presto_settings model.PrestoSettings
			if v, ok := configs["presto_settings"]; ok {
				prestoSettings := v.([]interface{})
				if len(prestoSettings) > 0 {
					configs := prestoSettings[0].(map[string]interface{})
					if v, ok := configs["presto_version"]; ok {
						presto_settings.Presto_version = v.(string)
					}
					if v, ok := configs["custom_presto_config"]; ok {
						presto_settings.Custom_presto_config = v.(string)
					}
					if v, ok := configs["enable_rubix"]; ok {
						presto_settings.Enable_rubix = v.(bool)
					}
					engine_config.Presto_settings = presto_settings
				}
			}
			//Read spark settings
			var spark_settings model.SparkSettings
			if v, ok := configs["spark_settings"]; ok {
				sparkSettings := v.([]interface{})
				if len(sparkSettings) > 0 {
					configs := sparkSettings[0].(map[string]interface{})
					if v, ok := configs["spark_version"]; ok {
						spark_settings.Spark_version = v.(string)
					}
					if v, ok := configs["custom_spark_config"]; ok {
						spark_settings.Custom_spark_config = v.(string)
					}
					if v, ok := configs["enable_rubix"]; ok {
						spark_settings.Enable_rubix = v.(bool)
					}
					engine_config.Spark_settings = spark_settings
				}
			}
			//Read hive settings
			var hive_settings model.HiveSettings
			if v, ok := configs["hive_settings"]; ok {
				hiveSettings := v.([]interface{})
				if len(hiveSettings) > 0 {
					configs := hiveSettings[0].(map[string]interface{})
					if v, ok := configs["is_hs2"]; ok {
						hive_settings.Is_hs2 = v.(bool)
					}
					if v, ok := configs["hive_version"]; ok {
						hive_settings.Hive_version = v.(string)
					}
					if v, ok := configs["is_metadata_cache_enabled"]; ok {
						hive_settings.Is_metadata_cache_enabled = v.(bool)
					}
					if v, ok := configs["hs2_thrift_port"]; ok {
						hive_settings.Hs2_thrift_port = v.(int)
					}
					if v, ok := configs["overrides"]; ok {
						hive_settings.Overrides = v.(string)
					}
					if v, ok := configs["execution_engine"]; ok {
						hive_settings.Execution_engine = v.(string)
					}
					engine_config.Hive_settings = hive_settings
				}
			}
			//Read airflow settings
			var airflow_settings model.AirflowSettings
			if v, ok := configs["airflow_settings"]; ok {
				airflowSettings := v.([]interface{})
				if len(airflowSettings) > 0 {
					configs := airflowSettings[0].(map[string]interface{})
					if v, ok := configs["dbtap_id"]; ok {
						airflow_settings.Dbtap_id = v.(int)
					}
					if v, ok := configs["fernet_key"]; ok {
						airflow_settings.Fernet_key = v.(string)
					}
					if v, ok := configs["overrides"]; ok {
						airflow_settings.Overrides = v.(string)
					}
					if v, ok := configs["version"]; ok {
						airflow_settings.Version = v.(string)
					}
					if v, ok := configs["airflow_python_version"]; ok {
						airflow_settings.Airflow_python_version = v.(string)
					}
					engine_config.Airflow_settings = airflow_settings
				}
			}
			return engine_config, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return engine_config, false
}

func readMonitoringFromTf(d *schema.ResourceData) (model.Monitoring, bool) {

	var monitoring model.Monitoring
	if v, ok := d.GetOk("monitoring"); ok {
		monitoringConfig := v.([]interface{})
		if len(monitoringConfig) > 0 {
			configs := monitoringConfig[0].(map[string]interface{})

			if v, ok := configs["ganglia"]; ok {
				monitoring.Ganglia = v.(bool)
			}

			//Read datadog settings
			var datadog model.Datadog
			if v, ok := configs["datadog"]; ok {
				datadogSettings := v.([]interface{})
				if len(datadogSettings) > 0 {
					configs := datadogSettings[0].(map[string]interface{})
					if v, ok := configs["datadog_api_token"]; ok {
						datadog.Datadog_api_token = v.(string)
					}
					if v, ok := configs["datadog_app_token"]; ok {
						datadog.Datadog_app_token = v.(string)
					}
					monitoring.Datadog = datadog
				}
			}

			return monitoring, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return monitoring, false
}

func readInternalFromTf(d *schema.ResourceData) (model.Internal, bool) {

	var internal model.Internal
	if v, ok := d.GetOk("internal"); ok {
		internalConfig := v.([]interface{})
		if len(internalConfig) > 0 {
			configs := internalConfig[0].(map[string]interface{})

			if v, ok := configs["zeppelin_interpreter_mode"]; ok {
				internal.Zeppelin_interpreter_mode = v.(string)
			}
			if v, ok := configs["spark_s3_package_name"]; ok {
				internal.Spark_s3_package_name = v.(string)
			}
			if v, ok := configs["zeppelin_s3_package_name"]; ok {
				internal.Zeppelin_s3_package_name = v.(string)
			}

			return internal, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return internal, false
}

func readClusterFromTf(d *schema.ResourceData) (model.Cluster, bool) {

	//Create the representative json object here
	var cluster model.Cluster

	//create nested datas structures
	//1. Cloud Config
	if cloud_config, ok := readCloudConfigFromTf(d); ok {
		cluster.Cloud_config = cloud_config
	} else {
		log.Printf("[WARN] No cloud_config seen.")
	}

	//2. Cluster Info
	if cluster_info, ok := readClusterInfoFromTf(d); ok {
		cluster.Cluster_info = cluster_info
	} else {
		log.Printf("[WARN] No cluster_info seen.")
	}

	//3. Engine Config
	if engine_config, ok := readEngineConfigFromTf(d); ok {
		cluster.Engine_config = engine_config
	} else {
		log.Printf("[WARN] No engine_config seen.")
	}

	//4. Monitoring
	if monitoring, ok := readMonitoringFromTf(d); ok {
		cluster.Monitoring = monitoring
	} else {
		log.Printf("[WARN] No monitoring seen.")
	}

	//5. Internal
	if internal, ok := readInternalFromTf(d); ok {
		cluster.Internal = internal
	} else {
		log.Printf("[WARN] No internal seen.")
	}

	//Finally, the cluster
	return cluster, true

}

/*
1. If the Create callback returns with or without an error without an ID set using SetId,
	the resource is assumed to not be created, and no state is saved.

2. If the Create callback returns with or without an error and an ID has been set,
	the resource is assumed created and all state is saved with it.
	Repeating because it is important: if there is an error, but the ID is set, the state is fully saved.
*/
//m holds the data returned by the provider configurer, in this it will be a struct with the configuration
func resourceQuboleClusterCreate(d *schema.ResourceData, meta interface{}) error {

	api_url := meta.(*Config).ConnectionString
	auth_token := meta.(*Config).AuthToken

	if cluster, ok := readClusterFromTf(d); ok {
		log.Printf("[INFO]constructing payload json representing the cluster.....")
		cluster_json, err := json.Marshal(cluster)
		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Error in marshalling json during create %s", err.Error())
		}
		log.Printf(string(cluster_json))

		//Make the http call to api here
		log.Printf("[INFO]Sending Create Cluster Request to URI %s", api_url)
		var payload = []byte(string(cluster_json))
		req, err := http.NewRequest(http.MethodPost, api_url, bytes.NewBuffer(payload))
		req.Header.Set("X-AUTH-TOKEN", auth_token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Error in cluster create API %s", err.Error())
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			log.Printf("Error in creation. Http Status Code %s %s", resp.StatusCode, string(body))
			return fmt.Errorf("Error in cluster create %s", string(body))
		}

		log.Printf("[INFO]response Body:", string(body))

		//Parse the response back to cluster object
		/*var clusterResponse model.Cluster
		unmarshallingError := json.Unmarshal(body, &clusterResponse)
		if unmarshallingError != nil {
			log.Printf("[ERR]There was an error:", unmarshallingError.Error())
			return fmt.Errorf("Error in unmarshalling json during update %s", err.Error())
		}*/
		//Parse the response using a custom un-marshaller
		var cluster_response *model.Cluster
		unmarshallingError := json.Unmarshal(body, &cluster_response)
		if unmarshallingError != nil {
			log.Printf("[ERR]There was an error:", unmarshallingError.Error())
			return fmt.Errorf("Error in unmarshalling json during update %s", unmarshallingError.Error())
		}
		log.Printf("[INFO]Pretty Printing Unmarshalled Response %#v", cluster_response)

		//Set Terraform ID; typecast the received ID to string for terraform
		d.SetId(strconv.Itoa(cluster_response.Id))

	} else {
		log.Printf("[WARN] No valid cluster definition seen.")
	}

	return resourceQuboleClusterRead(d, meta)
}

/*
The Read callback is used to sync the local state with the actual state (upstream). This is called at various points by Terraform and should be a read-only operation.
This callback should never modify the real resource.
*/
func resourceQuboleClusterRead(d *schema.ResourceData, meta interface{}) error {

	api_url := meta.(*Config).ConnectionString
	auth_token := meta.(*Config).AuthToken

	final_url := api_url + d.Id()

	//Make the http call to api here
	log.Printf("Sending GET Cluster Request to URI %s", final_url)
	req, err := http.NewRequest(http.MethodGet, final_url, nil)
	req.Header.Set("X-AUTH-TOKEN", auth_token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf(err.Error())
		d.SetId("")
		return fmt.Errorf("Error in cluster read %s", err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Printf("Error in read. Http Status Code %s %s", resp.StatusCode, string(body))
		return fmt.Errorf("Error in cluster read %s", string(body))
	}

	log.Printf("response Body:", string(body))

	//Unmarshal the response to a cluster object
	var cluster_response *model.Cluster
	unmarshallingError := json.Unmarshal(body, &cluster_response)
	if unmarshallingError != nil {
		log.Printf("[ERR]There was an error in unmarshalling during cluster read : %#v", unmarshallingError)
		return fmt.Errorf("Error in unmarshalling json during cluster read %#v", unmarshallingError)
	}
	log.Printf("[INFO]Pretty Printing Unmarshalled Response %#v", cluster_response)

	//Now start setting d with data from the unmarshalled object
	//Set cloud_config
	if err := d.Set("cloud_config", flattenCloudConfig(&cluster_response.Cloud_config)); err != nil {
		log.Printf("[ERR] Error setting cloud config: %#v", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting cloud config: %s", err)
	}
	//Set cluster_info
	if err := d.Set("cluster_info", flattenClusterInfo(&cluster_response.Cluster_info)); err != nil {
		log.Printf("[ERR] Error setting cluster info: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting cluster info: %s", err)
	}

	//Set engine_config
	if err := d.Set("engine_config", flattenEngineConfig(&cluster_response.Engine_config)); err != nil {
		log.Printf("[ERR] Error setting Engine Config: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Engine Config: %s", err)
	}

	//Set monitoring
	if err := d.Set("monitoring", flattenMonitoring(&cluster_response.Monitoring)); err != nil {
		log.Printf("[ERR] Error setting monitoring: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Monitoring: %s", err)
	}

	//Set internal
	if err := d.Set("internal", flattenInternal(&cluster_response.Internal)); err != nil {
		log.Printf("[ERR] Error setting Internal: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Internal: %s", err)
	}

	//Set rest of the simple objects
	d.Set("state", cluster_response.State)

	return nil
}

/*
1. If the Update callback returns with or without an error, the full state is saved.
	If the ID becomes blank, the resource is destroyed (even within an update, though this shouldn't happen except in error scenarios).
*/
func resourceQuboleClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	api_url := meta.(*Config).ConnectionString
	auth_token := meta.(*Config).AuthToken

	if cluster, ok := readClusterFromTf(d); ok {
		log.Printf("[INFO]constructing payload json representing the cluster.....")
		cluster_json, err := json.Marshal(cluster)
		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Error in marshalling json during update %s", err.Error())
		}
		log.Printf(string(cluster_json))

		//Make the http call to api here
		final_url := api_url + d.Id()
		log.Printf("[INFO]Sending Update Cluster Request to URI %s", final_url)
		var payload = []byte(string(cluster_json))
		req, err := http.NewRequest(http.MethodPut, final_url, bytes.NewBuffer(payload))
		req.Header.Set("X-AUTH-TOKEN", auth_token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Error in cluster update API cal %s", err.Error())
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			log.Printf("Error in update. Http Status Code %s %s", resp.StatusCode, string(body))
			return fmt.Errorf("Error in cluster update %s", string(body))
		}

		log.Printf("[INFO]response Body:", string(body))

		//Unmarshal the response to a cluster object
		var cluster_response *model.Cluster
		unmarshallingError := json.Unmarshal(body, &cluster_response)
		if unmarshallingError != nil {
			log.Printf("[ERR]There was an error in unmarshalling during cluster read : %#v", unmarshallingError)
			return fmt.Errorf("Error in unmarshalling json during cluster read %#v", unmarshallingError)
		}

		log.Printf("[INFO]Pretty Printing Unmarshalled Response %#v", cluster_response)

		//Set Terraform ID; typecast the received ID to string for terraform
		d.SetId(strconv.Itoa(cluster_response.Id))

	} else {
		log.Printf("[WARN] No valid cluster definition seen.")
	}

	return resourceQuboleClusterRead(d, meta)
}

/*
1. If the Destroy callback returns without an error, the resource is assumed to be destroyed, and all state is removed.

2. If the Destroy callback returns with an error, the resource is assumed to still exist, and all prior state is preserved.
*/
func resourceQuboleClusterDelete(d *schema.ResourceData, meta interface{}) error {
	// d.SetId("") is automatically called assuming delete returns no errors
	api_url := meta.(*Config).ConnectionString
	auth_token := meta.(*Config).AuthToken

	final_url := api_url + d.Id()

	//Make the http call to api here
	log.Printf("Sending Delete Cluster Request to URI %s", final_url)
	req, err := http.NewRequest(http.MethodDelete, final_url, nil)
	req.Header.Set("X-AUTH-TOKEN", auth_token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Error in cluster delete API call %s", err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Printf("Error in deletion. Http Status Code %s %s", resp.StatusCode, string(body))
		return fmt.Errorf("Error in cluster delete %s", string(body))
	}

	log.Printf("response Body:", string(body))

	//Nilling the terraform resource id explicitly
	d.SetId("")
	return nil
}

/*
function to flatten airflow settings to the schema that is defined
*/
func flattenAirflowSettings(ia *model.AirflowSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Dbtap_id != nil {
		attrs["dbtap_id"] = ia.Dbtap_id
	}
	if &ia.Fernet_key != nil {
		attrs["fernet_key"] = ia.Fernet_key
	}
	if &ia.Overrides != nil {
		attrs["overrides"] = ia.Overrides
	}
	if &ia.Version != nil {
		attrs["version"] = ia.Version
	}
	if &ia.Airflow_python_version != nil {
		attrs["airflow_python_version"] = ia.Airflow_python_version
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Cloud Config
*/
func flattenCloudConfig(ia *model.CloudConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Provider != nil {
		attrs["provider"] = ia.Provider
	}

	if &ia.Compute_config != nil {
		attrs["compute_config"] = flattenComputeConfig(&ia.Compute_config)
	}

	if &ia.Location != nil {
		attrs["location"] = flattenLocation(&ia.Location)
	}

	if &ia.Network_config != nil {
		attrs["network_config"] = flattenNetworkConfig(&ia.Network_config)
	}

	if &ia.Storage_config != nil {
		attrs["storage_config"] = flattenStorageConfig(&ia.Storage_config)
	}

	result = append(result, attrs)

	log.Print(result)
	return result
}

/*
function to Cluster Info
*/
func flattenClusterInfo(ia *model.ClusterInfo) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Master_instance_type != nil {
		attrs["master_instance_type"] = ia.Master_instance_type
	}

	if &ia.Slave_instance_type != nil {
		attrs["slave_instance_type"] = ia.Slave_instance_type
	}

	if &ia.Node_base_cooldown_period != nil {
		attrs["node_base_cooldown_period"] = ia.Node_base_cooldown_period
	}

	if &ia.Label != nil {
		log.Print(ia.Label)
		attrs["label"] = ia.Label
	}

	if &ia.Min_nodes != nil {
		attrs["min_nodes"] = ia.Min_nodes
	}

	if &ia.Max_nodes != nil {
		attrs["max_nodes"] = ia.Max_nodes
	}

	if &ia.Idle_cluster_timeout_in_secs != nil {
		attrs["idle_cluster_timeout_in_secs"] = ia.Idle_cluster_timeout_in_secs
	}

	if &ia.Cluster_name != nil {
		attrs["cluster_name"] = ia.Cluster_name
	}

	if &ia.Node_bootstrap != nil {
		attrs["node_bootstrap"] = ia.Node_bootstrap
	}

	if &ia.Disallow_cluster_termination != nil {
		attrs["disallow_cluster_termination"] = ia.Disallow_cluster_termination
	}

	if &ia.Force_tunnel != nil {
		attrs["force_tunnel"] = ia.Force_tunnel
	}

	if &ia.Customer_ssh_key != nil {
		attrs["customer_ssh_key"] = ia.Customer_ssh_key
	}

	if &ia.Child_hs2_cluster_id != nil {
		attrs["child_hs2_cluster_id"] = ia.Child_hs2_cluster_id
	}

	if &ia.Parent_cluster_id != nil {
		attrs["parent_cluster_id"] = ia.Parent_cluster_id
	}

	if &ia.Env_settings != nil {
		attrs["env_settings"] = flattenEnvSettings(&ia.Env_settings)
	}

	if &ia.Datadisk != nil {
		attrs["datadisk"] = flattenDatadisk(&ia.Datadisk)
	}

	if &ia.Heterogeneous_config != nil {
		attrs["heterogeneous_config"] = flattenHeterogeneousConfig(&ia.Heterogeneous_config)
	}

	if &ia.Slave_request_type != nil {
		attrs["slave_request_type"] = ia.Slave_request_type
	}

	if &ia.Custom_tags != nil {
		attrs["custom_tags"] = ia.Custom_tags
	}

	if &ia.Fallback_to_ondemand != nil {
		attrs["fallback_to_ondemand"] = ia.Fallback_to_ondemand
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Compute Config
*/
func flattenComputeConfig(ia *model.ComputeConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Compute_validated != nil {
		attrs["compute_validated"] = ia.Compute_validated
	}

	if &ia.Use_account_compute_creds != nil {
		attrs["use_account_compute_creds"] = ia.Use_account_compute_creds
	}

	if &ia.Instance_tenancy != nil {
		attrs["instance_tenancy"] = ia.Instance_tenancy
	}

	if &ia.Compute_role_arn != nil {
		attrs["compute_role_arn"] = ia.Compute_role_arn
	}

	if &ia.Compute_external_id != nil {
		attrs["compute_external_id"] = ia.Compute_external_id
	}

	if &ia.Role_instance_profile != nil {
		attrs["role_instance_profile"] = ia.Role_instance_profile
	}

	if &ia.Compute_client_id != nil {
		attrs["compute_client_id"] = ia.Compute_client_id
	}

	if &ia.Compute_client_secret != nil {
		attrs["compute_client_secret"] = ia.Compute_client_secret
	}

	if &ia.Compute_tenant_id != nil {
		attrs["compute_tenant_id"] = ia.Compute_tenant_id
	}

	if &ia.Compute_subscription_id != nil {
		attrs["compute_subscription_id"] = ia.Compute_subscription_id
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Datadisk
*/
func flattenDatadisk(ia *model.Datadisk) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Count != nil {
		attrs["count"] = ia.Count
	}

	if &ia.Disktype != nil {
		attrs["type"] = ia.Disktype
	}

	if &ia.Encryption != nil {
		attrs["encryption"] = ia.Encryption
	}

	if &ia.Size != nil {
		attrs["size"] = ia.Size
	}

	if &ia.Ebs_upscaling_config != nil {
		attrs["ebs_upscaling_config"] = flattenEbsUpscalingConfig(&ia.Ebs_upscaling_config)
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Datadog Settings
*/
func flattenDatadog(ia *model.Datadog) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Datadog_api_token != nil {
		attrs["datadog_api_token"] = ia.Datadog_api_token
	}

	if &ia.Datadog_app_token != nil {
		attrs["datadog_app_token"] = ia.Datadog_app_token
	}

	result = append(result, attrs)

	return result
}

/*
function to Disk Upscaling Settings
*/
func flattenDiskUpscalingConfig(ia *model.DiskUpscalingConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Percent_free_space_threshold != nil {
		attrs["percent_free_space_threshold"] = ia.Percent_free_space_threshold
	}
	if &ia.Max_data_disk_count != nil {
		attrs["max_data_disk_count"] = ia.Max_data_disk_count
	}
	if &ia.Absolute_free_space_threshold != nil {
		attrs["absolute_free_space_threshold"] = ia.Absolute_free_space_threshold
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Engine Config Settings
*/
func flattenEngineConfig(ia *model.EngineConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Flavour != nil {
		attrs["flavour"] = ia.Flavour
	}
	if &ia.Hadoop_settings != nil {
		attrs["hadoop_settings"] = flattenHadoopSettings(&ia.Hadoop_settings)
	}
	if &ia.Presto_settings != nil {
		attrs["presto_settings"] = flattenPrestoSettings(&ia.Presto_settings)
	}
	if &ia.Spark_settings != nil {
		attrs["spark_settings"] = flattenSparkSettings(&ia.Spark_settings)
	}
	if &ia.Hive_settings != nil {
		attrs["hive_settings"] = flattenHiveSettings(&ia.Hive_settings)
	}
	if &ia.Airflow_settings != nil {
		attrs["airflow_settings"] = flattenAirflowSettings(&ia.Airflow_settings)
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten fair scheduler Settings
*/
func flattenFairSchedulerSettings(ia *model.FairSchedulerSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Default_pool != nil {
		attrs["default_pool"] = ia.Default_pool
	}
	if &ia.Fairscheduler_config_xml != nil {
		attrs["fairscheduler_config_xml"] = ia.Fairscheduler_config_xml
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten EBS Upscaling Config
*/
func flattenEbsUpscalingConfig(ia *model.EbsUpscalingConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Max_ebs_volume_count != nil {
		attrs["max_ebs_volume_count"] = ia.Max_ebs_volume_count
	}
	if &ia.Percent_free_space_threshold != nil {
		attrs["percent_free_space_threshold"] = ia.Percent_free_space_threshold
	}
	if &ia.Absolute_free_space_threshold != nil {
		attrs["absolute_free_space_threshold"] = ia.Absolute_free_space_threshold
	}
	if &ia.Sampling_interval != nil {
		attrs["sampling_interval"] = ia.Sampling_interval
	}
	if &ia.Sampling_window != nil {
		attrs["sampling_window"] = ia.Sampling_window
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Env Settings
*/
func flattenEnvSettings(ia *model.EnvSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Python_version != nil {
		attrs["python_version"] = ia.Python_version
	}
	if &ia.R_version != nil {
		attrs["r_version"] = ia.R_version
	}
	if &ia.Name != nil {
		attrs["name"] = ia.Name
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Hadoop Settings
*/
func flattenHadoopSettings(ia *model.HadoopSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Use_qubole_placement_policy != nil {
		attrs["use_qubole_placement_policy"] = ia.Use_qubole_placement_policy
	}
	if &ia.Is_ha != nil {
		attrs["is_ha"] = ia.Is_ha
	}
	if &ia.Custom_hadoop_config != nil {
		attrs["custom_hadoop_config"] = ia.Custom_hadoop_config
	}
	if &ia.Fairscheduler_settings != nil {
		attrs["fairscheduler_settings"] = flattenFairSchedulerSettings(&ia.Fairscheduler_settings)
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Heterogeneous Configuration
*/
func flattenHeterogeneousConfig(ia *model.HeterogeneousConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Memory != nil {
		attrs["memory"] = ia.Memory
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Hive Settings
*/
func flattenHiveSettings(ia *model.HiveSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Is_hs2 != nil {
		attrs["is_hs2"] = ia.Is_hs2
	}
	if &ia.Hive_version != nil {
		attrs["hive_version"] = ia.Hive_version
	}
	if &ia.Is_metadata_cache_enabled != nil {
		attrs["is_metadata_cache_enabled"] = ia.Is_metadata_cache_enabled
	}
	if &ia.Hs2_thrift_port != nil {
		attrs["hs2_thrift_port"] = ia.Hs2_thrift_port
	}
	if &ia.Overrides != nil {
		attrs["overrides"] = ia.Overrides
	}
	if &ia.Execution_engine != nil {
		attrs["execution_engine"] = ia.Execution_engine
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Internal Config
*/
func flattenInternal(ia *model.Internal) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Zeppelin_interpreter_mode != nil {
		attrs["zeppelin_interpreter_mode"] = ia.Zeppelin_interpreter_mode
	}
	if &ia.Spark_s3_package_name != nil {
		attrs["spark_s3_package_name"] = ia.Spark_s3_package_name
	}
	if &ia.Zeppelin_s3_package_name != nil {
		attrs["zeppelin_s3_package_name"] = ia.Zeppelin_s3_package_name
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Location Config
*/
func flattenLocation(ia *model.Location) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Aws_region != nil {
		attrs["aws_region"] = ia.Aws_region
	}
	if &ia.Aws_availability_zone != nil {
		attrs["aws_availability_zone"] = ia.Aws_availability_zone
	}
	if &ia.Location != nil {
		attrs["location"] = ia.Location
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Monitoring Config
*/
func flattenMonitoring(ia *model.Monitoring) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Ganglia != nil {
		attrs["ganglia"] = ia.Ganglia
	}
	if &ia.Datadog != nil {
		attrs["datadog"] = flattenDatadog(&ia.Datadog)
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Network Config
*/
func flattenNetworkConfig(ia *model.NetworkConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Vpc_id != nil {
		attrs["vpc_id"] = ia.Vpc_id
	}
	if &ia.Subnet_id != nil {
		attrs["subnet_id"] = ia.Subnet_id
	}
	if &ia.Bastion_node_public_dns != nil {
		attrs["bastion_node_public_dns"] = ia.Bastion_node_public_dns
	}
	if &ia.Bastion_node_port != nil {
		attrs["bastion_node_port"] = ia.Bastion_node_port
	}
	if &ia.Bastion_node_user != nil {
		attrs["bastion_node_user"] = ia.Bastion_node_user
	}
	if &ia.Master_elastic_ip != nil {
		attrs["master_elastic_ip"] = ia.Master_elastic_ip
	}
	if &ia.Persistent_security_groups != nil {
		attrs["persistent_security_groups"] = ia.Persistent_security_groups
	}
	if &ia.Persistent_security_group_resource_group_name != nil {
		attrs["persistent_security_group_resource_group_name"] = ia.Persistent_security_group_resource_group_name
	}
	if &ia.Persistent_security_group_name != nil {
		attrs["persistent_security_group_name"] = ia.Persistent_security_group_name
	}
	if &ia.Vnet_name != nil {
		attrs["vnet_name"] = ia.Vnet_name
	}
	if &ia.Subnet_name != nil {
		attrs["subnet_name"] = ia.Subnet_name
	}
	if &ia.Vnet_resource_group_name != nil {
		attrs["vnet_resource_group_name"] = ia.Vnet_resource_group_name
	}
	if &ia.Master_static_nic_name != nil {
		attrs["master_static_nic_name"] = ia.Master_static_nic_name
	}
	if &ia.Master_static_public_ip_name != nil {
		attrs["master_static_public_ip_name"] = ia.Master_static_public_ip_name
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Presto Settings
*/
func flattenPrestoSettings(ia *model.PrestoSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Presto_version != nil {
		attrs["presto_version"] = ia.Presto_version
	}
	if &ia.Custom_presto_config != nil {
		attrs["custom_presto_config"] = ia.Custom_presto_config
	}
	if &ia.Enable_rubix != nil {
		attrs["enable_rubix"] = ia.Enable_rubix
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Spark Settings
*/
func flattenSparkSettings(ia *model.SparkSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Spark_version != nil {
		attrs["spark_version"] = ia.Spark_version
	}
	if &ia.Custom_spark_config != nil {
		attrs["custom_spark_config"] = ia.Custom_spark_config
	}
	if &ia.Enable_rubix != nil {
		attrs["enable_rubix"] = ia.Enable_rubix
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Spot Block Settings
*/
func flattenSpotBlockSettings(ia *model.SpotBlockSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Duration != nil {
		attrs["duration"] = ia.Duration
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Spot Instance Settings
*/
func flattenSpotInstanceSettings(ia *model.SpotInstanceSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Maximum_bid_price_percentage != nil {
		attrs["maximum_bid_price_percentage"] = ia.Maximum_bid_price_percentage
	}
	if &ia.Timeout_for_request != nil {
		attrs["timeout_for_request"] = ia.Timeout_for_request
	}
	if &ia.Maximum_spot_instance_percentage != nil {
		attrs["maximum_spot_instance_percentage"] = ia.Maximum_spot_instance_percentage
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Spot Settings
*/
func flattenSpotSettings(ia *model.SpotSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Spot_instance_settings != nil {
		attrs["spot_instance_settings"] = flattenSpotInstanceSettings(&ia.Spot_instance_settings)
	}
	if &ia.Spot_block_settings != nil {
		attrs["spot_block_settings"] = flattenSpotBlockSettings(&ia.Spot_block_settings)
	}
	if &ia.Stable_spot_instance_settings != nil {
		attrs["stable_spot_instance_settings"] = flattenStableSpotInstanceSettings(&ia.Stable_spot_instance_settings)
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Stable Spot Instance Settings
*/
func flattenStableSpotInstanceSettings(ia *model.StableSpotInstanceSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Maximum_bid_price_percentage != nil {
		attrs["maximum_bid_price_percentage"] = ia.Maximum_bid_price_percentage
	}
	if &ia.Timeout_for_request != nil {
		attrs["timeout_for_request"] = ia.Timeout_for_request
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Storage Config
*/
func flattenStorageConfig(ia *model.StorageConfig) []map[string]interface{} {
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
		attrs["disk_upscaling_config"] = flattenDiskUpscalingConfig(&ia.Disk_upscaling_config)
	}
	result = append(result, attrs)

	return result
}
