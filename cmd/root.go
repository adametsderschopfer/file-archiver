package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Go Archiver",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
