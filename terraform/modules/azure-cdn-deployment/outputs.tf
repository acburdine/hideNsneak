output "target_hostname" {
  value = "${var.azure_cdn_hostname}"
}

output "endpoint_name" {
  value = "${var.azure_cdn_endpoint_name}.azureedge.net"
}
