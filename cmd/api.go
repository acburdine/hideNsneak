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
	"regexp"

	"terraform-playground/deployer"

	"github.com/spf13/cobra"
)

var apiProvider string
var targetURI string

var api = &cobra.Command{
	Use:   "api",
	Short: "api child command",
	Long:  `API Gateway Command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'api --help' for usage.")
	},
}

var apiDeploy = &cobra.Command{
	//TODO: need to trim spaces
	Use:   "deploy",
	Short: "deploys an api gateway",
	Long:  `deploys an api gateway`,
	Args: func(cmd *cobra.Command, args []string) error {
		deployer.InitializeTerraformFiles()
		if !deployer.ProviderCheck(instanceProviders) {
			return fmt.Errorf("invalid providers specified: %v", instanceProviders)
		}
		r, _ := regexp.Compile(`http[s]{0,1}\:\/\/[a-zA-Z]+\.[a-zA-Z]+\/{1}[a-zA-Z]*`)
		if !r.MatchString(targetURI) {
			return fmt.Errorf("the target uri is formatted incorrectly")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {

		marshalledState := deployer.TerraformStateMarshaller()
		wrappers := deployer.CreateWrappersFromState(marshalledState)

		wrappers = deployer.APIDeploy(apiProvider, targetURI, wrappers)

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()
	},
}

var apiDestroy = &cobra.Command{
	// 	Use:   "destroy",
	Short: "destroy",
	Long:  `destroys an instance`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()
		apiList := deployer.ListAPIs(marshalledState)
		if !deployer.IsValidNumberInput(numberInput) {
			return fmt.Errorf("invalid formatting specified: %s", numberInput)
		}
		numsToDestroy := deployer.ExpandNumberInput(numberInput)
		largestNumToDestroy := deployer.FindLargestNumber(numsToDestroy)

		if largestNumToDestroy > len(apiList) {
			return errors.New("the number you entered is too big. try running `list` to see the number of apis you have")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListAPIs(marshalledState)
		numsToDelete := deployer.ExpandNumberInput(numberInput)

		var namesToDelete []string

		for _, numIndex := range numsToDelete {
			namesToDelete = append(namesToDelete, list[numIndex].Name)
		}

		deployer.TerraformDestroy(namesToDelete)
		return
	},
}

var apiList = &cobra.Command{
	Use:   "list",
	Short: "list api gateways",
	Long:  `list api gateways`,
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		apiList := deployer.ListAPIs(marshalledState)

		for index, item := range apiList {
			fmt.Print(index)
			fmt.Println(item.String())
		}
		return
	},
}

func init() {
	rootCmd.AddCommand(api)
	// instance.AddCommand(instanceDeploy, instanceDestroy, instanceInfo, instanceList)
	api.AddCommand(apiDeploy, apiList)

	apiDeploy.PersistentFlags().StringVarP(&apiProvider, "provider", "p", "", "the provider to use: i.e. AWS")
	apiDeploy.MarkPersistentFlagRequired("providers")

	apiDeploy.PersistentFlags().StringVarP(&apiProvider, "target", "t", "", "the target URI: i.e. https://google.com/")
	apiDeploy.MarkPersistentFlagRequired("target")

	// instanceDeploy.PersistentFlags().StringSliceVarP(&instanceProviders, "providers", "p", nil, "list of providers to enter")
	// instanceDeploy.MarkPersistentFlagRequired("providers")

	// instanceDeploy.PersistentFlags().IntVarP(&instanceCount, "count", "c", 0, "number of instances to deploy")
	// instanceDeploy.MarkPersistentFlagRequired("count")

	// instanceDeploy.PersistentFlags().StringVarP(&instancePrivateKey, "privatekey", "v", "", "full path to private key to connect to instances")
	// instanceDeploy.MarkPersistentFlagRequired("privatekey")

	// instanceDeploy.PersistentFlags().StringVarP(&instancePublicKey, "publickey", "b", "", "full path to public key corresponding to the private key")
	// instanceDeploy.MarkPersistentFlagRequired("publickey")

	// instanceDestroy.PersistentFlags().StringVarP(&numberInput, "input", "i", "", "number of instances to destroy")
	// instanceDestroy.MarkPersistentFlagRequired("input")

	// //TODO: default all regions
	// rootCmd.PersistentFlags().StringSliceVar(&regionAws, "region-aws", []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1", "eu-central-1", "eu-west-1", "eu-west-2", "eu-west-3", "ap-northeast-1", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2", "ap-south-1", "sa-east-1"}, "list of regions for aws. ex: us-east-1,us-west-2,ap-northeast-1")
	// rootCmd.PersistentFlags().StringSliceVar(&regionDo, "region-do", []string{"nyc1", "sgp1", "lon1", "nyc3", "ams3", "fra1", "tor1", "sfo2", "blr1"}, "list of regions for digital ocean. ex: AMS2,SFO2,NYC1")
	// rootCmd.PersistentFlags().StringSliceVar(&regionAzure, "region-azure", []string{"westus", "centralus"}, "list of regions for azure. ex: centralus, eastus, westus")
	// rootCmd.PersistentFlags().StringSliceVar(&regionGoogle, "region-google", []string{"us-west1", "us-east1"}, "list of regions for google. ex: us-east1, us-west1, us-central1")

}
