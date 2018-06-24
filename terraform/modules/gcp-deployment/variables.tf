variable "gcp_region" {}

variable "gcp_machine_type" {
  default = "f1-micro"
}

variable "gcp_image" {
  default = ""
}

variable "gcp_instance_count" {}

variable "gcp_project" {}
variable "gcp_ssh_user" {}
variable "gcp_ssh_pub_key_file" {}

variable "gcp_ssh_private_key_file" {}

variable "ansible_groups" {
  default = []
}
