package cmd

import "github.com/spf13/cobra"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Pack file by algorithm",
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
