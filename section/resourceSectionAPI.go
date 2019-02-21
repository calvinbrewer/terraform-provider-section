package section

import (
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
	"github.com/section-io/section-sdk-go/api"
)

func accountSchemaElement() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func applicationSchemaElement() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stack_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func environmentSchemaElement() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_environment_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSectionAPICreateAccount(d *schema.ResourceData, config providerConfig) error {
	accountName := d.Get("name")
	hostname := d.Get("hostname")
	origin := d.Get("origin")
	stackName := d.Get("stack_name")

	response, err := config.Client.AccountCreate(accountName, hostname, origin, stackName)

	if err != nil {
		return err
	}

	return response
}

func resourceSectionAPICreateApplication(d *schema.ResourceData, config providerConfig) error {
	hostname := d.Get("hostname")
	origin := d.Get("origin")
	stackName := d.Get("stack_name")

	response, err := config.Client.ApplicationCreate(accountName, hostname, origin, stackName)

	if err != nil {
		return err
	}

	return response
}

func resourceSectionAPICreateEnvironment(d *schema.ResourceData, config providerConfig) error {
	name := d.Get("name")
	sourceEnvironmentName := d.Get("source_environment_name")
	domainName := d.Get("domain_name")

	response, err := config.Client.AccountCreate(accountName, hostname, origin, stackName)

	if err != nil {
		return err
	}

	return response
}
