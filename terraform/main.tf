
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
	region_count         = "${map("",0)}"
	custom_ami           = ""
	aws_instance_type    = ""
	ec2_default_user     = ""
	aws_access_key       = "${var.aws_access_key}"
	aws_secret_key       = "${var.aws_secret_key}"
	aws_keypair_name     = "do_rsa"
	aws_private_key_file = ""
	aws_public_key_file  = ""
  }

  module "doDropletDeploy1" {
	  source              = "modules/droplet-deployment"
	  do_region_count     = "${map("AMS2",0, "AMS3",0, "BLR1",10, "FRA1",0, "LON1",0, "NYC1",2, "NYC2",0, "NYC3",0, "SFO1",0, "SFO2",0, "SGP1",0, "TOR1",0)}"
	  do_token            = "${var.do_token}"
	  do_image            = "ubuntu-16-04-x64"
	  do_private_key      = "/Users/mike.hodges/.ssh/do_rsa"
	  do_ssh_fingerprint  = "b3:b2:c7:b1:73:9e:28:c6:61:8d:15:e1:0e:61:7e:35"
	  do_size             = "512MB"
	  do_default_user     = "root"
  }
