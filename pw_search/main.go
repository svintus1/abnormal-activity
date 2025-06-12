package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"script/pkg/system"
	"script/pkg/tags"
	"strings"
	"sync"
)

var (
	searchWord    = "password"
	searchDirs    = []string{"home", "etc", "root", "opt", "tmp", "mnt", "media"}
	searchedFiles = []string{}
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	rootFS, err := parseRootFS()
	if err != nil {
		fmt.Printf("%sParsing error: %s\n", tags.Err, err.Error())
		return
	}

	fmt.Printf("%sT1552.001 Credentials In Files\n", tags.Info)

	var wg sync.WaitGroup
	wg.Add(len(searchDirs))

	for _, dir := range searchDirs {
		dirPath := filepath.Join(rootFS, dir)
		go findWord(searchWord, dirPath, &wg)
	}

	wg.Wait()

	printHeadOfSeachedFiles()

	fmt.Printf("%sThe script executed successfully\n", tags.Info)
}

func parseRootFS() (string, error) {
	rootFS := flag.String("rootfs", "/host", "Path to the host file system root")
	flag.Parse()
	if isDir(*rootFS) {
		return *rootFS, nil
	}
	return "", errors.New("incorrect path to the host root directory is specified")
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Записывает подходящие файлы в глобальную переменную searchedFiles
func findWord(word string, searchRoot string, wg *sync.WaitGroup) {
	err := filepath.Walk(searchRoot, func(path string, info os.FileInfo, err error) error {
		if shouldSkipFile(info, err) {
			return nil
		}

		if isFileContainSearchWord(path, word) {
			searchedFiles = append(searchedFiles, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("%sError when executing the script: %s\n", tags.Log, err.Error())
	}

	defer wg.Done()
}

func shouldSkipFile(info os.FileInfo, err error) bool {
	if err != nil {
		return true
	}

	if info.IsDir() {
		return true
	}

	if !info.Mode().IsRegular() {
		return true
	}

	return false
}

func isFileContainSearchWord(filePath string, searchWord string) bool {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}

	if strings.Contains(string(content), searchWord) {
		return true
	}

	return false
}

func printHeadOfSeachedFiles() {
	if len(searchedFiles) > 10 {
		printTenValues(searchedFiles)
	} else {
		printAllValues(searchedFiles)
	}
}

func printTenValues(elements []string) {
	for i := 0; i < 10; i++ {
		fmt.Println(elements[i])
	}
}

func printAllValues(elements []string) {
	for i := range elements {
		fmt.Println(elements[i])
	}
}
