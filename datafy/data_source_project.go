package datafy

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProject() *schema.Resource {
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
			"git_repo": {
				Type:     schema.TypeString,
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
			"last_activity": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: dataSourceProjectRead,
		Description: "Returns info about a project",
	}
}

func dataSourceProjectRead(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)
	id := data.Get("id").(string)

	proj, err := c.GetProject(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(proj.Id)

	if err = data.Set("name", proj.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("description", proj.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("git_repo", proj.GitRepo); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("state", proj.State); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tenant_id", proj.TenantId); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("created_at", proj.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("updated_at", proj.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("last_activity", proj.LastActivity); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
