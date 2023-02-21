remote_state {
  backend = "azurerm"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite"
  }

  config  = {
    resource_group_name  = "terratest-state"
    storage_account_name = "terratestremotestate"
    container_name       = "remotestate"
    key                  = "terraform.tfstate"
  }
}

generate "provider" {
  path      = "./provider.tf"
  if_exists = "overwrite_terragrunt"
  contents  = <<EOF
    provider "azurerm" {
      subscription_id = "6b651eb2-7854-43ce-89b8-5490b1a07783"
      features {}
    }
EOF
}