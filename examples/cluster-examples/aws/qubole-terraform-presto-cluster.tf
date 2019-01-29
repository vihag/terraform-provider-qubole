

resource "qubole_cluster" "qubole_terraform_presto_cluster" {
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
									label 							=	["tf-qb-managed-presto-cl"]
									node_base_cooldown_period		=	20
									min_nodes						=	2
									max_nodes						=	6
									idle_cluster_timeout_in_secs	=	7200
									node_bootstrap					=	"hoodie-presto-bootstrap.sh"
									disallow_cluster_termination	=	false
									datadisk						=	[
																			{
																				count		=	2,
																				type		=	"gp2"
																				size		=	100
																				encryption	=	true
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
