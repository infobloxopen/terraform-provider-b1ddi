terraform {
  required_providers {
    b1ddi = {
      version = "0.1"
      source  = "infobloxopen/b1ddi"
    }
  }
}

# Select all PTR DNS Records
data "b1ddi_dns_records" "ptr_dns_records" {
  filters = {
    "type" = "PTR"
  }
}

# Select all A DNS Records
data "b1ddi_dns_records" "a_dns_records" {
  filters = {
    "type" = "A"
  }
}

# Select DNS Record with specified name in zone
data "b1ddi_dns_records" "tf_example_a_dns_record" {
  filters = {
    "name_in_zone" = "tf_example_a_record"
  }
}

# Get all DNS Records
data "b1ddi_dns_records" "all_dns_records" {}

# Get all DNS Records with specific tags
data "b1ddi_dns_records" "all_dns_records_with_tags" {
       tfilters = {
            TestType  = "Acceptance"
          }
}

output "records" {
  value= data.b1ddi_dns_records.all_dns_records.results
}

output "records_with_tags" {
  value= data.b1ddi_dns_records.all_dns_records_with_tags.results
}