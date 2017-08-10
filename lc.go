package main

import (
	"os"
	"io"
	"log"
)

var logger = log.New(os.Stdout, "", 0)

func countFileLines (path string) int {
	f, err := os.Open(path)
	if err != nil {
		logger.Printf("%v: error opening the file: %v", path, err)
		return -1
	}
	defer f.Close()
	lineCount := 0
	b := make([]byte, 8)
	for {
		readNow, err := f.Read(b)
		if err == io.EOF {
			break
		}
		for i := 0; i < readNow; i++ {
			if b[i] == '\n' {
				lineCount++
			}
		}
	}
	return lineCount
}

func main() {
	relatedArgs := os.Args[1:]
	if len(relatedArgs) == 0 {
		logger.Fatalln("Wrong usage. Pass file paths or URLs to GitHub repos as arguments like so:\n" +
			"lc path1 [path2] ... [pathN] ")
	}

	sum := 0
	for _, path := range relatedArgs {
		stat, err := os.Stat(path)
		if err != nil {
			logger.Printf("Error getting info about %v", path)
			continue
		}
		if stat.Mode().IsDir() {
			logger.Printf("%v: Is a directory", path)
			logger.Printf("%6d\t%v", 0, path)
			continue
		}
		if res := countFileLines(path); res > 0 {
			sum += res
			logger.Printf("%6d\t%v", res, path)
		}
	}
	if len(relatedArgs) > 0 && sum != 0 {
		logger.Printf("%6d\ttotal", sum)
	}
}
