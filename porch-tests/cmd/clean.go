/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up after a failed test",
	Long:  `Clean up the cluster after a failed Porch test`,
	Run: func(cmd *cobra.Command, args []string) {
		testFile, err := cmd.Flags().GetString("test-file")
		if err == nil {
			porchTests.ParseTestFile(testFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
