package cmd

import (
	"fmt"
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
