package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func oldmain() {
	// TODO: make sure terraform and other important directories exist
	fmt.Print(ascii)
	fmt.Println(welcomeMessage)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		command, _ := reader.ReadString('\n')
		switch strings.TrimSpace(command) {
		case "help":
			fmt.Println(help)
		case "quit":
			fmt.Println(shutdown)
			os.Exit(1)
		case "exit":
			fmt.Println(shutdown)
			os.Exit(1)
		case "deploy":
			stillInLoop1 := true
			stillInLoop2 := true
			continueDeploy := true
			reader := bufio.NewReader(os.Stdin)
			var providerArray []string
			var providers string
			var count int
			var err error

			ec2Stuff := []ec2Deployer{}
			azureCdnStuff := []azureCdnDeployer{}
			azureStuff := []azureDeployer{}
			cloudFrontStuff := []cloudFrontDeployer{}
			doStuff := []digitalOceanDeployer{}
			gcpStuff := []googleCloudDeployer{}
			apiStuff := []apiGatewayDeployer{}
			userInput := readList{
				ec2Stuff,
				azureCdnStuff,
				azureStuff,
				cloudFrontStuff,
				doStuff,
				gcpStuff,
				apiStuff,
			}

			// if there is API Gateway in provider array, set up AG

			// if there is Google in provider array, set up GCP

			// if there is Azure in provider array, set up Azure
			// if there is AzureCDN in provider array, set up AzureCDN
			// if there is Digital Ocean in provider array, set up DO

			// for (stillInLoop2 == true) && (continueDeploy == true) {
			// 	fmt.Print(numServersToDeploy)
			// 	countString, _ := reader.ReadString('\n')
			// 	countString = strings.TrimSpace(countString)
			// 	count, err = strconv.Atoi(countString)
			// 	if err != nil {
			// 		fmt.Println("<hideNSneak/deploy> Error: Not an Integer.  ")
			// 		continue
			// 	}
			// 	break
			// }
			// providerMap := make(map[string]int)
			// division := count / len(providerArray)
			// remainder := count % len(providerArray)

			// for _, p := range providerArray {
			// 	providerMap[p] = division
			// }

			// if remainder != 0 {
			// 	for p := range providerMap {
			// 		providerMap[p] = providerMap[p] + 1
			// 		remainder = remainder - 1
			// 		if remainder == 0 {
			// 			break
			// 		}
			// 	}
			// }

			// instanceArray := cloud.DeployInstances(config, providerMap)
			// allInstances = append(allInstances, instanceArray...)
			// cloud.Initialize(allInstances, config)

			//TODO make user input
			//create an emoty struct from readList
			//assign each user input to values from struct
			//ass in that struct as userInput variable in createMasterList

			// parseUserInputIntoReadList()

			// tester1 := ec2Deployer{
			// 	Count:         1,
			// 	Region:        "us-east-1",
			// 	SecurityGroup: "tester1243",
			// 	PublicKeyFile: "/Users/mike.hodges/.ssh/do_rsa.pub",
			// 	KeypairName:   "do_rsa",
			// 	NewKeypair:    false,
			// }
			// tester2 := ec2Deployer{
			// 	Count:         1,
			// 	Region:        "us-west-1",
			// 	SecurityGroup: "tester1243",
			// 	PublicKeyFile: "/Users/mike.hodges/.ssh/do_rsa.pub",
			// 	KeypairName:   "do_rsa",
			// 	NewKeypair:    false,
			// }
			// tester3 := ec2Deployer{
			// 	Count:         1,
			// 	Region:        "eu-west-1",
			// 	SecurityGroup: "tester1243",
			// 	PublicKeyFile: "/Users/mike.hodges/.ssh/do_rsa.pub",
			// 	KeypairName:   "do_rsa",
			// 	NewKeypair:    false,
			// }

			createMasterList(userInput) //TODO: userInput is whatever masterStruct they want to pass in

			//Opening Main.tf to append parsed template
			file, err := os.OpenFile("terraform/main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			checkErr(err)

			//Writing the masterString to file. masterString was instantiated in master_list.go
			_, err = file.WriteString(masterString)
			checkErr(err)

			err = mainFile.Close()
			checkErr(err)

			//Perform all the terraform deployment
			terraformApply()

		case "destroy":
		case "start":
		case "stop":
		case "list":
		case "shell":
		case "socks-add":
		case "socks-kill":
		case "domainfront":
		case "domainfront-list":
		case "nmap":
		case "proxyconf":
		case "send":
		case "get":
		case "firewall":
		case "firewall-list":
		default:
			fmt.Println(doesntExist)
		}

	}

}
