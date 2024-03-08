/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"liamfallon/rough-work/porchtests/testrunner"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the porch test",
	Long:  `Run the porch test described in the test-file`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		testFile, err := cmd.Flags().GetString("test-file")
		if err != nil {
			return
		}
		ctx, err := testrunner.ParseTestFile(testFile)
		if err == nil {
			testrunner.Run(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
