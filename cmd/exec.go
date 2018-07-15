// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"hideNsneak/deployer"

	"github.com/spf13/cobra"
)

var execCommand string
var nmapPorts string
var nmapHostFile string
var nmapCommand string
var nmapOutput string
var nmapIndex string

var exec = &cobra.Command{
	Use:   "exec",
	Short: "execute custom command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'exec --help' for usage.")
	},
}

var command = &cobra.Command{
	Use:   "command",
	Short: "execute custom command",
	Long:  `executes the specified command on the specified remote system and returns both stdout and stderr`,
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"exec"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var nmap = &cobra.Command{
	Use:   "nmap",
	Short: "execute nmap",
	Long:  `executes nmap and splits up the job between all of the specified hosts returning the xml files to the specified directory`,
	Args: func(cmd *cobra.Command, args []string) error {
		_, err := deployer.ValidatePorts(nmapPorts)
		if err != nil {
			return err
		}
		_, err = deployer.ParseIPFile(nmapHostFile)
		if err != nil {
			return err
		}

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)
		if !deployer.IsValidNumberInput(nmapIndex) {
			return fmt.Errorf("invalid formatting specified: %s", nmapIndex)
		}

		largestInstanceNumToDestroy := deployer.FindLargestNumber(numsToDeploy)

		//make sure the largestInstanceNumToDestroy is not bigger than totalInstancesAvailable
		if len(list) < largestInstanceNumToDestroy {
			return errors.New("the number you entered is too big. Try running `list` to see the number of instances you have")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"nmap", "nmap-exec"}

		playbook := deployer.GeneratePlaybookFile(apps)

		numsToDeploy := deployer.ExpandNumberInput(nmapIndex)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []ListStruct

		for _, num := range numsToDeploy {
			instances = append(instances, list[num])
		}

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

func init() {
	rootCmd.AddCommand(exec)
	exec.AddCommand(command, nmap)

	command.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the remote server")
	command.MarkFlagRequired("id")
	command.PersistentFlags().StringVarP(&execCommand, "command", "c", "", "Specify the command you want to execute")
	command.MarkPersistentFlagRequired("command")

	nmap.PersistentFlags().IntVarP(&nmapIndex, "id", "i", 0, "Specify the ids for the scanning servers")
	nmap.MarkPersistentFlagRequired("id")
	nmap.PersistentFlags().StringVarP(&nmapHostFile, "hostFile", "f", "", "Specify filepath of the file containing the scope/hosts")
	nmap.MarkPersistentFlagRequired("hostFile")
	nmap.PersistentFlags().StringVarP(&nmapPorts, "ports", "p", "", "Specify the port range to be passed to nmap i.e 21-23,443-445,8080-8081,8443")
	nmap.MarkPersistentFlagRequired("ports")
	nmap.PersistentFlags().StringVarP(&nmapCommand, "nmapCommand", "n", "", "Specify the full nmap command to be run, excluding the -iL,-p, and -oA options options i.e. nmap -sV -sT --max-rate=250")
	nmap.MarkPersistentFlagRequired("nmapCommand")
	nmap.PersistentFlags().StringVarP(&nmapOutput, "nmapOutput", "o", "", "Specify the local directory for output to be saved")
	nmap.MarkPersistentFlagRequired("nmapOutput")

}
