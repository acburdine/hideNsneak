provider "azurerm" {
  subscription_id = "${var.azure_subscription_id}"
  client_id       = "${var.azure_client_id}"
  client_secret   = "${var.azure_client_secret}"
  tenant_id       = "${var.azure_tenant_id}"
}

resource "azurerm_resource_group" "test" {
  name     = "acctestrg"
  count    = 1
  location = "${var.azure_location}"
}

resource "azurerm_cdn_profile" "test" {
  name                = "${var.azure_cdn_profile_name}"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  sku                 = "Standard_Verizon"
}

resource "azurerm_cdn_endpoint" "test" {
  name                = "${var.azure_cdn_endpoint_name}"
  profile_name        = "${azurerm_cdn_profile.test.name}"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  querystring_caching_behaviour = "BypassCaching"

  optimization_type = "GeneralWebDelivery"

  origin {
    name      = "${var.azure_cdn_endpoint_name}"
    host_name = "${var.azure_cdn_hostname}"
  }
}
