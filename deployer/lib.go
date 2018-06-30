package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"golang.org/x/crypto/ssh"
)

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

//TerraformApply runs the init, plan, and apply commands for our
//generated terraform templates
func TerraformApply() {
	binary, err := exec.LookPath("terraform")

	checkErr(err)

	//Initializing Terraform
	fmt.Println("init")
	args := []string{"init", "-input=false", "terraform"}
	execCmd(binary, args, "terraform")

	//Planning Terraform changes and saving plan to file tfplan
	args = []string{"plan", "-out=tfplan", "-input=false", "-var-file=terraform.tfvars"}
	execCmd(binary, args, "terraform")

	//Applying Changes Identified in tfplan
	args = []string{"apply", "-input=false", "tfplan"}
	execCmd(binary, args, "terraform")

}

//TerraforrmOutputMarshaller runs the terraform output command
//and marshalls the resulting JSON into a TerraformOutput struct
func TerraformOutputMarshaller() (outputStruct TerraformOutput) {

	binary, err := exec.LookPath("terraform")

	checkErr(err)

	//Initializing Terraform
	args := []string{"output", "--json"}

	var stdout, stderr bytes.Buffer

	cmd := exec.Command(binary, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = "terraform"

	err = cmd.Run()
	if err != nil {
		return
	}

	bytes := stdout.Bytes()

	json.Unmarshal(bytes, &outputStruct)

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

func removeSpaces(input string) (newString string) {
	newString = strings.ToLower(input)
	newString = strings.Replace(newString, " ", "_", -1)

	return
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

func GenerateIPIDList() map[string]string {
	list := make(map[string]string)
	awsList := make(map[string]string)
	doList := make(map[string]string)
	googleList := make(map[string]string)
	azureList := make(map[string]string)
	marshalledOutput := TerraformOutputMarshaller()

	for i := range marshalledOutput.Master.ProviderValues.AWSProvider.Instances {
		awsList = mergeMap(awsList, marshalledOutput.Master.ProviderValues.AWSProvider.Instances[i].IPIDMap)
	}
	for i := range marshalledOutput.Master.ProviderValues.DOProvider.Instances {
		doList = mergeMap(doList, marshalledOutput.Master.ProviderValues.DOProvider.Instances[i].IPIDMap)
	}
	for i := range marshalledOutput.Master.ProviderValues.GoogleProvider.Instances {
		googleList = mergeMap(googleList, marshalledOutput.Master.ProviderValues.GoogleProvider.Instances[i].IPIDMap)
	}
	for i := range marshalledOutput.Master.ProviderValues.AzureProvider.Instances {
		azureList = mergeMap(azureList, marshalledOutput.Master.ProviderValues.AzureProvider.Instances[i].IPIDMap)
	}

	list = mergeMap(list, awsList)
	list = mergeMap(list, doList)
	list = mergeMap(list, googleList)
	list = mergeMap(list, azureList)

	return list
}

//checkEc2KeyExistence queries the Amazon EC2 API for the keypairs with the specified keyname
//Returns true if the resulting array is > 0, false otherwise
func checkEC2KeyExistance(secret string, accessID string, region string, privateKey string) (bool, string) {
	keyFingerprint := genEC2KeyFingerprint(privateKey)

	svc := ec2.New(session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, secret, ""),
	}))
	keyPairOutput, _ := svc.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("fingerprint"),
				Values: aws.StringSlice([]string{keyFingerprint}),
			},
		},
	})
	if len(keyPairOutput.KeyPairs) == 0 {
		return false, ""
	}
	return true, *keyPairOutput.KeyPairs[0].KeyName
}

func genEC2KeyFingerprint(privateKey string) (keyFingerprint string) {
	args1 := []string{"pkey", "-in", privateKey, "-pubout", "-outform", "DER"}
	args2 := []string{"md5", "-c"}

	pipeReader, pipeWriter := io.Pipe()

	cmd1 := exec.Command("openssl", args1...)
	cmd2 := exec.Command("openssl", args2...)

	cmd1.Stdout = pipeWriter
	cmd2.Stdin = pipeReader

	var cmd2Out bytes.Buffer

	cmd2.Stdout = &cmd2Out

	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	pipeWriter.Close()
	cmd2.Wait()

	keyFingerprint = strings.Split(strings.TrimSpace(cmd2Out.String()), " ")[1]

	return
}

func genDOKeyFingerprint(publicKey string) (keyFingerprint string) {
	key, err := ioutil.ReadFile(publicKey)

	if err != nil {
		fmt.Println("Unable to read")
	}
	pubKey, _, _, _, err := ssh.ParseAuthorizedKey(key)

	if err != nil {
		fmt.Println(err)
	}

	return ssh.FingerprintLegacyMD5(pubKey)
}

//checkEc2KeyExistence queries the Amazon EC2 API for the security groups
//with the specified name
//Returns true if the resulting array is > 0, false otherwise
func checkEC2SecurityGroupExistence(secret string, accessID string, region string, securityGroupName string) (bool, string) {
	svc := ec2.New(session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessID, secret, ""),
	}))
	securityGroupOutput, _ := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		GroupNames: aws.StringSlice([]string{securityGroupName}),
	})

	if len(securityGroupOutput.SecurityGroups) == 0 {
		return false, ""
	}

	return true, *securityGroupOutput.SecurityGroups[0].GroupId
}

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

func compareAWSConfig(initialRegion AWSRegionConfig, testRegion AWSRegionConfig) bool {
	if initialRegion.CustomAmi == testRegion.CustomAmi &&
		initialRegion.DefaultUser == testRegion.DefaultUser &&
		initialRegion.InstanceType == testRegion.InstanceType &&
		initialRegion.PrivateKeyFile == testRegion.PrivateKeyFile &&
		initialRegion.PublicKeyFile == testRegion.PublicKeyFile {
		return true
	}
	return false
}

func compareDOConfig(initialRegion DORegionConfig, testRegion DORegionConfig) bool {
	if initialRegion.Image == testRegion.Image &&
		initialRegion.Fingerprint == testRegion.Fingerprint &&
		initialRegion.PrivateKey == testRegion.PrivateKey &&
		initialRegion.Size == testRegion.Size &&
		initialRegion.Count == testRegion.Count &&
		initialRegion.DefaultUser == initialRegion.DefaultUser {
		return true
	}
	return false
}

//InstanceDeploy takes input from the user interface in order to divide and deploy appropriate regions
//it takes in a TerraformOutput struct, makes the appropriate edits, and returns that same struct
func InstanceDeploy(providers []string, awsRegions []string, doRegions []string, azureRegions []string,
	googleRegions []string, count int, privKey string, pubKey string, output TerraformOutput) TerraformOutput {

	var newRegionConfig AWSRegionConfig

	//Gather the count per provider and the remainder
	countPerProvider := count / len(providers)

	remainderForProviders := count % len(providers)

	for _, provider := range providers {
		switch strings.ToUpper(provider) {
		case "AWS":
			//Existing AWS Instances
			awsInstances := output.Master.ProviderValues.AWSProvider.Instances

			countPerAWSRegion := countPerProvider / len(awsRegions)

			remainderForAWSRegion := countPerProvider % len(awsRegions)

			//This if statement checks if the remainder for providers is 0
			//if it isnt, then we add 1 to the remainder for the region
			//It will result in 1 additional instance being added to the
			//next region in the list
			if remainderForProviders > 0 {
				remainderForAWSRegion = remainderForAWSRegion + 1
				remainderForProviders = remainderForProviders - 1
			}

			//Looping through the provided regions
			for _, region := range awsRegions {
				regionCount := countPerAWSRegion

				//TODO: Implement this, commented out due to broken functionality
				// keyCheckResult, keyName := checkEC2KeyExistance(awsSecretKey, awsAccessKey, region, privKey)
				// if !keyCheckResult {
				// 	keyName = "hideNsneak"
				// }

				if remainderForAWSRegion > 0 {
					regionCount = regionCount + 1
					remainderForAWSRegion = remainderForAWSRegion - 1
				}

				if regionCount > 0 {
					newRegionConfig = AWSRegionConfig{
						//TODO: Figure the security group thing out
						Count:          strconv.Itoa(regionCount),
						CustomAmi:      "",
						InstanceType:   "t2.micro",
						DefaultUser:    "ubuntu",
						Region:         region,
						PublicKeyFile:  pubKey,
						PrivateKeyFile: privKey,
					}

					if len(awsInstances) == 0 {
						awsInstances = append(awsInstances, AWSInstance{
							Config:  newRegionConfig,
							IPIDMap: make(map[string]string)})
						continue
					}

					for index := range awsInstances {
						if compareAWSConfig(awsInstances[index].Config, newRegionConfig) &&
							awsInstances[index].Config.Region == newRegionConfig.Region {

							//String conversion madness
							count1, _ := strconv.Atoi(awsInstances[index].Config.Count)
							count2, _ := strconv.Atoi(newRegionConfig.Count)

							awsInstances[index].Config.Count = strconv.Itoa(count1 + count2)
							break

						} else {
							awsInstances = append(awsInstances, AWSInstance{
								Config:  newRegionConfig,
								IPIDMap: make(map[string]string)})
						}

					}
				}

			}
			fmt.Println(awsInstances)
			output.Master.ProviderValues.AWSProvider.Instances = awsInstances
		case "DO":
			doInstances := output.Master.ProviderValues.DOProvider.Instances

			countPerDOregion := countPerProvider / len(doRegions)

			remainderForDORegion := countPerProvider % len(awsRegions)

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

				if regionCount > 0 {
					newDORegionConfig := DORegionConfig{
						Image:       "ubuntu-16-04-x64",
						Count:       regionCount,
						PrivateKey:  "/Users/mike.hodges/.ssh/do_rsa",
						Fingerprint: "b3:b2:c7:b1:73:9e:28:c6:61:8d:15:e1:0e:61:7e:35",
						Size:        "512MB",
						Region:      region,
						DefaultUser: "root",
					}

					if len(doInstances) == 0 {
						doInstances = append(doInstances, DOInstance{
							Config:  newDORegionConfig,
							IPIDMap: make(map[string]string),
						})
						continue
					}

					for index := range doInstances {
						if compareDOConfig(doInstances[index].Config, newDORegionConfig) &&
							doInstances[index].Config.Region == newDORegionConfig.Region {
							doInstances[index].Config.Count = doInstances[index].Config.Count + newDORegionConfig.Count
							break

						} else {
							doInstances = append(doInstances, DOInstance{
								Config:  newDORegionConfig,
								IPIDMap: make(map[string]string),
							})
						}
					}
				}
			}
			output.Master.ProviderValues.DOProvider.Instances = doInstances

			// var doDeployerList []digitalOceanDeployer

			// countPerDORegion := countPerProvider / len(doRegions)
			// remainderForDORegion := countPerProvider % len(doRegions)
			// if remainderForProviders != 0 {
			// 	remainderForDORegion = remainderForDORegion + 1
			// 	remainderForProviders = remainderForProviders - 1
			// }
			// for _, region := range doRegions {
			// 	regionCount := countPerDORegion
			// 	if remainderForDORegion > 0 {
			// 		regionCount = regionCount + 1
			// 		remainderForDORegion = remainderForDORegion - 1
			// 	}

			// 	if regionCount > 0 {
			// 		newDODeployer := digitalOceanDeployer{
			// 			Image:       "",
			// 			Fingerprint: genDOKeyFingerprint(pubKey),
			// 			PrivateKey:  privKey,
			// 			PublicKey:   pubKey,
			// 			Size:        "",
			// 			Count:       regionCount,
			// 			Region:      region,
			// 			DefaultUser: "",
			// 			Name:        "tester",
			// 		}
			// 		doDeployerList = append(doDeployerList, newDODeployer)
			// 	}

			// }
			// masterList.digitalOceanDeployerList = doDeployerList

		case "AZURE":
			// var azureDeployerList []azureDeployer
			// countPerAzureRegion := countPerProvider / len(azureRegions)
			// remainderForAzureRegion := countPerProvider % len(azureRegions)
			// if remainderForProviders != 0 {
			// 	remainderForAzureRegion = remainderForAzureRegion + 1
			// 	remainderForProviders = remainderForProviders - 1
			// }

			// for _, region := range awsRegions {
			// 	regionCount := countPerAzureRegion
			// 	//TODO check for existing keyname

			// 	if remainderForAzureRegion > 0 {
			// 		regionCount = regionCount + 1
			// 		remainderForAzureRegion = remainderForAzureRegion - 1
			// 	}

			// 	if regionCount > 0 {
			// 		newAzureDeployer := azureDeployer{
			// 			Location:    region,
			// 			Count:       regionCount,
			// 			VMSize:      "",
			// 			Environment: "",
			// 			PublicKey:   pubKey,
			// 			PrivateKey:  privKey,
			// 		}
			// 		azureDeployerList = append(azureDeployerList, newAzureDeployer)
			// 	}

			// }
			// masterList.azureDeployerList = azureDeployerList

		case "GOOGLE":

		// var googleDeployerList []googleCloudDeployer

		// countPerGoogleRegion := countPerProvider / len(googleRegions)
		// remainderForGoogleRegion := countPerProvider % len(googleRegions)
		// if remainderForProviders != 0 {
		// 	remainderForGoogleRegion = remainderForGoogleRegion + 1
		// 	remainderForProviders = remainderForProviders - 1
		// }

		// for _, region := range googleRegions {

		// 	regionCount := countPerGoogleRegion
		// 	if remainderForGoogleRegion > 0 {
		// 		regionCount = regionCount + 1
		// 		remainderForGoogleRegion = remainderForGoogleRegion - 1

		// 	}

		// 	if regionCount > 0 {
		// 		newGoogleDeployer := googleCloudDeployer{
		// 			Region:            region,
		// 			Project:           "inboxa90",
		// 			Count:             regionCount,
		// 			SSHUser:           "tester",
		// 			SSHPubKeyFile:     pubKey,
		// 			SSHPrivateKeyFile: privKey,
		// 			MachineType:       "",
		// 			Image:             "",
		// 		}
		// 		googleDeployerList = append(googleDeployerList, newGoogleDeployer)
		// 	}

		// }
		// masterList.googleCloudDeployerList = googleDeployerList

		default:
			continue
		}
	}
	return output
}
