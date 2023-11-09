terraform {
  required_providers {
    b1ddi = {
      version = "0.1.5"
      source  = "infobloxopen/b1ddi"
    }
  }
}

variable "internal_secondary_name" {
  type = string
  description = "Internal secondary DNS Host name for the DNS Auth Zone configuration"
}

# Get DNS Host by name
data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    name = var.internal_secondary_name
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
}

resource "b1ddi_dns_auth_zone" "tf_example_auth_zone" {
  internal_secondaries {
    host = data.b1ddi_dns_hosts.dns_host_by_name.results.0.id
  }
  fqdn = "tf-example.com."
  primary_type = "cloud"
  view = b1ddi_dns_view.tf_example_dns_view.id
    tags = {
      TestType  = "Acceptance"
    }
}

data "b1ddi_dns_auth_zones" "all_auth_zones" {}

data "b1ddi_dns_auth_zones" "all_auth_zones_with_tags" {
       tfilters = {
            TestType  = "Acceptance"
          }
}

output "zones" {
  value= data.b1ddi_dns_auth_zones.all_auth_zones.results
}

output "zones_with_tags" {
  value= data.b1ddi_dns_auth_zones.all_auth_zones_with_tags.results
}