data "azurerm_postgresql_server" "postgresql" {
  name                = "icd-cluster-test-mmfxtblz"
  resource_group_name = "icd-mips-pp"
}