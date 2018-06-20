resource "azure_hosted_service" "terraform-service" {
  name               = "terraform-service"
  location           = "North Europe"
  ephemeral_contents = false
  description        = "Hosted service created by Terraform."
  label              = "tf-hs-01"
}

resource "azure_instance" "default" {
  name     = "terraform-test"
  image    = "Ubuntu Server 14.04 LTS"
  size     = "Basic_A1"
  location = "West US"
  username = "${var.azure_username}"
  password = "${var.azure_password}"

  endpoint {
    name         = "SSH"
    protocol     = "tcp"
    public_port  = 22
    private_port = 22
  }
}
