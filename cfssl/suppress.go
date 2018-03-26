package cfssl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/structure"
)

func jsonDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	normalizedOld, err := structure.NormalizeJsonString(old)
	if err != nil {
		return false
	}
	normalizedNew, err := structure.NormalizeJsonString(new)
	if err != nil {
		return false
	}
	return normalizedOld == normalizedNew
}
