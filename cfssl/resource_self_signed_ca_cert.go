package cfssl

import (
	"encoding/json"
	"time"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceSelfSignedCACert() *schema.Resource {
	return &schema.Resource{
		Create: resourceSelfSignedCACertCreate,
		Read:   resourceSelfSignedCACertRead,
		Delete: resourceSelfSignedCACertDelete,

		Schema: map[string]*schema.Schema{
			"csr_json": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     validation.ValidateJsonString,
				DiffSuppressFunc: jsonDiffSuppress,
			},
			"cert": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"csr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSelfSignedCACertCreate(d *schema.ResourceData, meta interface{}) error {
	csrJson := []byte(d.Get("csr_json").(string))
	req := csr.CertificateRequest{
		KeyRequest: csr.NewBasicKeyRequest(),
	}
	err := json.Unmarshal(csrJson, &req)
	if err != nil {
		return err
	}

	cert, csrBytes, key, err := initca.New(&req)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("cert", string(cert))
	d.Set("csr", string(csrBytes))
	d.Set("key", string(key))

	return nil
}

func resourceSelfSignedCACertRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSelfSignedCACertDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
