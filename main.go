package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-datafy/datafy"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: datafy.Provider,
		// TODO: Investigate whether to use GRPC based implementation
		GRPCProviderFunc: nil,
		Logger:           nil,
		TestConfig:       nil,
	})
}
