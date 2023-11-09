terraform {
  required_providers {
    b1ddi = {
      version = "0.1.5"
      source  = "infobloxopen/b1ddi"
    }
  }
}

resource "b1ddi_dns_view" "tf_example_dns_view" {
  name = "example_tf_dns_view"
   tags = {
        TestType  = "Acceptance"
      }
}

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
