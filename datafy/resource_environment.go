package datafy

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnvironment() *schema.Resource {
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
			"deletion_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,
		Description:   "Resource for managing datafy Environments",
	}
}

func resourceEnvironmentCreate(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)

	envInput := EnvironmentInput{
		Name:               data.Get("name").(string),
		Description:        data.Get("description").(string),
		DeletionProtection: data.Get("deletion_protection").(bool),
	}

	env, err := c.CreateEnvironment(&envInput)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = mapEnvironmentData(data, env); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEnvironmentRead(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)
	id := data.Id()

	env, err := c.GetEnvironment(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = mapEnvironmentData(data, env); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEnvironmentUpdate(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	// TODO: FIXME does not work yet for some reason
	var diags diag.Diagnostics

	c := client.(*Client)

	envUpdate := EnvironmentUpdate{
		DeletionProtection: data.Get("deletion_protection").(bool),
	}

	env, err := c.UpdateEnvironment(data.Id(), &envUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = mapEnvironmentData(data, env); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEnvironmentDelete(_ context.Context, data *schema.ResourceData, client interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := client.(*Client)

	env, err := c.DeleteEnvironment(data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err = mapEnvironmentData(data, env); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func mapEnvironmentData(data *schema.ResourceData, env *Environment) error {
	var err error = nil

	data.SetId(env.Id)

	if err = data.Set("name", env.Name); err != nil {
		return err
	}
	if err = data.Set("description", env.Description); err != nil {
		return err
	}
	if err = data.Set("deletion_protection", env.DeletionProtection); err != nil {
		return err
	}
	if err = data.Set("state", env.State); err != nil {
		return err
	}
	if err = data.Set("tenant_id", env.TenantId); err != nil {
		return err
	}
	if err = data.Set("created_at", env.CreatedAt); err != nil {
		return err
	}
	if err = data.Set("updated_at", env.UpdatedAt); err != nil {
		return err
	}

	return err
}
