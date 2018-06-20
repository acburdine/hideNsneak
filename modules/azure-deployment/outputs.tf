output "azure_instance_id" {
  value = "${azurerm_virtual_machine.test.*.id}"
}

output "azure_ip_address" {
  value = "${azurerm_public_ip.public_ip.*.ip_address}"
}

output "azure_resource_group_id" {
  value = "${azurerm_resource_group.test.*.id}"
}
