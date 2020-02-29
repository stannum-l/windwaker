package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   string
	GitCommit string
	GoVersion string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Windwaker",
	Long:  "All software has versions. This is Windwaker's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:    %v\nGit Commit: %v\nGo Version: %v\n", Version, GitCommit, GoVersion)
	},
}
