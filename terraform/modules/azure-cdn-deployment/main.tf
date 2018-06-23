provider "azurerm" {
  subscription_id = "${var.azure_subscription_id}"
  client_id       = "${var.azure_client_id}"
  client_secret   = "${var.azure_client_secret}"
  tenant_id       = "${var.azure_tenant_id}"
}

resource "random_string" "resource_group_name" {
  length  = 8
  special = false
}

resource "random_string" "profile_name" {
  length = 8
}

//TODO: Resource group may not need to be created with each module
resource "azurerm_resource_group" "hideNsneak" {
  name     = "hideNsneak${random_string.resource_group_name.result}"
  count    = 1
  location = "${var.azure_location}"
}

resource "azurerm_cdn_profile" "hideNsneak" {
  name                = "${var.azure_cdn_profile_name}${random_string.profile_name.result}"
  location            = "${azurerm_resource_group.hideNsneak.location}"
  resource_group_name = "${azurerm_resource_group.hideNsneak.name}"
  sku                 = "Standard_Verizon"
}

resource "azurerm_cdn_endpoint" "hideNsneak" {
  name                = "${var.azure_cdn_endpoint_name}"
  profile_name        = "${azurerm_cdn_profile.hideNsneak.name}"
  location            = "${azurerm_resource_group.hideNsneak.location}"
  resource_group_name = "${azurerm_resource_group.hideNsneak.name}"

  querystring_caching_behaviour = "BypassCaching"

  optimization_type = "GeneralWebDelivery"

  origin {
    name      = "${var.azure_cdn_endpoint_name}"
    host_name = "${var.azure_cdn_hostname}"
  }
}
