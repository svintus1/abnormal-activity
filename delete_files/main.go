package main

import (
	"fmt"
	"os"
	"script/pkg/files"
	"script/pkg/system"
	"script/pkg/tags"
)

var configPaths = []string{
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
		fmt.Printf("%sError when creating a temporary file: %s\n", tags.Err, err)
		return
	}
	defer deleteTempFile(tempFile)

	err = writeSensitiveContent(tempFile.Name())
	if err != nil {
		fmt.Printf("%sError when writing to a file: %s\n", tags.Err, err)
		return
	}
	fmt.Printf("%sThe contents of the configurations are written to a temporary file %s\n", tags.Info, tempFile.Name())
}

func writeSensitiveContent(destinationFile string) error {
	sensitiveContent := consolidateAllFiles(configPaths)
	targetFile := files.NewFile(destinationFile)
	return targetFile.WriteFileLines(sensitiveContent, 0600)
}

func consolidateAllFiles(files []string) []string {
	var result []string
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("%sError reading file %s: %s\n", tags.Log, file, err.Error())
		}
		result = append(result, string(content))
	}
	return result
}

func deleteTempFile(tempFile *os.File) {
	tempFile.Close()
	err := os.Remove(tempFile.Name())
	if err != nil {
		fmt.Printf("%sError when deleting a file %s\n", tags.Log, tempFile.Name())
	} else {
		fmt.Printf("%sTemporary file %s deleted\n", tags.Log, tempFile.Name())
	}
}
