# `cfssl` Provider

The Terraform [cfssl](https://github.com/EvilSuperstars/terraform-provider-cfssl) provider implements functionality from [CloudFlare's PKI/TLS toolkit](https://github.com/cloudflare/cfssl).

This provider requires no configuration.

### Example Usage

```hcl
provider "cfssl" {}

resource "cfssl_ca_cert" "foo" {
}
```

## Resources

### `cfssl_self_signed_ca_cert`

Generate a self-signed root CA certificate and private key.
See [CloudFlare's documentation](https://github.com/cloudflare/cfssl#generating-self-signed-root-ca-certificate-and-private-key).

#### Argument Reference

The following arguments are supported:

* `csr` - (Required, string) The key request as a JSON string.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `yaml` - (string) The YAML string
