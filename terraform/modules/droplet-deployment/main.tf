provider "digitalocean" {
  token = "${var.do_token}"
}

locals {
  do_region_count_final = "${merge(var.do_region_count_template, var.do_region_count)}"
}

module "do-nyc1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["NYC1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "NYC1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

# //NYC2 is basically full
module "do-nyc2" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["NYC2"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "NYC2"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-nyc3" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["NYC3"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "NYC3"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

# Basically Full
module "do-ams2" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["AMS2"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "AMS2"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-ams3" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["AMS3"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "AMS3"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-blr1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["BLR1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "BLR1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-fra1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["FRA1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "FRA1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-lon1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["LON1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "LON1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

# Basically Full
module "do-sfo1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["SFO1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "SFO1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-sfo2" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["SFO2"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "SFO2"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-tor1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["TOR1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "TOR1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}

module "do-sgp1" {
  source             = "droplet-region-deployment"
  do_image           = "${var.do_image}"
  do_size            = "${var.do_size}"
  do_count           = "${local.do_region_count_final["SGP1"]}"
  do_ssh_fingerprint = "${var.do_ssh_fingerprint}"
  do_private_key     = "${var.do_private_key}"
  do_default_user    = "${var.do_default_user}"
  do_region          = "SGP1"

  //Think about seperating firewalls
  do_firewall_name = "${var.do_firewall_name}"
}
