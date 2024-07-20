package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"go-archiver/lib/compression"
	"go-archiver/lib/compression/vlc"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file by algorithm",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("Unknown method")
	}

	filePath := args[0]

	// ToDo: Add stream instead of read full file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		handleError(err)
	}

	packedFileName := createPackedFileName(filePath)
	packed := encoder.Encode(string(fileData))

	err = os.WriteFile(packedFileName, packed, 0644)

}

func createPackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "Compression methods: vlc")
	err := packCmd.MarkFlagRequired("method")
	if err != nil {
		panic(err)
	}
}
