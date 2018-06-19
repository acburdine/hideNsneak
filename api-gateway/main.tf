terraform {
  backend "s3" {
    bucket         = "hidensneak-terraform"
    key            = "filename.tfstate"
    dynamodb_table = "terraform-state-lock-dynamo"
    region         = "us-east-1"
    encrypt        = true
  }
}

provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-east-1"
}

resource "aws_api_gateway_rest_api" "gc-sync" {
  name        = "gc-sync"
  description = "Test of terraform API Gateway"
}

resource "aws_api_gateway_resource" "gc-sync" {
  rest_api_id = "${aws_api_gateway_rest_api.gc-sync.id}"
  parent_id   = "${aws_api_gateway_rest_api.gc-sync.root_resource_id}"
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "gc-sync" {
  rest_api_id   = "${aws_api_gateway_rest_api.gc-sync.id}"
  resource_id   = "${aws_api_gateway_resource.gc-sync.id}"
  http_method   = "ANY"
  authorization = "NONE"

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

resource "aws_api_gateway_integration" "gc-sync" {
  rest_api_id             = "${aws_api_gateway_rest_api.gc-sync.id}"
  resource_id             = "${aws_api_gateway_resource.gc-sync.id}"
  http_method             = "${aws_api_gateway_method.gc-sync.http_method}"
  integration_http_method = "ANY"
  type                    = "HTTP_PROXY"
  uri                     = "https://gmrcmail.gmrc.com/owa/{proxy}"

  request_parameters = {
    "integration.request.path.proxy" = "method.request.path.proxy"
  }
}

resource "aws_api_gateway_deployment" "instance" {
  rest_api_id = "${aws_api_gateway_rest_api.gc-sync.id}"
  stage_name  = "tester1"

  variables {
    deployed_at = "${timestamp()}"
  }

  depends_on = ["aws_api_gateway_integration.gc-sync", "aws_api_gateway_method.gc-sync", "aws_api_gateway_rest_api.gc-sync", "aws_api_gateway_resource.gc-sync"]
}
