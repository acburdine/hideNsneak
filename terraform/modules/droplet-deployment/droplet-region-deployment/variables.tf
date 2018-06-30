variable "do_private_key" {
  default = ""
}

variable "do_ssh_fingerprint" {}

variable "do_image" {}

variable "do_region" {}

variable "do_size" {}

variable "do_count" {}

variable "do_default_user" {
  default = "root"
}

variable "do_firewall_name" {
  default = "default-ssh"
}

variable "ansible_groups" {
  default = []
}
