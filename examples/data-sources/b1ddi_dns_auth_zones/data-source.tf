terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Get DNS Auth Zone with the specified FQDN
data "b1ddi_dns_auth_zones" "example_auth_zone" {
  filters = {
    fqdn = "tf-example.com."
  }
}

# Get all DNS Auth Zones
data "b1ddi_dns_auth_zones" "all_dns_auth_zones" {}

# Get DNS Auth Zone with the specific tags
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