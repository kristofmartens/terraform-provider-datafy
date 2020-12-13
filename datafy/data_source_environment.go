package datafy

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEnvironments() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Active",
				//Description:  "Defaults to Active, choose between Active, All, Inactive",
				//ExactlyOneOf: []string{"Active", "All", "Inactive"},
			},
			"environments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"description": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"tenant_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"created_at": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"updated_at": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"state": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"deletion_protection": {
						Type:     schema.TypeBool,
						Computed: true,
					},
				},
			},
		},
		ReadContext: dataSourceEnvironmentsRead,
		Description: "Returns all environments",
	}
}

func dataSourceEnvironmentsRead(ctx context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)
	state := data.Get("state").(string)
	envs, err := c.GetEnvironments(state)
	if err != nil {
		return diag.FromErr(err)
	}

	// TODO: map the output to the terraform schema
	fmt.Println(envs)

	return diags
}

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deletion_protection": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: dataSourceEnvironmentRead,
		Description: "Returns info about an environment",
	}
}

func dataSourceEnvironmentRead(ctx context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)
	id := data.Get("id").(string)

	env, err := c.GetEnvironment(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(env.Id)

	if err = data.Set("id", env.Id); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("name", env.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("description", env.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("deletion_protection", env.DeletionProtection); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("state", env.State); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tenant_id", env.TenantId); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("created_at", env.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("updated_at", env.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
