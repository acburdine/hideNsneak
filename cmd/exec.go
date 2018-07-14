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

	"github.com/spf13/cobra"
)

var exec = &cobra.Command{
	Use:   "exec",
	Short: "execute custom command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'exec --help' for usage.")
	},
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute custom command",
	Long:  `executes the specified command on the specified remote system and returns both stdout and stderr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

var execNmap = &cobra.Command{
	Use:   "exec",
	Short: "execute nmap",
	Long:  `executes nmap and splits up the job between all of the specified hosts returning the xml files to the specified directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

func init() {
	rootCmd.AddCommand(exec)
	exec.AddCommand(execCmd, execNmap)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
