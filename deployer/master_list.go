package deployer

import (
	"bytes"
	"html/template"
	"strconv"
	"strings"
)

func retrieveUserAndPrivateKey(module ModuleState) (privateKey string, user string) {
	for _, resource := range module.Resources {
		if resource.Type == "ansible_host" {
			privateKey = resource.Primary.Attributes["vars.ansible_ssh_private_key_file"].(string)
			user = resource.Primary.Attributes["vars.ansible_user"].(string)
			break
		}
	}
	return
}

////////////
//AWS API///
////////////
func createAWSAPIFromState(modules []ModuleState) (awsAPIConfigWrappers []AWSApiConfigWrapper, moduleCount int) {
	for _, module := range modules {
		if len(module.Path) > 1 && strings.Contains(module.Path[1], "awsAPIDeploy") {
			moduleCountString := strings.Split(module.Path[1], "awsAPIDeploy")[1]
			tempInt, _ := strconv.Atoi(moduleCountString)
			if moduleCount < tempInt {
				moduleCount = tempInt
			}
			var tempConfig AWSApiConfigWrapper
			for _, resource := range module.Resources {
				tempConfig.ModuleName = module.Path[1]
				switch resource.Type {
				case "aws_api_gateway_deployment":
					tempConfig.InvokeURI = resource.Primary.Attributes["invoke_url"].(string)
				case "aws_api_gateway_integration":
					tempConfig.TargetURI = resource.Primary.Attributes["uri"].(string)
				case "aws_api_gateway_rest_api":
					tempConfig.Name = resource.Primary.Attributes["name"].(string)
				default:
					continue
				}
			}
		}
	}
	return
}

/////////////////////
//Google Domain Front
/////////////////////

func createGooglefrontFromState(modules []ModuleState) (googlefrontConfigWrappers []GooglefrontConfigWrapper, moduleCount int) {
	for _, module := range modules {
		if len(module.Path) > 1 && strings.Contains(module.Path[1], "googlefrontDeploy") {
			moduleCountString := strings.Split(module.Path[1], "googlefrontDeploy")[1]
			tempInt, _ := strconv.Atoi(moduleCountString)
			if moduleCount < tempInt {
				moduleCount = tempInt
			}
			var tempConfig GooglefrontConfigWrapper
			tempConfig.ModuleName = module.Path[1]
			for _, resource := range module.Resources {
				if resource.Type == "google_cloudfunctions_function" {

					labels := resource.Primary.Attributes["labels"].(map[string]string)

					tempConfig.Enabled, _ = strconv.ParseBool(resource.Primary.Attributes["trigger_http"].(string))
					tempConfig.InvokeURI = resource.Primary.Attributes["https_trigger_url"].(string)
					tempConfig.FunctionName = resource.Primary.Attributes["name"].(string)
					tempConfig.HostURL = labels["target"]
					tempConfig.RestrictUA = labels["restrictua"]
					tempConfig.RestrictSubnet = labels["restrictsubnet"]
					tempConfig.RestrictHeader = labels["restrictheader"]
					tempConfig.RestrictHeaderValue = labels["restrictheadervalue"]

					googlefrontConfigWrappers = append(googlefrontConfigWrappers, tempConfig)
				}
			}
		}
	}
	return
}

////////////////
//AWS Cloudfront
////////////////

func createCloudfrontFromState(modules []ModuleState) (cloudfrontConfigWrappers []CloudfrontConfigWrapper, moduleCount int) {
	for _, module := range modules {
		if len(module.Path) > 1 && strings.Contains(module.Path[1], "cloudfrontDeploy") {
			moduleCountString := strings.Split(module.Path[1], "cloudfrontDeploy")[1]
			tempInt, _ := strconv.Atoi(moduleCountString)
			if moduleCount < tempInt {
				moduleCount = tempInt
			}
			var tempConfig CloudfrontConfigWrapper
			tempConfig.ModuleName = module.Path[1]
			for _, resource := range module.Resources {
				if resource.Type == "aws_cloudfront_distribution" {
					tempConfig.Status = resource.Primary.Attributes["status"].(string)
					tempConfig.ID = resource.Primary.Attributes["id"].(string)
					tempConfig.Etag = resource.Primary.Attributes["etag"].(string)
					for key, value := range resource.Primary.Attributes {
						if strings.Contains(key, "domain_name") {
							if strings.Contains(key, "origin") {
								tempConfig.Origin = value.(string)
							} else {
								tempConfig.URL = value.(string)
							}
						}

						tempConfig.Status = resource.Primary.Attributes["status"].(string)
						tempConfig.Enabled = resource.Primary.Attributes["enabled"].(string)
					}
					cloudfrontConfigWrappers = append(cloudfrontConfigWrappers, tempConfig)
				}
			}
		}
	}
	return
}

////////////
//EC2///////
////////////
func returnInitialEC2Config(module ModuleState) (tempConfig EC2ConfigWrapper) {
	privateKey, user := retrieveUserAndPrivateKey(module)

	tempConfig.RegionMap = make(map[string]int)
	for _, resource := range module.Resources {
		if resource.Type == "aws_instance" {
			availZone := resource.Primary.Attributes["availability_zone"].(string)
			region := availZone[:len(availZone)-1]
			tempConfig.KeyPairName = resource.Primary.Attributes["key_name"].(string)
			tempConfig.ModuleName = module.Path[1]
			tempConfig.InstanceType = resource.Primary.Attributes["instance_type"].(string)
			tempConfig.DefaultUser = user
			tempConfig.PrivateKey = privateKey
			tempConfig.RegionMap[region] = 1
			break
		}
	}
	return
}

func createEC2ConfigFromState(modules []ModuleState) (ec2Configs []EC2ConfigWrapper, maxModuleCount int) {
	for _, module := range modules {
		if len(module.Path) > 2 && strings.Contains(module.Path[1], "ec2Deploy") {
			for _, resource := range module.Resources {
				if resource.Type == "aws_instance" {
					availZone := resource.Primary.Attributes["availability_zone"].(string)
					region := availZone[:len(availZone)-1]

					countString := strings.Split(module.Path[1], "ec2Deploy")[1]
					countInt, _ := strconv.Atoi(countString)
					if countInt > maxModuleCount {
						maxModuleCount = countInt
					}
					//If the list is empty, return the first element found
					if len(ec2Configs) == 0 {
						ec2Configs = append(ec2Configs, returnInitialEC2Config(module))
					} else {
						privateKey, user := retrieveUserAndPrivateKey(module)

						tempConfig := EC2ConfigWrapper{
							ModuleName:   module.Path[1],
							KeyPairName:  resource.Primary.Attributes["key_name"].(string),
							InstanceType: resource.Primary.Attributes["instance_type"].(string),
							DefaultUser:  user,
							PrivateKey:   privateKey,
							RegionMap:    make(map[string]int),
						}
						tempConfig.RegionMap[region] = 1
						for index, config := range ec2Configs {
							if compareEC2Config(config, tempConfig) {
								if config.RegionMap[region] != 0 {
									config.RegionMap[region] = config.RegionMap[region] + 1
								} else {
									config.RegionMap[region] = 1
								}
							} else if index == len(ec2Configs)-1 {
								ec2Configs = append(ec2Configs, tempConfig)
							}
						}

					}

				}

			}
		}

	}
	return
}

//////////////
//DigitalOcean
/////////////
func returnInitialDOConfig(module ModuleState) (tempConfig DOConfigWrapper) {
	privateKey, user := retrieveUserAndPrivateKey(module)

	for _, resource := range module.Resources {
		if resource.Type == "digitalocean_droplet" {
			tempConfig.ModuleName = module.Path[1]
			tempConfig.Image = resource.Primary.Attributes["image"].(string)
			tempConfig.Fingerprint = resource.Primary.Attributes["ssh_keys.0"].(string)
			tempConfig.Size = resource.Primary.Attributes["size"].(string)
			tempConfig.RegionMap = make(map[string]int)
			tempConfig.PrivateKey = privateKey
			tempConfig.DefaultUser = user
			tempConfig.RegionMap[resource.Primary.Attributes["region"].(string)] = 1
			break
		}
	}
	return
}

func createDOConfigFromState(modules []ModuleState) (doConfigs []DOConfigWrapper, maxModuleCount int) {
	for _, module := range modules {
		if len(module.Path) > 2 && strings.Contains(module.Path[1], "doDropletDeploy") {
			for _, resource := range module.Resources {
				if resource.Type == "digitalocean_droplet" {
					countString := strings.Split(module.Path[1], "doDropletDeploy")[1]
					countInt, _ := strconv.Atoi(countString)
					if countInt > maxModuleCount {
						maxModuleCount = countInt
					}
					//If the list is empty, return the first element found
					if len(doConfigs) == 0 {
						doConfigs = append(doConfigs, returnInitialDOConfig(module))
					} else {
						privateKey, user := retrieveUserAndPrivateKey(module)
						tempConfig := DOConfigWrapper{
							ModuleName:  module.Path[1],
							Image:       resource.Primary.Attributes["image"].(string),
							Fingerprint: resource.Primary.Attributes["ssh_keys.0"].(string),
							Size:        resource.Primary.Attributes["size"].(string),
							DefaultUser: user,
							PrivateKey:  privateKey,
							RegionMap:   make(map[string]int),
						}
						tempConfig.RegionMap[resource.Primary.Attributes["region"].(string)] = 1
						for index, config := range doConfigs {
							if compareDOConfig(config, tempConfig) {
								if config.RegionMap[resource.Primary.Attributes["region"].(string)] != 0 {
									config.RegionMap[resource.Primary.Attributes["region"].(string)] = config.RegionMap[resource.Primary.Attributes["region"].(string)] + 1
								} else {
									config.RegionMap[resource.Primary.Attributes["region"].(string)] = 1
								}
							} else if index == len(doConfigs)-1 {
								doConfigs = append(doConfigs, tempConfig)
							}
						}

					}

				}

			}
		}

	}
	return
}

func CreateWrappersFromState(state State) (wrappers ConfigWrappers) {
	wrappers.DO, wrappers.DropletModuleCount = createDOConfigFromState(state.Modules)
	wrappers.EC2, wrappers.EC2ModuleCount = createEC2ConfigFromState(state.Modules)
	wrappers.AWSAPI, wrappers.AWSAPIModuleCount = createAWSAPIFromState(state.Modules)
	wrappers.Cloudfront, wrappers.CloudfrontModuleCount = createCloudfrontFromState(state.Modules)
	wrappers.Googlefront, wrappers.GooglefrontModuleCount = createGooglefrontFromState(state.Modules)
	return
}

//CreateMasterList takes a MasterList object as input
//and maps it to the corresponding templates, executes them,
//then adds the resulting string to a complete string
//containing the main.tf file for terraform
func CreateMasterFile(wrappers ConfigWrappers) (masterString string) {
	for _, config := range wrappers.EC2 {
		templ := template.Must(template.New("ec2").Funcs(template.FuncMap{"counter": templateCounter}).Parse(mainEc2Module))

		var templBuffer bytes.Buffer
		err := templ.Execute(&templBuffer, config)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, config := range wrappers.DO {
		templ := template.Must(template.New("droplet").Funcs(template.FuncMap{"counter": templateCounter}).Parse(mainDropletModule))

		var templBuffer bytes.Buffer
		err := templ.Execute(&templBuffer, config)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, config := range wrappers.AWSAPI {
		templ := template.Must(template.New("awsapi").Funcs(template.FuncMap{"counter": templateCounter}).Parse(mainAWSAPIModule))

		var templBuffer bytes.Buffer
		err := templ.Execute(&templBuffer, config)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, config := range wrappers.Cloudfront {
		templ := template.Must(template.New("cloudfront").Funcs(template.FuncMap{"counter": templateCounter}).Parse(mainCloudfrontModule))
		var templBuffer bytes.Buffer
		err := templ.Execute(&templBuffer, config)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	return masterString
}
