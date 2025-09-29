package search

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
)

func collectPaths(root, pattern string) ([]string, error) {
	re := regexp.MustCompile(pattern)
	files := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // propagate error
		}

		if !d.IsDir() {
			if re.MatchString(path) {
				files = append(files, path)
			}
		}

		return nil
	})

	return files, err
}

func Search(path, pattern string) string {
	files, err := collectPaths(path, pattern)
	if err != nil {
		log.Fatalf("error walking the path: %v", err)
	}
	fmt.Println(files)
	return "ciao"
}
