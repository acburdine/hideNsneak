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

var installArgs string
var burpCmd string
var installIndex int
var numberInput string
var fqdn string
var domain string
var burpDir string

var install = &cobra.Command{
	Use:   "install",
	Short: "Main install command",
	Long:  `Main install command, with subcommands for Burp, Cobalt Strike, GoPhish, LetsEncrypt, Nmap, Socat, SQLMap`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'install --help' for usage.")
	},
}

var burpInstall = &cobra.Command{
	Use:   "burp",
	Short: "Installs Burp Suite Collaborator Server",
	Long:  `Installs and starts a Burp Suite collaborator with the specified domain on the specified remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"burp"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var cobaltStrikeInstall = &cobra.Command{
	Use:   "cobaltstrike",
	Short: "Installs Cobalt Strike",
	Long:  `Installs,starts, and optionally licenses Cobaltstrike on the remote server with the specified malleable C2 profile and password`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"cobalstrike"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var goPhishInstall = &cobra.Command{
	Use:   "gophish",
	Short: "Installs Gophish",
	Long:  `Installs and starts Gophish on the remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"gophish"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var letsEncryptInstall = &cobra.Command{
	Use:   "letsencrypt",
	Short: "Installs Letsencrypt",
	Long:  `Installs Letsencrypt with the specified domain on the specified server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"letsencrypt"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var nmapInstall = &cobra.Command{
	Use:   "nmap",
	Short: "Installs Nmap",
	Long:  `Installs Nmap to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"nmap"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var socatInstall = &cobra.Command{
	Use:   "socat",
	Short: "Installs Socat",
	Long:  `Installs Socat to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"socat"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var sqlMapInstall = &cobra.Command{
	Use:   "sqlmap",
	Short: "Installs SQLmap",
	Long:  `Installs SQLmap to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"sqlmap"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instances := list[installIndex : installIndex+1]

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath, execCommand, nmapOutput, nmapCommands)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

func init() {
	rootCmd.AddCommand(install)
	install.AddCommand(burpInstall, cobaltStrikeInstall, goPhishInstall, letsEncryptInstall, nmapInstall, socatInstall, sqlMapInstall)

	burpInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	burpInstall.MarkPersistentFlagRequired("id")
	burpInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	burpInstall.MarkPersistentFlagRequired("fqdn")
	burpInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	burpInstall.MarkPersistentFlagRequired("domain")
	burpInstall.PersistentFlags().StringVarP(&burpDir, "burpDir", "b", "", "Specify the directory where burp is located")
	burpInstall.MarkPersistentFlagRequired("burpDir")

	cobaltStrikeInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	cobaltStrikeInstall.MarkFlagRequired("id")
	cobaltStrikeInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	cobaltStrikeInstall.MarkPersistentFlagRequired("fqdn")
	cobaltStrikeInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	cobaltStrikeInstall.MarkPersistentFlagRequired("domain")

	goPhishInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	goPhishInstall.MarkFlagRequired("id")
	goPhishInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	goPhishInstall.MarkPersistentFlagRequired("fqdn")
	goPhishInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	goPhishInstall.MarkPersistentFlagRequired("domain")

	letsEncryptInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	letsEncryptInstall.MarkFlagRequired("id")
	letsEncryptInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	letsEncryptInstall.MarkPersistentFlagRequired("fqdn")
	letsEncryptInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	letsEncryptInstall.MarkPersistentFlagRequired("domain")

	nmapInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	nmapInstall.MarkFlagRequired("id")

	socatInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	socatInstall.MarkFlagRequired("id")

	sqlMapInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	sqlMapInstall.MarkFlagRequired("id")
}
