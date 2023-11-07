package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	size string
)

var SplitCmd = &cobra.Command{
	Use:     "split",
	Short:   "split file",
	Example: "fs split -s 4M src.txt dst.txt or sf split -s 4M src.txt",
	Run: func(cmd *cobra.Command, args []string) {
		var srcFile string
		var dstPrefix string
		switch len(args) {
		case 1:
			srcFile = args[0]
			dstPrefix = args[0]
		case 2:
			srcFile = args[0]
			dstPrefix = args[1]
		default:
			fmt.Println("not expect input")
			_ = cmd.Help()
			return
		}
		chunkSize, err := parseSizeString(size)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = splitFile(srcFile, dstPrefix, chunkSize)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	SplitCmd.Flags().StringVarP(&size, "size", "s", "4M", "split file size, unit [K|k|M|m|G|g]")
}

func parseSizeString(sizeString string) (int, error) {
	sizeString = strings.TrimSpace(sizeString)
	unit := strings.ToUpper(sizeString[len(sizeString)-1:])
	sizeValueStr := sizeString[:len(sizeString)-1]

	sizeValue, err := strconv.Atoi(sizeValueStr)
	if err != nil {
		return 0, err
	}

	switch unit {
	case "K":
		return sizeValue << 10, nil
	case "M":
		return sizeValue << 20, nil
	case "G":
		return sizeValue << 30, nil
	default:
		return 0, errors.New("not valid unit")
	}
}

func splitFile(inputFilePath, dstPrefix string, chunkSize int) error {
	// open input file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// ensure dst dir exist
	if err := os.MkdirAll(path.Dir(dstPrefix), os.ModePerm); err != nil {
		return err
	}

	// read the input file and cut it into multiple file segments
	buffer := make([]byte, chunkSize)
	fileNumber := 1

	for {
		n, err := inputFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// create output file
		outputFilePath := fmt.Sprintf("%s%d", dstPrefix, fileNumber)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		// write to output file
		_, err = outputFile.Write(buffer[:n])
		if err != nil {
			return err
		}

		fileNumber++
	}

	return nil
}
