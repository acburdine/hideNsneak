output "allRegions" {
  value = "${list(module.aws-us-east-1.region_info,
        module.aws-us-east-2.region_info,
        module.aws-us-west-1.region_info,
        module.aws-us-west-2.region_info,
        module.aws-ca-central-1.region_info,
        module.aws-eu-west-1.region_info,
        module.aws-eu-west-2.region_info,
        module.aws-eu-west-3.region_info,
        module.aws-ap-northeast-1.region_info,
        module.aws-ap-northeast-2.region_info,
        module.aws-ap-southeast-1.region_info,
         module.aws-ap-southeast-2.region_info,
        module.aws-ap-south-1.region_info,
         module.aws-sa-east-1.region_info,
        )}"
}
