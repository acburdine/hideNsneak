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
	"fmt"
	"hideNsneak/deployer"

	"github.com/spf13/cobra"
)

var ufwAction string
var ufwTCPPort string
var ufwUDPPort string
var ufwIndex int

// helloCmd represents the hello command
var firewall = &cobra.Command{
	Use:   "file",
	Short: "file",
	Long:  `file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'file --help' for usage.")
	},
}

var firewallAdd = &cobra.Command{
	Use:   "push",
	Short: "send a file",
	Long:  `send a file from your local host to a remote server via absolute filepath`,
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"sync-push"}
		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath,
			remoteFilePath, execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			ufwAction, ufwTCPPort, ufwUDPPort)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var firewallDelete = &cobra.Command{
	Use:   "pull",
	Short: "get a file",
	Long:  `get a file from your remote server to your local host via absolute filepath`,
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"sync-pull"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath,
			remoteFilePath, execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			ufwAction, ufwTCPPort, ufwUDPPort)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var firewallList = &cobra.Command{
	Use:   "pull",
	Short: "get a file",
	Long:  `get a file from your remote server to your local host via absolute filepath`,
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"sync-pull"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			ufwAction, ufwTCPPort, ufwUDPPort)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

func init() {
	rootCmd.AddCommand(firewall)
	firewall.AddCommand(firewallAdd, firewallDelete, firewallList)

	firewallAdd.PersistentFlags().IntVarP(&ufwIndex, "id", "i", 0, "Specify the id for the remote server")
	firewallAdd.MarkFlagRequired("id")

	firewallAdd.PersistentFlags().StringVarP(&ufwTCPPort, "tcp", "t", "", "Specify the id for the remote server")
	firewallAdd.PersistentFlags().StringVarP(&ufwTCPPort, "udp", "u", "", "Specify the id for the remote server")

	firewallDelete.PersistentFlags().IntVarP(&ufwIndex, "id", "i", 0, "Specify the id for the remote server")
	firewallDelete.MarkFlagRequired("id")

	firewallDelete.PersistentFlags().StringVarP(&ufwUDPPort, "tcp", "t", "", "Specify the id for the remote server")
	firewallDelete.PersistentFlags().StringVarP(&ufwUDPPort, "udp", "u", "", "Specify the id for the remote server")

}
