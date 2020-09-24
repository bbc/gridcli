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
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

var env string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configures the Grid runtime dependencies",
	Long:  `Configures the Grid runtime dependencies listed on .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch env {
		case "docker":
			fmt.Println("selected AWS")
			fmt.Printf("\t%s\t Successfully installed in AWS\n", Success)
		case "aws":
			fmt.Println("selected AWS")
			fmt.Printf("\t%s\t Successfully installed in AWS\n", Success)
		default:
			s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)  // Build our new spinner
			s.Suffix = "configuring local environment..."
			s.Start()
			selectGOOS()
			s.Stop()

		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&env, "env", "e", "", "The deployment environment [local,AWS,GCP]. Default is [local].") // Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//installBrewLocally know how install brewfile.
func installBrewLocally() ([]byte, error) {
	out, er := exec.Command("brew", "bundle").Output()
	if er != nil {
		return nil, er
	}
	return out, nil
}

//selectGOOS knows how to configure to a specific operating system.
func selectGOOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		out, er := installBrewLocally()
		if er != nil {
			fmt.Printf("\t%s\t Failed configuraring grid local environment: %s bug: %s\n", Failed, out, er)
		} else {
			fmt.Printf("\t%s\t %s\n", Success, out)

		}
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
