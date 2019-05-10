package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	GitCommit string
	Version   string
	GoVersion string
	OsArchi   string
	BuildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:    %s\n", Version)
		fmt.Printf("Commit:     %s\n", GitCommit)
		fmt.Printf("Build date: %s\n", BuildDate)
		fmt.Printf("Go version: %s\n", GoVersion)
		fmt.Printf("OS/Arch::   %s\n", OsArchi)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
