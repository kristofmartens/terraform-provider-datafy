package datafy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	// TODO: Implement
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATAFY_HOST", DefaultDatafyHost),
			},
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATAFY_PROFILE", DefaultDatafyProfile),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"datafy_environment": resourceEnvironment(),
			"datafy_project":     resourceProject(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"datafy_environment": dataSourceEnvironment(),
			"datafy_project":     dataSourceProject(),
		},
		ProviderMetaSchema:   nil,
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	host := d.Get("host").(string)
	profileName := d.Get("profile").(string)

	c, err := NewClient(host, profileName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Datafy client",
			Detail:   "Unable to retrieve token information",
		})
		return nil, diags
	}

	return c, diags
}
