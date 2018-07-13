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
	Short: "API Gateway parent command",
	Long:  `parent command for deploying API gateways, via a target parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'api --help' for usage.")
	},
}

var apiDeploy = &cobra.Command{
	//TODO: need to trim spaces
	Use:   "deploy",
	Short: "deploys an API Gateway",
	Long:  `deploys an API Gateway for AWS only at this time`,
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
	Use:   "destroy",
	Short: "destroys an API Gateway",
	Long:  `destroys an API Gateway by choosing an index`,
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
	Short: "detailed list of API Gateways",
	Long:  `list API Gateways and show their target URIs, invoke URIs, providers and names`,
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
	api.AddCommand(apiDeploy, apiList, apiDestroy)

	apiDeploy.PersistentFlags().StringVarP(&apiProvider, "provider", "p", "", "the provider to use: i.e. AWS")
	apiDeploy.MarkPersistentFlagRequired("providers")

	apiDeploy.PersistentFlags().StringVarP(&apiProvider, "target", "t", "", "the target URI: i.e. https://google.com/")
	apiDeploy.MarkPersistentFlagRequired("target")

	apiDestroy.PersistentFlags().StringVarP(&numberInput, "input", "i", "", "number of instances to destroy")
	apiDestroy.MarkPersistentFlagRequired("input")
}
