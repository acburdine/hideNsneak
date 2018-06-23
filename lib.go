package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
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
	args := []string{"init", "-input=false"}
	execCmd(binary, args)

	//Planning Terraform changes and saving plan to file tfplan
	args = []string{"plan", "-out=tfplan", "-input=false"}
	execCmd(binary, args)

	//Applying Changes Identified in tfplan
	args = []string{"apply", "-input=false", "tfplan"}
	execCmd(binary, args)

}
