package deployer

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

func retrieveUserAndPrivateKey(module ModuleState) (privateKey string, user string) {
	for _, resource := range module.Resources {
		if resource.Type == "ansible_host" {
			privateKey = resource.Primary.Attributes["vars.ansible_ssh_private_key_file"]
			user = resource.Primary.Attributes["vars.ansible_user"]
			break
		}
	}
	return
}


////////////
//EC2///////
////////////
func returnInitialEC2Config(module ModuleState) (tempConfig EC2ConfigWrapper) {
	privateKey, user := retrieveUserAndPrivateKey(module)

	for name, resource := range module.Resources {
		if strings.Contains(name, "digitalocean_droplet") {
			tempConfig.ModuleName = module.Path[1]
			tempConfig. = resource.Primary.Attributes["image"]
			tempConfig.Fingerprint = resource.Primary.Attributes["ssh_keys.0"]
			tempConfig.Size = resource.Primary.Attributes["size"]
			tempConfig.RegionMap = make(map[string]int)
			tempConfig.PrivateKey = privateKey
			tempConfig.DefaultUser = user
			tempConfig.RegionMap[resource.Primary.Attributes["region"]] = 1
			break
		}
	}
	return
}

func createEC2ConfigFromState(modules []ModuleState) (doConfigs []EC2ConfigWrapper, maxModuleCount int) {
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
							Image:       resource.Primary.Attributes["image"],
							Fingerprint: resource.Primary.Attributes["ssh_keys.0"],
							Size:        resource.Primary.Attributes["size"],
							DefaultUser: user,
							PrivateKey:  privateKey,
							RegionMap:   make(map[string]int),
						}
						tempConfig.RegionMap[resource.Primary.Attributes["region"]] = 1
						for index, config := range doConfigs {
							if compareDOConfig(config, tempConfig) {
								if config.RegionMap[resource.Primary.Attributes["region"]] != 0 {
									config.RegionMap[resource.Primary.Attributes["region"]] = config.RegionMap[resource.Primary.Attributes["region"]] + 1
								} else {
									config.RegionMap[resource.Primary.Attributes["region"]] = 1
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

//////////////
//DigitalOcean
/////////////
func returnInitialDOConfig(module ModuleState) (tempConfig DOConfigWrapper) {
	privateKey, user := retrieveUserAndPrivateKey(module)

	for name, resource := range module.Resources {
		if strings.Contains(name, "digitalocean_droplet") {
			tempConfig.ModuleName = module.Path[1]
			tempConfig.Image = resource.Primary.Attributes["image"]
			tempConfig.Fingerprint = resource.Primary.Attributes["ssh_keys.0"]
			tempConfig.Size = resource.Primary.Attributes["size"]
			tempConfig.RegionMap = make(map[string]int)
			tempConfig.PrivateKey = privateKey
			tempConfig.DefaultUser = user
			tempConfig.RegionMap[resource.Primary.Attributes["region"]] = 1
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
							Image:       resource.Primary.Attributes["image"],
							Fingerprint: resource.Primary.Attributes["ssh_keys.0"],
							Size:        resource.Primary.Attributes["size"],
							DefaultUser: user,
							PrivateKey:  privateKey,
							RegionMap:   make(map[string]int),
						}
						tempConfig.RegionMap[resource.Primary.Attributes["region"]] = 1
						for index, config := range doConfigs {
							if compareDOConfig(config, tempConfig) {
								if config.RegionMap[resource.Primary.Attributes["region"]] != 0 {
									config.RegionMap[resource.Primary.Attributes["region"]] = config.RegionMap[resource.Primary.Attributes["region"]] + 1
								} else {
									config.RegionMap[resource.Primary.Attributes["region"]] = 1
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

//CreateMasterList takes a MasterList object as input
//and maps it to the corresponding templates, executes them,
//then adds the resulting string to a complete string
//containing the main.tf file for terraform
func CreateMasterFile(wrappers ConfigWrappers) (masterString string) {
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
