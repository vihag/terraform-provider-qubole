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
									master_instance_type			=	"m1.xlarge"
									label 							=	["tf-qb-managed-airflowcl"]
									min_nodes						=	1		
									max_nodes						=	1
									node_bootstrap					=	"hoodie-presto-bootstrap.sh"
									disallow_cluster_termination	=	false
									datadisk						=	[
																			{
																				encryption	=	true
																			}
									]
									slave_request_type				=	"ondemand"
									custom_tags						=	{
																			"Owner"			=	"Vihag Gupta"
																			"Environment"	=	"Dev"
																			"Project"		=	"Terraform Provider"
																		}
								}								
	]
	engine_config		=	[
								{
									flavour			=	"airflow"
									airflow_settings	=	[
															{
																dbtap_id				=	907
																fernet_key				=	"a2DbzaMVWidhZOVuNDJ0aKE/Nvw//AGUkTtQ1c6tI7Q="
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
