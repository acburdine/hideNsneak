output "gcp_instance_id" {
  value = "${google_compute_instance.ubuntu-xenial.instance.*.id}"
}

output "gcp_public_ip" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.network_interface.0.access_config.0.assigned_nat_ip}"
}

output "gcp_private_ip" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.network_interface.0.address}"
}

output "gcp_tags_fingerprint" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.tags_fingerprint}"
}

output "gcp_metadata" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.metadata_fingerprint}"
}
