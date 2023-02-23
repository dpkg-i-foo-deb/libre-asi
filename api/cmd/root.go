package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "libre-asi-api",
	Short: "Manage the Addiction Severity Index (ASI) form concepts with ease using JSON files, and implement any frontend with a flexible contract.",
	Run:   root,
}

func root(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to Libre-ASI API")
}

func Execute() {
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
