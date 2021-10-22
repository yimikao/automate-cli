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
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A command to take quick notes",
	Long:  `This command helps you take quick notes by opening a file .txt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Homer is opening a new note")
		_ = os.Mkdir("./notes", 0755)

		editorCmd := exec.Command("nano", fmt.Sprintf("./notes/%v.txt", args[0]))
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		err := editorCmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
