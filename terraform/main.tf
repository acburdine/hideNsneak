terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

module "ec2deploy1" {
  source          = "modules/ec2-deployment"
  default_sg_name = ""
  aws_sg_id       = ""

  #Example of region_count
  region_count         = "${map("us-east-1",0, "us-east-2",0)}"
  custom_ami           = ""
  aws_instance_type    = "t2.micro"
  ec2_default_user     = "ubuntu"
  aws_access_key       = "${var.aws_access_key}"
  aws_secret_key       = "${var.aws_secret_key}"
  aws_keypair_name     = "do_rsa"
  aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
  aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa.pub"
}

#   module "doDropletDeploy1" {
# 	  source              = "modules/droplet-deployment"
# 	  do_region_count     = "${map("blr1",2, "nyc1",8)}"
# 	  do_token            = "${var.do_token}"
# 	  do_image            = "ubuntu-16-04-x64"
# 	  do_private_key      = "/Users/mike.hodges/.ssh/do_rsa"
# 	  do_ssh_fingerprint  = "b3:b2:c7:b1:73:9e:28:c6:61:8d:15:e1:0e:61:7e:35"
# 	  do_size             = "512mb"
# 	  do_default_user     = "root"
#   }


#   module "doDropletDeploy2" {
# 	  source              = "modules/droplet-deployment"
# 	  do_region_count     = "${map("blr1",1, "nyc1",2)}"
# 	  do_token            = "${var.do_token}"
# 	  do_image            = "ubuntu-16-04-x64"
# 	  do_private_key      = "/Users/mike.hodges/.ssh/do_rsa"
# 	  do_ssh_fingerprint  = "b3:b2:c7:b1:73:9e:28:c6:61:8d:15:e1:0e:61:7e:35"
# 	  do_size             = "512mb"
# 	  do_default_user     = "rooter"
#   }

