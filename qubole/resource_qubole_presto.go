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
						"use_qubole_placement_policy": {
							Type:     schema.TypeBool,
							Optional: true,
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
		},
	}
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

	//Create the representative json object here
	var cluster entity.Cluster

	//create nested datas structures
	//1. EC2 Settings
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
			cluster.Ec2_settings = ec2_settings
		}
	}

	//2. Node configuration, but this will require constructing other sub-objects
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
						spot_instance_settings.Maximum_bid_price_percentage = v.(int)
					}
					if v, ok := spotInstSettings["timeout_for_request"]; ok {
						spot_instance_settings.Timeout_for_request = v.(int)
					}
					if v, ok := spotInstSettings["maximum_spot_instance_percentage"]; ok {
						spot_instance_settings.Maximum_spot_instance_percentage = v.(int)
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
						stable_spot_instance_settings.Maximum_bid_price_percentage = v.(int)
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
			cluster.Node_configuration = node_configuration
		}
	}

	//3. Hadoop Settings, but this will require constructing other sub-objects
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
			cluster.Hadoop_settings = hadoop_settings
		}
	}

	//4. SecuritySettings
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
			cluster.Security_settings = security_settings
		}
	}

	//5. PrestoSettings
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
			cluster.Presto_settings = presto_settings
		}
	}

	//6. SparkSettings
	var spark_settings entity.SparkSettings
	if v, ok := d.GetOk("spark_settings"); ok {
		sparkSettings := v.([]interface{})
		if len(sparkSettings) > 0 {
			sSettings := sparkSettings[0].(map[string]interface{})

			if v, ok := sSettings["custom_config"]; ok {
				spark_settings.Custom_config = v.(string)
			}
			cluster.Spark_settings = spark_settings
		}
	}

	//7. DatadogSettings
	var datadog_settings entity.DatadogSettings
	if v, ok := d.GetOk("presto_settings"); ok {
		datadogSettings := v.([]interface{})
		if len(datadogSettings) > 0 {
			ddSettings := datadogSettings[0].(map[string]interface{})

			if v, ok := ddSettings["datadog_api_token"]; ok {
				datadog_settings.Datadog_api_token = v.(string)
			}

			if v, ok := ddSettings["datadog_app_token"]; ok {
				datadog_settings.Datadog_app_token = v.(string)
			}
			cluster.Datadog_settings = datadog_settings
		}
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

	//log the json constructed
	log.Printf("printing json.....")
	cluster_json, err := json.Marshal(cluster)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	log.Printf(string(cluster_json))

	//Make the http call to api here
	log.Printf("Sending Create Cluster Request to URI %s", api_url)
	var payload = []byte(string(cluster_json))
	req, err := http.NewRequest("POST", api_url, bytes.NewBuffer(payload))
	req.Header.Set("X-AUTH-TOKEN", auth_token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf(err.Error())
	}
	defer resp.Body.Close()

	log.Printf("response Status:", resp.Status)
	log.Printf("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("response Body:", string(body))

	//Parse the response back to cluster object

	return resourceQubolePrestoRead(d, meta)
}

/*
The Read callback is used to sync the local state with the actual state (upstream). This is called at various points by Terraform and should be a read-only operation.
This callback should never modify the real resource.
client := m.(*MyClient)
// Attempt to read from an upstream API
  obj, ok := client.Get(d.Id())
// If the resource does not exist, inform Terraform. We want to immediately
// return here to prevent further processing.
  if !ok {
    d.SetId("")
    return nil
  }

  d.Set("address", obj.Address)
  return nil
*/
func resourceQubolePrestoRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

/*
1. If the Update callback returns with or without an error, the full state is saved.
	If the ID becomes blank, the resource is destroyed (even within an update, though this shouldn't happen except in error scenarios).
*/
func resourceQubolePrestoUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceQubolePrestoRead(d, m)
}

/*
1. If the Destroy callback returns without an error, the resource is assumed to be destroyed, and all state is removed.

2. If the Destroy callback returns with an error, the resource is assumed to still exist, and all prior state is preserved.
*/
func resourceQubolePrestoDelete(d *schema.ResourceData, m interface{}) error {
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
