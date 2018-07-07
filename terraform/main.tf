terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

module "cloudfrontDeploy1" {
  source = "modules/cloudfront-deployment"

  aws_access_key = "${var.aws_access_key}"
  aws_secret_key = "${var.aws_secret_key}"

  cloudfront_origin  = "google.com"
  cloudfront_enabled = false
}

resource "null_resource" "test" {
  provisioner "local-exec" {
    command = "ssh -i /Users/mike.hodges/.ssh/do_rsa -D 8081 -o StrictHostKeyChecking=no -f -n ubuntu@18.232.107.249"
  }
}
