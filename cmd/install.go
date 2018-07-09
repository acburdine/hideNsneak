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
	"terraform-playground/deployer"

	"github.com/spf13/cobra"
)

var installArgs string
var burpCmd string
var installIndex int
var numberInput string
var fqdn string
var domain string

var install = &cobra.Command{
	Use:   "install",
	Short: "Main install command",
	Long:  `Main install command, with subcommands for Burp, Cobalt Strike, GoPhish, LetsEncrypt, Nmap, Socat, SQLMap`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Install called")
	},
}

var burpInstall = &cobra.Command{
	Use:   "burp",
	Short: "Installs Burp Suite",
	Long:  `Installs Burp Suite to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		playbook := deployer.GeneratePlaybookFile("burp")
		//TODO: right now we can only do one host, but need to do multiples
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		//run burp installation here

		//1. open up hostFile and edit
		//		get host IP address
		//		get burp_dir: /Users/mischy/Downloads/
		//		get burp_server_domain: swansonmedical.com
		//		get burp_local_address: 127.0.0.1
		//		get burp_public_address: 35.171.8.53
		//		close and save
		//2. open up main.yml and add burp to roles
		//instance index, look into how destroy was done
		//all the stuff they need for butp
	},
}

var cobaltStrikeInstall = &cobra.Command{
	Use:   "cobaltstrike",
	Short: "Installs Cobalt Strike",
	Long:  `Installs Cobalt Strike to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		playbook := deployer.GeneratePlaybookFile("cobaltstrike")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
	},
}

var goPhishInstall = &cobra.Command{
	Use:   "gophish",
	Short: "Installs Gophish",
	Long:  `Installs Gophish to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		playbook := deployer.GeneratePlaybookFile("gophish")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
	},
}

var letsEncryptInstall = &cobra.Command{
	Use:   "letsencrypt",
	Short: "Installs Letsencrypt",
	Long:  `Installs Letsencrypt to remote server`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.ValidateListOfInstances(numberInput)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		playbook := deployer.GeneratePlaybookFile("letsencrypt")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
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
		playbook := deployer.GeneratePlaybookFile("nmap")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
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
		playbook := deployer.GeneratePlaybookFile("socat")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
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
		playbook := deployer.GeneratePlaybookFile("sqlmap")

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		instance := list[installIndex]

		hostFile := deployer.GenerateHostFile(instance.IP, instance.Username, instance.PrivateKey, fqdn, domain)

		deployer.WriteToFile("../ansible/hosts.yml", hostFile)
		deployer.WriteToFile("../ansible/main.yml", playbook)

		fmt.Println(deployer.ExecAnsible("hosts.yml", "main.yml", "../ansible"))
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
	nmapInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	nmapInstall.MarkPersistentFlagRequired("fqdn")
	nmapInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	nmapInstall.MarkPersistentFlagRequired("domain")

	socatInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	socatInstall.MarkFlagRequired("id")
	socatInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	socatInstall.MarkPersistentFlagRequired("fqdn")
	socatInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	socatInstall.MarkPersistentFlagRequired("domain")

	sqlMapInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
	sqlMapInstall.MarkFlagRequired("id")
	sqlMapInstall.PersistentFlags().StringVarP(&fqdn, "fqdn", "f", "", "Specify the FQDN for the instance's service")
	sqlMapInstall.MarkPersistentFlagRequired("fqdn")
	sqlMapInstall.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Specify the domain for the instance")
	sqlMapInstall.MarkPersistentFlagRequired("domain")
}
