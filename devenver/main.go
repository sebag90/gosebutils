package main

import (
	"devenver/units"
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

func DirSize(path string) (int64, error) {
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
	return size, err
}

func main() {
	var totalReclaimedSpace int64
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() && filepath.Base(path) == ".venv" {
			size, _ := DirSize(path)
			totalReclaimedSpace += size
			fmt.Printf("%s%s%s\n", PURPLE, path, END)
			err := os.RemoveAll(path)

			if err != nil {
				fmt.Println(err)
				return err
			}
			return fs.SkipDir
		}
		return nil
	})
	fmt.Printf("Total reclaimed space: %s%s%s\n", GREEN, units.HumanSize(float64(totalReclaimedSpace)), END)
}
