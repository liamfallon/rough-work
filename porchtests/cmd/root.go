package cmd

import (
	"os"

	porchTests "github.com/liamfallon/rough-work/tree/main/porchTests"

	"github.com/spf13/cobra"
)

var TestFile string

var rootCmd = &cobra.Command{
	Use:   "porch-tests",
	Short: "Run tests on Porch",
	Long:  `This command runs or cleans up after the porch test specified in the test yaml file`,
}

func Execute(ctx porchTests.testContext) {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	porchTests.DeleteAllPackages()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&TestFile, "test-file", "f", "", "The required test in a Yaml file")
	rootCmd.MarkPersistentFlagRequired("test-file")
}
