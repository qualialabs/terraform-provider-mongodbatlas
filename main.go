package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
    "github.com/qualialabs/terraform-provider-mongodbatlas/mongodbatlas"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return mongodbatlas.Provider()
        },
    })
}