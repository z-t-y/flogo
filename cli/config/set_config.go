/*
Copyright © 2021 Andy Zhou

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
package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/z-t-y/flogo/utils"
)

// setConfigCmd represents the setConfig command
var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "Write a config to your config file",
	Long:  `Write a config to your config file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("accepts 2 args, received", len(args))
			os.Exit(1)
		}
		err := utils.WriteToConfig(args[0], args[1])
		cobra.CheckErr(err)
	},
}

func init() {
	configCmd.AddCommand(setConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
