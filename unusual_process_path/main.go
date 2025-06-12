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
	unusualPath = "/var/log"
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
	defer deleteFile(newFilePath)

	err = launchExecFile(newFilePath)
	if err != nil {
		fmt.Printf("%sError when executing a binary file %s: %s\n", tags.Err, filePath, err.Error())
		return
	}
	fmt.Printf("%sFile %s successfully launched\n", tags.Log, filePath)
}

func deleteFile(filePath string) {
	os.Remove(filePath)
	fmt.Printf("%sFile %s removed\n", tags.Log, filePath)
}

func launchExecFile(filePath string) error {
	err := exec.Command(filePath).Run()
	if err != nil {
		return fmt.Errorf("error when executing file %s: %s", filePath, err.Error())
	}
	return nil
}
