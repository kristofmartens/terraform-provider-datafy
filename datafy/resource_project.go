package datafy

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"git_repo": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Description:   "Resource for managing Datafu project resources",
	}
}

func resourceProjectCreate(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)

	input := ProjectInput{
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		GitRepo:     data.Get("git_repo").(string),
	}

	project, err := c.CreateProject(&input)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(project.Id)

	if err = data.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("description", project.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("git_repo", project.GitRepo); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tenant_id", project.TenantId); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("created_at", project.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("updated_at", project.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("last_activity", project.LastActivity); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProjectRead(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)

	id := data.Id()

	project, err := c.GetProject(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(project.Id)

	if err = data.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("description", project.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("git_repo", project.GitRepo); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tenant_id", project.TenantId); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("created_at", project.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("updated_at", project.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("last_activity", project.LastActivity); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProjectUpdate(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	// TODO: FIXME does not work yet for some reason
	var diags diag.Diagnostics

	c := client.(*Client)

	id := data.Id()
	input := ProjectUpdate{
		Description: data.Get("description").(string),
		GitRepo:     data.Get("git_repo").(string),
	}

	project, err := c.UpdateProject(id, &input)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(project.Id)

	if err = data.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("description", project.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("git_repo", project.GitRepo); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tenant_id", project.TenantId); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("created_at", project.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("updated_at", project.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("last_activity", project.LastActivity); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProjectDelete(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)

	id := data.Id()

	project, err := c.DeleteProject(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(project.Id)

	if err = data.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("description", project.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("git_repo", project.GitRepo); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("tenant_id", project.TenantId); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("created_at", project.CreatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("updated_at", project.UpdatedAt); err != nil {
		return diag.FromErr(err)
	}
	if err = data.Set("last_activity", project.LastActivity); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
