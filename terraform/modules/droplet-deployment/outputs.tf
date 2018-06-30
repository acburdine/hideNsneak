output allRegions {
  value = "${list(
      module.do-nyc1.region_info,
module.do-nyc2.region_info,
module.do-nyc3.region_info,
module.do-sfo1.region_info,
module.do-sfo2.region_info,
module.do-sgp1.region_info,
module.do-tor1.region_info,
module.do-ams2.region_info,
module.do-ams3.region_info,
module.do-blr1.region_info,
module.do-fra1.region_info,
module.do-lon1.region_info,
  )}"
}
