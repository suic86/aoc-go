package util

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFileToStringArray(pathFromCaller string) []string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}
	absPath := path.Join(path.Dir(filename), pathFromCaller)
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner:= bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func ReadFile(pathFromCaller string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not find Caller of util.ReadFile")
	}
	absPath := path.Join(path.Dir(filename), pathFromCaller)
	content, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}

func Dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}
