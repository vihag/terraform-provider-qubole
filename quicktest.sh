#!/bin/bash

export TF_LOG=DEBUG
export TF_LOG_PATH=/tmp/tflog.log


clear
>/tmp/tflog.log
go build -o terraform-provider-qubole
terraform init
terraform plan
terraform apply