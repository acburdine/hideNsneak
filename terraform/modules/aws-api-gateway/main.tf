provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_api_region}"
}

resource "random_string" "api_name" {
  length  = 8
  special = false
}

resource "aws_api_gateway_rest_api" "hideNsneak" {
  name        = "hideNsneak-${random_string.api_name.result}"
  description = "hideNsneak gateway"
}

resource "aws_api_gateway_resource" "hideNsneak" {
  rest_api_id = "${aws_api_gateway_rest_api.hideNsneak.id}"
  parent_id   = "${aws_api_gateway_rest_api.hideNsneak.root_resource_id}"
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "hideNsneak" {
  rest_api_id   = "${aws_api_gateway_rest_api.hideNsneak.id}"
  resource_id   = "${aws_api_gateway_resource.hideNsneak.id}"
  http_method   = "ANY"
  authorization = "NONE"

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

resource "aws_api_gateway_integration" "hideNsneak" {
  rest_api_id             = "${aws_api_gateway_rest_api.hideNsneak.id}"
  resource_id             = "${aws_api_gateway_resource.hideNsneak.id}"
  http_method             = "${aws_api_gateway_method.hideNsneak.http_method}"
  integration_http_method = "ANY"
  type                    = "HTTP_PROXY"
  uri                     = "${var.aws_api_target_uri}{proxy}"

  request_parameters = {
    "integration.request.path.proxy" = "method.request.path.proxy"
  }
}

resource "aws_api_gateway_deployment" "hideNsneak" {
  rest_api_id = "${aws_api_gateway_rest_api.hideNsneak.id}"
  stage_name  = "${var.aws_api_stage_name}"

  variables {
    deployed_at = "${timestamp()}"
  }

  depends_on = ["aws_api_gateway_integration.hideNsneak", "aws_api_gateway_method.hideNsneak", "aws_api_gateway_rest_api.hideNsneak", "aws_api_gateway_resource.hideNsneak"]
}
