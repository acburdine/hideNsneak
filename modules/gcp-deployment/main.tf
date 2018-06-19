provider "google" {
  credentials = "${file("modules/gcp-deployment/inboxa90-f05affa600c1.json")}"
  project     = "inboxa90"
  region      = "${var.gcp_region}"
}

data "google_compute_zones" "available" {
  region = "${var.gcp_region}"
  status = "UP"
}

data "google_compute_image" "my_image" {
  name    = "debian-7-wheezy-v20140606"
  project = "debian-cloud"
}

resource "google_compute_instance" "ubuntu-xenial" {
  name         = "${var.gcp_region}-${count.index+1}"
  machine_type = "f1-micro"
  zone         = "${data.google_compute_zones.available.names[count.index % length(data.google_compute_zones.available.names)]}"
  count        = 0

  boot_disk {
    initialize_params {
      image = "ubuntu-1604-lts"
    }
  }

  network_interface {
    network       = "default"
    access_config = {}
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}
