package main

import (
	"bytes"
	"html/template"
)

var masterString string

func createMasterList(inputList readList) (err error) {
	ec2List := inputList.ec2DeployerList
	azureCdnList := inputList.azureCdnDeployerList
	azureList := inputList.azureDeployerList
	cloudFrontList := inputList.cloudFrontDeployerList
	digitalOceanList := inputList.digitalOceanDeployerList
	googleCloudList := inputList.googleCloudDeployerList
	apiGatewayList := inputList.apiGatewayDeployerList

	for _, ec2Struct := range ec2List {
		templ, err := template.New("ec2").Parse(ec2Module)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, ec2Struct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, azureCdnStruct := range azureCdnList {
		templ, err := template.New("azureCdn").Parse(azureCdnModule)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, azureCdnStruct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, azureStruct := range azureList {
		templ, err := template.New("azureCdn").Parse(azureModule)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, azureStruct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, cloudFrontStruct := range cloudFrontList {
		templ, err := template.New("cloudFront").Parse(cloudfrontModule)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, cloudFrontStruct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	for _, digitalOceanStruct := range digitalOceanList {
		templ, err := template.New("digitalOcean").Parse(digitalOceanModule)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, digitalOceanStruct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}
	for _, googleCloudStruct := range googleCloudList {
		templ, err := template.New("azureCdn").Parse(googleCloudModule)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, googleCloudStruct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}
	for _, apiGatewayStruct := range apiGatewayList {
		templ, err := template.New("azureCdn").Parse(apiGatewayModule)
		checkErr(err)

		var templBuffer bytes.Buffer
		err = templ.Execute(&templBuffer, apiGatewayStruct)
		masterString = masterString + templBuffer.String()
		checkErr(err)
	}

	return
}
