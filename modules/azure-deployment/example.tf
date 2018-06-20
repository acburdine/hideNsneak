provider "azurerm" {}

resource "azurerm_virtual_machine" "helloterraformvm" {
  name                  = "terraformvm"
  location              = "West US"
  resource_group_name   = "${azurerm_resource_group.helloterraform.name}"
  network_interface_ids = ["${azurerm_network_interface.helloterraformnic.id}"]
  vm_size               = "Standard_A0"

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "14.04.2-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name          = "myosdisk"
    vhd_uri       = "${azurerm_storage_account.helloterraformstorage.primary_blob_endpoint}${azurerm_storage_container.helloterraformstoragestoragecontainer.name}/myosdisk.vhd"
    caching       = "ReadWrite"
    create_option = "FromImage"
  }

  os_profile {
    computer_name  = "hostname"
    admin_username = "${var.azure_username}"
    admin_password = "${var.azure_password}"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }
}
