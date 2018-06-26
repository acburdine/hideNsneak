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

	"terraform-playground/deployer"

	"github.com/spf13/cobra"
)

var instanceProviders []string
var instancePrivateKey string
var instancePublicKey string
var instanceCount int

// helloCmd represents the hello command
var instance = &cobra.Command{
	Use:   "instance",
	Short: "instance parent command",
	Long:  `Domain Front Command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Instance Called")
	},
}

var instanceDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "deploys an instance",
	Long:  `deploys an instance`,
	Args: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deploy Called")
		fmt.Println(instanceProviders)
		deployer.InitializeTerraformFiles()
		if deployer.ProviderCheck(instanceProviders) {
			return nil
		}
		if len(instanceProviders) < 1 {
			return errors.New("you need to enter at least one provider")
		}
		return fmt.Errorf("invalid providers specified: &s", instanceProviders)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("got in!")
	},
}

var instanceDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy",
	Long:  `Destroys an instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Destroy Called")
	},
}

var instanceList = &cobra.Command{
	Use:   "list",
	Short: "list instances",
	Long:  `list instances`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List called")
	},
}

var instanceInfo = &cobra.Command{
	Use:   "info",
	Short: "info",
	Long:  `provides information on specific instance`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Info Called")
	},
}

func init() {
	rootCmd.AddCommand(instance)
	instance.AddCommand(instanceDeploy, instanceDestroy, instanceInfo, instanceList)

	// rootCmd.PersistentFlags().StringSliceVarP(&region, "region", "r", "", "")
	instanceDeploy.PersistentFlags().StringSliceVarP(&instanceProviders, "providers", "p", []string{}, "List of providers to enter")
	instanceDeploy.MarkFlagRequired("providers")

	instanceDeploy.PersistentFlags().IntVarP(&instanceCount, "count", "c", 0, "Number of Instances to Deploy")
	instanceDeploy.MarkFlagRequired("count")

	// instanceDeploy.PersistentFlags().StringVarP(&instancePrivateKey, "privateKey", "priv", "Full Path to Private Key to Connect to Instances")
	// instanceDeploy.MarkFlagRequired("count")

	// instanceDeploy.PersistentFlags().StringVarP(&instanceCount, "publicKey", "pub", 0, "Full Path to Public Key Corresponding to the Private Key")
	// instanceDeploy.MarkFlagRequired("count")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
