/*
Copyright © 2024 Kvs Vishnu Kumar <kvsvishnukumar@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var n int
var headers string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hcurl",
	Short: "hcurl is a command line tool primarily built for testing http(s) url endpoints.",
	Long:  `hcurl, short for "HTTP cURL", stands as a simple command-line utility designed with a primary focus on the testing and analysis of http(s) url endpoints`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is hcurl. Run `hcurl -h` for more info")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hcurl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().IntVarP(&n, "count", "n", 1, "Number of requests to send")
	rootCmd.PersistentFlags().StringVarP(&headers, "header", "H", "", "Pass headers here. Ex:- \"Content-Type: application/json, Authorization: Bearer xxx, x-api-key: xxxx\" etc")
}
