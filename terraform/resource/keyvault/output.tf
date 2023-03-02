output "value" {
  value = data.azurerm_key_vault_secret.key_vault_secret.value
  sensitive = true
}

output "version" {
  value = data.azurerm_key_vault_secret.key_vault_secret.version
}