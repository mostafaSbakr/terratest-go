terraform {
  required_version = ">= 1.3.7"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.42.0"
    }
  }
}

provider "azurerm" {
  subscription_id = "6b651eb2-7854-43ce-89b8-5490b1a07783"
  features {}
}

resource "azurerm_resource_group" "terratesting" {
  location = var.location
  name     = var.resource_group_name
}

