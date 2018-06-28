output "allRegions" {
  value = "${list(
        module.aws-us-east-1.region-map,
        module.aws-us-west-1.region-map)}"
}
