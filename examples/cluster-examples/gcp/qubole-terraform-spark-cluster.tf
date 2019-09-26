provider "qubole" {
  auth_token	=	"${var.auth_token}"
  api_endpoint	=	"${var.api_endpoint}"
}

resource "qubole_cluster" "qubole_terraform_spark_cluster" {
	cloud_config		=	[
								{
									provider 		= 	"gcp"
									compute_config	=	[
															{
																use_account_compute_creds	=	true
															}
									]
									location		=	[
															{
																region					=	"asia-southeast1"
																zone					=	"asia-southeast1-b"
															}
									]
									network_config	=	[
															{
																network						=	"projects/sa-demo-project-1/global/networks/jlaw-vpc-sg"
																subnet_id					=	"projects/sa-demo-project-1/regions/asia-southeast1/subnetworks/jlaw-vpc-sg-network"
															}
									]
									storage_config	=	[
															{
																disk_type					=	"pd-ssd"
																disk_size_in_gb				=	150
																disk_count					=	1
																disk_upscaling_config		=	[
																									{
																										percent_free_space_threshold	=	25
																										max_persistent_disk_count		=	3
																										absolute_free_space_threshold	=	90
																									}
																]
															}
									]
									cluster_composition	=	[
																{
																	master						=	[
																										{
																											preemptible		=	false
																										}
																	]
																	min_nodes					=	[
																										{
																											preemptible		=	false
																											percentage 		= 	0.0
																										}
																	]
																	autoscaling_nodes			=	[
																										{
																											preemptible		=	true
																											percentage 		= 	75.0
																										}
																	]
																}
									]
								}								
	]
	cluster_info		=	[
								{
									master_instance_type			=	"n1-highmem-4"
									slave_instance_type				=	"n1-highmem-8"
									label 							=	["tf-qb-managed-spark-cl"]
									node_base_cooldown_period		=	5
									node_volatile_cooldown_period	=	5
									min_nodes						=	1
									max_nodes						=	4
									idle_cluster_timeout_in_secs	=	1800
									node_bootstrap					=	"empty-bootstrap.sh"
									disallow_cluster_termination	=	false
									env_settings					=	[
																			{
																				python_version	=	"3.5"
																				r_version		=	"3.3"
																			}
									
									] 
									heterogeneous_config			=	[
																			{
																				memory	=	[
																								{
																									instance_type	=	"n1-highmem-8"
																									weight			=	1.0
																								}, 
																								{
																									instance_type 	=	"n1-standard-16"
																									weight			=	1.15385
																								}, 
																								{
																									instance_type 	=	"n1-standard-32"
																									weight			=	2.30769
																								}
																				]
																			}
									]
									custom_tags						=	{
																			"owner"			=	"vihag_gupta"
																			"environment"	=	"dev"
																			"project"		=	"terraform_provider_spark"
																		}
								}								
	]
	engine_config		=	[
								{
									flavour			=	"spark"
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
									spark_settings	=	[
															{
																custom_spark_config		=	"spark-defaults.conf:\nspark.driver.extraLibraryPath /usr/lib/hadoop2/lib/native"
																spark_version				=	"2.3-latest"
																enable_rubix				=	true
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
