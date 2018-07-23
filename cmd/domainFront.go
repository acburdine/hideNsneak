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
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var domainFrontProvider string
var domainFrontIndex int
var domainFrontOrigin string
var restrictUA string
var functionName string
var frontedDomain string

// helloCmd represents the hello command
var domainFront = &cobra.Command{
	Use:   "domainfront",
	Short: "domain front parent command",
	Long:  `domain front parent command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'domainfront --help' for usage.")
	},
}

var domainFrontDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "deploys a domain front",
	Long:  `initializes and deploys a domain front to either AWS Cloudfront or Azure where origin is the your target C2`,
	Args: func(cmd *cobra.Command, args []string) error {
		switch strings.ToUpper(domainFrontProvider) {
		case "AWS":
		case "GOOGLE":
			match, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9]+", functionName)

			if functionName == "" {
				return fmt.Errorf("Google Domain Fronts must have a function name (-n)")
			} else if !match {
				return fmt.Errorf("Google Domain Front function must begin with a letter and can only contain letters and numbers")
			}
		case "AZURE":
		default:
			return fmt.Errorf("Unknown provider")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()
		wrappers := deployer.CreateWrappersFromState(marshalledState)
		wrappers = deployer.DomainFrontDeploy(domainFrontProvider, domainFrontOrigin,
			restrictUA, functionName, frontedDomain, wrappers)

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()
	},
}

var domainFrontDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "destroy",
	Long:  `destroys an existing domain front`,
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
		} else {
			deployer.TerraformDestroy([]string{list[domainFrontIndex].Name})
		}
	},
}

var domainFrontDisable = &cobra.Command{
	Use:   "disable",
	Short: "disable domainfront",
	Long:  `disables an enabled Cloudfront domainfront`,
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
		switch list[domainFrontIndex].Provider {
		case "AWS":
			for index, front := range wrappers.Cloudfront {
				if list[domainFrontIndex].ID == front.ID {
					wrappers.Cloudfront[index].Enabled = "false"
				}
			}
		case "AZURE":
			//TODO: Implement Azure CDN domain fronting
		case "GOOGLE":
			fmt.Println("Disabling Google Domain Fronts is not currently supported")
			fmt.Println("Exiting....")
			// for index, front := range wrappers.Googlefront {
			// 	if list[domainFrontIndex].Invoke == front.InvokeURI {
			// 		wrappers.Googlefront[index].Enabled = false
			// 	}
			// }
			return
		default:
		}
		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()

	},
}

var domainFrontEnable = &cobra.Command{
	Use:   "enable",
	Short: "enable",
	Long:  `enables a disabled Cloudfront domain front`,
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

		switch list[domainFrontIndex].Provider {
		case "AWS":
			for index, front := range wrappers.Cloudfront {
				if list[domainFrontIndex].ID == front.ID {
					wrappers.Cloudfront[index].Enabled = "true"
				}
			}
		case "AZURE":
		case "GOOGLE":
			fmt.Println("Enabling Google Domain Fronts is not currently supported")
			fmt.Println("Exiting....")
			// for index, front := range wrappers.Googlefront {
			// 	if list[domainFrontIndex].Invoke == front.InvokeURI {
			// 		wrappers.Googlefront[index].Enabled = false
			// 	}
			// }
			return
		default:
		}

		mainFile := deployer.CreateMasterFile(wrappers)

		deployer.CreateTerraformMain(mainFile)

		deployer.TerraformApply()
	},
}

var domainFrontList = &cobra.Command{
	Use:   "list",
	Short: "list domain fronts",
	Long:  `list all domain fronts and their index, origin domain, invoke url, provider, and terraform name`,
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

func init() {
	rootCmd.AddCommand(domainFront)
	domainFront.AddCommand(domainFrontDeploy, domainFrontDestroy, domainFrontList, domainFrontEnable, domainFrontDisable)

	domainFrontDeploy.PersistentFlags().StringVarP(&domainFrontProvider, "provider", "p", "", "Specify the provider. i.e. AWS or Azure")
	domainFrontDeploy.MarkPersistentFlagRequired("provider")

	domainFrontDeploy.PersistentFlags().StringVarP(&domainFrontOrigin, "target", "t", "", "Specify the target domain or IP. i.e. yourc2example.com")
	domainFrontDeploy.MarkPersistentFlagRequired("target")

	domainFrontDeploy.PersistentFlags().StringVarP(&frontedDomain, "frontedDomain", "d", "", "Specify the Google domain to front i.e inbox.google.com")
	domainFrontDeploy.PersistentFlags().StringVarP(&functionName, "name", "n", "", "Specify the function name of the Google Domain front i.e /functionname1")
	domainFrontDeploy.PersistentFlags().StringVar(&restrictUA, "restrictua", "", "Specify the User Agent header to filter on for Google Domain Front")

	domainFrontEnable.PersistentFlags().IntVarP(&domainFrontIndex, "id", "i", 0, "Specify the id of the domain front")
	domainFrontEnable.MarkPersistentFlagRequired("id")

	domainFrontDisable.PersistentFlags().IntVarP(&domainFrontIndex, "id", "i", 0, "Specify the id of the domain front")
	domainFrontDisable.MarkPersistentFlagRequired("id")

	domainFrontDestroy.PersistentFlags().IntVarP(&domainFrontIndex, "id", "i", 0, "Specify the id of the domain front")
	domainFrontDestroy.MarkPersistentFlagRequired("id")

}
