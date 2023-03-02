output "server_name" {
  value = data.azurerm_postgresql_server.postgresql.name
}

output "admin_login" {
  value = data.azurerm_postgresql_server.postgresql.administrator_login
}
