package util

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

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
