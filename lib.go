package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal()
	}
}

func execCmd(binary string, args []string) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(binary, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(stdout.String())
}

func terraformApply() {
	binary, err := exec.LookPath("terraform")

	checkErr(err)

	//Initializing Terraform
	fmt.Println("init")
	args := []string{"init", "-input=false", "terraform"}
	execCmd(binary, args)

	//Planning Terraform changes and saving plan to file tfplan
	args = []string{"plan", "-out=terraform/tfplan", "-input=false", "-var-file=terraform/terraform.tfvars", "terraform"}
	execCmd(binary, args)

	//Applying Changes Identified in tfplan
	args = []string{"apply", "-input=false", "terraform/tfplan"}
	execCmd(binary, args)

}

func removeSpaces(input string) (newString string) {
	newString = strings.ToLower(input)
	newString = strings.Replace(newString, " ", "_", -1)

	return
}

func providerCheck(providerArray []string) bool {
	for _, p := range providerArray {
		if strings.ToUpper(p) != "EC2" &&
			strings.ToUpper(p) != "DO" &&
			strings.ToUpper(p) != "GOOGLE" &&
			strings.ToUpper(p) != "AZURE" &&
			strings.ToUpper(p) != "AZURECDN" &&
			strings.ToUpper(p) != "APIGATEWAY" {
			fmt.Println(unknownProvider)
			return false
		}
	}
	return true
}
