package main // import "github.com/section-io/terraform-provider-section"

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/section-io/terraform-provider-section/section"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: section.Provider})
}
