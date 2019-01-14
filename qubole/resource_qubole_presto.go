package qubole

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceQubolePresto() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"label": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
			},
			"presto_version": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
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
							Required: true,
						},
						"compute_secret_key": {
							Type:     schema.TypeString,
							Required: true,
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
						"role_instance_profile": {
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
							Type:     schema.TypeBoolean,
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
							Type:     schema.TypeString,
							Optional: true,
						},
						"ebs_volume_count": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ebs_volume_count": { //ebs_upscaling_config
							Type:     schema.TypeInt,
							Optional: true,
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
							Type:     schema.TypeBoolean,
							Optional: true,
						},
						"use_spark": {
							Type:     schema.TypeBoolean,
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
							Type:     schema.TypeBoolean,
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
							Type:     schema.TypeBoolean,
							Optional: true,
						},
						"customer_ssh_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"persistent-security-groups": {
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
							Type:     schema.TypeBoolean,
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
				Type:     schema.TypeBoolean,
				Optional: true,
			},
			"enable_ganglia_monitoring": &schema.Schema{
				Type:     schema.TypeBoolean,
				Optional: true,
			},
			"node_bootstrap_file": &schema.Schema{
				Type:     schema.TypeBoolean,
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
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceServerRead(d, m)
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
	return resourceServerRead(d, m)
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
