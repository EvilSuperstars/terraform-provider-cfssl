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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.ValidateJsonString,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"out_json": {
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
	out := map[string]string{}
	if cert != nil {
		out["cert"] = string(cert)
	}
	if csrBytes != nil {
		out["csr"] = string(csrBytes)
	}
	if key != nil {
		out["key"] = string(key)
	}
	outJson, err := json.Marshal(out)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("out_json", string(outJson))

	return nil
}

func resourceSelfSignedCACertRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSelfSignedCACertDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
