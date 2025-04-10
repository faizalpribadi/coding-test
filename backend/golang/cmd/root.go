package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() error {
	// Add subcommands to the root command
	rootCmd.AddCommand(serverCmd)

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
