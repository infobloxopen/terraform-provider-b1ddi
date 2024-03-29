// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcInheritedASMConfig InheritedASMConfig
//
// The inheritance configuration for the __ASMConfig__ object.
//
// swagger:model ipamsvcInheritedASMConfig
func schemaIpamsvcInheritedASMConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The block of ASM fields: _enable_, _enable_notification_, _reenable_date_.
			"asm_enable_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedAsmEnableBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The block of ASM fields: _enable_, _enable_notification_, _reenable_date_.",
			},

			// The block of ASM fields: _growth_factor_, _growth_type_.
			"asm_growth_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedAsmGrowthBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The block of ASM fields: _growth_factor_, _growth_type_.",
			},

			// ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_percent_ * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged.
			"asm_threshold": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_percent_ * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged.",
			},

			// The forecast period in days.
			"forecast_period": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "The forecast period in days.",
			},

			// The minimum amount of history needed before ASM can run on this subnet.
			"history": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "The minimum amount of history needed before ASM can run on this subnet.",
			},

			// The minimum size of range needed for ASM to run on this subnet.
			"min_total": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "The minimum size of range needed for ASM to run on this subnet.",
			},

			// The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change.
			"min_unused": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change.",
			},
		},
	}
}

func flattenIpamsvcInheritedASMConfig(r *models.IpamsvcInheritedASMConfig) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"asm_enable_block": flattenIpamsvcInheritedAsmEnableBlock(r.AsmEnableBlock),
			"asm_growth_block": flattenIpamsvcInheritedAsmGrowthBlock(r.AsmGrowthBlock),
			"asm_threshold":    flattenInheritanceInheritedUInt32(r.AsmThreshold),
			"forecast_period":  flattenInheritanceInheritedUInt32(r.ForecastPeriod),
			"history":          flattenInheritanceInheritedUInt32(r.History),
			"min_total":        flattenInheritanceInheritedUInt32(r.MinTotal),
			"min_unused":       flattenInheritanceInheritedUInt32(r.MinUnused),
		},
	}
}

func expandIpamsvcInheritedASMConfig(ctx context.Context, d []interface{}) (*models.IpamsvcInheritedASMConfig, error) {
	if len(d) == 0 || d[0] == nil {
		return nil, nil
	}
	in := d[0].(map[string]interface{})

	asmEnableBlock, err := expandIpamsvcInheritedAsmEnableBlock(ctx, in["asm_enable_block"].([]interface{}))
	if err != nil {
		tflog.Error(ctx, "Failed to parse 'asm_enable_block' field. The underlying expand function returned an error.")
		return nil, err
	}

	return &models.IpamsvcInheritedASMConfig{
		AsmEnableBlock: asmEnableBlock,
		AsmGrowthBlock: expandIpamsvcInheritedAsmGrowthBlock(in["asm_growth_block"].([]interface{})),
		AsmThreshold:   expandInheritanceInheritedUInt32(in["asm_threshold"].([]interface{})),
		ForecastPeriod: expandInheritanceInheritedUInt32(in["forecast_period"].([]interface{})),
		History:        expandInheritanceInheritedUInt32(in["history"].([]interface{})),
		MinTotal:       expandInheritanceInheritedUInt32(in["min_total"].([]interface{})),
		MinUnused:      expandInheritanceInheritedUInt32(in["min_unused"].([]interface{})),
	}, nil
}
