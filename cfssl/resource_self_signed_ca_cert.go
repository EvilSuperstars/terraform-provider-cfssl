package cfssl

import (
	"encoding/json"
	"time"

	"github.com/cloudflare/cfssl/csr"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceSelfSignedCACert() *schema.Resource {
	return &schema.Resource{
		Create: resourceSelfSignedCACertCreate,

		Schema: map[string]*schema.Schema{
			"csr": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.ValidateJsonString,
			},
		},
	}
}

func resourceSelfSignedCACertCreate(d *schema.ResourceData, meta interface{}) error {
	csrBytes := []byte(d.Get("csr").(string))
	req := csr.CertificateRequest{
		KeyRequest: csr.NewBasicKeyRequest(),
	}
	err := json.Unmarshal(csrBytes, &req)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())

	return nil
}
