package cmd

import (
	"github.com/spf13/cobra"
	"go-archiver/lib/compression"
	"go-archiver/lib/compression/vlc"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Pack file by algorithm",
	Run:   unpack,
}

// TODO: Take extension from file
const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("Unknown method")
	}

	filePath := args[0]

	// ToDo: Add stream instead of read full file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		handleError(err)
	}

	unpackedFileName := createPackedFileName(filePath)
	packed := decoder.Decode(fileData)

	err = os.WriteFile(unpackedFileName, []byte(packed), 0644)

}

// todo: refactor - remove duplicate and add extension to argument
func createUnPackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "Compression methods: vlc")
	err := unpackCmd.MarkFlagRequired("method")
	if err != nil {
		panic(err)
	}
}
