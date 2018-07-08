
terraform {
	backend "s3" {
		bucket         = "hidensneak-terraform"
		key            = "filename.tfstate"
		dynamodb_table = "terraform-state-lock-dynamo"
		region         = "us-east-1"
		encrypt        = true
		access_key     = "AKIAJWJ5PH35DWANFUAA"
		secret_key     = "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"
		}
	  }
