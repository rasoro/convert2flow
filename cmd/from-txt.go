package cmd

import (
	"encoding/json"
	"log"

	"github.com/rasoro/convert2flow/flow"
	"github.com/rasoro/convert2flow/io"
	"github.com/spf13/cobra"
)

// fromTxtCmd represents the from-txt command
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
