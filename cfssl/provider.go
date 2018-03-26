package cfssl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"cfssl_cert":                resourceCert(),
			"cfssl_self_signed_ca_cert": resourceSelfSignedCACert(),
		},

		DataSourcesMap: map[string]*schema.Resource{},
	}
}
