package section

import (
	"log"
	"strconv"
	"time"

	"github.com/calvinbrewer/section-sdk-go.git"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
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
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
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
			"account_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"application_id": {
				Type:     schema.TypeInt,
				Computed: true,
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
			"account_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"application_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"environment_id": {
				Type:     schema.TypeInt,
				Computed: true,
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

	panicOnError(d.Set("account_id", response.id))

	if err != nil {
		return err
	}

	return response
}

func resourceSectionAPICreateApplication(d *schema.ResourceData, config providerConfig) error {
	hostname := d.Get("hostname")
	origin := d.Get("origin")
	stackName := d.Get("stack_name")
	accountId := d.Get("account_id")

	response, err := config.Client.ApplicationCreate(accountId, hostname, origin, stackName)

	panicOnError(d.Set("application_id", response.id))

	if err != nil {
		return err
	}

	return response
}

func resourceSectionAPICreateEnvironment(d *schema.ResourceData, config providerConfig) error {
	name := d.Get("name")
	sourceEnvironmentName := d.Get("source_environment_name")
	domainName := d.Get("domain_name")
	accountId := d.Get("account_id")
	applicationId := d.Get("application_id")

	response, err := config.Client.EnvironmentCreate(accountId, applicationId, name, sourceEnvironmentName, domainName)

	panicOnError(d.Set("environment_id", response.id))

	if err != nil {
		return err
	}

	return response
}
