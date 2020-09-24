/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configures the Grid runtime dependencies",
	Long:  `Configures the Grid runtime dependencies listed on .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		envFlag, er := cmd.PersistentFlags().GetString("env")
		if er != nil {
			b := installBrewLocally(ctx)
			if b {
				fmt.Printf("\t%s\t Successfully installed brew Brewfile", Success)
			} else {
				fmt.Printf("\t%s\t Failed installing brew Brewfile", Failed)

			}
		}
		switch envFlag {
		case "docker":
			fmt.Println("selected docker")
			fmt.Printf("\t%s\t Successfully installed in Docker", Success)

		case "aws":
			fmt.Println("selected AWS")
			fmt.Printf("\t%s\t Successfully installed in AWS", Success)

		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().String("help", "", "Used to set up runtime env. Must set --env flag [local,AWS,GCP]. Default is [local].")
	configCmd.PersistentFlags().String("env", "", "Set environment flag [docker,local,AWS,GCP]. Default is [local].")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//checkIfPresent know how to check if brew exists in the local environment
func installBrewLocally(ctx context.Context) bool {
	b, er := exec.CommandContext(ctx, "brew", "bundle", "&&", "brew install").Output()
	if er != nil || len(b) == 0 {
		return false
	}
	return true
}
