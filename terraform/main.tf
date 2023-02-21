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
  location = "westeurope"
  name     = "terratesting"
}

output "resource_group_name" {
  value = azurerm_resource_group.terratesting.name
}

provider "azurerm" {
  features {}
}
