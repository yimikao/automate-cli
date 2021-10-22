/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var getgopherCmd = &cobra.Command{
	Use:   "getgopher",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = "dr-who.png"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

		fmt.Println("Try to get '" + gopherName + "' Gopher...")
		//Get the data
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			//create the file
			file, err := os.Create(gopherName + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			//write the body to file
			_, err = io.Copy(file, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just saved in " + file.Name() + "!")
		} else {
			fmt.Println("Error: " + gopherName + " doesn't exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getgopherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getgopherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getgopherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
