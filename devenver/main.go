package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	PURPLE = "\u001b[95m"
	GREEN  = "\u001b[92m"
	RED    = "\u001b[91m"
	END    = "\u001b[0m"
	YELLOW = "\u001b[93m"
)

func DirSize(path string) (int, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return int(float64(size) / 1024.0 / 1024.0), err
}

func main() {
	totalReclaimedSpace := 0
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, _ := os.Stat(path)
		if info.IsDir() {
			if filepath.Base(path) == ".venv" {
				size, _ := DirSize(path)
				totalReclaimedSpace += size
				fmt.Printf("%s%s%s\n", PURPLE, path, END)

				os.RemoveAll(path)
			}

		}
		return err
	})
	fmt.Printf("Total reclaimed space: %s%d MB%s\n", GREEN, totalReclaimedSpace, END)
}
