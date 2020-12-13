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

// Environment Data source
data "datafy_environment" "env" {
  id = "80342423-7538-4620-a7d0-fece6d279864"
}

output "env" {
  value = data.datafy_environment.env
}

// Environments Data source
data "datafy_environments" "envs" {}

output "envs" {
  value = data.datafy_environments.envs
}

data "datafy_project" "project" {
  id = "b1955c1f-d6bb-4154-8baa-3de5b8963792"
}

output "project" {
  value = data.datafy_project.project
}