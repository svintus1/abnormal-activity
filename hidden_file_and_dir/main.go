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
	fileName  = "myapp"
	hiddenDir = "/tmp/.hidden"
)

func main() {
	filePath := getFilePath()
	hiddenFilePath := getHiddenFilePath()

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1564.001 Hidden Files and Directories\n", tags.Info)

	createDirectory()
	defer deleteDirectory(hiddenDir)

	err := system.CopyFile(filePath, hiddenFilePath, 0755)
	if err != nil {
		fmt.Printf("%sError when copying a file %s: %s\n", tags.Err, fileName, err.Error())
		return
	}
	fmt.Printf("%sFile successfully copied %s\n", tags.Info, fileName)

	out, err := exec.Command(hiddenFilePath).Output()
	if err != nil {
		fmt.Printf("%sError during file execution %s: %s\n", tags.Err, hiddenFilePath, err.Error())
		return
	}

	fmt.Print(string(out))
	fmt.Printf("%sFile %s successfully executed\n", tags.Info, hiddenFilePath)
}

func getHiddenFilePath() string {
	return filepath.Join(hiddenDir, "."+fileName)
}

func getFilePath() string {
	return filepath.Join("assets", fileName)
}

func createDirectory() {
	err := os.Mkdir(hiddenDir, 0755)
	if err != nil {
		fmt.Printf("%sError when creating a directory %s: %s\n", tags.Log, hiddenDir, err.Error())
	} else {
		fmt.Printf("%sA directory has been created %s\n", tags.Log, hiddenDir)
	}
}

func deleteDirectory(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("%sError when deleting a directory %s: %s\n", tags.Log, path, err.Error())
	} else {
		fmt.Printf("%sDirectory %s recursively deleted\n", tags.Log, path)
	}
}
