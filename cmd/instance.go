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
	"strings"
	"time"

	"hideNsneak/deployer"

	"github.com/schollz/progressbar"
	"github.com/spf13/cobra"
)

var instanceProviders []string
var instancePrivateKey string
var instancePublicKey string
var instanceCount int
var regionAws []string
var regionDo []string
var regionAzure []string
var regionGoogle []string
var instanceDestroyIndices []int

var instance = &cobra.Command{
	Use:   "instance",
	Short: "instance parent command",
	Long:  `parent command for managing instances`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'instance --help' for usage.")
	},
}

var instanceDeploy = &cobra.Command{
	//TODO: need to trim spaces
	Use:   "deploy",
	Short: "deploys instances",
	Long:  `deploys instances for AWS, Azure, Digital Ocean, or Google Cloud`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !deployer.ProviderCheck(instanceProviders) {
			return fmt.Errorf("invalid providers specified: %v", instanceProviders)
		}
		if deployer.ContainsString(instanceProviders, "DO") {
			availableDORegions := deployer.GetDoRegions()
			var unavailableRegions []string
			for _, region := range regionDo {
				if !deployer.ContainsString(availableDORegions, strings.ToLower(region)) {
					unavailableRegions = append(unavailableRegions, region)
				}
			}
			if len(unavailableRegions) != 0 {
				return fmt.Errorf("digitalocean region(s) not available: %s", strings.Join(unavailableRegions, ","))
			}
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {

		marshalledState := deployer.TerraformStateMarshaller()
		wrappers := deployer.CreateWrappersFromState(marshalledState)

		oldList := deployer.ListInstances(marshalledState)

		wrappers = deployer.InstanceDeploy(instanceProviders, regionAws, regionDo, regionAzure, regionGoogle, instanceCount, instancePrivateKey, instancePublicKey, "hidensneak", wrappers)

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()

		fmt.Println("Waiting for instances to initialize...")

		bar := progressbar.New(120)
		for i := 0; i < 120; i++ {
			bar.Add(1)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("")
		fmt.Println("Restricting Ports to only port 22...")

		marshalledState = deployer.TerraformStateMarshaller()
		newList := deployer.ListInstances(marshalledState)
		firewallList := deployer.InstanceDiff(oldList, newList)

		apps := []string{"firewall"}
		playbook := deployer.GeneratePlaybookFile(apps)

		ufwTCPPorts = []string{"22"}
		ufwAction = "add"

		hostFile := deployer.GenerateHostFile(firewallList, fqdn, domain, burpFile, localFilePath, remoteFilePath,
			execCommand, socatPort, socatIP, nmapOutput, nmapCommands,
			cobaltStrikeLicense, cobaltStrikePassword, cobaltStrikeC2Path, cobaltStrikeFile, cobaltStrikeKillDate,
			ufwAction, ufwTCPPorts, ufwUDPPorts)

		deployer.WriteToFile("ansible/hosts.yml", hostFile)
		deployer.WriteToFile("ansible/main.yml", playbook)

		deployer.ExecAnsible("hosts.yml", "main.yml", "ansible")

	},
}

var instanceDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroys instances",
	Long:  `destroys instances by choosing an index`,
	Args: func(cmd *cobra.Command, args []string) error {
		return deployer.ValidateNumberOfInstances(instanceDestroyIndices)
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListInstances(marshalledState)

		var namesToDelete []string

		for _, numIndex := range instanceDestroyIndices {
			namesToDelete = append(namesToDelete, list[numIndex].Name)
		}

		emptyEC2Modules := deployer.CheckForEmptyEC2Module(namesToDelete, marshalledState)

		namesToDelete = append(namesToDelete, emptyEC2Modules...)

		deployer.TerraformDestroy(namesToDelete)
		return
	},
}

var instanceList = &cobra.Command{
	Use:   "list",
	Short: "detailed list of instances",
	Long:  `list instances and shows their index, IP, provider, region, and name`,
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListInstances(marshalledState)

		for index, item := range list {
			fmt.Print(index)
			fmt.Println(" : " + item.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(instance)
	instance.AddCommand(instanceDeploy, instanceDestroy, instanceList)

	instanceDeploy.PersistentFlags().StringSliceVarP(&instanceProviders, "providers", "p", nil, "list of providers to enter")
	instanceDeploy.MarkPersistentFlagRequired("providers")

	instanceDeploy.PersistentFlags().IntVarP(&instanceCount, "count", "c", 0, "number of instances to deploy")
	instanceDeploy.MarkPersistentFlagRequired("count")

	instanceDeploy.PersistentFlags().StringVarP(&instancePrivateKey, "privatekey", "v", "", "full path to private key to connect to instances")
	instanceDeploy.MarkPersistentFlagRequired("privatekey")

	instanceDeploy.PersistentFlags().StringVarP(&instancePublicKey, "publickey", "b", "", "full path to public key corresponding to the private key")
	instanceDeploy.MarkPersistentFlagRequired("publickey")

	instanceDestroy.PersistentFlags().IntSliceVarP(&instanceDestroyIndices, "input", "i", []int{}, "indices of instances to destroy")
	instanceDestroy.MarkPersistentFlagRequired("input")

	//TODO: default all regions
	instanceDeploy.PersistentFlags().StringSliceVar(&regionAws, "region-aws", []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1", "eu-central-1", "eu-west-1", "eu-west-2", "eu-west-3", "ap-northeast-1", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2", "ap-south-1", "sa-east-1"}, "list of regions for aws. ex: us-east-1,us-west-2,ap-northeast-1")
	instanceDeploy.PersistentFlags().StringSliceVar(&regionDo, "region-do", []string{"nyc1", "sgp1", "lon1", "nyc3", "ams3", "fra1", "tor1", "sfo2", "blr1"}, "list of regions for digital ocean. ex: AMS2,SFO2,NYC1")
	instanceDeploy.PersistentFlags().StringSliceVar(&regionAzure, "region-azure", []string{"westus", "centralus"}, "list of regions for azure. ex: centralus, eastus, westus")
	instanceDeploy.PersistentFlags().StringSliceVar(&regionGoogle, "region-google", []string{"us-west1", "us-east1"}, "list of regions for google. ex: us-east1, us-west1, us-central1")

}
