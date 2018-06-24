variable "azure_tenant_id" {}

variable "azure_client_id" {}

variable "azure_client_secret" {}

variable "azure_location" {}

variable "azure_subscription_id" {}

variable "azure_instance_count" {
  default = 0
}

variable "azure_default_username" {
  default = "hidensneak"
}

variable "azure_vm_size" {
  default = "Standard_F2"
}

variable "azure_environment" {
  default = "hideNsneak"
}

variable "azure_public_key_file" {}

variable "azure_private_key_file" {}

variable "ansible_groups" {
  default = []
}
