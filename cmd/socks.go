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
	"strings"
	"terraform-playground/deployer"

	"github.com/spf13/cobra"
)

var socksPort int
var socksInstanceInput string

// helloCmd represents the hello command
var socks = &cobra.Command{
	Use:   "socks",
	Short: "socks",
	Long:  `socks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("socks called")
	},
}

var socksDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy SOCKS Proxy",
	Long:  `Deploy SOCKS Proxy`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)
		if !deployer.IsValidNumberInput(socksInstanceInput) {
			return fmt.Errorf("invalid formatting specified: %s", numberInput)
		}
		numsToDeploy := deployer.ExpandNumberInput(socksInstanceInput)
		largestInstanceNumToDestroy := deployer.FindLargestNumber(numsToDeploy)

		//make sure the largestInstanceNumToDestroy is not bigger than totalInstancesAvailable
		if len(list) < largestInstanceNumToDestroy {
			return errors.New("the number you entered is too big. Try running `list` to see the number of instances you have.")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		numsToDeploy := deployer.ExpandNumberInput(socksInstanceInput)
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		for _, num := range numsToDeploy {
			err := deployer.CreateSingleSOCKS(list[num].PrivateKey, list[num].Username, list[num].IP, socksPort)
			if err != nil {
				fmt.Println("SOCKS creation failed for " + list[num].IP)
			}
			socksPort = socksPort + 1
		}

	},
}

var socksDestroy = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy a SOCKS Proxy",
	Long:  `Destroy a SOCKS Proxy`,
	Args: func(cmd *cobra.Command, args []string) error {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)
		if !deployer.IsValidNumberInput(socksInstanceInput) {
			return fmt.Errorf("invalid formatting specified: %s", numberInput)
		}
		numsToDestroy := deployer.ExpandNumberInput(socksInstanceInput)
		largestInstanceNumToDestroy := deployer.FindLargestNumber(numsToDestroy)

		//make sure the largestInstanceNumToDestroy is not bigger than totalInstancesAvailable
		if len(list) < largestInstanceNumToDestroy {
			return errors.New("The number you entered is too big. Try running `list` to see the number of instances you have.")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		numsToDestroy := deployer.ExpandNumberInput(socksInstanceInput)

		for _, num := range numsToDestroy {
			deployer.DestroySOCKS(list[num].IP)
		}

	},
}

var socksList = &cobra.Command{
	Use:   "list",
	Short: "List available SOCKS Proxies",
	Long:  `List available SOCKS Proxies`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pulling Terraform State...")
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		output := deployer.ListProxies(list)

		fmt.Println(output)
	},
}

var proxychains = &cobra.Command{
	Use:   "proxychains",
	Short: "Proxychains Config",
	Long:  `Prints out the proper proxychains configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		output := deployer.ListProxies(list)

		fmt.Println(deployer.PrintProxyChains(output))
	},
}

var socksd = &cobra.Command{
	Use:   "socksd",
	Short: "SOCKSd config",
	Long:  `Prints out the proper socksd config`,
	Run: func(cmd *cobra.Command, args []string) {
		marshalledState := deployer.TerraformStateMarshaller()

		list := deployer.ListIPAddresses(marshalledState)

		output := deployer.ListProxies(list)

		output = strings.TrimSpace(output)

		fmt.Println(deployer.PrintSocksd(output))
	},
}

func init() {
	rootCmd.AddCommand(socks)
	socks.AddCommand(socksDeploy, socksDestroy, socksList, proxychains, socksd)

	socksDeploy.PersistentFlags().IntVarP(&socksPort, "port", "p", 8081, "Start port for socks proxy")
	socksDeploy.MarkPersistentFlagRequired("port")

	socksDeploy.PersistentFlags().StringVarP(&socksInstanceInput, "index", "i", "", "Indices of the instances to deploy")
	socksDeploy.MarkPersistentFlagRequired("index")

	socksDestroy.PersistentFlags().StringVarP(&socksInstanceInput, "index", "i", "", "Indices of the instances to deploy")
	socksDestroy.MarkPersistentFlagRequired("index")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
