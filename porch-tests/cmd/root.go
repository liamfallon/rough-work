package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var TestFile string

var rootCmd = &cobra.Command{
	Use:   "porch-tests",
	Short: "Run tests on Porch",
	Long:  `This command runs or cleans up after the porch test specified in the test yaml file`,
}

func Execute(ctx testContext) {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&TestFile, "test-file", "f", "", "The required test in a Yaml file")
	rootCmd.MarkPersistentFlagRequired("test-file")
}
