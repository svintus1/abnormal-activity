package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"script/pkg/files"
	"script/pkg/system"
	"script/pkg/tags"
	"strings"
	"sync"
)

var (
	searchString = "password"
	dirs         = []string{"/home", "/etc", "/root", "/opt", "/tmp", "/mnt", "/media"}
)

func main() {
	prefixPath := flag.String("src", "/host", "Путь к корню файловой системы хоста")
	flag.Parse()
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(dirs))

	for _, d := range dirs {
		dir := d
		go findPw(filepath.Join(*prefixPath, dir), &wg)
	}
	wg.Wait()
}

func findPw(dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		file := files.NewFile(path)
		content, err := file.ReadFileLines()
		if err != nil {
			return nil
		}

		for _, line := range content {
			if strings.Contains(line, searchString) {
				fmt.Printf("%s%s\n", tags.Info, path)
				break
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("%sError when executing the script: %s\n", tags.Err, err.Error())
	}
}
