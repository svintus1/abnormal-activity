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
	unusualPath = "/dev/shm"
	fileName    = "myapp"
)

func main() {
	filePath := filepath.Join("assets", fileName)
	newFilePath := filepath.Join(unusualPath, fileName)

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sProcesses in Unusual Paths\n", tags.Info)

	err := system.CopyFile(filePath, newFilePath, 0755)
	if err != nil {
		fmt.Printf("%sError when copying a binary file\n", tags.Err)
		return
	}
	fmt.Printf("%sFile %s successfully copied to %s\n", tags.Info, filePath, newFilePath)
	defer deleteFile(newFilePath)

	err = exec.Command(newFilePath).Run()
	if err != nil {
		fmt.Printf("%sError when executing a binary file %s: %s\n", tags.Err, newFilePath, err.Error())
		return
	}
	fmt.Printf("%sFile %s successfully launched\n", tags.Info, newFilePath)
}

func deleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("%sError when deleting file %s: %s\n", tags.Log, filePath, err.Error())
	} else {
		fmt.Printf("%sFile %s removed\n", tags.Log, filePath)
	}
}
