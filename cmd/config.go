package cmd

import (
	"github.com/AdFabConnect/ovh-cli/ovh"
	"github.com/spf13/cobra"
)

var displayCurrentConfig = &cobra.Command{
	Use:   "current",
	Short: "Display current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		ovh.DisplayCurrentOvhConfig()
	},
}

var displayConfig = &cobra.Command{
	Use:   "all",
	Short: "Display current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		ovh.DisplayOvhConfig()
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Ovh configuration",
}

func init() {
	configCmd.AddCommand(displayConfig)
	configCmd.AddCommand(displayCurrentConfig)
	rootCmd.AddCommand(configCmd)
}
