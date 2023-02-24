terraform {
  required_version = ">= 1.3.7"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.42.0"
    }
  }
}

resource "azurerm_eventhub_namespace" "eventhub_namespace" {
  location            = var.location
  name                = "terratesting-eventhub-namespace"
  resource_group_name = var.resource_group_name
  sku                 = "Standard"
}

resource "azurerm_eventhub" "eventhub" {
  message_retention   = 7
  name                = "terratesting-eventhub"
  namespace_name      = azurerm_eventhub_namespace.eventhub_namespace.name
  partition_count     = 1
  resource_group_name = var.resource_group_name
}

resource "azurerm_eventhub_authorization_rule" "eventhub_authorization_rule" {
  eventhub_name       = azurerm_eventhub.eventhub.name
  name                = "eventhub-authorization-rule"
  namespace_name      = azurerm_eventhub_namespace.eventhub_namespace.name
  resource_group_name = var.resource_group_name

  listen = true
  send   = false
  manage = false
}