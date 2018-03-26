# `cfssl` Provider

The Terraform [cfssl](https://github.com/EvilSuperstars/terraform-provider-cfssl) provider implements functionality from [CloudFlare's PKI/TLS toolkit](https://github.com/cloudflare/cfssl).

This provider requires no configuration.

### Example Usage

```hcl
provider "cfssl" {}

resource "cfssl_self_signed_ca_cert" "foo" {
}
```

## Resources

### `cfssl_self_signed_ca_cert`

Generate a self-signed root CA certificate and private key.
See [CloudFlare's documentation](https://github.com/cloudflare/cfssl#generating-self-signed-root-ca-certificate-and-private-key).

#### Argument Reference

The following arguments are supported:

* `csr_json` - (Required, string) The request as a JSON string.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `cert` - (string) The output CA certificate
* `csr` - (string) The output CSR in PEM format
* `key` - (string) The output CA private key

### `cfssl_cert`

Generate a certificate and private key signed by a CA.
See [CloudFlare's documentation](https://github.com/cloudflare/cfssl#generating-a-local-issued-certificate-and-private-key).

#### Argument Reference

The following arguments are supported:

* `csr_json` - (Required, string) The request as a JSON string.
* `ca_cert` - (Required, string) The CA certificate.
* `ca_key` - (Required, string) The CA private key.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `cert` - (string) The output ertificate
* `csr` - (string) The output CSR in PEM format
* `key` - (string) The output private key
