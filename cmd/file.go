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

var hostFilePath string
var remoteFilePath string
var fqdn string
var domain string
var burpDir string

// helloCmd represents the hello command
var file = &cobra.Command{
	Use:   "file",
	Short: "file",
	Long:  `file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'file --help' for usage.")
	},
}

var filePush = &cobra.Command{
	Use:   "push",
	Short: "send a file",
	Long:  `send a file from your local host to a remote server via absolute filepath`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateHostFilePathExists(hostFilePath)
		deployer.ValidateRemoteFilePathExists(remoteFilePath)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		playbook := deployer.GeneratePlaybookFile("sync-push")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := GenerateHostFile(instances, fqdn, domain, burpDir, hostFilePath, remoteFilePath)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
	},
}

var filePull = &cobra.Command{
	Use:   "pull",
	Short: "get a file",
	Long:  `get a file from your remote server to your local host via absolute filepath`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateHostFilePathExists(hostFilePath)
		deployer.ValidateRemoteFilePathExists(remoteFilePath)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		playbook := deployer.GeneratePlaybookFile("sync-pull")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := GenerateHostFile(instances, fqdn, domain, burpDir, hostFilePath, remoteFilePath)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
	},
}

func init() {
	rootCmd.AddCommand(file)
	file.AddCommand(filePush, filePull)

	filePush.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the remote server")
	filePush.MarkFlagRequired("id")
	filePush.PersistentFlags().StringVarP(&fqdn, "host", "h", "", "Specify the host file's absolute path")
	filePush.MarkPersistentFlagRequired("host")
	filePush.PersistentFlags().StringVarP(&domain, "remote", "r", "", "Specify the remote file's absolute path")
	filePush.MarkPersistentFlagRequired("remote")

	filePull.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the remote server")
	filePull.MarkFlagRequired("id")
	filePull.PersistentFlags().StringVarP(&hostFilePath, "host", "h", "", "Specify the host file's absolute path")
	filePull.MarkPersistentFlagRequired("host")
	filePull.PersistentFlags().StringVarP(&remoteFilePath, "remote", "r", "", "Specify the remote file's absolute path")
	filePull.MarkPersistentFlagRequired("remote")
}
