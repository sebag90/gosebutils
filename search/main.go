package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var (
		path        string
		filePattern string
		windowSize  int
		ignoreCase  bool
	)

	var rootCmd = &cobra.Command{
		Use:   "sebsearch",
		Short: "Search files in a directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			searchPattern := args[0]
			Search(path, filePattern, searchPattern, windowSize, ignoreCase)
		},
	}

	rootCmd.Flags().StringVar(&path, "path", ".", "Directory to search")
	rootCmd.Flags().StringVar(&filePattern, "file-regex", ".", "regex pattern to filter files")
	rootCmd.Flags().IntVar(&windowSize, "winsize", 10, "Size of the print window")
	rootCmd.Flags().BoolVar(&ignoreCase, "ignore-case", false, "ignore case")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
