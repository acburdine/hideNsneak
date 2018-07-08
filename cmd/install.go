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
	},
	Run: func(cmd *cobra.Command, args []string) {
		// generate host file
		// generate playbook file

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

// var cobaltStrikeInstall = &cobra.Command{
// 	Use:   "cobaltstrike",
// 	Short: "Installs Cobalt Strike",
// 	Long:  `Installs Cobalt Strike to remote server`,
// Args: func(cmd *cobra.Command, args []string) error {
// 	deployer.ValidateListOfInstances(numberInput)
// },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//run burp installation here
// 	},
// }

// var goPhishInstall = &cobra.Command{
// 	Use:   "burp",
// 	Short: "Installs burp suite",
// 	Long:  `Installs burp suite to remote server`,
// Args: func(cmd *cobra.Command, args []string) error {
// 	deployer.ValidateListOfInstances(numberInput)
// },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//run burp installation here
// 	},
// }

// var letsEncryptInstall = &cobra.Command{
// 	Use:   "burp",
// 	Short: "Installs burp suite",
// 	Long:  `Installs burp suite to remote server`,
// Args: func(cmd *cobra.Command, args []string) error {
// 	deployer.ValidateListOfInstances(numberInput)
// },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//run burp installation here
// 	},
// }

// var nmapInstall = &cobra.Command{
// 	Use:   "burp",
// 	Short: "Installs burp suite",
// 	Long:  `Installs burp suite to remote server`,
// Args: func(cmd *cobra.Command, args []string) error {
// 	deployer.ValidateListOfInstances(numberInput)
// },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//run burp installation here
// 	},
// }

// var socatInstall = &cobra.Command{
// 	Use:   "burp",
// 	Short: "Installs burp suite",
// 	Long:  `Installs burp suite to remote server`,
// Args: func(cmd *cobra.Command, args []string) error {
// 	deployer.ValidateListOfInstances(numberInput)
// },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//run burp installation here
// 	},
// }

// var sqlMapInstall = &cobra.Command{
// 	Use:   "burp",
// 	Short: "Installs burp suite",
// 	Long:  `Installs burp suite to remote server`,
// Args: func(cmd *cobra.Command, args []string) error {
// 	deployer.ValidateListOfInstances(numberInput)
// },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		//run burp installation here
// 	},
// }

func init() {
	rootCmd.AddCommand(install)
	install.AddCommand(burpInstall /*, cobaltStrikeInstall, goPhishInstall, letsEncryptInstall, nmapInstall, socatInstall, sqlMapInstall*/)

	burpInstall.PersistentFlags().IntVarP(&installIndex, "id", "i", 0, "Specify the id for the install")
}
