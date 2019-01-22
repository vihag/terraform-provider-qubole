package qubole

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	entity "github.com/terraform-providers/terraform-provider-qubole/qubole/entity"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func resourceQubolePresto() *schema.Resource {
	return &schema.Resource{
		Create: resourceQubolePrestoCreate,
		Read:   resourceQubolePrestoRead,
		Update: resourceQubolePrestoUpdate,
		Delete: resourceQubolePrestoDelete,

		Schema: map[string]*schema.Schema{
			"label": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"presto_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"spark_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zeppelin_interpreter_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ec2_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compute_access_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"compute_secret_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"aws_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"aws_preferred_availability_zone": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vpc_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"master_elastic_ip": {
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
						"role_instance_profile": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"use_account_compute_creds": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"instance_tenancy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"compute_validated": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"compute_role_arn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"compute_external_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"node_configuration": {
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
						"slave_request_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cluster_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"heterogeneous_instance_config": {
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
						"initial_nodes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_nodes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
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
						"fallback_to_ondemand": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						/*
							1. The default EBS volume type is standard (magnetic). The other possible values are
							2. ssd (gp2, General Purpose SSD),
							3. st1 (Throughput Optimized HDD), and
							4. sc1 (Cold HDD)
						*/
						"ebs_volume_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						/*
							1. For standard (magnetic) EBS volume type, the supported value range is 100 GB to 1 TB.
							2. For ssd/gp2 (General Purpose SSD) EBS volume type, the supported value range is 100 GB to 16 GB.
							3. For st1 (Throughput Optimized HDD) and sc1 (Cold HDD), the supported value range is 500 GB to 16 TB.
						*/
						"ebs_volume_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ebs_volume_count": {
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
						/*
							1. Json of key value pairs.
						*/
						//TODO better schema representation of this JSON instead of String
						"custom_ec2_tags": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"idle_cluster_timeout_in_secs": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"node_base_cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"node_spot_cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"parent_cluster_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"child_hs2_cluster_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"hadoop_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_hadoop2": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"use_spark": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"use_hbase": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"use_qubole_placement_policy": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"is_ha": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_rubix": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"node_bootstrap_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"custom_config": {
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
			"security_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encrypted_ephemerals": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"customer_ssh_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"persistent_security_group": {
							Type:     schema.TypeString,
							Optional: true,
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
						"enable_presto": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"custom_config": {
							Type:     schema.TypeString,
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
						"custom_config": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"datadog_settings": {
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
			"disallow_cluster_termination": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_ganglia_monitoring": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"node_bootstrap_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idle_cluster_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tunnel_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_config": {
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
						"engine_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"overrides": {
							Type:     schema.TypeString,
							Optional: true,
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
									"hive_qubole_metadata_cache": {
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
								},
							},
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
func readEc2SettingsFromTf(d *schema.ResourceData) (entity.Ec2Settings, bool) {

	var ec2_settings entity.Ec2Settings
	if v, ok := d.GetOk("ec2_settings"); ok {
		ec2Settings := v.([]interface{})
		if len(ec2Settings) > 0 {
			settings := ec2Settings[0].(map[string]interface{})
			if v, ok := settings["compute_access_key"]; ok {
				ec2_settings.Compute_access_key = v.(string)
			}
			if v, ok := settings["compute_secret_key"]; ok {
				ec2_settings.Compute_secret_key = v.(string)
			}
			if v, ok := settings["aws_region"]; ok {
				ec2_settings.Aws_region = v.(string)
			}
			if v, ok := settings["aws_preferred_availability_zone"]; ok {
				ec2_settings.Aws_preferred_availability_zone = v.(string)
			}
			if v, ok := settings["vpc_id"]; ok {
				ec2_settings.Vpc_id = v.(string)
			}
			if v, ok := settings["subnet_id"]; ok {
				ec2_settings.Subnet_id = v.(string)
			}
			if v, ok := settings["bastion_node_public_dns"]; ok {
				ec2_settings.Bastion_node_public_dns = v.(string)
			}
			if v, ok := settings["bastion_node_port"]; ok {
				ec2_settings.Bastion_node_port = v.(int)
			}
			if v, ok := settings["bastion_node_user"]; ok {
				ec2_settings.Bastion_node_user = v.(string)
			}
			if v, ok := settings["role_instance_profile"]; ok {
				ec2_settings.Role_instance_profile = v.(string)
			}
			if v, ok := settings["use_account_compute_creds"]; ok {
				ec2_settings.Use_account_compute_creds = v.(bool)
			}
			return ec2_settings, true
		}
	}
	//the reading method needs to check for the boolean variable to see if all was okay
	return ec2_settings, false
}

func readNodeConfigurationFromTf(d *schema.ResourceData) (entity.NodeConfiguration, bool) {
	var node_configuration entity.NodeConfiguration
	if v, ok := d.GetOk("node_configuration"); ok {
		nodeConfigurations := v.([]interface{})
		if len(nodeConfigurations) > 0 {
			configurations := nodeConfigurations[0].(map[string]interface{})

			if v, ok := configurations["master_instance_type"]; ok {
				node_configuration.Master_instance_type = v.(string)
			}

			if v, ok := configurations["slave_instance_type"]; ok {
				node_configuration.Slave_instance_type = v.(string)
			}

			//Get Heterogeneous Instance Config
			var heterogeneous_instance_config entity.HeterogeneousInstanceConfig
			if v, ok := configurations["heterogeneous_instance_config"]; ok {
				heterogeneousInstanceConfigs := v.([]interface{})
				if len(heterogeneousInstanceConfigs) > 0 {
					heteroInstanceConfigs := heterogeneousInstanceConfigs[0].(map[string]interface{})
					if v, ok := heteroInstanceConfigs["memory"]; ok {
						heterogeneous_instance_config.Memory = v.(string)
					}
					node_configuration.Heterogeneous_instance_config = heterogeneous_instance_config
				}

			}

			//Get Spot Instance Settings
			var spot_instance_settings entity.SpotInstanceSettings
			if v, ok := configurations["spot_instance_settings"]; ok {
				spotInstanceSettings := v.([]interface{})
				if len(spotInstanceSettings) > 0 {
					spotInstSettings := spotInstanceSettings[0].(map[string]interface{})
					if v, ok := spotInstSettings["maximum_bid_price_percentage"]; ok {
						spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
					}
					if v, ok := spotInstSettings["timeout_for_request"]; ok {
						spot_instance_settings.Timeout_for_request = v.(int)
					}
					if v, ok := spotInstSettings["maximum_spot_instance_percentage"]; ok {
						spot_instance_settings.Maximum_spot_instance_percentage = float32(v.(int))
					}
					node_configuration.Spot_instance_settings = spot_instance_settings
				}
			}

			//Get Stable Spot Settings
			var stable_spot_instance_settings entity.StableSpotInstanceSettings
			if v, ok := configurations["stable_spot_instance_settings"]; ok {
				stableSpotInstanceSettings := v.([]interface{})
				if len(stableSpotInstanceSettings) > 0 {
					stableSpotInstSettings := stableSpotInstanceSettings[0].(map[string]interface{})
					if v, ok := stableSpotInstSettings["maximum_bid_price_percentage"]; ok {
						stable_spot_instance_settings.Maximum_bid_price_percentage = float32(v.(int))
					}
					if v, ok := stableSpotInstSettings["timeout_for_request"]; ok {
						stable_spot_instance_settings.Timeout_for_request = v.(int)
					}
					node_configuration.Stable_spot_instance_settings = stable_spot_instance_settings
				}
			}

			//Get Spot Block Settings
			var spot_block_settings entity.SpotBlockSettings
			if v, ok := configurations["spot_block_settings"]; ok {
				spotBlockSettings := v.([]interface{})
				if len(spotBlockSettings) > 0 {
					blockSettings := spotBlockSettings[0].(map[string]interface{})
					if v, ok := blockSettings["duration"]; ok {
						spot_block_settings.Duration = v.(int)
					}
					node_configuration.Spot_block_settings = spot_block_settings
				}
			}

			//Ebs Upscaling Config
			var ebs_upscaling_config entity.EbsUpscalingConfig
			if v, ok := configurations["ebs_upscaling_config"]; ok {
				ebsUpscalingConfig := v.([]interface{})
				if len(ebsUpscalingConfig) > 0 {
					upscalingConfig := ebsUpscalingConfig[0].(map[string]interface{})
					if v, ok := upscalingConfig["max_ebs_volume_count"]; ok {
						ebs_upscaling_config.Max_ebs_volume_count = v.(int)
					}
					if v, ok := upscalingConfig["percent_free_space_threshold"]; ok {
						ebs_upscaling_config.Percent_free_space_threshold = v.(int)
					}
					if v, ok := upscalingConfig["absolute_free_space_threshold"]; ok {
						ebs_upscaling_config.Absolute_free_space_threshold = v.(int)
					}
					if v, ok := upscalingConfig["sampling_interval"]; ok {
						ebs_upscaling_config.Sampling_interval = v.(int)
					}
					if v, ok := upscalingConfig["sampling_window"]; ok {
						ebs_upscaling_config.Sampling_window = v.(int)
					}
					node_configuration.Ebs_upscaling_config = ebs_upscaling_config
				}
			}

			if v, ok := configurations["initial_nodes"]; ok {
				node_configuration.Initial_nodes = v.(int)
			}

			if v, ok := configurations["max_nodes"]; ok {
				node_configuration.Max_nodes = v.(int)
			}

			if v, ok := configurations["slave_request_type"]; ok {
				node_configuration.Slave_request_type = v.(string)
			}

			if v, ok := configurations["fallback_to_ondemand"]; ok {
				node_configuration.Fallback_to_ondemand = v.(bool)
			}

			if v, ok := configurations["ebs_volume_type"]; ok {
				node_configuration.Ebs_volume_type = v.(string)
			}

			if v, ok := configurations["ebs_volume_size"]; ok {
				node_configuration.Ebs_volume_size = v.(int)
			}

			if v, ok := configurations["ebs_volume_count"]; ok {
				node_configuration.Ebs_volume_count = v.(int)
			}

			if v, ok := configurations["custom_ec2_tags"]; ok {
				//node_configuration.Custom_ec2_tags = v.(map[string]string)
				ec2_tags := v.(map[string]interface{})
				custom_ec2_tags := make(map[string]string)
				for key, value := range ec2_tags {
					strKey := fmt.Sprintf("%v", key)
					strValue := fmt.Sprintf("%v", value)

					custom_ec2_tags[strKey] = strValue
				}
				node_configuration.Custom_ec2_tags = custom_ec2_tags
			}

			if v, ok := configurations["idle_cluster_timeout_in_secs"]; ok {
				node_configuration.Idle_cluster_timeout_in_secs = v.(int)
			}

			if v, ok := configurations["node_base_cooldown_period"]; ok {
				node_configuration.Node_base_cooldown_period = v.(int)
			}

			if v, ok := configurations["node_spot_cooldown_period"]; ok {
				node_configuration.Node_spot_cooldown_period = v.(int)
			}
			return node_configuration, true
		}
	}
	return node_configuration, false
}

func readHadoopSettingsFromTf(d *schema.ResourceData) (entity.HadoopSettings, bool) {
	var hadoop_settings entity.HadoopSettings
	if v, ok := d.GetOk("hadoop_settings"); ok {
		hadoopSettings := v.([]interface{})
		if len(hadoopSettings) > 0 {
			hdSettings := hadoopSettings[0].(map[string]interface{})

			if v, ok := hdSettings["use_hadoop2"]; ok {
				hadoop_settings.Use_hadoop2 = v.(bool)
			}

			if v, ok := hdSettings["use_spark"]; ok {
				hadoop_settings.Use_spark = v.(bool)
			}

			if v, ok := hdSettings["use_hbase"]; ok {
				hadoop_settings.Use_hbase = v.(bool)
			}

			if v, ok := hdSettings["is_ha"]; ok {
				hadoop_settings.Is_ha = v.(bool)
			}

			if v, ok := hdSettings["enable_rubix"]; ok {
				hadoop_settings.Enable_rubix = v.(bool)
			}

			if v, ok := hdSettings["node_bootstrap_timeout"]; ok {
				hadoop_settings.Node_bootstrap_timeout = v.(int)
			}

			if v, ok := hdSettings["use_qubole_placement_policy"]; ok {
				hadoop_settings.Use_qubole_placement_policy = v.(bool)
			}

			//Fair Scheduler Settings
			var fairscheduler_settings entity.FairSchedulerSettings
			if v, ok := hdSettings["spot_block_settings"]; ok {
				fairSchedulerSettings := v.([]interface{})
				if len(fairSchedulerSettings) > 0 {
					fsSettings := fairSchedulerSettings[0].(map[string]interface{})
					if v, ok := fsSettings["default_pool"]; ok {
						fairscheduler_settings.Default_pool = v.(string)
					}
					if v, ok := fsSettings["fairscheduler_config_xml"]; ok {
						fairscheduler_settings.Fairscheduler_config_xml = v.(string)
					}
					hadoop_settings.Fairscheduler_settings = fairscheduler_settings
				}
			}
			return hadoop_settings, true
		}
	}
	return hadoop_settings, false
}

func readSecuritySettingsFromTf(d *schema.ResourceData) (entity.SecuritySettings, bool) {
	var security_settings entity.SecuritySettings
	if v, ok := d.GetOk("security_settings"); ok {
		securitySettings := v.([]interface{})
		if len(securitySettings) > 0 {
			secSettings := securitySettings[0].(map[string]interface{})

			if v, ok := secSettings["encrypted_ephemerals"]; ok {
				security_settings.Encrypted_ephemerals = v.(bool)
			}

			if v, ok := secSettings["customer_ssh_key"]; ok {
				security_settings.Customer_ssh_key = v.(string)
			}

			if v, ok := secSettings["persistent_security_group"]; ok {
				security_settings.Persistent_security_group = v.(string)
			}
			return security_settings, true
		}
	}
	return security_settings, false
}

func readPrestoSettingsFromTf(d *schema.ResourceData) (entity.PrestoSettings, bool) {
	var presto_settings entity.PrestoSettings
	if v, ok := d.GetOk("presto_settings"); ok {
		prestoSettings := v.([]interface{})
		if len(prestoSettings) > 0 {
			pSettings := prestoSettings[0].(map[string]interface{})

			if v, ok := pSettings["enable_presto"]; ok {
				presto_settings.Enable_presto = v.(bool)
			}

			if v, ok := pSettings["custom_config"]; ok {
				presto_settings.Custom_config = v.(string)
			}
			return presto_settings, true
		}
	}
	return presto_settings, false
}

func readSparkSettingsFromTf(d *schema.ResourceData) (entity.SparkSettings, bool) {
	var spark_settings entity.SparkSettings
	if v, ok := d.GetOk("spark_settings"); ok {
		sparkSettings := v.([]interface{})
		if len(sparkSettings) > 0 {
			sSettings := sparkSettings[0].(map[string]interface{})

			if v, ok := sSettings["custom_config"]; ok {
				spark_settings.Custom_config = v.(string)
			}
			return spark_settings, true
		}
	}
	return spark_settings, false
}

func readDatadogSettingsFromTf(d *schema.ResourceData) (entity.DatadogSettings, bool) {
	var datadog_settings entity.DatadogSettings
	if v, ok := d.GetOk("datadog_settings"); ok {
		datadogSettings := v.([]interface{})
		if len(datadogSettings) > 0 {
			ddSettings := datadogSettings[0].(map[string]interface{})

			if v, ok := ddSettings["datadog_api_token"]; ok {
				datadog_settings.Datadog_api_token = v.(string)
			}

			if v, ok := ddSettings["datadog_app_token"]; ok {
				datadog_settings.Datadog_app_token = v.(string)
			}
			return datadog_settings, true
		}
	}
	return datadog_settings, false
}

func readEngineConfigFromTf(d *schema.ResourceData) (entity.EngineConfig, bool) {
	var engine_config entity.EngineConfig
	if v, ok := d.GetOk("engine_config"); ok {
		engineConfigs := v.([]interface{})
		if len(engineConfigs) > 0 {
			eConfigs := engineConfigs[0].(map[string]interface{})

			if v, ok := eConfigs["dbtap_id"]; ok {
				engine_config.Dbtap_id = v.(int)
			}
			if v, ok := eConfigs["fernet_key"]; ok {
				engine_config.Fernet_key = v.(string)
			}
			if v, ok := eConfigs["engine_type"]; ok {
				engine_config.Engine_type = v.(string)
			}
			if v, ok := eConfigs["version"]; ok {
				engine_config.Version = v.(string)
			}
			if v, ok := eConfigs["overrides"]; ok {
				engine_config.Overrides = v.(string)
			}
			//Hive Settings
			var hive_settings entity.HiveSettings
			if v, ok := eConfigs["hive_settings"]; ok {
				hiveSettings := v.([]interface{})
				if len(hiveSettings) > 0 {
					hvSettings := hiveSettings[0].(map[string]interface{})
					if v, ok := hvSettings["is_hs2"]; ok {
						hive_settings.Is_hs2 = v.(bool)
					}
					if v, ok := hvSettings["hive_version"]; ok {
						hive_settings.Hive_version = v.(string)
					}
					if v, ok := hvSettings["hive_qubole_metadata_cache"]; ok {
						hive_settings.Hive_qubole_metadata_cache = v.(bool)
					}
					if v, ok := hvSettings["hs2_thrift_port"]; ok {
						hive_settings.Hs2_thrift_port = v.(int)
					}
					if v, ok := hvSettings["overrides"]; ok {
						hive_settings.Overrides = v.(string)
					}
					engine_config.Hive_settings = hive_settings
				}
			}
			return engine_config, true
		}
	}
	return engine_config, false
}

func readClusterFromTf(d *schema.ResourceData) (entity.Cluster, bool) {

	//Create the representative json object here
	var cluster entity.Cluster

	//create nested datas structures
	//1. EC2 Settings
	if ec2_settings, ok := readEc2SettingsFromTf(d); ok {
		cluster.Ec2_settings = ec2_settings
	} else {
		log.Printf("[WARN] No ec2_settings seen.")
	}

	//2. Node configuration, but this will require constructing other sub-objects
	if node_configuration, ok := readNodeConfigurationFromTf(d); ok {
		cluster.Node_configuration = node_configuration
	} else {
		log.Printf("[WARN] No node_configuration seen.")
	}

	//3. Hadoop Settings, but this will require constructing other sub-objects
	if hadoop_settings, ok := readHadoopSettingsFromTf(d); ok {
		cluster.Hadoop_settings = hadoop_settings
	} else {
		log.Printf("[WARN] No hadoop_settings seen.")
	}

	//4. SecuritySettings
	if security_settings, ok := readSecuritySettingsFromTf(d); ok {
		cluster.Security_settings = security_settings
	} else {
		log.Printf("[WARN] No security_settings seen.")
	}

	//5. PrestoSettings
	if presto_settings, ok := readPrestoSettingsFromTf(d); ok {
		cluster.Presto_settings = presto_settings
	} else {
		log.Printf("[WARN] No presto_settings seen.")
	}

	//6. SparkSettings
	if spark_settings, ok := readSparkSettingsFromTf(d); ok {
		cluster.Spark_settings = spark_settings
	} else {
		log.Printf("[WARN] No spark_settings seen.")
	}

	//7. DatadogSettings
	if datadog_settings, ok := readDatadogSettingsFromTf(d); ok {
		cluster.Datadog_settings = datadog_settings
	} else {
		log.Printf("[WARN] No datadog_settings seen.")
	}

	//8. EngineConfig
	if engine_config, ok := readEngineConfigFromTf(d); ok {
		cluster.Engine_config = engine_config
	} else {
		log.Printf("[WARN] No engine_config seen.")
	}

	//Finally, the cluster

	if v, ok := d.GetOk("label"); ok {
		labelSet := v.(*schema.Set)
		labels := make([]*string, labelSet.Len())
		for i, label := range labelSet.List() {
			labels[i] = String(label.(string))
		}
		cluster.Label = labels
	}
	if v, ok := d.GetOk("presto_version"); ok {
		cluster.Presto_version = v.(string)
	}
	if v, ok := d.GetOk("spark_version"); ok {
		cluster.Spark_version = v.(string)
	}
	if v, ok := d.GetOk("zeppelin_interpreter_mode"); ok {
		cluster.Zeppelin_interpreter_mode = v.(string)
	}
	if v, ok := d.GetOk("disallow_cluster_termination"); ok {
		cluster.Disallow_cluster_termination = v.(bool)
	}
	if v, ok := d.GetOk("enable_ganglia_monitoring"); ok {
		cluster.Enable_ganglia_monitoring = v.(bool)
	}
	if v, ok := d.GetOk("node_bootstrap_file"); ok {
		cluster.Node_bootstrap_file = v.(string)
	}
	if v, ok := d.GetOk("idle_cluster_timeout"); ok {
		cluster.Idle_cluster_timeout = v.(int)
	}

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
func resourceQubolePrestoCreate(d *schema.ResourceData, meta interface{}) error {

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
		var clusterResponse entity.Cluster
		unmarshallingError := json.Unmarshal(body, &clusterResponse)
		if unmarshallingError != nil {
			log.Printf("[ERR]There was an error:", unmarshallingError.Error())
			return fmt.Errorf("Error in unmarshalling json during update %s", err.Error())
		}
		log.Printf("[INFO]Pretty Printing Unmarshalled Response %#v", clusterResponse)

		//Set Terraform ID; typecast the received ID to string for terraform
		d.SetId(strconv.Itoa(clusterResponse.Id))

	} else {
		log.Printf("[WARN] No valid cluster definition seen.")
	}

	return resourceQubolePrestoRead(d, meta)
}

/*
The Read callback is used to sync the local state with the actual state (upstream). This is called at various points by Terraform and should be a read-only operation.
This callback should never modify the real resource.
*/
func resourceQubolePrestoRead(d *schema.ResourceData, meta interface{}) error {

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
	var cluster entity.Cluster
	unmarshallingError := json.Unmarshal(body, &cluster)
	if unmarshallingError != nil {
		log.Printf("There was an error:", unmarshallingError.Error())
	}
	log.Printf(">>>Pretty Printing Unmarshalled Response")
	log.Printf("%#v", cluster)

	//Now start setting d with data from the unmarshalled object
	//Set EC2 settings
	if err := d.Set("ec2_settings", flattenEc2Settings(&cluster.Ec2_settings)); err != nil {
		log.Printf("[ERR] Error setting EC2 Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting EC2 Settings: %s", err)
	}
	//Set Hadoop settings
	if err := d.Set("hadoop_settings", flattenHadoopSettings(&cluster.Hadoop_settings)); err != nil {
		log.Printf("[ERR] Error setting Hadoop Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Hadoop Settings: %s", err)
	}

	//Set Security settings
	if err := d.Set("security_settings", flattenSecuritySettings(&cluster.Security_settings)); err != nil {
		log.Printf("[ERR] Error setting Security Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Security Settings: %s", err)
	}

	//Set Presto settings
	if err := d.Set("presto_settings", flattenPrestoSettings(&cluster.Presto_settings)); err != nil {
		log.Printf("[ERR] Error setting Presto Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Presto Settings: %s", err)
	}

	//Set Spark settings
	if err := d.Set("spark_settings", flattenSparkSettings(&cluster.Spark_settings)); err != nil {
		log.Printf("[ERR] Error setting Spark Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Spark Settings: %s", err)
	}

	//Set Datadog settings
	if err := d.Set("datadog_settings", flattenDatadogSettings(&cluster.Datadog_settings)); err != nil {
		log.Printf("[ERR] Error setting Datadog Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Datadog Settings: %s", err)
	}

	//Set Node Configuration
	if err := d.Set("node_configuration", flattenNodeConfiguration(&cluster.Node_configuration)); err != nil {
		log.Printf("[ERR] Error setting Node Configuration Settings: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Node Configurations: %s", err)
	}

	//Set Engine Configuration
	if err := d.Set("engine_config", flattenEngineConfig(&cluster.Engine_config)); err != nil {
		log.Printf("[ERR] Error setting Engine Config: %s", err)
		d.SetId("")
		return fmt.Errorf("[ERR] Error setting Engine Config: %s", err)
	}

	//Set rest of the simple objects
	d.Set("state", cluster.State)
	d.Set("force_tunnel", cluster.Force_tunnel)
	d.Set("tunnel_server_ip", cluster.Tunnel_server_ip)
	d.Set("label", cluster.Label)
	d.Set("presto_version", cluster.Presto_version)
	d.Set("spark_version", cluster.Spark_version)
	d.Set("zeppelin_interpreter_mode", cluster.Zeppelin_interpreter_mode)
	d.Set("disallow_cluster_termination", cluster.Disallow_cluster_termination)
	d.Set("enable_ganglia_monitoring", cluster.Enable_ganglia_monitoring)
	d.Set("node_bootstrap_file", cluster.Spark_version)
	d.Set("idle_cluster_timeout", cluster.Zeppelin_interpreter_mode)
	d.Set("spark_s3_package_name", cluster.Disallow_cluster_termination)
	d.Set("zeppelin_s3_package_name", cluster.Enable_ganglia_monitoring)

	return nil
}

/*
1. If the Update callback returns with or without an error, the full state is saved.
	If the ID becomes blank, the resource is destroyed (even within an update, though this shouldn't happen except in error scenarios).
*/
func resourceQubolePrestoUpdate(d *schema.ResourceData, meta interface{}) error {
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

		//Parse the response back to cluster object
		var clusterResponse entity.Cluster
		unmarshallingError := json.Unmarshal(body, &clusterResponse)
		if unmarshallingError != nil {
			log.Printf("[ERR]There was an error:", unmarshallingError.Error())
			return fmt.Errorf("Error in unmarshalling json during update %s", err.Error())
		}
		log.Printf("[INFO]Pretty Printing Unmarshalled Response %#v", clusterResponse)

		//Set Terraform ID; typecast the received ID to string for terraform
		d.SetId(strconv.Itoa(clusterResponse.Id))

	} else {
		log.Printf("[WARN] No valid cluster definition seen.")
	}

	return resourceQubolePrestoRead(d, meta)
}

/*
1. If the Destroy callback returns without an error, the resource is assumed to be destroyed, and all state is removed.

2. If the Destroy callback returns with an error, the resource is assumed to still exist, and all prior state is preserved.
*/
func resourceQubolePrestoDelete(d *schema.ResourceData, meta interface{}) error {
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
function to flatten EC2 settings to the schema that is defined
*/
func flattenEc2Settings(ia *entity.Ec2Settings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Compute_access_key != nil {
		attrs["compute_access_key"] = ia.Compute_access_key
	}
	if &ia.Compute_secret_key != nil {
		attrs["compute_secret_key"] = ia.Compute_secret_key
	}
	if &ia.Aws_region != nil {
		attrs["aws_region"] = ia.Aws_region
	}
	if &ia.Aws_preferred_availability_zone != nil {
		attrs["aws_preferred_availability_zone"] = ia.Aws_preferred_availability_zone
	}
	if &ia.Vpc_id != nil {
		attrs["vpc_id"] = ia.Vpc_id
	}
	if &ia.Subnet_id != nil {
		attrs["subnet_id"] = ia.Subnet_id
	}
	if &ia.Master_elastic_ip != nil {
		attrs["master_elastic_ip"] = ia.Master_elastic_ip
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
	if &ia.Role_instance_profile != nil {
		attrs["role_instance_profile"] = ia.Role_instance_profile
	}
	if &ia.Use_account_compute_creds != nil {
		attrs["use_account_compute_creds"] = ia.Use_account_compute_creds
	}
	if &ia.Compute_validated != nil {
		attrs["compute_validated"] = ia.Compute_validated
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

	result = append(result, attrs)

	return result
}

/*
function to flatten Heterogenous Instance Config
*/
func flattenHeterogenousInstanceConfig(ia *entity.HeterogeneousInstanceConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Memory != nil {
		attrs["memory"] = ia.Memory
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Spot Instance Settings
*/
func flattenSpotInstanceSettings(ia *entity.SpotInstanceSettings) []map[string]interface{} {
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
function to flatten Stable Spot Instance Settings
*/
func flattenStableSpotInstanceSettings(ia *entity.StableSpotInstanceSettings) []map[string]interface{} {
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
function to flatten Stable Spot Instance Settings
*/
func flattenSpotBlockSettings(ia *entity.SpotBlockSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Duration != nil {
		attrs["duration"] = ia.Duration
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Spark Settings
*/
func flattenSparkSettings(ia *entity.SparkSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Custom_config != nil {
		attrs["custom_config"] = ia.Custom_config
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Security Settings
*/
func flattenSecuritySettings(ia *entity.SecuritySettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Encrypted_ephemerals != nil {
		attrs["encrypted_ephemerals"] = ia.Encrypted_ephemerals
	}
	if &ia.Customer_ssh_key != nil {
		attrs["customer_ssh_key"] = ia.Customer_ssh_key
	}
	if &ia.Persistent_security_group != nil {
		attrs["persistent_security_group"] = ia.Persistent_security_group
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Presto Settings
*/
func flattenPrestoSettings(ia *entity.PrestoSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Enable_presto != nil {
		attrs["enable_presto"] = ia.Enable_presto
	}
	if &ia.Custom_config != nil {
		attrs["custom_config"] = ia.Custom_config
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Hive Settings
*/
func flattenHiveSettings(ia *entity.HiveSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Is_hs2 != nil {
		attrs["is_hs2"] = ia.Is_hs2
	}
	if &ia.Hive_version != nil {
		attrs["hive_version"] = ia.Hive_version
	}
	if &ia.Hive_qubole_metadata_cache != nil {
		attrs["hive_qubole_metadata_cache"] = ia.Hive_qubole_metadata_cache
	}
	if &ia.Hs2_thrift_port != nil {
		attrs["hs2_thrift_port"] = ia.Hs2_thrift_port
	}
	if &ia.Overrides != nil {
		attrs["overrides"] = ia.Overrides
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Fair Scheduler Settings
*/
func flattenFairSchedulerSettings(ia *entity.FairSchedulerSettings) []map[string]interface{} {
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
func flattenEbsUpscalingConfig(ia *entity.EbsUpscalingConfig) []map[string]interface{} {
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
function to flatten Datadog Settings
*/
func flattenDatadogSettings(ia *entity.DatadogSettings) []map[string]interface{} {
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
function to flatten Hadoop Settings
*/
func flattenHadoopSettings(ia *entity.HadoopSettings) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Use_hadoop2 != nil {
		attrs["use_hadoop2"] = ia.Use_hadoop2
	}
	if &ia.Use_spark != nil {
		attrs["use_spark"] = ia.Use_spark
	}
	if &ia.Use_hbase != nil {
		attrs["use_hbase"] = ia.Use_hbase
	}
	if &ia.Use_qubole_placement_policy != nil {
		attrs["use_qubole_placement_policy"] = ia.Use_qubole_placement_policy
	}
	if &ia.Is_ha != nil {
		attrs["is_ha"] = ia.Is_ha
	}
	if &ia.Enable_rubix != nil {
		attrs["enable_rubix"] = ia.Enable_rubix
	}
	if &ia.Node_bootstrap_timeout != nil {
		attrs["node_bootstrap_timeout"] = ia.Node_bootstrap_timeout
	}
	if &ia.Custom_config != nil {
		attrs["custom_config"] = ia.Custom_config
	}
	if &ia.Fairscheduler_settings != nil {
		attrs["fairscheduler_settings"] = flattenFairSchedulerSettings(&ia.Fairscheduler_settings)
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Node Configuration
*/
func flattenNodeConfiguration(ia *entity.NodeConfiguration) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Master_instance_type != nil {
		attrs["master_instance_type"] = ia.Master_instance_type
	}
	if &ia.Slave_instance_type != nil {
		attrs["slave_instance_type"] = ia.Slave_instance_type
	}
	if &ia.Heterogeneous_instance_config != nil {
		attrs["heterogeneous_instance_config"] = flattenHeterogenousInstanceConfig(&ia.Heterogeneous_instance_config)
	}
	if &ia.Initial_nodes != nil {
		attrs["initial_nodes"] = ia.Initial_nodes
	}
	if &ia.Max_nodes != nil {
		attrs["max_nodes"] = ia.Max_nodes
	}
	if &ia.Slave_request_type != nil {
		attrs["slave_request_type"] = ia.Slave_request_type
	}
	if &ia.Spot_instance_settings != nil {
		attrs["spot_instance_settings"] = flattenSpotInstanceSettings(&ia.Spot_instance_settings)
	}
	if &ia.Stable_spot_instance_settings != nil {
		attrs["stable_spot_instance_settings"] = flattenStableSpotInstanceSettings(&ia.Stable_spot_instance_settings)
	}
	if &ia.Spot_block_settings != nil {
		attrs["spot_block_settings"] = flattenSpotBlockSettings(&ia.Spot_block_settings)
	}
	if &ia.Fallback_to_ondemand != nil {
		attrs["fallback_to_ondemand"] = ia.Fallback_to_ondemand
	}
	if &ia.Ebs_volume_type != nil {
		attrs["ebs_volume_type"] = ia.Ebs_volume_type
	}
	if &ia.Ebs_volume_size != nil {
		attrs["ebs_volume_size"] = ia.Ebs_volume_size
	}
	if &ia.Ebs_volume_count != nil {
		attrs["ebs_volume_count"] = ia.Ebs_volume_count
	}
	if &ia.Ebs_upscaling_config != nil {
		attrs["ebs_upscaling_config"] = flattenEbsUpscalingConfig(&ia.Ebs_upscaling_config)
	}
	if &ia.Custom_ec2_tags != nil {
		attrs["custom_ec2_tags"] = ia.Custom_ec2_tags
	}
	if &ia.Idle_cluster_timeout_in_secs != nil {
		attrs["idle_cluster_timeout_in_secs"] = ia.Idle_cluster_timeout_in_secs
	}
	if &ia.Node_base_cooldown_period != nil {
		attrs["node_base_cooldown_period"] = ia.Node_base_cooldown_period
	}
	if &ia.Node_spot_cooldown_period != nil {
		attrs["node_spot_cooldown_period"] = ia.Node_spot_cooldown_period
	}
	if &ia.Child_hs2_cluster_id != nil {
		attrs["child_hs2_cluster_id"] = ia.Child_hs2_cluster_id
	}
	if &ia.Parent_cluster_id != nil {
		attrs["parent_cluster_id"] = ia.Parent_cluster_id
	}
	if &ia.Cluster_name != nil {
		attrs["cluster_name"] = ia.Cluster_name
	}

	result = append(result, attrs)

	return result
}

/*
function to flatten Engine Config
*/
func flattenEngineConfig(ia *entity.EngineConfig) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Dbtap_id != nil {
		attrs["dbtap_id"] = ia.Dbtap_id
	}
	if &ia.Fernet_key != nil {
		attrs["fernet_key"] = ia.Fernet_key
	}
	if &ia.Engine_type != nil {
		attrs["engine_type"] = ia.Engine_type
	}
	if &ia.Version != nil {
		attrs["version"] = ia.Version
	}
	if &ia.Overrides != nil {
		attrs["overrides"] = ia.Overrides
	}
	if &ia.Hive_settings != nil {
		attrs["hive_settings"] = flattenHiveSettings(&ia.Hive_settings)
	}

	result = append(result, attrs)

	return result
}
