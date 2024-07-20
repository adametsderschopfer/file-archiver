package cmd

import (
	"github.com/spf13/cobra"
	"go-archiver/lib/vlc"
	"os"
	"path/filepath"
	"strings"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code algorithm",
	Run:   unpack,
}

// TODO: Take extension from file
const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	filePath := args[0]

	// ToDo: Add stream instead of read full file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		handleError(err)
	}

	unpackedFileName := createPackedFileName(filePath)
	packed := vlc.Decode(fileData)

	err = os.WriteFile(unpackedFileName, []byte(packed), 0644)

}

// todo: refactor - remove duplicate and add extension to argument
func createUnPackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
