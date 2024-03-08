package cmd

import (
	"liamfallon/rough-work/porchtests/testrunner"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up after a failed test",
	Long:  `Clean up the cluster after a failed Porch test`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		testFile, err := cmd.Flags().GetString("test-file")
		if err != nil {
			return
		}
		ctx, err := testrunner.ParseTestFile(testFile)
		if err == nil {
			testrunner.DeleteAllPackages(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
