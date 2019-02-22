package section

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func panicOnError(err error) {
	if err == nil {
		return
	}
	panic(err)
}

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
	accountName := d.Get("name").(string)
	hostname := d.Get("hostname").(string)
	origin := d.Get("origin").(string)
	stackName := d.Get("stack_name").(string)

	response, err := config.Client.AccountCreate(accountName, hostname, origin, stackName)

	panicOnError(d.Set("account_id", response.AccountID))

	return err
}

func resourceSectionAPICreateApplication(d *schema.ResourceData, config providerConfig) error {
	hostname := d.Get("hostname").(string)
	origin := d.Get("origin").(string)
	stackName := d.Get("stack_name").(string)
	accountId := d.Get("account_id").(int)

	response, err := config.Client.ApplicationCreate(accountId, hostname, origin, stackName)

	panicOnError(d.Set("application_id", response.ApplicationID))

	return err
}

func resourceSectionAPICreateEnvironment(d *schema.ResourceData, config providerConfig) error {
	name := d.Get("name").(string)
	sourceEnvironmentName := d.Get("source_environment_name").(string)
	domainName := d.Get("domain_name").(string)
	accountId := d.Get("account_id").(int)
	applicationId := d.Get("application_id").(int)

	response, err := config.Client.EnvironmentCreate(accountId, applicationId, name, sourceEnvironmentName, domainName)

	panicOnError(d.Set("environment_id", response.EnvironmentID))

	return err
}
