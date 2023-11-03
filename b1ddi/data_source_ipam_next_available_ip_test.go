package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcNaIP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetNAIPBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_address" "tf_acc_test_address" {
						address = b1ddi_subnet.tf_acc_test_subnet.id
						comment = "This Address is created by subnet NAIP terraform provider acceptance test"
						space = b1ddi_ip_space.tf_acc_test_space.id
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_addresses.tf_acc_addresses", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_addresses.tf_acc_addresses", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_addresses.tf_acc_addresses", "results.0.comment", "This Address is created by subnet NAIP terraform provider acceptance test"),
				),
			},
		},
	})
}
