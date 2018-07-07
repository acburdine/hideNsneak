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

var domainFrontProvider string
var domainFrontIndex int
var domainFrontOrigin string

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
	Args: func(cmd *cobra.Command, args []string) error {
		if domainFrontProvider != "AWS" || domainFrontProvider != "AZURE" {
			return fmt.Errorf("Unknown provider")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()
		wrappers := deployer.CreateWrappersFromState(marshalledState)
		wrappers = deployer.DomainFrontDeploy(domainFrontProvider, domainFrontOrigin, wrappers)

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()
	},
}

var domainFrontDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy",
	Long:  `Destroys an existing domain front`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListDomainFronts(marshalledState)

		if domainFrontIndex > len(list)-1 {
			return fmt.Errorf("domainfront index not in bounds")
		}
		if list[domainFrontIndex].Provider == "AWS" {
			if list[domainFrontIndex].Status == "Enabled" {
				return fmt.Errorf("domainfront must be disabled to destroy")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListDomainFronts(marshalledState)
		currentDomainfront := list[domainFrontIndex]

		if list[domainFrontIndex].Provider == "AWS" {
			fmt.Println(deployer.AWSCloudFrontDestroy(currentDomainfront))
		}
	},
}

var domainFrontDisable = &cobra.Command{
	Use:   "disable",
	Short: "disable",
	Long:  `Disables an enabled domain front`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListDomainFronts(marshalledState)
		if domainFrontIndex > len(list)-1 {
			return fmt.Errorf("domainfront index not in bounds")
		}

		if list[domainFrontIndex].Provider == "AWS" {
			if list[domainFrontIndex].Status == "Disabled" {
				return fmt.Errorf("domainfront is already disabled")
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListDomainFronts(marshalledState)
		wrappers := deployer.CreateWrappersFromState(marshalledState)

		if list[domainFrontIndex].Provider == "AWS" {
			for index, front := range wrappers.Cloudfront {
				if list[domainFrontIndex].ID == front.ID {
					wrappers.Cloudfront[index].Enabled = "false"
				}
			}
		} else if domainFrontProvider == "Azure" {

		}

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()

	},
}

var domainFrontEnable = &cobra.Command{
	Use:   "enable",
	Short: "enable",
	Long:  `Enables a disabled domain front`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListDomainFronts(marshalledState)
		if domainFrontIndex > len(list)-1 {
			return fmt.Errorf("domainfront index not in bounds")
		}

		if list[domainFrontIndex].Provider == "AWS" {
			if list[domainFrontIndex].Status == "Enabled" {
				return fmt.Errorf("domainfront is already enabled")
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()
		list := deployer.ListDomainFronts(marshalledState)
		wrappers := deployer.CreateWrappersFromState(marshalledState)

		if list[domainFrontIndex].Provider == "AWS" {
			for index, front := range wrappers.Cloudfront {
				if list[domainFrontIndex].ID == front.ID {
					wrappers.Cloudfront[index].Enabled = "true"
				}
			}
		} else if list[domainFrontIndex].Provider == "Azure" {

		}

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()
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
			fmt.Println(front.String())
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
	domainFront.AddCommand(domainFrontDeploy, domainFrontDestroy, domainFrontInfo, domainFrontList, domainFrontEnable, domainFrontDisable)

	domainFrontDeploy.PersistentFlags().StringVarP(&domainFrontProvider, "provider", "p", "", "Specify the provider. i.e. AWS or Azure")
	domainFrontDeploy.MarkPersistentFlagRequired("provider")

	domainFrontDeploy.PersistentFlags().StringVarP(&domainFrontOrigin, "target", "t", "", "Specify the target domain. i.e. google.com")
	domainFrontDeploy.MarkPersistentFlagRequired("target")

	domainFrontEnable.PersistentFlags().IntVarP(&domainFrontIndex, "id", "i", 0, "Specify the id of the domain front")
	domainFrontEnable.MarkPersistentFlagRequired("id")

	domainFrontDisable.PersistentFlags().IntVarP(&domainFrontIndex, "id", "i", 0, "Specify the id of the domain front")
	domainFrontDisable.MarkPersistentFlagRequired("id")

	domainFrontDestroy.PersistentFlags().IntVarP(&domainFrontIndex, "id", "i", 0, "Specify the id of the domain front")
	domainFrontDestroy.MarkPersistentFlagRequired("id")
}
