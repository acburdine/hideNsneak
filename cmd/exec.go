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

var execCommand string

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
		playbook := deployer.GeneratePlaybookFile("exec")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, hostFilePath, remoteFilePath, execCommand)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var nmap = &cobra.Command{
	Use:   "nmap",
	Short: "execute nmap",
	Long:  `executes nmap and splits up the job between all of the specified hosts returning the xml files to the specified directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

func init() {
	rootCmd.AddCommand(exec)
	exec.AddCommand(command, nmap)

	command.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the remote server")
	command.MarkFlagRequired("id")
	command.PersistentFlags().StringVarP(&execCommand, "command", "c", "", "Specify the command you want to execute")
	command.MarkPersistentFlagRequired("command")

	// nmap.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the remote server")
	// nmap.MarkFlagRequired("id")
	// nmap.PersistentFlags().StringVarP(&fqdn, "command", "c", "", "Specify the command you want to execute")
	// nmap.MarkPersistentFlagRequired("command")
}
