package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code algorithm",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	filePath := args[0]

	// ToDo: Add stream instead of read full file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		handleError(err)
	}

	packedFileName := createPackedFileName(filePath)
	// packed := Encode(fileData)
	packed := ""
	_ = fileData

	err = os.WriteFile(packedFileName, []byte(packed), 0644)

}

func createPackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
