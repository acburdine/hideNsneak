output "instance_id" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance_id}"
}

output "public_ip" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.network_interface.0.access_config.0.assigned_nat_ip}"
}

output "private_ip" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.network_interface.0.address}"
}

output "tags_fingerprint" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.tags_fingerprint}"
}

output "metadata" {
  value = "${google_compute_instance.ubuntu-xenial.*.instance.metadata_fingerprint}"
}
