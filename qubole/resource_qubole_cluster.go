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

	if cluster, ok := model.ReadClusterFromTf(d); ok {
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

	if cluster, ok := model.ReadClusterFromTf(d); ok {
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
