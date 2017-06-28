package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/idmsubs/terraform-provider-vra/vra"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vra.Provider})
}