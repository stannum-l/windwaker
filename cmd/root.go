package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	ERR = 1
)

var (
	rootCmd = &cobra.Command{
		Use:   "windwaker",
		Short: "Wind Waker",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(ERR)
	}
}
