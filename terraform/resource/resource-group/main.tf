terraform {
  required_version = ">= 1.3.7"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.42.0"
    }
  }
}

resource "azurerm_resource_group" "terratesting" {
  location = var.location
  name     = var.resource_group_name
}

