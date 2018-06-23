output "aws_api_target_uri" {
  value = "${var.aws_api_target_uri}"
}

output "aws_api_invoke_uri" {
  value = "${aws_api_gateway_deployment.hideNsneak.invoke_url}"
}

output "aws_api_name" {
  value = "${aws_api_gateway_rest_api.name}"
}
