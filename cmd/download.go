package cmd

import (
	"cliApp/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download command",
	Long:  `download command let's you download file from a url`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName, err := utils.DownloadAndSave(Source, Dest)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully downloaded", fileName)
	},
}

var Source string
var Dest string
