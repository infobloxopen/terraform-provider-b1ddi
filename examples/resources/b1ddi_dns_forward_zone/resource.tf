terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}

resource "b1ddi_dns_forward_zone" "tf_example_forward_zone" {
  fqdn = "tf-example.com."
  view = b1ddi_dns_view.tf_example_dns_view.id
}
