variable "do_region_count_template" {
  type = "map"

  default = {
    "nyc1" = 0
    "nyc2" = 0
    "nyc3" = 0
    "sfo1" = 0
    "sfo2" = 0
    "sgp1" = 0
    "tor1" = 0
    "ams2" = 0
    "ams3" = 0
    "blr1" = 0
    "fra1" = 0
    "lon1" = 0
  }
}

variable "do_token" {}

variable "do_private_key" {}
variable "do_ssh_fingerprint" {}

variable "do_image" {}

variable "do_size" {}

variable "do_default_user" {
  default = "root"
}

# variable "do_firewall_name" {
#   default = "default-ssh"
# }

variable "do_region_count" {
  type = "map"
}
