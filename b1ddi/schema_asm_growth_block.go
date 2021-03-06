// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcAsmGrowthBlock AsmGrowthBlock
//
// ASM growth group of fields.
//
// swagger:model ipamsvcAsmGrowthBlock
func schemaIpamsvcAsmGrowthBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Either the number or percentage of addresses to grow by.
			"growth_factor": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Either the number or percentage of addresses to grow by.",
			},

			// The type of factor to use: _percent_ or _count_.
			"growth_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of factor to use: _percent_ or _count_.",
			},
		},
	}
}

func flattenIpamsvcAsmGrowthBlock(r *models.IpamsvcAsmGrowthBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"growth_factor": r.GrowthFactor,
			"growth_type":   r.GrowthType,
		},
	}
}

func expandIpamsvcAsmGrowthBlock(d []interface{}) *models.IpamsvcAsmGrowthBlock {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})
	return &models.IpamsvcAsmGrowthBlock{
		GrowthFactor: int64(in["growth_factor"].(int)),
		GrowthType:   in["growth_type"].(string),
	}
}
