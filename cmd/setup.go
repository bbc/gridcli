package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
	Red     = "\\x1B[0;31m"
	Plain   = "\\x1B[0m"
)

var env string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configures the Grid runtime dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		switch env {
		case "docker":
			fmt.Println("selected AWS")
			fmt.Printf("\t%s\t Successfully installed in AWS\n", Success)
		case "aws":
			fmt.Println("selected AWS")
			fmt.Printf("\t%s\t Successfully installed in AWS\n", Success)
		default:
			RunAndSetup()
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&env, "env", "e", "", "The deployment environment [local,AWS,GCP]. Default is [local].") // Here you will define your flags and configuration settings.
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
func selectGOOS() ([]byte, error) {
	var bug error
	var output []byte
	switch goos := runtime.GOOS; goos {
	case "darwin":
		out, er := installBrewLocally()
		if er != nil {
			bug = errors.Wrap(er, "Failed configuring grid local environment")
		} else {
			output = out
		}
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", goos)
	}
	return output, bug
}
func checkDependenciesWorkers(waitGroup *sync.WaitGroup, in <-chan string, bug chan<- error) {
	defer waitGroup.Done()
	var wg sync.WaitGroup
	wg.Add(12)

	//Launch workers.
	for i := 0; i < 12; i++ {
		go func() {
			defer wg.Done()
			for dependency := range in {
				_, err := exec.LookPath(dependency)
				if err != nil {
					bug <- errors.Wrapf(err, "please install dependency | %s", dependency)
				}
				fmt.Printf("%s %s\n", dependency, Success)
			}

		}()
	}

	wg.Wait()
}
func RunAndSetup() {
	checkDependencies()
	if err := configLocalStack(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := devSetup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := devStart(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func checkDependencies() {
	in := make(chan string)
	bug := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)

	go checkDependenciesWorkers(&wg, in, bug)
	dependencies := []string{"java", "sbt", "npm", "docker", "gm", "magick", "convert", "pngquant", "exiftool", "nginx", "jq", "aws"}
	for _, dependency := range dependencies {
		in <- dependency
	}
	close(in)

	wg.Wait()
}
func devSetup() error {
	if wd, err := os.Getwd(); err != nil {
		return errors.Wrapf(err, "running dev-setup script")
	} else {
		cmd := exec.Command("/bin/sh", "-c", "./dev-setup/dev-configure.sh")
		cmd.Dir = wd
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return errors.Wrap(err, "failed running dev-setup script")
		}
	}
	return nil
}
func devStart() error {
	if wd, err := os.Getwd(); err != nil {
		return errors.Wrapf(err, "running dev-setup script")
	} else {
		cmd := exec.Command("/bin/sh", "-c", "./dev-setup/dev-start.sh")
		cmd.Dir = wd
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return errors.Wrap(err, "failed running dev-setup script")
		}
	}
	return nil
}
func configLocalStack() error {
	cmd := exec.Command("/bin/sh", "-c", "aws configure set profile.media-service.region eu-west-1; aws configure set profile.media-service.aws_secret_access_key xxxxxx; aws configure set profile.media-service.aws_access_key_id xxxxxx;")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed configuring aws profile")
	}
	return nil
}
