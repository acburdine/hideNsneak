provider "azurerm" {
  subscription_id = "${var.azure_subscription_id}"
  client_id       = "${var.azure_client_id}"
  client_secret   = "${var.azure_client_secret}"
  tenant_id       = "${var.azure_tenant_id}"
}

resource "azurerm_resource_group" "test" {
  name     = "acctestrg2"
  count    = 1
  location = "${var.azure_location}"
}

resource "azurerm_public_ip" "public_ip" {
  count                        = "${azurerm_virtual_machine.test.count}"
  name                         = "tester"
  location                     = "${azurerm_resource_group.test.location}"
  resource_group_name          = "${azurerm_resource_group.test.name}"
  public_ip_address_allocation = "static"
}

resource "azurerm_virtual_network" "test" {
  name                = "acctvn"
  count               = "${var.azure_instance_count > 0 ? 1 : 0}"
  address_space       = ["10.0.0.0/16"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}

resource "azurerm_subnet" "test" {
  name                 = "acctsub"
  count                = 0
  resource_group_name  = "${azurerm_resource_group.test.name}"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_network_interface" "test" {
  name                = "acctni"
  count               = "${var.azure_instance_count > 0 ? 1 : 0}"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  ip_configuration {
    name                          = "testconfiguration1"
    subnet_id                     = "${azurerm_subnet.test.id}"
    private_ip_address_allocation = "dynamic"
  }
}

resource "azurerm_managed_disk" "test" {
  name                 = "datadisk_existing"
  count                = "${var.azure_instance_count > 0 ? 1 : 0}"
  location             = "${azurerm_resource_group.test.location}"
  resource_group_name  = "${azurerm_resource_group.test.name}"
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "10"
}

#azurerm_virtual_machine.test.*.ip_adress
resource "azurerm_virtual_machine" "test" {
  name                  = "acctvm"
  count                 = "${var.azure_instance_count}"
  location              = "${azurerm_resource_group.test.location}"
  resource_group_name   = "${azurerm_resource_group.test.name}"
  network_interface_ids = ["${azurerm_network_interface.test.id}"]
  vm_size               = "${var.azure_vm_size}"

  # Uncomment this line to delete the OS disk automatically when deleting the VM
  # delete_os_disk_on_termination = true


  # Uncomment this line to delete the data disks automatically when deleting the VM
  # delete_data_disks_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
  storage_os_disk {
    name              = "myosdisk1"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }
  os_profile {
    computer_name  = "hostname"
    admin_username = "testadmin"
    admin_password = "Password1234!"
  }
  os_profile_linux_config {
    disable_password_authentication = true

    ssh_keys {
      path     = "/home/testadmin/.ssh/authorized_keys"
      key_data = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDI38w64qILMmXfcZHyDc6h0ApN+XMSbRH69dHY9JeDKpsJeIsaI0L1GdWEXJl0stetQ3wjjnKQF5D9nNrZ4H9nusTtE2D65Zek/W2JlUFLo2ayji2MBQ0uh4Rn6MR9/TnD/PdcB6z52a5SvCv7ngytab7Lhnx416kya6zRwiBkJYbHarDAer6i5edA7XO7nHfqWFjzgWS3scQdmxhHTdQ+Keg4BM7VHa2xTLB7BaH2POwlBbM9UdFxhQbnj9ErQokAPI1mVlE1CkYP4d6SPU+UEzf0rVujhGdSek9H2EnoqKHMN2yDyHgk13NiBVux9pM3pjkhCVlZ+Wn3su0JhscN"
    }
  }
  tags {
    environment = "test"
  }
}
