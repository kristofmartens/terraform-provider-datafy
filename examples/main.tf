// Provider config
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

// Environment config
resource "datafy_environment" "test" {
  name = "test"
}

output "terraform_env" {
  value = datafy_environment.test
}

data "datafy_environment" "env" {
  id = datafy_environment.test.id
}

output "env" {
  value = data.datafy_environment.env
}

// Project config
resource "datafy_project" "test" {
  name        = "test"
  description = "some description"
}

output "proj" {
  value = datafy_project.test
}

data "datafy_project" "project" {
  id = datafy_project.test.id
}

output "project" {
  value = data.datafy_project.project
}
