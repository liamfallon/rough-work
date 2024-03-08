package cmd

import (
	"os"

	porchtests "github.com/liamfallon/rough-work/porchtests"

	"github.com/spf13/cobra"
)

var TestFile string

var rootCmd = &cobra.Command{
	Use:   "porch-tests",
	Short: "Run tests on Porch",
	Long:  `This command runs or cleans up after the porch test specified in the test yaml file`,
}

func Execute(ctx porchtests.testContext) {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	porchtests.DeleteAllPackages()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&TestFile, "test-file", "f", "", "The required test in a Yaml file")
	rootCmd.MarkPersistentFlagRequired("test-file")
}
