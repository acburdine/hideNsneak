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

// helloCmd represents the hello command
var domainFront = &cobra.Command{
	Use:   "domainfront",
	Short: "Domain Front Command",
	Long:  `Domain Front Command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Domain Front Called")
	},
}

var domainFrontDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "deploys a domain front",
	Long:  `Initializes and Deploys a domain front`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deploy Called")
	},
}

var domainFrontDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy",
	Long:  `Destroys an existing domain front`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Destroy Called")
	},
}

var domainFrontList = &cobra.Command{
	Use:   "list",
	Short: "list domain fronts",
	Long:  `list domain fronts`,
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListDomainFronts(marshalledState)

		for index, front := range list {
			fmt.Print(index)
			fmt.Println(front)
		}

		return
	},
}

var domainFrontInfo = &cobra.Command{
	Use:   "info",
	Short: "info",
	Long:  `provides information on specific domain front`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Info Called")
	},
}

func init() {
	rootCmd.AddCommand(domainFront)
	domainFront.AddCommand(domainFrontDeploy, domainFrontDestroy, domainFrontInfo, domainFrontList)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
