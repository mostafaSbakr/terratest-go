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
  subscription_id = "9eabba05-660d-46ac-8b27-3b30c7c9420c"
  features {}
}

data "azurerm_key_vault_secret" "key_vault_secret" {
  key_vault_id = "/subscriptions/9eabba05-660d-46ac-8b27-3b30c7c9420c/resourceGroups/icd-mips-pp/providers/Microsoft.KeyVault/vaults/mipspp-icd-keyvault-6KeG"
  name         = "icd-automount-test-2-apicurio-admin"
}