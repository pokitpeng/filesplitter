package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	prefix string
)

var MergeCmd = &cobra.Command{
	Use:     "merge",
	Short:   "merge files",
	Example: "fs merge -p src.txt dst.txt",
	Run: func(cmd *cobra.Command, args []string) {
		var out string
		if len(args) < 1 {
			fmt.Println("not expect input")
			_ = cmd.Help()
			return
		}
		out = args[0]

		err := mergeFiles(prefix, out)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	MergeCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "merge files prefix")
}

func mergeFiles(filePrefix string, outputFilePath string) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	fileNumber := 1
	for {
		inputFilePath := fmt.Sprintf("%s%d", filePrefix, fileNumber)
		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			if os.IsNotExist(err) {
				break // no more file to merge
			}
			return err
		}
		defer inputFile.Close()

		_, err = io.Copy(outputFile, inputFile)
		if err != nil {
			return err
		}

		fileNumber++
	}

	return nil
}
