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
package main

import (
	"github.com/gridcli/cmd"
)

func main() {
	//if wd, err := os.Getwd(); err != nil {
	//	fmt.Println("error pop",err)
	//}else{
	//	cmd := exec.Command("/bin/sh", "-c", "cat .nvmrc")
	//	cmd.Dir = wd
	//	cmd.Stdout = os.Stdout
	//	cmd.Stderr = os.Stderr
	//	if err := cmd.Run(); err != nil {
	//		fmt.Println("error 2",er)
	//	}
	//}
	//path, err := exec.LookPath("nvm")
	//if err != nil {
	//	fmt.Println("we found an error", err)
	//} else {
	//	fmt.Println("GOT", path)
	//}
	//cmd.RunAndSetup()
	//cmd := exec.Command("/bin/sh", "-c", "./dev-setup/dev-configure.sh")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//if err := cmd.Run(); err != nil {
	//	fmt.Println(err)
	//}
	//comment
	cmd.Execute()
}
