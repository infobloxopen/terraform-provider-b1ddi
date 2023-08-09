package b1ddi

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var (
	ParseError = "unable to parse key '%s': %v"
)

/*
	updateDataRecordRData helps convert rDATA record fields of type integer from string value
    Introduced to fix issues where terraform converts integer values to string in rendered config
*/
func updateDataRecordRData(d interface{}, recordType string) (interface{}, diag.Diagnostics) {
	if d == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	in := d.(map[string]interface{})
	switch recordType {
	case "CAA":
		if val, ok := in["flags"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "flags", err)))...)
			} else {
				in["flags"] = i
			}
		}

	case "MX":
		if val, ok := in["preference"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "preference", err)))...)
			} else {
				in["preference"] = i
			}
		}

	case "NAPTR":
		if val, ok := in["order"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "order", err)))...)
			} else {
				in["order"] = i
			}
		}
		if val, ok := in["preference"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "preference", err)))...)
			} else {
				in["preference"] = i
			}
		}

	case "SOA":
		if val, ok := in["serial"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "serial", err)))...)
			} else {
				in["serial"] = i
			}
		}

	case "SRV":
		if val, ok := in["port"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "port", err)))...)
			} else {
				in["port"] = i
			}
		}
		if val, ok := in["priority"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "priority", err)))...)
			} else {
				in["priority"] = i
			}
		}
		if val, ok := in["weight"]; ok {
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				diags = append(diags, diag.FromErr(errors.New(fmt.Sprintf(ParseError, "weight", err)))...)
			} else {
				in["weight"] = i
			}
		}
	default:
		return d, nil
	}

	return in, diags
}
