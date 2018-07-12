package deployer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
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

//IsValidNumberInput takes in a string and checks if the numbers are valid
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

//ExpandNumberInput expands input string and returns a list of ints
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

//WriteToFile opens, clears and writes to file
func WriteToFile(path string, content string) {
	fmt.Println(path)
	fmt.Println(content)
	file, err := os.Create(path)
	checkErr(err)

	_, err = file.Write([]byte(content))
	checkErr(err)
	defer file.Close()
}

//ValidateListOfInstances makes sure that the number input is actually available in our list of active instances
func ValidateListOfInstances(numberInput string) error {
	marshalledState := TerraformStateMarshaller()
	list := ListIPAddresses(marshalledState)
	if !IsValidNumberInput(numberInput) {
		return fmt.Errorf("invalid formatting specified: %s", numberInput)
	}
	numsToInstall := ExpandNumberInput(numberInput)
	largestInstanceNumToInstall := FindLargestNumber(numsToInstall)

	//make sure the largestInstanceNumToInstall is not bigger than totalInstancesAvailable
	if len(list) < largestInstanceNumToInstall {
		return errors.New("the number you entered is too big; try running `list` to see the number of instances you have")
	}
	return nil
}

/////////////////////
//Ansible Functions
/////////////////////

//GeneratePlaybookFile generates an ansible playbook
func GeneratePlaybookFile(app string) string {
	var playbookStruct ansiblePlaybook

	playbookStruct.GenerateDefault()

	playbookStruct.Roles = append(playbookStruct.Roles, app)

	playbookList := []ansiblePlaybook{playbookStruct}

	playbook, err := yaml.Marshal(playbookList)

	if err != nil {
		fmt.Println("Error marshalling playbook")
	}

	return string(playbook)
}

//GenerateHostsFile generates an ansible host file
func GenerateHostFile(instances []ListStruct, domain string, fqdn string, burpDir string) string {
	var inventory ansibleInventory

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	inventory.All.Hosts = make(map[string]ansibleHost)
	for _, instance := range instances {
		inventory.All.Hosts[instance.IP] = ansibleHost{
			AnsibleHost:       instance.IP,
			AnsiblePrivateKey: usr.HomeDir + "/.ssh/" + instance.PrivateKey,
			AnsibleUser:       instance.Username,
			AnsibleFQDN:       fqdn,
			AnsibleDomain:     domain,
			BurpDir:           burpDir,
		}
	}

	hostFile, err := yaml.Marshal(inventory)

	if err != nil {
		fmt.Println("problem marshalling inventory file")
	}

	return string(hostFile)
}

func ExecAnsible(hostsFile string, playbook string, filepath string) string {
	var stdout, stderr bytes.Buffer
	binary, err := exec.LookPath("ansible-playbook")

	checkErr(err)

	args := []string{"-i", hostsFile, playbook}
	cmd := exec.Command(binary, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = filepath

	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
	}

	return stdout.String()
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

	return stdout.String()
}

//InitializeTerraformFiles Creates the base templates for
//the terraform infrastructure
func InitializeTerraformFiles() {
	var config configStruct

	var configContents, _ = ioutil.ReadFile(configFile)

	json.Unmarshal(configContents, &config)

	secrets, err := template.New("secrets").Parse(templateSecrets)

	if err != nil {
		fmt.Println(err)
	}

	secretBuff := new(bytes.Buffer)

	err = secrets.Execute(secretBuff, &config)

	if err != nil {
		fmt.Println(err)
	}

	noEscapeSecrets := template.HTML(secretBuff.String())

	mainFile, err := os.Create(tfMainFile)
	checkErr(err)
	defer mainFile.Close()

	varFile, err := os.Create(tfVariablesFile)
	checkErr(err)
	defer varFile.Close()

	tfvarsFile, err := os.Create(tfVarsFile)
	checkErr(err)
	defer tfvarsFile.Close()

	fmt.Println("Creating Terraform Files...")

	mainFile.Write([]byte(backend))
	varFile.Write([]byte(variables))
	tfvarsFile.Write([]byte(noEscapeSecrets))
}

//TerraformApply runs the init, plan, and apply commands for our
//generated terraform templates
func TerraformApply() {

	//Initializing Terraform
	fmt.Println("Initializing Terraform...")
	args := []string{"init", "-backend-config=../config/backend.txt"}
	execTerraform(args, "terraform")

	//Applying Changes Identified in tfplan
	fmt.Println("Applying Terraform Changes...")
	args = []string{"apply", "-input=false", "-auto-approve"}
	execTerraform(args, "terraform")

}

func TerraformDestroy(nameList []string) {

	//Initializing Terraform
	args := []string{"init", "-backend-config=../config/backend.txt"}
	execTerraform(args, "terraform")

	args = []string{"destroy", "-auto-approve"}

	for _, name := range nameList {
		args = append(args, "-target", name)
	}
	fmt.Println("Destroying Terraform Targets...")

	execTerraform(args, "terraform")
}

//TerraforrmOutputMarshaller runs the terraform output command
//and marshalls the resulting JSON into a TerraformOutput struct
func TerraformStateMarshaller() (outputStruct State) {

	args := []string{"state", "pull"}
	output := execTerraform(args, "terraform")

	json.Unmarshal([]byte(output), &outputStruct)

	return
}

//CreateTerraformMain takes in a string containing all the necessary calls
//for the main.tf file
func CreateTerraformMain(masterString string) {

	InitializeTerraformFiles()

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

func PrintProxyChains(socksList string) (proxies string) {
	socksArray := strings.Split(socksList, "\n")
	for _, command := range socksArray {
		args := strings.Split(command, " ")
		for index, arg := range args {
			var port string
			if arg == "-D" {
				port = args[index+1]
				proxies = proxies + fmt.Sprintf("socks5 127.0.0.1 %s\n", port)
			}
		}
	}
	return
}

func PrintSocksd(socksList string) (proxies string) {
	socksArray := strings.Split(socksList, "\n")
	proxies = "{\"proxies\":[\n\t\"upstreams\":["
	for index, command := range socksArray {
		var port string
		var ip string
		args := strings.Split(command, " ")
		for index, arg := range args {

			if arg == "-D" {
				port = args[index+1]
			}
			if strings.Contains(arg, "@") {
				ip = strings.Split(arg, "@")[1]
			}
		}
		upstream := fmt.Sprintf("\t\t{\"type\": \"socks5\", \"address\": \"127.0.0.1:%s\", \"target\": \"%s\"}", port, ip)
		if index != len(socksArray)-1 {
			upstream = upstream + ","
		}
		proxies = proxies + "\n" + upstream

	}
	proxies = proxies + "\n\t\t]\n\t}\n\t]\n}"
	return
}

func DestroySOCKS(ip string) {
	args := []string{"-f", "ssh.*-D [0-9]{4,}.*@" + ip}
	cmd := exec.Command("pkill", args...)

	cmd.Run()
}

//createSingleSOCKS initiates a SOCKS Proxy on the local host with the specifed ipv4 address
func CreateSingleSOCKS(privateKey string, username string, ipv4 string, port int) (err error) {
	portString := strconv.Itoa(port)
	args := []string{"-D", portString, "-o", "StrictHostKeyChecking=no", "-N", "-f", "-i", privateKey, username + "@" + ipv4}
	cmd := exec.Command("ssh", args...)
	err = cmd.Start()
	if err != nil {
		return
	}
	return

}

func (listStruct *ListStruct) String() string {
	return ("IP: " + listStruct.IP + " - Provider: " + listStruct.Provider + " - Region: " + listStruct.Region + " - Name: " + listStruct.Name)
}

//listStructSOrt takes a list of listStructs and goes ahead and
//sorts them based on position. This is important because it will
//ensure that the order of the elements remain the same on each list call
func listStructSort(listStructs []ListStruct) (finalList []ListStruct) {
	for index := range listStructs {
		for _, list := range listStructs {
			if list.Place == index {
				finalList = append(finalList, list)
				break
			}
		}
	}
	return
}

func ListProxies(instances []ListStruct) (output string) {
	for _, instance := range instances {
		ip := instance.IP
		args := []string{"-f", "ssh.*-D.*" + ip}
		cmd := exec.Command("pgrep", args...)
		out, err := cmd.Output()

		if len(out) > 0 {
			pid := strings.TrimSpace(string(out))

			args = []string{"-o", "command", pid}
			cmd = exec.Command("ps", args...)
			out, err = cmd.Output()
			if err != nil {
				fmt.Println(err)
			}

			commandOutput := strings.Split(string(out), "COMMAND")[1]

			output = output + "PID: " + pid + " Command: " + strings.TrimSpace(commandOutput) + "\n"
		}

	}
	return
}

func ListDomainFronts(state State) (domainFronts []DomainFrontOutput) {
	for _, module := range state.Modules {
		var domainFrontOutput DomainFrontOutput
		if len(module.Path) > 1 {
			for name, resource := range module.Resources {
				if strings.Contains(module.Path[1], "cloudfrontDeploy") {
					domainFrontOutput.Provider = "AWS"
					domainFrontOutput.ID = resource.Primary.Attributes["id"]
					domainFrontOutput.Etag = resource.Primary.Attributes["etag"]
					domainFrontOutput.Status = resource.Primary.Attributes["status"]
					domainFrontOutput.Name = "module." + strings.Join(module.Path[1:], ".module.") + "." + name
					for key, value := range resource.Primary.Attributes {
						if strings.Contains(key, "domain_name") {
							if strings.Contains(key, "origin") {
								domainFrontOutput.Origin = value
							} else {
								domainFrontOutput.Invoke = value
							}
						}
					}
					domainFronts = append(domainFronts, domainFrontOutput)
				} else if strings.Contains(module.Path[1], "azurefrontDeploy") {
					domainFrontOutput.Provider = "AZURE"
					// domainFronts = append(domainFronts, domainFrontOutput)
				}
			}

		}
	}
	return
}

func ListAPIs(state State) (apiOutputs []APIOutput) {
	for _, module := range state.Modules {
		var apiOutput APIOutput
		if len(module.Path) > 1 && strings.Contains(module.Path[1], "apiDeploy") {
			apiOutput.Provider = "AWS"
			for name, resource := range module.Resources {
				switch resource.Type {
				case "aws_api_gateway_deployment":
					apiOutput.InvokeURI = resource.Primary.Attributes["invoke_url"]
				case "aws_api_gateway_integration":
					apiOutput.TargetURI = resource.Primary.Attributes["uri"]
				case "aws_api_gateway_rest_api":
					apiOutput.Name = "module." + strings.Join(module.Path[1:], ".module.") + "." + name
				default:
					continue
				}
			}
		}
		apiOutputs = append(apiOutputs, apiOutput)
	}
	return
}

func ListIPAddresses(state State) (hostOutput []ListStruct) {
	for _, module := range state.Modules {
		var tempOutput []ListStruct
		if len(module.Path) > 1 {
			privatekey, username := retrieveUserAndPrivateKey(module)
			for name, resource := range module.Resources {
				fullName := "module." + strings.Join(module.Path[1:], ".module.") + "." + name
				nameSlice := strings.Split(name, ".")
				finalString := nameSlice[len(nameSlice)-1]
				count, err := strconv.Atoi(finalString)
				if err == nil {

					index := "[" + finalString + "]"

					newName := strings.Join(nameSlice[:len(nameSlice)-1], ".")

					fullName = "module." + strings.Join(module.Path[1:], ".module.") + "." + newName + index
				}
				switch resource.Type {
				case "digitalocean_droplet":
					tempOutput = append(tempOutput, ListStruct{
						IP:         resource.Primary.Attributes["ipv4_address"],
						Provider:   "DigitalOcean",
						Region:     resource.Primary.Attributes["region"],
						Name:       fullName,
						Place:      count,
						PrivateKey: privatekey,
						Username:   username,
					})
				case "aws_instance":
					tempOutput = append(tempOutput, ListStruct{
						IP:         resource.Primary.Attributes["public_ip"],
						Provider:   "AWS",
						Region:     resource.Primary.Attributes["availability_zone"],
						Name:       fullName,
						Place:      count,
						PrivateKey: privatekey,
						Username:   username,
					})
				default:
					continue
				}
			}
			hostOutput = append(hostOutput, listStructSort(tempOutput)...)
		}
	}
	return
}

//InstanceDeploy takes input from the user interface in order to divide and deploy appropriate regions
//it takes in a TerraformOutput struct, makes the appropriate edits, and returns that same struct
func InstanceDeploy(providers []string, awsRegions []string, doRegions []string, azureRegions []string,
	googleRegions []string, count int, privKey string, pubKey string, keyName string, wrappers ConfigWrappers) ConfigWrappers {

	doModuleCount := wrappers.DropletModuleCount
	awsModuleCount := wrappers.EC2ModuleCount

	//Strip Directories from key name
	//Identical Keypairs must be named the same
	privKey = filepath.Base(privKey)
	pubKey = filepath.Base(pubKey)

	//Gather the count per provider and the remainder
	countPerProvider := count / len(providers)

	remainderForProviders := count % len(providers)

	for _, provider := range providers {
		switch strings.ToUpper(provider) {
		case "AWS":
			countPerAWSregion := countPerProvider / len(awsRegions)

			remainderForAWSRegion := countPerProvider % len(awsRegions)

			if remainderForProviders > 0 {
				remainderForAWSRegion = remainderForAWSRegion + 1
				remainderForProviders = remainderForProviders - 1
			}

			for _, region := range awsRegions {

				regionCount := countPerAWSregion

				if remainderForAWSRegion > 0 {
					regionCount = regionCount + 1
					remainderForAWSRegion = remainderForAWSRegion - 1
				}
				//TODO: Add custom input
				if regionCount > 0 {
					//TODO: Ensure private key is the same
					//Check this functionality between two clients
					// result := checkEC2KeyExistence(config.AwsSecretKey, config.AwsAccessID, region, keyName)

					// if !result {
					// 	publicKeyBytes, _ := ioutil.ReadFile(pubKey)

					// 	err := importEC2Key(config.AwsSecretKey, config.AwsAccessID, region, publicKeyBytes, pubKey)
					// 	if err != nil {
					// 		fmt.Printf("There was an errror importing your key to EC2: %s", err)
					// 	}
					// }

					newEC2RegionConfig := EC2ConfigWrapper{
						InstanceType: "t2.micro",
						PrivateKey:   privKey,
						PublicKey:    pubKey,
						KeyPairName:  keyName,
						DefaultUser:  "ubuntu",
						RegionMap:    make(map[string]int),
					}
					newEC2RegionConfig.RegionMap[region] = regionCount

					if len(wrappers.EC2) == 0 {
						awsModuleCount = 1
						newEC2RegionConfig.ModuleName = "ec2Deploy" + strconv.Itoa(awsModuleCount)
						awsModuleCount = awsModuleCount + 1
						wrappers.EC2 = append(wrappers.EC2, newEC2RegionConfig)
						continue
					}
					for index, config := range wrappers.EC2 {
						if compareEC2Config(config, newEC2RegionConfig) {
							if config.RegionMap[region] > 0 {
								config.RegionMap[region] = config.RegionMap[region] + regionCount
							} else {
								config.RegionMap[region] = regionCount
							}
							break
						} else if index == len(wrappers.DO)-1 {
							awsModuleCount = awsModuleCount + 1
							newEC2RegionConfig.ModuleName = "ec2Deploy" + strconv.Itoa(awsModuleCount)
							wrappers.EC2 = append(wrappers.EC2, newEC2RegionConfig)
							awsModuleCount = awsModuleCount + 1
						}
					}
				}
			}
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
	return wrappers
}

//APIDeploy takes argruments to deploy an API Gateway
func APIDeploy(provider string, targetURI string, wrappers ConfigWrappers) ConfigWrappers {
	moduleCount := wrappers.AWSAPIModuleCount

	if strings.ToUpper(provider) == "AWS" {
		if len(wrappers.AWSAPI) > 0 {
			for _, wrapper := range wrappers.AWSAPI {
				if targetURI == wrapper.TargetURI {
					continue
				}
				wrappers.AWSAPI = append(wrappers.AWSAPI, AWSApiConfigWrapper{
					ModuleName: "awsAPIDeploy" + strconv.Itoa(moduleCount+1),
					TargetURI:  targetURI,
				})
				moduleCount = moduleCount + 1
			}
		} else {
			wrappers.AWSAPI = append(wrappers.AWSAPI, AWSApiConfigWrapper{
				ModuleName: "awsAPIDeploy" + strconv.Itoa(moduleCount+1),
				TargetURI:  targetURI,
			})
		}
	} else if strings.ToUpper(provider) == "ALIBABA" {
	}

	return wrappers
}

func DomainFrontDeploy(provider string, origin string, wrappers ConfigWrappers) ConfigWrappers {
	moduleCount := wrappers.CloudfrontModuleCount

	if strings.ToUpper(provider) == "AWS" {
		if len(wrappers.Cloudfront) > 0 {
			for _, wrapper := range wrappers.Cloudfront {
				if origin == wrapper.Origin {
					continue
				}
				wrappers.Cloudfront = append(wrappers.Cloudfront, CloudfrontConfigWrapper{
					ModuleName: "cloudfrontDeploy" + strconv.Itoa(moduleCount+1),
					Origin:     origin,
					Enabled:    "true",
				})
			}
		} else {
			wrappers.Cloudfront = append(wrappers.Cloudfront, CloudfrontConfigWrapper{
				ModuleName: "cloudfrontDeploy" + strconv.Itoa(moduleCount+1),
				Origin:     origin,
				Enabled:    "true",
			})
		}
	}

	return wrappers
}

//AWSCloufFrontDestroy uses the deleteCloudFront function to delete
//the specified cloudfront due to the problems with terraforms destruction process
func AWSCloudFrontDestroy(output DomainFrontOutput) error {
	//TODO catch the error here
	err := deleteCloudFront(output.ID, output.Etag, config.AwsAccessID, config.AwsSecretKey)
	if err != nil {
		return err
	}

	args := []string{"state", "rm", output.Name}
	execTerraform(args, "terraform")
	return nil
}
