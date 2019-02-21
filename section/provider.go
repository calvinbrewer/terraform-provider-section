package section

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/section-io/section-sdk-go/api"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECTION_USERNAME", nil),
				Description: "The username for Section account",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECTION_PASSWORD", nil),
				Description: "The password for Section Account",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"section_api": resourceSectionAPI(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client, err := api.NewClient(d.Get("username").(string), d.Get("password").(string))
	if err != nil {
		return nil, err
	}

	config := providerConfig{
		Client:     client
	}

	return config, nil
}

type providerConfig struct {
	Client     api.Client
}
