package cmd

import (
	"encoding/json"
	"log"

	"github.com/rasoro/convert2flow/flow"
	"github.com/rasoro/convert2flow/io"
	"github.com/spf13/cobra"
)

// fromDocxCmd represents the from-docx command
var fromDocxCmd = &cobra.Command{
	Use:   "from-docx",
	Short: "convert from docx",
	Long:  `convert from docx`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			convert2FlowFromDocx(args[0])
		} else {
			log.Fatal("please pass the file")
		}
	},
}

func init() {
	rootCmd.AddCommand(fromDocxCmd)
}

// TODO: Add destination file name param
func convert2FlowFromDocx(filePath string) {
	file := io.NewDocxFile()
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
