terraform {
  required_providers {
    datafy = {
      source = "brightml.org/terraform/datafy"
    }
  }
}

// Default provider configuration
provider "datafy" {

}

data "datafy_environment" "test" {
  id = "80342423-7538-4620-a7d0-fece6d279864"
}

output "test" {
  value = data.datafy_environment.test
}