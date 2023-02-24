module "storage-account" {
  source              = "../resource/storage-account"
  resource_group_name = module.resource-group.resource_group_name
}

module "resource-group" {
  source = "../resource/resource-group"
}

module "eventhub" {
  source              = "../resource/eventhub"
  resource_group_name = module.resource-group.resource_group_name
}