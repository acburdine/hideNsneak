variable "do_private_key" {
  default = ""
}

variable "do_ssh_fingerprint" {}

variable "do_image" {}

variable "do_region" {}

variable "do_size" {}

variable "name" {
  default = "hidensneak"
}

variable "do_count" {
  default = 0
}

variable "do_default_user" {
  default = "root"
}

# variable "do_firewall_name" {
#   default = "default-ssh"
# }

variable "ansible_groups" {
  default = []
}
