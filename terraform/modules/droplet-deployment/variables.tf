variable "do_region_count_template" {
  type = "map"

  default = {
    "NYC1" = 0
    "NYC2" = 0
    "NYC3" = 0
    "SFO1" = 0
    "SFO2" = 0
    "SGP1" = 0
    "TOR1" = 0
    "AMS2" = 0
    "AMS3" = 0
    "BLR1" = 0
    "FRA1" = 0
    "LON1" = 0
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

variable "do_firewall_name" {
  default = "default-ssh"
}

variable "do_region_count" {
  type = "map"
}
