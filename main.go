package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pokitpeng/filesplitter/cmd"
	"github.com/spf13/cobra"
)

var (
	Version   string = ""
	BuildTime string = time.Now().Format(time.RFC3339)
)

var rootCmd = &cobra.Command{
	Use:   "fs",
	Short: "split and merge file tool",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}
	},
	Version: fmt.Sprintf("version: %s\nbuild time: %s\n", Version, BuildTime),
}

func init() {
	rootCmd.AddCommand(cmd.SplitCmd)
	rootCmd.AddCommand(cmd.MergeCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
