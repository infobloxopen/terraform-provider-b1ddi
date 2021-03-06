// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcHostnameRewriteBlock HostnameRewriteBlock
//
// Hostname Rewrite grouping fields.
//
// swagger:model ipamsvcHostnameRewriteBlock
func schemaIpamsvcHostnameRewriteBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance configuration for _hostname_rewrite_char_ field.
			"hostname_rewrite_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The inheritance configuration for _hostname_rewrite_char_ field.",
			},

			// The inheritance configuration for _hostname_rewrite_enabled_ field.
			"hostname_rewrite_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The inheritance configuration for _hostname_rewrite_enabled_ field.",
			},

			// The inheritance configuration for _hostname_rewrite_regex_ field.
			"hostname_rewrite_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The inheritance configuration for _hostname_rewrite_regex_ field.",
			},
		},
	}
}

func flattenIpamsvcHostnameRewriteBlock(r *models.IpamsvcHostnameRewriteBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"hostname_rewrite_char":    r.HostnameRewriteChar,
			"hostname_rewrite_enabled": r.HostnameRewriteEnabled,
			"hostname_rewrite_regex":   r.HostnameRewriteRegex,
		},
	}
}

func expandIpamsvcHostnameRewriteBlock(d []interface{}) *models.IpamsvcHostnameRewriteBlock {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.IpamsvcHostnameRewriteBlock{
		HostnameRewriteChar:    in["hostname_rewrite_char"].(string),
		HostnameRewriteEnabled: in["hostname_rewrite_enabled"].(bool),
		HostnameRewriteRegex:   in["hostname_rewrite_regex"].(string),
	}
}
