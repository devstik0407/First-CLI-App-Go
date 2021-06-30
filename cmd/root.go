package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swd",
	Short: "A command line application for downloading files at ease",
	Long: `A simple Command Line Application
Download any file from any url
Save it in any folder of you wish`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Yes I am here")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&Source, "source", "s", "", "source URL to download from")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	downloadCmd.Flags().StringVarP(&Dest, "destination", "d", wd, "directory to save the file")
	downloadCmd.MarkFlagRequired("source")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
