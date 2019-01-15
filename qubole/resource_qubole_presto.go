package qubole

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	entity "github.com/terraform-providers/terraform-provider-qubole/qubole/entity"
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
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
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
							Type:     schema.TypeString,
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
										Type:     schema.TypeString,
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
func resourceQubolePrestoCreate(d *schema.ResourceData, m interface{}) error {

	//Create the representative json object here
	//First create nested datas structures
	//1. EC2 Settings
	ec2_settings := entity.Ec2Settings{
		Compute_access_key:              d.Get("compute_access_key").(string),
		Compute_secret_key:              d.Get("compute_secret_key").(string),
		Aws_region:                      d.Get("aws_region").(string),
		Aws_preferred_availability_zone: d.Get("aws_preferred_availability_zone").(string),
		Vpc_id:                          d.Get("vpc_id").(string),
		Subnet_id:                       d.Get("subnet_id").(string),
		Master_elastic_ip:               d.Get("master_elastic_ip").(string),
		Bastion_node_public_dns:         d.Get("bastion_node_public_dns").(string),
		Bastion_node_port:               d.Get("bastion_node_port").(int),
		Bastion_node_user:               d.Get("bastion_node_user").(string),
		Role_instance_profile:           d.Get("role_instance_profile").(string),
		Use_account_compute_creds:       d.Get("use_account_compute_creds").(bool),
	}

	//2. Node configuration, but this will require constructing other sub-objects
	//2.1 HeterogeneousInstanceConfig
	heterogeneous_instance_config := entity.HeterogeneousInstanceConfig{
		Memory: d.Get("memory").(string),
	}

	//2.2 SpotInstanceSettings
	spot_instance_settings := entity.SpotInstanceSettings{
		Maximum_bid_price_percentage:     d.Get("maximum_bid_price_percentage").(string),
		Timeout_for_request:              d.Get("timeout_for_request").(int),
		Maximum_spot_instance_percentage: d.Get("maximum_spot_instance_percentage").(int),
	}

	//2.3 StableSpotInstanceSettings
	stable_spot_instance_settings := entity.StableSpotInstanceSettings{
		Maximum_bid_price_percentage: d.Get("maximum_bid_price_percentage").(string),
		Timeout_for_request:          d.Get("timeout_for_request").(int),
	}

	//2.4 SpotBlockSettings
	spot_block_settings := entity.SpotBlockSettings{
		Duration: d.Get("duration").(int),
	}

	//2.5 EbsUpscalingConfig
	ebs_upscaling_config := entity.EbsUpscalingConfig{
		Max_ebs_volume_count:          d.Get("max_ebs_volume_count").(int),
		Percent_free_space_threshold:  d.Get("percent_free_space_threshold").(int),
		Absolute_free_space_threshold: d.Get("absolute_free_space_threshold").(int),
		Sampling_interval:             d.Get("sampling_interval").(int),
		Sampling_window:               d.Get("sampling_window").(int),
	}

	node_configuration := entity.NodeConfiguration{
		Master_instance_type:          d.Get("master_instance_type").(string),
		Slave_instance_type:           d.Get("slave_instance_type").(string),
		Heterogeneous_instance_config: heterogeneous_instance_config,
		Initial_nodes:                 d.Get("initial_nodes").(int),
		Max_nodes:                     d.Get("max_nodes").(int),
		Slave_request_type:            d.Get("slave_request_type").(string),
		Spot_instance_settings:        spot_instance_settings,
		Stable_spot_instance_settings: stable_spot_instance_settings,
		Spot_block_settings:           spot_block_settings,
		Fallback_to_ondemand:          d.Get("fallback_to_ondemand").(bool),
		Ebs_volume_type:               d.Get("ebs_volume_type").(string),
		Ebs_volume_size:               d.Get("ebs_volume_size").(int),
		Ebs_volume_count:              d.Get("ebs_volume_count").(int),
		Ebs_upscaling_config:          ebs_upscaling_config,
		Custom_ec2_tags:               d.Get("custom_ec2_tags").(map[string]string),
		Idle_cluster_timeout_in_secs:  d.Get("idle_cluster_timeout_in_secs").(int),
		Node_base_cooldown_period:     d.Get("node_base_cooldown_period").(int),
		Node_spot_cooldown_period:     d.Get("node_spot_cooldown_period").(int),
	}

	//3. Hadoop Settings, but this will require constructing other sub-objects
	//3.1 FairSchedulerSettings
	fairscheduler_settings := entity.FairSchedulerSettings{
		Default_pool:             d.Get("default_pool").(string),
		Fairscheduler_config_xml: d.Get("fairscheduler_config_xml").(string),
	}

	hadoop_settings := entity.HadoopSettings{
		Use_hadoop2:                 d.Get("use_hadoop2").(bool),
		Use_spark:                   d.Get("use_spark").(bool),
		Custom_config:               d.Get("custom_config").(string),
		Fairscheduler_settings:      fairscheduler_settings,
		Use_qubole_placement_policy: d.Get("use_qubole_placement_policy").(bool),
	}

	//4. SecuritySettings
	security_settings := entity.SecuritySettings{
		Encrypted_ephemerals:      d.Get("encrypted_ephemerals").(bool),
		Customer_ssh_key:          d.Get("customer_ssh_key").(string),
		Persistent_security_group: d.Get("persistent_security_group").(string),
	}

	//5. PrestoSettings
	presto_settings := entity.PrestoSettings{
		Enable_presto: d.Get("enable_presto").(bool),
		Custom_config: d.Get("custom_config").(string),
	}

	//6. SparkSettings
	spark_settings := entity.SparkSettings{
		Custom_config: d.Get("custom_config").(string),
	}

	//7. DatadogSettings
	datadog_settings := entity.DatadogSettings{
		Datadog_api_token: d.Get("datadog_api_token").(string),
		Datadog_app_token: d.Get("datadog_app_token").(string),
	}

	//Finally, the cluster
	cluster := entity.Cluster{
		Label:                        d.Get("label").([]string),
		Presto_version:               d.Get("presto_version").(string),
		Spark_version:                d.Get("spark_version").(string),
		Zeppelin_interpreter_mode:    d.Get("zeppelin_interpreter_mode").(string),
		Ec2_settings:                 ec2_settings,
		Node_configuration:           node_configuration,
		Hadoop_settings:              hadoop_settings,
		Security_settings:            security_settings,
		Presto_settings:              presto_settings,
		Spark_settings:               spark_settings,
		Datadog_settings:             datadog_settings,
		Disallow_cluster_termination: d.Get("disallow_cluster_termination").(bool),
		Enable_ganglia_monitoring:    d.Get("enable_ganglia_monitoring").(bool),
		Node_bootstrap_file:          d.Get("node_bootstrap_file").(string),
		Idle_cluster_timeout:         d.Get("idle_cluster_timeout").(int),
	}

	//log the json constructed
	cluster_json, err := json.Marshal(cluster)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(cluster_json))

	//Make the http call to api here

	//Parse the response back to cluster object

	return resourceQubolePrestoRead(d, m)
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
