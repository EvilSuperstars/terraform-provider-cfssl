package cfssl

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertCreate,
		Read:   resourceCertRead,
		Delete: resourceCertDelete,

		Schema: map[string]*schema.Schema{
			"csr_json": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     validation.ValidateJsonString,
				DiffSuppressFunc: jsonDiffSuppress,
			},
			"ca_cert": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ca_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceCertCreate(d *schema.ResourceData, meta interface{}) error {
	csrJson := []byte(d.Get("csr_json").(string))
	req := csr.CertificateRequest{
		KeyRequest: csr.NewBasicKeyRequest(),
	}
	err := json.Unmarshal(csrJson, &req)
	if err != nil {
		return err
	}

	tmpCAFile, err := ioutil.TempFile("", "ca")
	if err != nil {
		return err
	}
	defer os.Remove(tmpCAFile.Name())
	if _, err := tmpCAFile.Write([]byte(d.Get("ca_cert").(string))); err != nil {
		return err
	}
	tmpCAKeyFile, err := ioutil.TempFile("", "ca-key")
	if err != nil {
		return err
	}
	defer os.Remove(tmpCAKeyFile.Name())
	if _, err := tmpCAKeyFile.Write([]byte(d.Get("ca_key").(string))); err != nil {
		return err
	}

	g := &csr.Generator{Validator: genkey.Validator}
	csrBytes, key, err := g.ProcessRequest(&req)
	if err != nil {
		return err
	}

	c := cli.Config{
		CAFile:    tmpCAFile.Name(),
		CAKeyFile: tmpCAKeyFile.Name(),
	}
	s, err := sign.SignerFromConfig(c)
	if err != nil {
		return err
	}
	signReq := signer.SignRequest{
		Request: string(csrBytes),
	}
	cert, err := s.Sign(signReq)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("cert", string(cert))
	d.Set("csr", string(csrBytes))
	d.Set("key", string(key))

	return nil
}

func resourceCertRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCertDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
