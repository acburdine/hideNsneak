provider "google" {
  credentials = "${file("modules/gcp-deployment/inboxa90-f05affa600c1.json")}"
  project     = "${var.gcp_project}"
  region      = "${var.gcp_region}"
}

data "google_compute_zones" "available" {
  region = "${var.gcp_region}"
  status = "UP"
}

data "google_compute_image" "ubuntu_image" {
  family  = "ubuntu-1604-lts"
  project = "ubuntu-os-cloud"
}

resource "google_compute_instance" "ubuntu-xenial" {
  name         = "${var.gcp_region}-${count.index+1}"
  machine_type = "${var.gcp_machine_type}"
  zone         = "${data.google_compute_zones.available.names[count.index % length(data.google_compute_zones.available.names)]}"
  count        = "${var.gcp_instance_count}"

  boot_disk {
    initialize_params {
      image = "${var.gcp_image == "" ? data.google_compute_image.ubuntu_image.name : var.gcp_image}"
    }
  }

  network_interface {
    network       = "default"
    access_config = {}
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }

  metadata {
    sshKeys = "${var.gcp_ssh_user}:${file(var.gcp_ssh_pub_key_file)}"
  }
}

##This may need to be broken out into its own module
##if the changes are network-wideresource "google_compute_firewall" "default" {
resource "google_compute_firewall" "default" {
  name    = "test-firewall"
  network = "default"

  count = "${google_compute_instance.ubuntu-xenial.count > 0 ? 1 : 0}"

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_tags = ["ssh"]
}
