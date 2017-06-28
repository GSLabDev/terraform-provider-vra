package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/idmsubs/gslab-provider-vra/vra"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vra.Provider})
}