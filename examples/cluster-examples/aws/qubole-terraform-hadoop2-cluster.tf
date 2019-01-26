provider "qubole" {
  auth_token	=	"${var.auth_token}"
  api_endpoint	=	"${var.api_endpoint}"
}

resource "qubole_cluster" "qubole_terraform_hive_cluster" {
	cloud_config		=	[
								{
									provider 		= 	"aws"
									compute_config	=	[
															{
																use_account_compute_creds	=	true
															}
									]
									location		=	[
															{
																aws_region				=	"ap-southeast-1"
																aws_availability_zone	=	"Any"
															}
									]
									network_config	=	[
															{
																vpc_id						=	"vpc-0e2be045e5fe137cd"
																subnet_id					=	"subnet-05cf662ed46f9cde8"
																bastion_node_public_dns		=	"13.229.7.25"
																bastion_node_port			=	22
																bastion_node_user			=	"ec2-user"
																persistent_security_groups	=	"persistent-security-group"
															}
									]
								}								
	]
	cluster_info		=	[
								{
									master_instance_type			=	"r4.xlarge"
									slave_instance_type				=	"r4.2xlarge"
									label 							=	["tf-qb-managed-hivecl"]
									node_base_cooldown_period		=	20
									min_nodes						=	2
									max_nodes						=	5
									idle_cluster_timeout_in_secs	=	3600
									node_bootstrap					=	"hoodie-presto-bootstrap.sh"
									disallow_cluster_termination	=	false
									datadisk						=	[
																			{
																				count		=	1,
																				type		=	"gp2"
																				size		=	100
																				encryption	=	true
																				ebs_upscaling_config	=	[
																												{
																													max_ebs_volume_count 			= 2
																													percent_free_space_threshold 	= 25
																													sampling_interval 				= 30
																													sampling_window 				= 5
																													absolute_free_space_threshold 	= 100
																												}
																				]
																			}
									]
									heterogeneous_config			=	[
																			{
																				memory	=	[
																								{
																									instance_type	=	"r4.2xlarge"
																									weight			=	1.0
																								}, 
																								{
																									instance_type 	=	"r4.4xlarge"
																									weight			=	0.5
																								}, 
																								{
																									instance_type 	=	"r4.xlarge"
																									weight			=	0.7
																								}
																				]
																			}
									]
									slave_request_type				=	"spot"
									spot_settings					=	[
																			{
																				spot_instance_settings	=	[
																												{
																													maximum_bid_price_percentage		=	70
																													timeout_for_request					=	5
																													maximum_spot_instance_percentage	=	40
																												}																										
																				]
																			}																		
									]
									custom_tags						=	{
																			"Owner"			=	"Vihag Gupta"
																			"Environment"	=	"Dev"
																			"Project"		=	"Terraform Provider"
																		}
								}								
	]
	engine_config		=	[
								{
									flavour			=	"hadoop2"
									hadoop_settings	=	[
															{
																custom_hadoop_config		=	"mapreduce.map.memory.mb=4704\nmapreduce.reduce.java.opts=-Xmx3763m"
																use_qubole_placement_policy	=	true
																fairscheduler_settings		=	[
																									{
																										default_pool	=	"pool"
																									}
																]
															}
									]
									hive_settings	=	[
															{
																is_metadata_cache_enabled		=	true
																hive_version					=	"1.2"
																is_hs2							=	true
																execution_engine				=	"tez"
																hs2_thrift_port					=	10003
																overrides						=	"hive.auto.convert.join=true"
															}
									]
								}			
	]
	monitoring			=	[
								{
									ganglia		=	true
									datadog	=	[
													{
														datadog_api_token	=	"api-token"
														datadog_app_token	=	"app-token"
													}
									]
								}
	]
}
