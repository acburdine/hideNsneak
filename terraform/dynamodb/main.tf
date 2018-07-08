provider "aws" {
  # access_key = "${var.aws_access_key}"
  # secret_key = "${var.aws_secret_key}"
  region = "us-east-1"
}

resource "aws_dynamodb_table" "terraform_statelock" {
  name           = "terraform-state-lock-dynamo"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
