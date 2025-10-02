package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var (
		path               string
		filePattern        string
		excludeFilePattern string
		windowSize         int
		ignoreCase         bool
	)

	var rootCmd = &cobra.Command{
		Use:   "search",
		Short: "Search files recursively in a directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			searchPattern := args[0]
			Search(path, searchPattern, filePattern, excludeFilePattern, windowSize, ignoreCase)
		},
	}

	rootCmd.Flags().StringVarP(&path, "path", "p", ".", "Directory to search")
	rootCmd.Flags().StringVarP(&filePattern, "file-regex", "f", ".", "regex pattern to filter files")
	rootCmd.Flags().StringVarP(&excludeFilePattern, "exclude-file", "e", "a^", "regex pattern to exclude files from search")
	rootCmd.Flags().IntVarP(&windowSize, "winsize", "w", 10, "Size of the print window")
	rootCmd.Flags().BoolVarP(&ignoreCase, "ignore-case", "i", false, "ignore case")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
