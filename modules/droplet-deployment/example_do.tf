provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_droplet" "default" {
  image  = "ubuntu-14-04-x64"
  name   = "example-droplet"
  region = "nyc2"
  size   = "512mb"

  #   ssh_keys = [
  #     "${var.ssh_fingerprint}",
  #   ]

  #   connection {
  #     user        = "root"
  #     type        = "ssh"
  #     private_key = "${file(var.pvt_key)}"
  #     timeout     = "2m"
  #   }
}
