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

// fromCsvCmd represents the fromCsv command
var fromCsvCmd = &cobra.Command{
	Use:   "from-csv",
	Short: "convert from csv",
	Long:  `convert from csv`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			convert2FlowFromCsv(args[0])
		} else {
			log.Fatal("please pass the file")
		}
	},
}

func init() {
	rootCmd.AddCommand(fromCsvCmd)
}

func convert2FlowFromCsv(filePath string) {
	file := io.NewCsvFile()
	file.ReadFile(filePath)
	newImport := flow.NewImport()
	newImport.Flows = []flow.Flow{
		flow.MessagesToFlow(file.Messages),
	}
	importJson, err := json.MarshalIndent(newImport, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteJsonToFile(string(importJson))
}
