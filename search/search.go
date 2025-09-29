package search

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	PURPLE = "\u001b[95m"
	GREEN  = "\u001b[92m"
	RED    = "\u001b[91m"
	END    = "\u001b[0m"
	YELLOW = "\u001b[93m"
)

func collectPaths(root string, pattern *regexp.Regexp) ([]string, error) {
	files := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // propagate error
		}

		if !d.IsDir() {
			if pattern.MatchString(path) {
				files = append(files, path)
			}
		}

		return nil
	})

	return files, err
}

func printResult(line, filePath string, indeces [][]int, lineNum int) {
	fileName := fmt.Sprintf("%s%s%s", PURPLE, filePath, END)
	coloredLine := ""
	lastEnd := 0
	starts := []int{}

	for _, m := range indeces {
		start, end := m[0], m[1]
		coloredLine += line[lastEnd:start]
		coloredLine += fmt.Sprintf("%s%s%s", RED, line[start:end], END)
		lastEnd = end
		starts = append(starts, start)
	}

	coloredLine += line[lastEnd:]
	coloredLine = strings.TrimSpace(coloredLine)
	fmt.Printf("%s:%s:%s\t%s\n",
		fileName,
		fmt.Sprintf("%s%d%s", YELLOW, lineNum, END),
		fmt.Sprintf("%s%d%s", GREEN, starts, END),
		coloredLine,
	)
}

func searchInFile(filePath string, searchPattern *regexp.Regexp) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineIndex := 1
	for scanner.Scan() {
		line := scanner.Text()
		indeces := searchPattern.FindAllStringIndex(line, -1)
		if indeces != nil {
			printResult(line, filePath, indeces, lineIndex)
		}
		lineIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Search(path, filePattern, searchPattern string) {
	fileRegex := regexp.MustCompile(filePattern)
	searchRegex := regexp.MustCompile(searchPattern)

	files, err := collectPaths(path, fileRegex)
	if err != nil {
		log.Fatalf("error walking the path: %v", err)
	}
	for _, filePath := range files {
		searchInFile(filePath, searchRegex)
	}
}
