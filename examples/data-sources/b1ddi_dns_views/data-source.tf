terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Select DNS View by name
data "b1ddi_dns_views" "example_dns_view" {
  filters = {
    "name" = "example_tf_dns_view"
  }
}

# Get all DNS Views
data "b1ddi_dns_views" "all_dns_views" {}

# Get all DNS Views with specific tags
data "b1ddi_dns_views" "all_dns_views" {}

data "b1ddi_dns_views" "all_dns_views_with_tags" {
       tfilters = {
            TestType  = "Acceptance"
          }
}

output "views" {
  value= data.b1ddi_dns_views.all_dns_views.results
}

output "views_with_tags" {
  value= data.b1ddi_dns_views.all_dns_views_with_tags.results
}