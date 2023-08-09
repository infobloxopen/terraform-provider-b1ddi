package b1ddi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateDataRecordRData(t *testing.T) {
	testData := map[string]struct {
		input       map[string]interface{}
		output      map[string]interface{}
		recordType  string
		errExpected bool
	}{
		"Valid Rdata - MX Record": {
			input: map[string]interface{}{
				"preference": "1",
				"exchange":   "mail.infoblox.com",
			},
			output: map[string]interface{}{
				"preference": 1,
				"exchange":   "mail.infoblox.com",
			},
			recordType:  "MX",
			errExpected: false,
		},
		"Invalid Rdata - MX Record": {
			input: map[string]interface{}{
				"preference": "1qw",
				"exchange":   "mail.infoblox.com",
			},
			output: map[string]interface{}{
				"preference": "1qw",
				"exchange":   "mail.infoblox.com",
			},
			recordType:  "MX",
			errExpected: true,
		},
		"Valid Rdata - CAA Record": {
			input: map[string]interface{}{
				"flags": "0",
				"tag":   "issue",
				"value": "infoblox",
			},
			output: map[string]interface{}{
				"flags": 0,
				"tag":   "issue",
				"value": "infoblox",
			},
			recordType:  "CAA",
			errExpected: false,
		},
		"Invalid Rdata - CAA Record": {
			input: map[string]interface{}{
				"flags": "0qq",
				"tag":   "issue",
				"value": "infoblox",
			},
			output: map[string]interface{}{
				"flags": "0qq",
				"tag":   "issue",
				"value": "infoblox",
			},
			recordType:  "CAA",
			errExpected: true,
		},
		"Valid Rdata - SRV Record": {
			input: map[string]interface{}{
				"port":     "1234",
				"priority": "1",
				"target":   "infoblox",
				"weight":   "100",
			},
			output: map[string]interface{}{
				"port":     1234,
				"priority": 1,
				"target":   "infoblox",
				"weight":   100,
			},
			recordType:  "SRV",
			errExpected: false,
		},
		"Invalid Rdata - SRV Record": {
			input: map[string]interface{}{
				"port":     "1234qq",
				"priority": "1",
				"target":   "infoblox",
				"weight":   "100",
			},
			output: map[string]interface{}{
				"port":     "1234qq",
				"priority": 1,
				"target":   "infoblox",
				"weight":   100,
			},
			recordType:  "SRV",
			errExpected: true,
		},
		"Valid Rdata - SOA Record": {
			input: map[string]interface{}{
				"serial": "100",
				"mname":  "infoblox.com",
			},
			output: map[string]interface{}{
				"serial": 100,
				"mname":  "infoblox.com",
			},
			recordType:  "SOA",
			errExpected: false,
		},
		"Invalid Rdata - SOA Record": {
			input: map[string]interface{}{
				"serial": "1234qq",
			},
			output: map[string]interface{}{
				"serial": "1234qq",
			},
			recordType:  "SOA",
			errExpected: true,
		},
	}

	for tn, tc := range testData {
		t.Run(tn, func(t *testing.T) {
			output, err := updateDataRecordRData(tc.input, tc.recordType)
			assert.Equal(t, tc.errExpected, err.HasError())
			assert.Equal(t, tc.output, output)
		})
	}
}
