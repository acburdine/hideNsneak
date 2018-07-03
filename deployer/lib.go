package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

////////////////////////
//Miscellaneous Functions
////////////////////////

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}
}
func templateCounter() func() int {
	i := -1
	return func() int {
		i++
		return i
	}
}

func removeSpaces(input string) (newString string) {
	newString = strings.ToLower(input)
	newString = strings.Replace(newString, " ", "_", -1)

	return
}

//ContainsString checks to see if the array contains the target string
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//ContainsInt checks to see if the array contains the target int
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func execCmd(binary string, args []string, filepath string) string {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(binary, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = filepath

	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return stdout.String()
}

func IsValidNumberInput(input string) bool {
	sliceToParse := strings.Split(input, ",")

	for _, num := range sliceToParse {
		_, err := strconv.Atoi(num)
		if err != nil {
			dashSlice := strings.Split(num, "-")
			if len(dashSlice) != 2 {
				return false
			} else {
				_, err := strconv.Atoi(dashSlice[0])
				if err != nil {
					return false
				}
				_, err = strconv.Atoi(dashSlice[1])
				if err != nil {
					return false
				}
			}
			continue
		}
	}
	return true
}

func ExpandNumberInput(input string) []int {
	var result []int
	sliceToParse := strings.Split(input, ",")

	for _, num := range sliceToParse {
		getInt, err := strconv.Atoi(num)
		if err != nil {
			sliceToSplit := strings.Split(num, "-")
			firstNum, err := strconv.Atoi(sliceToSplit[0])
			if err != nil {
				continue
			}
			secondNum, err := strconv.Atoi(sliceToSplit[1])
			if err != nil {
				continue
			}
			for i := firstNum; i <= secondNum; i++ {
				result = append(result, i)
			}
		} else {
			result = append(result, getInt)
		}
	}
	return result
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

/////////////////////
//Terraform Functions
/////////////////////

func execTerraform(args []string, filepath string) string {
	var stdout, stderr bytes.Buffer

	binary, err := exec.LookPath("terraform")

	checkErr(err)

	cmd := exec.Command(binary, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = filepath

	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
	}

	fmt.Println(stderr.String())

	return stdout.String()
}

//InitializeTerraformFiles Creates the base templates for
//the terraform infrastructure
func InitializeTerraformFiles() {
	mainFile, err := os.Create(tfMainFile)
	checkErr(err)
	defer mainFile.Close()

	varFile, err := os.Create(tfVariablesFile)
	checkErr(err)
	defer varFile.Close()

	tfvarsFile, err := os.Create(tfVarsFile)
	checkErr(err)
	defer tfvarsFile.Close()

	mainFile.Write([]byte(state))
	varFile.Write([]byte(variables))
	tfvarsFile.Write([]byte(templateSecrets))
}

//TerraformApply runs the init, plan, and apply commands for our
//generated terraform templates
func TerraformApply() {

	//Initializing Terraform
	args := []string{"init"}
	execTerraform(args, "terraform")

	//Applying Changes Identified in tfplan
	args = []string{"apply", "-input=false", "-auto-approve"}
	execTerraform(args, "terraform")

}

//Types for our destruction magic
//TODO: Rename and put in correct place

type ConcurrentSlice struct {
	sync.RWMutex
	items []interface{}
}

// Concurrent slice item
type ConcurrentSliceItem struct {
	Index int
	Value interface{}
}

func (cs *ConcurrentSlice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

func (cs *ConcurrentSlice) Iter() <-chan ConcurrentSliceItem {
	c := make(chan ConcurrentSliceItem)

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

func NewConcurrentSlice() *ConcurrentSlice {
	cs := &ConcurrentSlice{
		items: make([]interface{}, 0),
	}

	return cs
}

func terraformRetrieveNames(IDList []string) (nameList []string) {
	var wg sync.WaitGroup

	concurrentSlice := NewConcurrentSlice()
	for _, id := range IDList {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			args := []string{"state", "list", "-id=" + id}
			name := strings.TrimSpace(execTerraform(args, "terraform"))
			concurrentSlice.Append(name)
		}(id)
	}
	wg.Wait()
	for i := range concurrentSlice.Iter() {
		nameList = append(nameList, i.Value.(string))
	}

	return nameList
}

func TerraformDestroy(nameList []string) {

	//Initializing Terraform
	args := []string{"init", "-input=false", "terraform"}
	execTerraform(args, "terraform")

	args = []string{"destroy", "-auto-approve"}

	for _, name := range nameList {
		args = append(args, "-target", name)
	}
	fmt.Println(args)

	execTerraform(args, "terraform")
}

//TerraforrmOutputMarshaller runs the terraform output command
//and marshalls the resulting JSON into a TerraformOutput struct
func TerraformStateMarshaller() (outputStruct State) {

	//Initializing Terraform
	args := []string{"state", "pull"}
	output := execTerraform(args, "terraform")

	json.Unmarshal([]byte(output), &outputStruct)

	return
}

//CreateTerraformMain takes in a string containing all the necessary calls
//for the main.tf file
func CreateTerraformMain(masterString string) {
	//Opening Main.tf to append parsed template
	mainFile, err := os.OpenFile("terraform/main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	//Writing the masterString to file. masterString was instantiated in master_list.go
	_, err = mainFile.WriteString(masterString)
	checkErr(err)

	err = mainFile.Close()
	checkErr(err)
}

//ProviderCheck takes in a user-defined array of
//providers and validates they are supported
func ProviderCheck(providerArray []string) bool {
	for _, p := range providerArray {
		if strings.ToUpper(p) != "AWS" &&
			strings.ToUpper(p) != "DO" &&
			strings.ToUpper(p) != "GOOGLE" &&
			strings.ToUpper(p) != "AZURE" {
			return false
		}
	}
	return true
}

func mergeMap(map1 map[string]string, map2 map[string]string) map[string]string {
	for key, value := range map2 {
		map1[key] = value
	}
	return map1
}

func FindLargestNumber(nums []int) int {
	var n, largest int
	for _, i := range nums {
		if i > n {
			n = i
			largest = n
		}
	}
	return largest
}

// func GenerateIPIDList() IPID {

// 	var masterIPID IPID
// 	marshalledOutput := TerraformStateMarshaller()

// 	for _, instance := range marshalledOutput.Master.ProviderValues.AWSProvider.Instances {
// 		masterIPID.IDList = append(masterIPID.IDList, instance.IPID.IDList...)
// 		masterIPID.IPList = append(masterIPID.IPList, instance.IPID.IPList...)
// 	}
// 	for _, instance := range marshalledOutput.Master.ProviderValues.DOProvider.Instances {
// 		masterIPID.IDList = append(masterIPID.IDList, instance.IPID.IDList...)
// 		masterIPID.IPList = append(masterIPID.IPList, instance.IPID.IPList...)
// 	}
// 	for _, instance := range marshalledOutput.Master.ProviderValues.GoogleProvider.Instances {
// 		masterIPID.IDList = append(masterIPID.IDList, instance.IPID.IDList...)
// 		masterIPID.IPList = append(masterIPID.IPList, instance.IPID.IPList...)
// 	}
// 	for _, instance := range marshalledOutput.Master.ProviderValues.AzureProvider.Instances {
// 		masterIPID.IDList = append(masterIPID.IDList, instance.IPID.IDList...)
// 		masterIPID.IPList = append(masterIPID.IPList, instance.IPID.IPList...)
// 	}

// 	return masterIPID
// }

//createSingleSOCKS initiates a SOCKS Proxy on the local host with the specifed ipv4 address
//returns a pointer to the specified OS process so that we can kill it effictively
func createSingleSOCKS(privateKey string, username string, ipv4 string, port int) *os.Process {
	portString := strconv.Itoa(port)
	cmd := exec.Command("ssh", "-N", "-D", portString, "-o", "StrictHostKeyChecking=no", "-i", privateKey, fmt.Sprintf(username+"@%s", ipv4))
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return nil
	}
	return cmd.Process
}

func (listStruct *ListStruct) String() string {
	return ("IP: " + listStruct.IP + " - Provider: " + listStruct.Provider + " - Region: " + listStruct.Region + " - Name: " + listStruct.Name)
}

func ListIPAddresses(state State) (hostOutput []ListStruct) {
	for _, module := range state.Modules {
		for name, resource := range module.Resources {
			fullName := "module." + strings.Join(module.Path[1:], ".module.") + "." + name
			nameSlice := strings.Split(name, ".")
			finalString := nameSlice[len(nameSlice)-1]
			_, err := strconv.Atoi(finalString)
			if err == nil {
				fullName = "module." + strings.Join(module.Path[1:], ".module.") + ".[" + finalString + "]"
			}
			switch {
			case strings.Contains(name, "digitalocean_droplet"):
				hostOutput = append(hostOutput, ListStruct{
					IP:       resource.Primary.Attributes["ipv4_address"],
					Provider: "DigitalOcean",
					Region:   resource.Primary.Attributes["region"],
					Name:     fullName,
				})
			default:
				continue
			}
		}
	}
	return
}

//InstanceDeploy takes input from the user interface in order to divide and deploy appropriate regions
//it takes in a TerraformOutput struct, makes the appropriate edits, and returns that same struct
func InstanceDeploy(providers []string, awsRegions []string, doRegions []string, azureRegions []string,
	googleRegions []string, count int, privKey string, pubKey string, state State) (wrappers ConfigWrappers) {

	var doModuleCount int

	//Gather the count per provider and the remainder
	countPerProvider := count / len(providers)

	remainderForProviders := count % len(providers)

	wrappers.DO, doModuleCount = createDOConfigFromState(state.Modules)

	for _, provider := range providers {
		switch strings.ToUpper(provider) {
		case "AWS":
		case "DO":
			countPerDOregion := countPerProvider / len(doRegions)

			remainderForDORegion := countPerProvider % len(doRegions)

			if remainderForProviders > 0 {
				remainderForDORegion = remainderForDORegion + 1
				remainderForProviders = remainderForProviders - 1
			}

			for _, region := range doRegions {
				regionCount := countPerDOregion

				if remainderForDORegion > 0 {
					regionCount = regionCount + 1
					remainderForDORegion = remainderForDORegion - 1
				}
				//TODO: Add custom input
				if regionCount > 0 {
					newDORegionConfig := DOConfigWrapper{
						Image:       "ubuntu-16-04-x64",
						PrivateKey:  privKey,
						Fingerprint: genDOKeyFingerprint(pubKey),
						Size:        "512mb",
						DefaultUser: "root",
						RegionMap:   make(map[string]int),
					}
					newDORegionConfig.RegionMap[region] = regionCount

					if len(wrappers.DO) == 0 {
						doModuleCount = 1
						newDORegionConfig.ModuleName = "doDropletDeploy" + strconv.Itoa(doModuleCount)
						doModuleCount = doModuleCount + 1
						wrappers.DO = append(wrappers.DO, newDORegionConfig)
						continue
					}
					for index, config := range wrappers.DO {
						if compareDOConfig(config, newDORegionConfig) {
							if config.RegionMap[region] > 0 {
								config.RegionMap[region] = config.RegionMap[region] + regionCount
							} else {
								config.RegionMap[region] = regionCount
							}
							break
						} else if index == len(wrappers.DO)-1 {
							doModuleCount = doModuleCount + 1
							newDORegionConfig.ModuleName = "doDropletDeploy" + strconv.Itoa(doModuleCount)
							wrappers.DO = append(wrappers.DO, newDORegionConfig)
							doModuleCount = doModuleCount + 1
						}
					}
				}
			}
		default:
			continue
		}
	}
	return
}

///Deprecated Deploy
// for _, provider := range providers {
// 	switch strings.ToUpper(provider) {
// 	case "AWS":
// 		//Existing AWS Instances
// 		awsInstances := output.Master.ProviderValues.AWSProvider.Instances

// 		countPerAWSRegion := countPerProvider / len(awsRegions)

// 		remainderForAWSRegion := countPerProvider % len(awsRegions)

// 		//This if statement checks if the remainder for providers is 0
// 		//if it isnt, then we add 1 to the remainder for the region
// 		//It will result in 1 additional instance being added to the
// 		//next region in the list
// 		if remainderForProviders > 0 {
// 			remainderForAWSRegion = remainderForAWSRegion + 1
// 			remainderForProviders = remainderForProviders - 1
// 		}

// 		//Looping through the provided regions
// 		for _, region := range awsRegions {
// 			regionCount := countPerAWSRegion

// 			//TODO: Implement this, commented out due to broken functionality
// 			// keyCheckResult, keyName := checkEC2KeyExistance(awsSecretKey, awsAccessKey, region, privKey)
// 			// if !keyCheckResult {
// 			// 	keyName = "hideNsneak"
// 			// }

// 			if remainderForAWSRegion > 0 {
// 				regionCount = regionCount + 1
// 				remainderForAWSRegion = remainderForAWSRegion - 1
// 			}

// 			if regionCount > 0 {
// 				newRegionConfig = AWSRegionConfig{
// 					//TODO: Figure the security group thing out
// 					Count:          regionCount,
// 					CustomAmi:      "",
// 					InstanceType:   "t2.micro",
// 					DefaultUser:    "ubuntu",
// 					Region:         region,
// 					PublicKeyFile:  pubKey,
// 					PrivateKeyFile: privKey,
// 				}

// 				if len(awsInstances) == 0 {
// 					awsInstances = append(awsInstances, AWSInstance{
// 						Config: newRegionConfig,
// 						IPID: IPID{
// 							IPList: []string{},
// 							IDList: []string{},
// 						}})
// 					continue
// 				}

// 				for index := range awsInstances {
// 					if compareAWSConfig(awsInstances[index].Config, newRegionConfig) &&
// 						awsInstances[index].Config.Region == newRegionConfig.Region {

// 						awsInstances[index].Config.Count = awsInstances[index].Config.Count + newRegionConfig.Count

// 					} else if index == len(awsInstances)-1 {
// 						awsInstances = append(awsInstances, AWSInstance{
// 							Config: newRegionConfig,
// 							IPID: IPID{
// 								IPList: []string{},
// 								IDList: []string{},
// 							}})
// 					}

// 				}
// 			}

// 		}
// 		output.Master.ProviderValues.AWSProvider.Instances = awsInstances
// 	case "DO":
// 		doInstances := output.Master.ProviderValues.DOProvider.Instances

// 		countPerDOregion := countPerProvider / len(doRegions)

// 		remainderForDORegion := countPerProvider % len(awsRegions)

// 		if remainderForProviders > 0 {
// 			remainderForDORegion = remainderForDORegion + 1
// 			remainderForProviders = remainderForProviders - 1
// 		}

// 		for _, region := range doRegions {
// 			regionCount := countPerDOregion

// 			if remainderForDORegion > 0 {
// 				regionCount = regionCount + 1
// 				remainderForDORegion = remainderForDORegion - 1
// 			}

// 			if regionCount > 0 {
// 				newDORegionConfig := DORegionConfig{
// 					Image:       "ubuntu-16-04-x64",
// 					Count:       regionCount,
// 					PrivateKey:  privKey,
// 					Fingerprint: genDOKeyFingerprint(pubKey),
// 					Size:        "512MB",
// 					Region:      region,
// 					DefaultUser: "root",
// 				}

// 				if len(doInstances) == 0 {
// 					doInstances = append(doInstances, DOInstance{
// 						Config: newDORegionConfig,
// 						IPID: IPID{
// 							IPList: []string{},
// 							IDList: []string{},
// 						}})
// 					continue
// 				}

// 				for index := range doInstances {
// 					if compareDOConfig(doInstances[index].Config, newDORegionConfig) &&
// 						doInstances[index].Config.Region == newDORegionConfig.Region {
// 						doInstances[index].Config.Count = doInstances[index].Config.Count + newDORegionConfig.Count
// 					} else if index == len(doInstances)-1 {
// 						doInstances = append(doInstances, DOInstance{
// 							Config: newDORegionConfig,
// 							IPID: IPID{
// 								IPList: []string{},
// 								IDList: []string{},
// 							}})
// 					}
// 				}
// 			}
// 		}
// 		fmt.Println(doInstances)
// 		output.Master.ProviderValues.DOProvider.Instances = doInstances

// 		// var doDeployerList []digitalOceanDeployer

// 		// countPerDORegion := countPerProvider / len(doRegions)
// 		// remainderForDORegion := countPerProvider % len(doRegions)
// 		// if remainderForProviders != 0 {
// 		// 	remainderForDORegion = remainderForDORegion + 1
// 		// 	remainderForProviders = remainderForProviders - 1
// 		// }
// 		// for _, region := range doRegions {
// 		// 	regionCount := countPerDORegion
// 		// 	if remainderForDORegion > 0 {
// 		// 		regionCount = regionCount + 1
// 		// 		remainderForDORegion = remainderForDORegion - 1
// 		// 	}

// 		// 	if regionCount > 0 {
// 		// 		newDODeployer := digitalOceanDeployer{
// 		// 			Image:       "",
// 		// 			Fingerprint: genDOKeyFingerprint(pubKey),
// 		// 			PrivateKey:  privKey,
// 		// 			PublicKey:   pubKey,
// 		// 			Size:        "",
// 		// 			Count:       regionCount,
// 		// 			Region:      region,
// 		// 			DefaultUser: "",
// 		// 			Name:        "tester",
// 		// 		}
// 		// 		doDeployerList = append(doDeployerList, newDODeployer)
// 		// 	}

// 		// }
// 		// masterList.digitalOceanDeployerList = doDeployerList

// 	case "AZURE":
// 		// var azureDeployerList []azureDeployer
// 		// countPerAzureRegion := countPerProvider / len(azureRegions)
// 		// remainderForAzureRegion := countPerProvider % len(azureRegions)
// 		// if remainderForProviders != 0 {
// 		// 	remainderForAzureRegion = remainderForAzureRegion + 1
// 		// 	remainderForProviders = remainderForProviders - 1
// 		// }

// 		// for _, region := range awsRegions {
// 		// 	regionCount := countPerAzureRegion
// 		// 	//TODO check for existing keyname

// 		// 	if remainderForAzureRegion > 0 {
// 		// 		regionCount = regionCount + 1
// 		// 		remainderForAzureRegion = remainderForAzureRegion - 1
// 		// 	}

// 		// 	if regionCount > 0 {
// 		// 		newAzureDeployer := azureDeployer{
// 		// 			Location:    region,
// 		// 			Count:       regionCount,
// 		// 			VMSize:      "",
// 		// 			Environment: "",
// 		// 			PublicKey:   pubKey,
// 		// 			PrivateKey:  privKey,
// 		// 		}
// 		// 		azureDeployerList = append(azureDeployerList, newAzureDeployer)
// 		// 	}

// 		// }
// 		// masterList.azureDeployerList = azureDeployerList

// 	case "GOOGLE":

// 	// var googleDeployerList []googleCloudDeployer

// 	// countPerGoogleRegion := countPerProvider / len(googleRegions)
// 	// remainderForGoogleRegion := countPerProvider % len(googleRegions)
// 	// if remainderForProviders != 0 {
// 	// 	remainderForGoogleRegion = remainderForGoogleRegion + 1
// 	// 	remainderForProviders = remainderForProviders - 1
// 	// }

// 	// for _, region := range googleRegions {

// 	// 	regionCount := countPerGoogleRegion
// 	// 	if remainderForGoogleRegion > 0 {
// 	// 		regionCount = regionCount + 1
// 	// 		remainderForGoogleRegion = remainderForGoogleRegion - 1

// 	// 	}

// 	// 	if regionCount > 0 {
// 	// 		newGoogleDeployer := googleCloudDeployer{
// 	// 			Region:            region,
// 	// 			Project:           "inboxa90",
// 	// 			Count:             regionCount,
// 	// 			SSHUser:           "tester",
// 	// 			SSHPubKeyFile:     pubKey,
// 	// 			SSHPrivateKeyFile: privKey,
// 	// 			MachineType:       "",
// 	// 			Image:             "",
// 	// 		}
// 	// 		googleDeployerList = append(googleDeployerList, newGoogleDeployer)
// 	// 	}

// 	// }
// 	// masterList.googleCloudDeployerList = googleDeployerList
