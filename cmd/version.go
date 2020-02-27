package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Info struct {
	Version string `json:"Version,omitempty"`
	Commit  string `json:"Commit,omitempty"`
}

func New(version string, commit string) *Info {
	return &Info{
		Version: version,
		Commit:  commit,
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Windwaker",
	Long:  "All software has versions. This is Windwaker's.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version -- TEST xx - HEAD")
	},
}
