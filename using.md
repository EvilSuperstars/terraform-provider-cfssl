# `cfssl` Provider

The Terraform [cfssl](https://github.com/EvilSuperstars/terraform-provider-cfssl) provider implements functionality from [CloudFlare's PKI/TLS toolkit](https://github.com/cloudflare/cfssl).

This provider requires no configuration.

### Example Usage

```hcl
provider "cfssl" {}

resource "cfssl_xxx" "foo" {
}
```

## Resources

### `cfssl_xxx`

Converts JSON to YAML.

#### Argument Reference

The following arguments are supported:

* `json` - (Required, string) The JSON string that is to be converted to YAML.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `yaml` - (string) The YAML string
