provider "qubole" {
  auth_token	=	"${var.auth_token}"
  api_endpoint	=	"${var.api_endpoint}"
}

resource "qubole_cluster" "qubole_terraform_airflow_cluster" {
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
																bastion_node_public_dns		=	"13.76.216.161"
																bastion_node_port			=	22
																bastion_node_user			=	"ec2-user"
															}
									]
									storage_config	=	[
															{
																managed_disk_account_type	=	"standard_lrs"
																data_disk_count				=	0
																data_disk_size				=	256
															}
									]
								}								
	]
	cluster_info		=	[
								{
									master_instance_type			=	"Standard_A6"
									label 							=	["tf-qb-managed-airflow-cl"]
									min_nodes						=	1
									max_nodes						=	1
									node_bootstrap					=	"hoodie-airflow-bootstrap.sh"
									disallow_cluster_termination	=	true
									datadisk						=	[
																			{
																				encryption	=	false
																			}
									]
									slave_request_type				=	"ondemand"
								}								
	]
	engine_config		=	[
								{
									flavour			=	"airflow"
									airflow_settings	=	[
															{
																dbtap_id				=	-1
																fernet_key				=	"Wk607UZxIATjFfvIr/r12xAUTb/CRlKrXAJwUSEl4YQ="
																overrides				=	"core.executor=CeleryExecutor\ncore.parallelism=32"
																version					=	"1.8.2"
																airflow_python_version	=	"2.7"
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