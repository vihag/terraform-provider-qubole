
resource "qubole_cluster" "qubole_terraform_hive_cluster" {
	cloud_config		=	[
								{
									provider 			= 	"azure"
									resource_group_name	=	"qubole-apac-azure-readiness-secure-rg",
									compute_config	=	[
															{
																use_account_compute_creds	=	true
															}
									]
									location		=	[
															{
																location				=	"southeastasia"
															}
									]
									network_config	=	[
															{
																vnet_name					=	"qubole-apac-azure-readniness-secure-vnet"
																subnet_name					=	"private-subnet-for-secure-app"
																vnet_resource_group_name	=	"qubole-apac-azure-readiness-secure-rg"
																bastion_node_public_dns		=	"137.116.133.46"
																bastion_node_port			=	22
																bastion_node_user			=	"ec2-user"
															}
									]
									storage_config	=	[
															{
																managed_disk_account_type	=	"standard_lrs"
																data_disk_count				=	1
																data_disk_size				=	256
																disk_upscaling_config		=	 [
																									{
																										percent_free_space_threshold 	=	25
																										max_data_disk_count 			=	2
																										absolute_free_space_threshold	=	100.0
																									}
																]
															}
									]
								}								
	]
	cluster_info		=	[
								{
									master_instance_type			=	"Standard_A6"
									slave_instance_type				=	"Standard_A7"
									label 							=	["tf-qb-managed-hive-cl"]
									min_nodes						=	2
									max_nodes						=	5
									idle_cluster_timeout_in_secs	=	3600
									node_bootstrap					=	"empty-bootstrap.sh"
									disallow_cluster_termination	=	false
									datadisk						=	[
																			{
																				encryption	=	false
																			}
									]
									slave_request_type				=	"ondemand"
									env_settings					=	[
																			{
																				python_version	=	"3.5"
																				r_version		=	"3.3"
																				name			=	"ql-tf-env"
																			}
									
									] 
									heterogeneous_config			=	[
																			{
																				memory	=	[
																								{
																									instance_type	=	"Standard_A7"
																									weight			=	1.0
																								}
																				]
																			}
									]
									custom_tags						=	{
																			"Owner"			=	"Vihag Gupta"
																			"Environment"	=	"Dev"
																			"Project"		=	"Terraform Provider Hive"
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