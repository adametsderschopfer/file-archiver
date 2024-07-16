package cmd

import "github.com/spf13/cobra"

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file by algorithm",
}

func init() {
	rootCmd.AddCommand(packCmd)
}
