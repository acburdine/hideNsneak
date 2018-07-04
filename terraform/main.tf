
	terraform {
		backend "s3" {
		  bucket         = "hidensneak-terraform"
		  key            = "filename.tfstate"
		  dynamodb_table = "terraform-state-lock-dynamo"
		  region         = "us-east-1"
		  encrypt        = true
		}
	  }

	module "ec2Deploy1" {
	source          = "modules/ec2-deployment"
  
	default_sg           = ""
	aws_sg_id            = ""

	region_count         = "${map("us-east-1",4, "us-east-2",2, "us-west-1",1, "us-west-2",1)}"
	aws_instance_type    = "t2.micro"
	ec2_default_user     = "ubuntu"
	aws_access_key       = "${var.aws_access_key}"
	aws_secret_key       = "${var.aws_secret_key}"
	aws_keypair_name     = ""
	aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
	aws_public_key_file  = ""
  }
