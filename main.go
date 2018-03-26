package main

import (
	"github.com/EvilSuperstars/terraform-provider-cfssl/cfssl"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cfssl.Provider,
	})
}
