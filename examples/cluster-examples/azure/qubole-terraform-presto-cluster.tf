provider "qubole" {
  auth_token	=	"${var.auth_token}"
  api_endpoint	=	"${var.api_endpoint}"
}

resource "qubole_cluster" "qubole_terraform_presto_cluster" {
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
																data_disk_count				=	2
																data_disk_size				=	256
															}
									]
								}								
	]
	cluster_info		=	[
								{
									master_instance_type			=	"Standard_A6"
									slave_instance_type				=	"Standard_A7"
									label 							=	["tf-qb-managed-prestocl"]
									min_nodes						=	2
									max_nodes						=	5
									idle_cluster_timeout_in_secs	=	3600
									node_bootstrap					=	"hoodie-presto-bootstrap.sh"
									disallow_cluster_termination	=	false
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
									flavour			=	"presto"
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
									presto_settings	=	[
															{
																custom_presto_config		=	"config.properties:\nascm.enabled=true\nascm.bds.target-latency=10s\ncatalog/hive.properties:\nhive.parquet.use-column-names=true\nbootstrap-file-path:\ns3://vihagg.us/qubole/us/scripts/hadoop/node_bootstrap_uber_hudi_presto.sh"
																presto_version				=	"0.193"
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