// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hidensneak",
	Short: "hideNsneak is an application that will help you automate cloud management.",
	Long: `
	__     __     __         _______                              __    
	|  |--.|__|.--|  |.-----.|    |  |.-----..-----..-----..---.-.|  |--.
	|     ||  ||  _  ||  -__||       ||__ --||     ||  -__||  _  ||    < 
	|__|__||__||_____||_____||__|____||_____||__|__||_____||___._||__|__|
																		 

hideNsneak is a CLI that empowers red teamers during penetration testing.
This application is a tool that automates deployment, management, and destruction
of cloud infrastructure.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$GOPATH/src/github.com/rmikehodges/hideNsneak/config/config.json", "config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig expands the filepath for the user
func initConfig() {
	goPath := os.Getenv("GOPATH")
	cfgFile = goPath + "/src/github.com/rmikehodges/hideNsneak/config/config.json"
}
