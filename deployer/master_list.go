package deployer

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

////THIS IS NEW EXPeriment

func returnInitialDOConfig(module ModuleState) (tempConfig DOConfigWrapper) {
	for name, resource := range module.Resources {
		if strings.Contains(name, "ansible_host") {
			tempConfig.PrivateKey = resource.Primary.Attributes["vars.ansible_ssh_private_key_file"]
			tempConfig.DefaultUser = resource.Primary.Attributes["vars.ansible_user"]
			break
		}
	}
	for name, resource := range module.Resources {
		if strings.Contains(name, "digitalocean_droplet") {
			tempConfig.ModuleName = module.Path[1]
			tempConfig.Image = resource.Primary.Attributes["image"]
			tempConfig.Fingerprint = resource.Primary.Attributes["ssh_keys.0"]
			tempConfig.Size = resource.Primary.Attributes["size"]
			tempConfig.RegionMap = make(map[string]int)
			tempConfig.RegionMap[resource.Primary.Attributes["region"]] = 0
			break
		}
	}
	return
}

func createDOConfigFromState(modules []ModuleState) (doConfigs []DOConfigWrapper, maxModuleCount int) {
	for _, module := range modules {
		if len(module.Path) > 2 && strings.Contains(module.Path[1], "doDropletDeploy") {
			for _, resource := range module.Resources {
				if strings.Contains(module.Path[1], "doDropletDeploy") {
					countString := strings.Split(module.Path[1], "doDropletDeploy")[1]
					countInt, _ := strconv.Atoi(countString)
					if countInt > maxModuleCount {
						maxModuleCount = countInt
					}
					//If the list is empty, return the first element found
					if len(doConfigs) == 0 {
						doConfigs = append(doConfigs, returnInitialDOConfig(module))
					} else {
						for _, config := range doConfigs {
							tempConfig := DOConfigWrapper{
								ModuleName:  module.Path[1],
								Image:       resource.Primary.Attributes["image"],
								Fingerprint: resource.Primary.Attributes["ssh_keys.0"],
								Size:        resource.Primary.Attributes["size"],
								DefaultUser: config.DefaultUser,
								PrivateKey:  config.PrivateKey,
							}
							if compareDOConfig(config, tempConfig) {
								if config.RegionMap[resource.Primary.Attributes["region"]] != 0 {
									config.RegionMap[resource.Primary.Attributes["region"]] = config.RegionMap[resource.Primary.Attributes["region"]] + 1
								} else {
									config.RegionMap[resource.Primary.Attributes["region"]] = 1
								}
							}
						}
					}

				}

			}
		}

	}
	return
}

//CreateMasterList takes a MasterList object as input
//and maps it to the corresponding templates, executes them,
//then adds the resulting string to a complete string
//containing the main.tf file for terraform
func CreateMasterFile(wrappers ConfigWrappers) (masterString string) {
	// var masterString string

	// awsApiGateways := terraformOutput.Master.ProviderValues.AWSProvider.API
	// cloudFronts := terraformOutput.Master.ProviderValues.AWSProvider.DomainFront

	// azureCDNs := terraformOutput.Master.ProviderValues.AzureProvider.DomainFront
	// azureInstances := terraformOutput.Master.ProviderValues.AzureProvider.Instances

	// DOInstances := terraformOutput.Master.ProviderValues.DOProvider.Instances

	// GoogleInstances := terraformOutput.Master.ProviderValues.GoogleProvider.Instances

	//EC2 Creation

	//EC2
	// for _, config := range ec2ConfigWrappers {
	// 	fmt.Println(config)
	// 	templ := template.Must(template.New("ec2").Funcs(template.FuncMap{"counter": templateCounter}).Parse(mainEc2Module))

	// 	var templBuffer bytes.Buffer
	// 	err := templ.Execute(&templBuffer, config)
	// 	masterString = masterString + templBuffer.String()
	// 	checkErr(err)
	// }

	//DigitalOcean Droplets
	for _, config := range wrappers.DO {
		fmt.Println(config)
		templ := template.Must(template.New("droplet").Funcs(template.FuncMap{"counter": templateCounter}).Parse(mainDropletModule))

		var templBuffer bytes.Buffer
		err := templ.Execute(&templBuffer, config)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	return masterString
}

///////Deprecated
// func createDOConfigWrapper(doInstances []DOInstance) (wrapperList []DOConfigWrapper) {
// 	var moduleCounter = 1

// 	for outer := range doInstances {
// 		count := doInstances[outer].Count

// 		if outer == 0 {
// 			doInstances[outer].ModuleName = "doDropletDeploy" + strconv.Itoa(moduleCounter)
// 			tempMap := make(map[string]int)
// 			tempMap[doInstances[outer].Region] = count
// 			configWrapper := DOConfigWrapper{Config: doInstances[outer], RegionMap: tempMap}

// 			wrapperList = append(wrapperList, configWrapper)
// 		} else {
// 			for i := range wrapperList {
// 				if compareDOConfig(doInstances[outer], wrapperList[i]) {
// 					if wrapperList[i].RegionMap[doInstances[outer].Region] == 0 {
// 						wrapperList[i].RegionMap[doInstances[outer].Region] = count
// 					}
// 				} else if i == len(wrapperList)-1 {
// 					tempMap := make(map[string]int)
// 					tempMap[doInstances[outer].Region] = count
// 					configWrapper := DOConfigWrapper{Config: doInstances[outer], RegionMap: tempMap}
// 					wrapperList = append(wrapperList, configWrapper)
// 				}
// 			}
// 		}
// 		for inner := range doInstances {

// 			if compareDOConfig(doInstances[outer], doInstances[inner]) {
// 				doInstances[inner].ModuleName = doInstances[outer].ModuleName
// 			} else if doInstances[inner].ModuleName == "" {
// 				doInstances[inner].ModuleName = "doDropletDeploy" + strconv.Itoa(moduleCounter)
// 				moduleCounter = moduleCounter + 1
// 			}
// 		}
// 	}
// 	return
// }

// func CreateMasterList(inputList ReadList) (masterString string, err error) {
// 	ec2List := inputList.ec2DeployerList
// 	azureCdnList := inputList.azureCdnDeployerList
// 	azureList := inputList.azureDeployerList
// 	cloudFrontList := inputList.cloudFrontDeployerList
// 	digitalOceanList := inputList.digitalOceanDeployerList
// 	googleCloudList := inputList.googleCloudDeployerList
// 	apiGatewayList := inputList.apiGatewayDeployerList

// 	for _, ec2Struct := range ec2List {
// 		templ, err := template.New("ec2").Parse(ec2Module)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, ec2Struct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}

// 	for _, azureCdnStruct := range azureCdnList {
// 		templ, err := template.New("azureCdn").Parse(azureCdnModule)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, azureCdnStruct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}

// 	for _, azureStruct := range azureList {
// 		templ, err := template.New("azureCdn").Parse(azureModule)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, azureStruct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}

// 	for _, cloudFrontStruct := range cloudFrontList {
// 		templ, err := template.New("cloudFront").Parse(cloudfrontModule)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, cloudFrontStruct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}

// 	for _, digitalOceanStruct := range digitalOceanList {
// 		templ, err := template.New("digitalOcean").Parse(digitalOceanModule)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, digitalOceanStruct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}
// 	for _, googleCloudStruct := range googleCloudList {
// 		templ, err := template.New("azureCdn").Parse(googleCloudModule)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, googleCloudStruct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}
// 	for _, apiGatewayStruct := range apiGatewayList {
// 		templ, err := template.New("azureCdn").Parse(apiGatewayModule)
// 		checkErr(err)

// 		var templBuffer bytes.Buffer
// 		err = templ.Execute(&templBuffer, apiGatewayStruct)
// 		masterString = masterString + templBuffer.String()
// 		checkErr(err)
// 	}

// 	return
// }

//This is necessary to group instances together by Module name
// func createEc2ConfigWrapper(ec2Instances []AWSInstance) (wrapperList []AWSConfigWrapper) {
// 	var moduleCounter = 1

// 	for outer := range ec2Instances {
// 		count := ec2Instances[outer].Count
// 		//On First iteration create the initial module name and configWrapper for our template
// 		if outer == 0 {
// 			//Module Name
// 			ec2Instances[outer].ModuleName = "ec2deploy" + strconv.Itoa(moduleCounter)
// 			moduleCounter = moduleCounter + 1

// 			//Initializing Template Wrapper
// 			tempMap := make(map[string]int)
// 			tempMap[ec2Instances[outer].Region] = count
// 			configWrapper := AWSConfigWrapper{Config: ec2Instances[outer], RegionMap: tempMap}

// 			wrapperList = append(wrapperList, configWrapper)
// 		} else {
// 			//For each configWrapper, check if the config matches, if it does then
// 			//add it to the RegionMap by either adding to the existing value or creating
// 			//a new map key. If it doesn't match we initialize a new wrapper and add it to list.
// 			for i := range wrapperList {
// 				if compareAWSConfig(ec2Instances[outer], wrapperList[i]) {
// 					if wrapperList[i].RegionMap[ec2Instances[outer].Region] < 1 {
// 						wrapperList[i].RegionMap[ec2Instances[outer].Region] = count
// 					}
// 				} else if i == len(wrapperList)-1 {
// 					tempMap := make(map[string]int)
// 					tempMap[ec2Instances[outer].Region] = count
// 					configWrapper := AWSConfigWrapper{Config: ec2Instances[outer], RegionMap: tempMap}
// 					wrapperList = append(wrapperList, configWrapper)
// 				}

// 			}
// 		}
// 		//For each instance, compare the configs. If they match then
// 		//add create the same module name as the outer instance in the
// 		//top for loop
// 		for inner := range ec2Instances {

// 			if compareAWSConfig(ec2Instances[outer], ec2Instances[inner]) {
// 				ec2Instances[inner].ModuleName = ec2Instances[outer].ModuleName
// 			} else if ec2Instances[inner].ModuleName == "" {
// 				ec2Instances[inner].ModuleName = "ec2deploy" + strconv.Itoa(moduleCounter)
// 				moduleCounter = moduleCounter + 1
// 			}
// 		}

// 	}

// 	return
// }
