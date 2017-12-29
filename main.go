package main

import (
	"github.com/GSLabDev/terraform-provider-vra/vra"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vra.Provider})
}
