package main

// Build the plugin by the following command: `go build -o terraform-provider-qubole`

// For subsequent testing on the in-folder test tf, first terraform init, second terraform plan, third terraform apply

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/hashicorp/terraform/terraform"
        "github.com/terraform-providers/terraform-provider-qubole/qubole"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return qubole.Provider()
                },
        })
}
