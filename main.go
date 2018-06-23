package main

import (
	"os"
)

func main() {
	//Creating Files
	mainFile, err := os.Create("main.tf")
	checkErr(err)
	defer mainFile.Close()

	varFile, err := os.Create("variables.tf")
	checkErr(err)
	defer varFile.Close()

	tfvarsFile, err := os.Create("terraform.tfvars")
	checkErr(err)
	defer tfvarsFile.Close()

	//Writing Constants
	mainFile.Write([]byte(state))
	varFile.Write([]byte(variables))
	tfvarsFile.Write([]byte(tfvars))

	//TODO: creater master list (string) by createMasterList(input the master struct here)

	//Opening Main.tf to append parsed template
	file, err := os.OpenFile("main.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)

	//Writing the masterString to file. masterString was instantiated in master_list.go
	_, err = file.WriteString(masterString)
	checkErr(err)

	err = mainFile.Close()
	checkErr(err)

	return

	//Perform all the terraform deployment
	terraformApply()
}
