package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"script/pkg/system"
	"script/pkg/tags"
)

var (
	fileName   = "myapp"
	dirPath    = "/tmp"
	extensions = []string{".png", ".txt", ".com", ".scr", ".cmd"}
)

func main() {
	filePath := filepath.Join("assets", fileName)
	newFilePath := filepath.Join(dirPath, fileName)

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1036.008  Masquerade File Type\n", tags.Info)

	for _, extension := range extensions {
		newFileName := newFilePath + extension
		copyFile(filePath, newFileName)
		execCommand(newFileName)
		removeFile(newFileName)
	}
}

func copyFile(filePath string, fileName string) {
	err := system.CopyFile(filePath, fileName, 0755)
	if err != nil {
		fmt.Printf("%sError while copying a %s: %s\n", tags.Log, filePath, err.Error())
	} else {
		fmt.Printf("%sSuccessfully copied %s\n", tags.Log, fileName)
	}
}

func execCommand(fileName string) {
	err := exec.Command(fileName).Run()
	if err != nil {
		fmt.Printf("%sError while executing a %s: %s\n", tags.Log, fileName, err.Error())
	} else {
		fmt.Printf("%sSuccessfully executed %s\n", tags.Log, fileName)
	}
}

func removeFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("%sError when deleting a %s: %s\n", tags.Log, fileName, err.Error())
	} else {
		fmt.Printf("%sSuccessfully removed %s\n", tags.Log, fileName)
	}
}
