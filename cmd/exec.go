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

var nmapPorts []string
var nmapHostFile string
var nmapCommand string
var nmapOutput string
var nmapEvasive bool
var nmapCommands map[int][]string
var execCommand string
var socatPort string
var socatIP string
var cobaltStrikeLicense string
var cobaltStrikeIp string
var cobaltStrikePassword string
var cobaltStrikeC2Path string

var commandIndices []int

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

		var instances []deployer.ListStruct

		for _, num := range commandIndices {
			instances = append(instances, list[num])
		}

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			cobaltStrikeLicense, cobaltStrikePassword, cobaltStrikeC2Path,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

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

		err = deployer.ValidateNumberOfInstances(commandIndices)
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"nmap", "nmap-exec"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []deployer.ListStruct

		for _, num := range commandIndices {
			instances = append(instances, list[num])
		}

		nmapCommands := deployer.SplitNmapCommandsIntoHosts(nmapPorts, nmapHostFile, nmapCommand, len(instances), nmapEvasive)

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			cobaltStrikeLicense, cobaltStrikePassword, cobaltStrikeC2Path,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var socatRedirect = &cobra.Command{
	Use:   "socat-redirect",
	Short: "redirects ports to target hosts",
	Long:  "initializes scat redirector that sends all traffic from the specified port to the specified target",
	Args: func(cmd *cobra.Command, args []string) error {
		return deployer.ValidateNumberOfInstances(commandIndices)
	},
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"socat", "socat-exec"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []deployer.ListStruct

		for _, num := range commandIndices {
			instances = append(instances, list[num])
		}

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			cobaltStrikeLicense, cobaltStrikePassword, cobaltStrikeC2Path,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

var cobaltStrikeRun = &cobra.Command{
	Use:   "cobaltstrike-run",
	Short: "updates and runs cobalt strike teamserver",
	Long:  "updates the cobalt strike teamserver with the licensse and starts the teamserver with specified profile and password ",
	Run: func(cmd *cobra.Command, args []string) {
		apps := []string{"cobaltstrike", "cobaltstrike-exec"}

		playbook := deployer.GeneratePlaybookFile(apps)

		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		var instances []deployer.ListStruct

		for _, num := range commandIndices {
			instances = append(instances, list[num])
		}

		remoteFilePath = "/opt/cobaltstrike"

		hostFile := deployer.GenerateHostFile(instances, fqdn, domain, burpDir, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			cobaltStrikeLicense, cobaltStrikePassword, cobaltStrikeC2Path,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")
	},
}

// ---
// # Starts up Cobalt Strike
// cobaltstrike_license
// public_ip
// password
// path_to_malleable_c2

// ---
// # Synchronize two directories on one remote host.
// remote_absolute_path
// host_absolute_path

func init() {
	rootCmd.AddCommand(exec)
	exec.AddCommand(command, nmap, socatRedirect, cobaltStrikeRun)

	command.PersistentFlags().IntSliceVarP(&commandIndices, "id", "i", []int{}, "Specify the id(s) for the remote server")
	command.MarkFlagRequired("id")
	command.PersistentFlags().StringVarP(&execCommand, "command", "c", "", "Specify the command you want to execute")
	command.MarkPersistentFlagRequired("command")

	nmap.PersistentFlags().IntSliceVarP(&commandIndices, "id", "i", []int{}, "Specify the id(s) for the scanning servers")
	nmap.MarkPersistentFlagRequired("id")
	nmap.PersistentFlags().StringVarP(&nmapHostFile, "hostFile", "f", "", "Specify filepath of the file containing the scope/hosts")
	nmap.MarkPersistentFlagRequired("hostFile")
	nmap.PersistentFlags().StringSliceVarP(&nmapPorts, "ports", "p", []string{}, "Specify the port range to be passed to nmap i.e 21-23,443-445,8080-8081,8443")
	nmap.MarkPersistentFlagRequired("ports")
	nmap.PersistentFlags().StringVarP(&nmapCommand, "nmapCommand", "n", "", "Specify the full nmap command to be run, excluding the -iL,-p, and -oA options options i.e. nmap -sV -sT --max-rate=250")
	nmap.MarkPersistentFlagRequired("nmapCommand")
	nmap.PersistentFlags().StringVarP(&nmapOutput, "nmapOutput", "o", "", "Specify the local directory for output to be saved")
	nmap.MarkPersistentFlagRequired("nmapOutput")
	nmap.PersistentFlags().BoolVarP(&nmapEvasive, "nmapEvasion", "e", false, "Specify whether or not you want nmap to be evasive i.e. true or false")

	socatRedirect.PersistentFlags().IntSliceVarP(&commandIndices, "id", "i", []int{}, "Specify the id(s) for the remote server")
	socatRedirect.MarkFlagRequired("id")
	socatRedirect.PersistentFlags().StringVarP(&socatPort, "port", "p", "", "Specify the port you want to forward")
	socatRedirect.MarkPersistentFlagRequired("port")
	socatRedirect.PersistentFlags().StringVarP(&socatIP, "target", "t", "", "Specify the target ip address for the socat redirector")
	socatRedirect.MarkPersistentFlagRequired("target")

	cobaltStrikeRun.PersistentFlags().IntSliceVarP(&commandIndices, "id", "i", []int{}, "Specify the id for the remote server")
	cobaltStrikeRun.MarkFlagRequired("id")
	cobaltStrikeRun.PersistentFlags().StringVarP(&cobaltStrikeLicense, "license", "l", "", "Specify the cobalt strike license")
	cobaltStrikeRun.MarkPersistentFlagRequired("license")
	cobaltStrikeRun.PersistentFlags().StringVarP(&cobaltStrikeIp, "target", "t", "", "Specify the target ip address")
	cobaltStrikeRun.MarkPersistentFlagRequired("target")
	cobaltStrikeRun.PersistentFlags().StringVarP(&cobaltStrikePassword, "password", "p", "", "Enter your password")
	cobaltStrikeRun.MarkPersistentFlagRequired("password")
	cobaltStrikeRun.PersistentFlags().StringVarP(&cobaltStrikeC2Path, "c2", "c", "", "Specify the malleable C2 path")
	cobaltStrikeRun.MarkPersistentFlagRequired("c2")
	cobaltStrikeRun.PersistentFlags().StringVarP(&localFilePath, "file", "f", "", "cobaltstrike ")
	cobaltStrikeRun.MarkPersistentFlagRequired("localhost")
}
