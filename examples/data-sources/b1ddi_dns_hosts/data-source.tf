terraform {
  required_providers {
    b1ddi = {
      source = "infobloxopen/b1ddi"
    }
  }
}

# Get all DNS hosts
data "b1ddi_dns_hosts" "all_hosts" {}

# Get DNS Host by name
data "b1ddi_dns_hosts" "dns_host_by_name" {
  filters = {
    name = "dns_host_name"
  }
}

data "b1ddi_dns_hosts" "all_dns_host_by_tags" {
  tfilters = {
    location = "site1"
  }
}