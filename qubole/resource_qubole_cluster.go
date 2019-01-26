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
						"resource_group_name": {
							Type:     schema.TypeString,
							Optional: true,
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
									"upscaling_config": {
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
									"memory": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_type": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"weight": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
											},
										},
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
			
			//Read cloud provider resource group for azure
			if v, ok := configs["resource_group_name"]; ok {
				cloud_config.Resource_group_name = v.(string)
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
					var upscaling_config model.UpscalingConfig
					if v, ok := configs["upscaling_config"]; ok {
						ebsUpscalingConfigs := v.([]interface{})
						if len(ebsUpscalingConfigs) > 0 {
							configs := ebsUpscalingConfigs[0].(map[string]interface{})
							if v, ok := configs["max_ebs_volume_count"]; ok {
								upscaling_config.Max_ebs_volume_count = v.(int)
							}
							if v, ok := configs["percent_free_space_threshold"]; ok {
								upscaling_config.Percent_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["absolute_free_space_threshold"]; ok {
								upscaling_config.Absolute_free_space_threshold = float32(v.(int))
							}
							if v, ok := configs["sampling_interval"]; ok {
								upscaling_config.Sampling_interval = v.(int)
							}
							if v, ok := configs["sampling_window"]; ok {
								upscaling_config.Sampling_window = v.(int)
							}
							datadisk.Upscaling_config = upscaling_config
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
					//level 1 of hetro config is the memory sub-object . there will be only one of these

					if v, ok := configs["memory"]; ok {
						inst_type_array := v.([]interface{})
						if len(inst_type_array) > 0 {
							insts := make([]map[string]interface{}, len(inst_type_array))
							for i, ins := range inst_type_array {
								datamap := ins.(map[string]interface{})
								log.Printf("[DEBUG] PRINTING DATAMAP %s", datamap)

								insts[i] = datamap

							}

							heterogeneous_config.Memory = insts
						}
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
	if err := d.Set("cloud_config", model.FlattenCloudConfig(&cluster_response.Cloud_config)); err != nil {
		log.Printf("[ERR] Error setting cloud config: %#v", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting cloud config: %s", err)
	}
	//Set cluster_info
	if err := d.Set("cluster_info", model.FlattenClusterInfo(&cluster_response.Cluster_info)); err != nil {
		log.Printf("[ERR] Error setting cluster info: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting cluster info: %s", err)
	}

	//Set engine_config
	if err := d.Set("engine_config", model.FlattenEngineConfig(&cluster_response.Engine_config)); err != nil {
		log.Printf("[ERR] Error setting Engine Config: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Engine Config: %s", err)
	}

	//Set monitoring
	if err := d.Set("monitoring", model.FlattenMonitoring(&cluster_response.Monitoring)); err != nil {
		log.Printf("[ERR] Error setting monitoring: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Monitoring: %s", err)
	}

	//Set internal
	if err := d.Set("internal", model.FlattenInternal(&cluster_response.Internal)); err != nil {
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
