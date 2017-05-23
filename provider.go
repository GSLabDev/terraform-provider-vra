package vra

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider ... provides scvmm capability to terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SCVMM Server IP",
				DefaultFunc: schema.EnvDefaultFunc("VRA_URL", nil),
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User name",
				DefaultFunc: schema.EnvDefaultFunc("VRA_USER", nil),
			},
			"user_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password for provided user_name",
				DefaultFunc: schema.EnvDefaultFunc("VRA_PASSWORD", nil),
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Tennant for provided user_name",
				DefaultFunc: schema.EnvDefaultFunc("VRA_TENANT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"vra_execute_blueprint": ExecuteBlueprint(),
		
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
    	Host:     d.Get("host_url").(string),
		Username: d.Get("user_name").(string),
		Password: d.Get("user_password").(string),
		Tenant:   d.Get("tenant").(string),
	}

	log.Println("[INFO] Initializing Winrm Connection")
	return config.Connection()
}