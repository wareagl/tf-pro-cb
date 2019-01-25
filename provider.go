package main

import (
        "github.com/hashicorp/terraform/helper/schema"
        "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
        return &schema.Provider{
                Schema: map[string]*schema.Schema{
                	"server": &schema.Schema{
                		Type:		schema.TypeString,
                		Required:	true,
                		DefaultFunc:	schema.EnvDefaultFunc("CB_SERVER", nil),
                		Description:	"The Cloudbolt server",
                	},
                	"username": &schema.Schema{
                		Type:		schema.TypeString,
                		Required:	true,
                		DefaultFunc:	schema.EnvDefaultFunc("CB_USERNAME", nil),
                		Description:	"Your Cloudbolt username",
                	},
                	"password": &schema.Schema{
                		Type:		schema.TypeString,
                		Required:	true,
                		DefaultFunc:	schema.EnvDefaultFunc("CB_PASSWORD", nil),
                		Description:	"Your Cloudbolt password",
                	},
                },

                ResourcesMap: map[string]*schema.Resource{
                	"cb_server":	resourceCBServer(),
                },

                ConfigureFunc:	providerConfigure,
        }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	user := d.Get("username").(string)
	pass := d.Get("password").(string)
	// server := d.Get("server").(string)

	return user, nil
}