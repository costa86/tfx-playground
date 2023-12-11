terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

variable "name_prefix" {
  type    = string
  default = "hashi"
}

variable "name_count" {
  type    = number
  default = 4
}

resource "random_pet" "name" {
  prefix = var.name_prefix
  length = var.name_count
}

output "name" {
  value = random_pet.name.id
}
