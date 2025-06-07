package main

import (
	"fmt"
	"os"
	"script/pkg/files"
	"script/pkg/system"
	"script/pkg/tags"
)

var configs = []string{
	"/etc/passwd",
	"/etc/shadow",
}

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1070.004 File Deletion\n", tags.Info)

	tempFile, err := os.CreateTemp("", "exfil-*")
	if err != nil {
		fmt.Printf("%sОшибка при создании временного файла: %s\n", tags.Err, err)
		return
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
		fmt.Printf("%sВременный файл %s удалён\n", tags.Log, tempFile.Name())
	}()

	err = copySensitiveContent(tempFile.Name())
	if err != nil {
		fmt.Printf("%sОшибка при записи в файл: %s\n", tags.Err, err)
		return
	}

	fmt.Printf("%sСодержимое конфигураций записано во временный файл %s\n", tags.Info, tempFile.Name())
}

func copySensitiveContent(destPath string) error {
	var allLines []string
	for _, confPath := range configs {
		confFile := files.NewFile(confPath)
		lines, err := confFile.ReadFileLines()
		if err != nil {
			return fmt.Errorf("не удалось прочитать %s: %w", confPath, err)
		}
		allLines = append(allLines, lines...)
	}

	destFile := files.NewFile(destPath)
	return destFile.WriteFileLines(allLines, 0600)
}
