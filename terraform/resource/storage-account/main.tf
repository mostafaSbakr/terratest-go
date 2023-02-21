terraform {
  required_version = ">= 1.3.7"
  required_providers {
    azurerm = {
      version = "=3.42.0"
    }
  }
}

provider "azurerm" {
  subscription_id = "6b651eb2-7854-43ce-89b8-5490b1a07783"
  features {}
}

resource "azurerm_storage_account" "storage_account" {
  name                      = var.storage_account_name
  location                  = var.location
  resource_group_name       = var.resource_group_name
  account_tier              = "Standard"
  account_replication_type  = "LRS"
  min_tls_version           = "TLS1_2"
  enable_https_traffic_only = true
  large_file_share_enabled  = true

  account_kind = "StorageV2"
  access_tier  = var.access_tier

  blob_properties {
    versioning_enabled = true

    delete_retention_policy {
      days = 365
    }
  }

  tags = {
    environment = "${var.resource_group_name}-storage-account"
  }
}