# ########AZURE###########
# output "azure_instance_id" {
#   value = "${module.azure-example-1.instance_id}"
# }
# output "azure_ip_address" {
#   value = "${module.azure-example-1.ip_address}"
# }
# output "azure_resource_group_id" {
#   value = "${module.azure-example-1.resource_group_id}"
# }
# ########DO###########
# output "do_instance_id" {
#   value = "${module.do-example-1.instance_id}"
# }
# output "do_region" {
#   value = "${module.do-example-1.region}"
# }
# output "do_ipv4_address" {
#   value = "${module.do-example-1.ipv4_address}"
# }
# output "do_status" {
#   value = "${module.do-example-1.status}"
# }
########AWS###########
# output "domainfront_url" {
#   value = "${module.cloudfront.domainfront_url}"
# }
output "providers" {
  value = "${map(
    "AWS", map(
      "instances", list(map()),
      "security_group", list(map()), 
      "api", list(map()),
      "domain_front", list(map())),
    "DO", map(
      "instances", concat(module.doDropletDeploy1.allRegions),
      "firewalls", list(map())),
    "GOOGLE", map(
      "instances", list(map())),
    "AZURE", map(
      "instances", list(map())))
    }"
}
