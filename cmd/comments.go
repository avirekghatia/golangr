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

	"github.com/avirekghatia/golangr/comments"
	"github.com/spf13/cobra"
)

// commentsCmd represents the comments command
var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "Returns go code string with comments in it",
	Long: `As part of the exercise golangr.com/comments,
this command returns a string that contains the go code
as string that can be copied to a main.go file and run 
as an example in the exercise above.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(comments.CreateComments())
	},
}

func init() {
	rootCmd.AddCommand(commentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
