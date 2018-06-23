provider "aws" {
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

