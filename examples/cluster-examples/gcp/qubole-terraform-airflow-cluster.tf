

resource "qubole_cluster" "qubole_terraform_airflow_cluster" {
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
																master_static_ip			=	"1.2.3.4"
															}
									]
									storage_config	=	[
															{
																disk_type					=	"pd-ssd"
																disk_size_in_gb				=	150
																disk_count					=	1
															}
									]
								}								
	]
	cluster_info		=	[
								{
									master_instance_type			=	"n1-highmem-4"
									label 							=	["tf-qb-managed-airflow-cl"]
									min_nodes						=	1		
									max_nodes						=	1
									node_bootstrap					=	"empty-bootstrap.sh"
									disallow_cluster_termination	=	false
									datadisk						=	[
																			{
																				encryption	=	true
																			}
									]
									slave_request_type				=	"ondemand"
									custom_tags						=	{
																			"owner"			=	"vihag_gupta"
																			"environment"	=	"dev"
																			"project"		=	"terraform_provider_airflow"
																		}
								}								
	]
	engine_config		=	[
								{
									flavour			=	"airflow"
									airflow_settings	=	[
															{
																dbtap_id				= 	-1
																fernet_key				= 	"my-fernet-key"
																overrides				=	"core.executor=CeleryExecutor\ncore.parallelism=32"
																version					=	"1.10.0"
																airflow_python_version	=	"3.5"
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
