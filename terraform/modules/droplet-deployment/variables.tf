variable "do_token" {}

variable "pvt_key" {}
variable "ssh_fingerprint" {}

variable "do_image" {}

variable "do_region" {}

variable "do_size" {}

variable "do_count" {}

variable "do_ssh_source_ip" {
  default = "0.0.0.0/0"
}
