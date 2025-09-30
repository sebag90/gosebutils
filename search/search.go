package search

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
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

		info, _ := os.Stat(path)
		if !info.IsDir() {
			if pattern.MatchString(path) {
				files = append(files, path)
			}
		}
		return nil
	})

	return files, err
}

func printResult(line string, indeces [][]int, lineNum, windowSize int) {
	for _, m := range indeces {
		start, end := m[0], m[1]
		leftMargin := line[max(0, start-windowSize):start]
		rightMargin := line[end:min(len(line), end+windowSize)]
		coloredWord := fmt.Sprintf("%s%s%s", RED, line[start:end], END)
		linetoDisplay := fmt.Sprintf("%s%s%s", leftMargin, coloredWord, rightMargin)

		fmt.Printf("\t%s:%s\t%s\n",
			fmt.Sprintf("%s%d%s", YELLOW, lineNum, END),
			fmt.Sprintf("%s%d%s", GREEN, start, END),
			strings.TrimSpace(linetoDisplay),
		)
	}

}

func searchInFile(filePath string, searchPattern *regexp.Regexp, windowSize int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("error opening file %s: %v", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024 * 100 // 1MB (adjust as needed)
	buf := make([]byte, 0, 64*1024)       // initial buffer size
	scanner.Buffer(buf, maxCapacity)

	lineIndex := 1
	filenamePrinted := false
	for scanner.Scan() {
		line := scanner.Text()
		indeces := searchPattern.FindAllStringIndex(line, -1)
		if indeces != nil {
			if !filenamePrinted {
				fmt.Printf("%s%s%s\n", PURPLE, filePath, END)
				filenamePrinted = true
			}
			printResult(line, indeces, lineIndex, windowSize)
		}
		lineIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Search(path, filePattern, searchPattern string, windowSize int) {
	fileRegex := regexp.MustCompile(filePattern)
	searchRegex := regexp.MustCompile(searchPattern)

	files, err := collectPaths(path, fileRegex)
	if err != nil {
		log.Printf("error walking the path: %v", err)
	}
	numWorkers := runtime.NumCPU() * 2
	runtime.GOMAXPROCS(numWorkers)

	var wg sync.WaitGroup
	jobs := make(chan string, numWorkers*2)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for filePath := range jobs {
				searchInFile(filePath, searchRegex, windowSize)
			}
		}()
	}

	for _, filePath := range files {
		jobs <- filePath
	}

	close(jobs)
	wg.Wait()
}
