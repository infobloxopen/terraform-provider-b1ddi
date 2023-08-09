package b1ddi

import (
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
		diags = inPlaceUpdater(in, "flags")

	case "MX":
		diags = inPlaceUpdater(in, "preference")

	case "NAPTR":
		diagErr := inPlaceUpdater(in, "order")
		if diagErr != nil {
			diags = append(diags, diagErr...)
		}
		diagErr = inPlaceUpdater(in, "preference")
		if diagErr != nil {
			diags = append(diags, diagErr...)
		}

	case "SOA":
		diags = inPlaceUpdater(in, "serial")

	case "SRV":
		diagErr := inPlaceUpdater(in, "port")
		if diagErr != nil {
			diags = append(diags, diagErr...)
		}
		diagErr = inPlaceUpdater(in, "priority")
		if diagErr != nil {
			diags = append(diags, diagErr...)
		}
		diagErr = inPlaceUpdater(in, "weight")
		if diagErr != nil {
			diags = append(diags, diagErr...)
		}
	default:
		return d, nil
	}

	return in, diags
}

func inPlaceUpdater(in map[string]interface{}, key string) diag.Diagnostics {
	if val, ok := in[key]; ok {
		i, err := strconv.Atoi(val.(string))
		if err != nil {
			return diag.Errorf(ParseError, key, err)
		} else {
			in[key] = i
		}
	}
	return nil
}
