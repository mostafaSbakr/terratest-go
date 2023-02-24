variable "storage_account_name" {
  default = "terratestingstorage"
}
variable "location" {
  default = "westeurope"
}
variable "resource_group_name" {
  type = string
}
variable "access_tier" {
  default = "Hot"
}