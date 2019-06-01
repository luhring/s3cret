package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "s3cret",
	Short: "Copy files to/from AWS S3 with automatic client-side encryption/decryption",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
