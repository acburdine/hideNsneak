output "instance_id" {
  value = "${azurerm_virtual_machine.test.*.id}"
}

output "ip_address" {
  value = "${azurerm_public_ip.public_ip.*.ip_address}"
}

output "resource_group_id" {
  value = "${azurerm_resource_group.test.*.id}"
}
