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

var ufwAction string
var ufwTCPPorts []string
var ufwUDPPorts []string
var ufwIndices []int

// helloCmd represents the hello command
var firewall = &cobra.Command{
	Use:   "firewall",
	Short: "firewall",
	Long:  `firewall`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'firewall --help' for usage.")
	},
}

var firewallAdd = &cobra.Command{
	Use:   "add",
	Short: "add a ufw firewall rule",
	Long:  `adds a ufw firewall rules to target host containing the tcp and udp port specifications set out by the user`,
	Args: func(cmd *cobra.Command, args []string) error {
		_, err := deployer.ValidatePorts(ufwTCPPorts)
		if err != nil {
			return err
		}
		_, err = deployer.ValidatePorts(ufwUDPPorts)
		if err != nil {
			return err
		}

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		largestInstanceNum := deployer.FindLargestNumber(ufwIndices)

		//make sure the largestInstanceNumToDestroy is not bigger than totalInstancesAvailable
		if len(list) < largestInstanceNum+1 {
			return errors.New("the number you entered is too big. Try running `list` to see the number of instances you have")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ufwTCPPorts, _ := deployer.ValidatePorts(ufwTCPPorts)

		ufwUDPPorts, _ := deployer.ValidatePorts(ufwUDPPorts)

		apps := []string{"firewall"}
		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []deployer.ListStruct

		for _, num := range ufwIndices {
			instances = append(instances, list[num])
		}

		ufwAction = "add"

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath,
			remoteFilePath, execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var firewallDelete = &cobra.Command{
	Use:   "delete",
	Short: "delete a ufw firewall rule",
	Long:  `adds a ufw firewall rules to target host containing the tcp and udp port specifications set out by the user`,
	Args: func(cmd *cobra.Command, args []string) error {
		_, err := deployer.ValidatePorts(ufwTCPPorts)
		if err != nil {
			return err
		}
		_, err = deployer.ValidatePorts(ufwUDPPorts)
		if err != nil {
			return err
		}

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		largestInstanceNum := deployer.FindLargestNumber(ufwIndices)

		//make sure the largestInstanceNumToDestroy is not bigger than totalInstancesAvailable
		if len(list) < largestInstanceNum+1 {
			return errors.New("the number you entered is too big. Try running `list` to see the number of instances you have")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ufwTCPPorts, _ = deployer.ValidatePorts(ufwTCPPorts)

		ufwUDPPorts, _ := deployer.ValidatePorts(ufwUDPPorts)

		apps := []string{"firewall"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []deployer.ListStruct

		for _, num := range ufwIndices {
			instances = append(instances, list[num])
		}

		ufwAction = "delete"

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath,
			remoteFilePath, execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var firewallList = &cobra.Command{
	Use:   "list",
	Short: "list ufw firewall rules",
	Long:  `lists all of the ufw firewall rules on the specifiec host`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		largestInstanceNum := deployer.FindLargestNumber(ufwIndices)

		//make sure the largestInstanceNumToDestroy is not bigger than totalInstancesAvailable
		if len(list) < largestInstanceNum+1 {
			return errors.New("the number you entered is too big. Try running `list` to see the number of instances you have")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"firewall"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []deployer.ListStruct

		for _, num := range ufwIndices {
			instances = append(instances, list[num])
		}

		ufwAction = "list"

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

func init() {
	rootCmd.AddCommand(firewall)
	firewall.AddCommand(firewallAdd, firewallDelete, firewallList)

	firewallAdd.PersistentFlags().IntSliceVarP(&ufwIndices, "id", "i", []int{}, "Specify the id for the remote server")
	firewallAdd.MarkFlagRequired("id")

	firewallAdd.PersistentFlags().StringSliceVarP(&ufwTCPPorts, "tcp", "t", []string{}, "Specify the tcp ports to add i.e. 22,23")
	firewallAdd.PersistentFlags().StringSliceVarP(&ufwUDPPorts, "udp", "u", []string{}, "Specify the udp ports to add i.e. 500,53")

	firewallDelete.PersistentFlags().IntSliceVarP(&ufwIndices, "id", "i", []int{}, "Specify the id for the remote server")
	firewallDelete.MarkFlagRequired("id")

	firewallDelete.PersistentFlags().StringSliceVarP(&ufwTCPPorts, "tcp", "t", []string{}, "Specify the tcp ports to delete i.e. 22,23")
	firewallDelete.PersistentFlags().StringSliceVarP(&ufwUDPPorts, "udp", "u", []string{}, "Specify the udp ports to delete i.e. 500,53")

	firewallList.PersistentFlags().IntSliceVarP(&ufwIndices, "id", "i", []int{}, "Specify the id for the remote server")
	firewallList.MarkFlagRequired("id")

}
