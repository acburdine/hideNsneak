output "allRegions" {
  value = "${map("us-east-1", module.aws-us-east-1.region_info,
        "us-east-2",module.aws-us-east-2.region_info,
        "us-west-1",module.aws-us-west-1.region_info,
        "us-west-2",module.aws-us-west-2.region_info,
        "ca-central-1",module.aws-ca-central-1.region_info,
        "eu-west-1",module.aws-eu-west-1.region_info,
        "eu-west-2",module.aws-eu-west-2.region_info,
        "eu-west-3",module.aws-eu-west-3.region_info,
        "ap-northeast-1",module.aws-ap-northeast-1.region_info,
        "ap-northeast-2",module.aws-ap-northeast-2.region_info,
        "ap-southeast-1",module.aws-ap-southeast-1.region_info,
        "ap-southeast-2",module.aws-ap-southeast-2.region_info,
        "ap-south-1",module.aws-ap-south-1.region_info,
        "sa-east-1", module.aws-sa-east-1.region_info,
        )}"
}
