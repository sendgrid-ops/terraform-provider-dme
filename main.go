package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sendgrid-ops/terraform-provider-dme/dme"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dme.Provider})
}
