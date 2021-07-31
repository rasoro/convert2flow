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
	"encoding/json"
	"log"

	"github.com/rasoro/convert2flow/flow"
	"github.com/rasoro/convert2flow/io"
	"github.com/spf13/cobra"
)

// fromTxtCmd represents the fromTxt command
var fromTxtCmd = &cobra.Command{
	Use:   "from-txt",
	Short: "convert from txt",
	Long:  `convert from txt`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			convert2FlowFromTxt(args[0])
		} else {
			log.Fatal("please pass the file")
		}
	},
}

func init() {
	rootCmd.AddCommand(fromTxtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fromTxtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fromTxtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO: Add destination filename param
func convert2FlowFromTxt(filePath string) {
	file := io.NewTxtFile()
	file.ReadFile(filePath)
	newImport := flow.NewImport()
	newImport.Flows = []flow.Flow{
		flow.MessagesToFlow(file.Messages),
	}
	importJson, err := json.MarshalIndent(newImport, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteJsonToFile(string(importJson))
}
