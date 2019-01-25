package main

import (
        // "fmt"
        // "log"
        // "strings"

        "github.com/hashicorp/terraform/helper/schema"
)

func resourceCBServer() *schema.Resource {
        return &schema.Resource{
                Create: resourceCBServerCreate,
                Read:   resourceCBServerRead,
                Update: resourceCBServerUpdate,
                Delete: resourceCBServerDelete,

                Schema: map[string]*schema.Schema{
                        "quantity": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "cb_group": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "blueprint": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "cb_environment": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "osbuild": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "cpu": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "memory_gb": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "chef_environment": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "wid": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "function": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "parameters": &schema.Schema{
                                Type:     schema.TypeList,
                                Optional: true,
                                Elem:     &schema.Schema{Type: schema.TypeMap},
                        },
                },
        }
}

func resourceCBServerCreate(d *schema.ResourceData, meta interface{}) error {
        return nil
}

func resourceCBServerRead(d *schema.ResourceData, meta interface{}) error {
        return nil
}

func resourceCBServerUpdate(d *schema.ResourceData, meta interface{}) error {
        return nil
}

func resourceCBServerDelete(d *schema.ResourceData, meta interface{}) error {
        return nil
}