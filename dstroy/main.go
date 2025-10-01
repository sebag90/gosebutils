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

func main() {
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Base(path) == ".DS_Store" {
			fmt.Printf("%s%s%s\n", PURPLE, path, END)
			os.Remove(path)
		}
		return err
	})
}
