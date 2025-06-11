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
	fileName    = "myapp"
	tempDirPath = "/tmp/bin"
	fakeCmd     = "su"
)

func main() {
	filePath := filepath.Join("assets", fileName)

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1574.007 Path Interception by PATH Environment Variable\n", tags.Info)

	createDir(tempDirPath)
	defer deleteDir(tempDirPath)

	fakeCmdPath := filepath.Join(tempDirPath, fakeCmd)
	system.UpdatePath([]string{tempDirPath})

	err := system.CopyFile(filePath, fakeCmdPath, 0755)
	if err != nil {
		fmt.Printf("%sError when copying a file %s: %s\n", tags.Err, fileName, err.Error())
		return
	}
	fmt.Printf("%sFile %s has been successfully copied to %s\n", tags.Info, fileName, fakeCmdPath)

	out, err := exec.Command(fakeCmd).Output()
	if err != nil {
		fmt.Printf("%sError while executing the %s command: %s\n", tags.Err, fakeCmd, err.Error())
		return
	}
	fmt.Print(string(out))
}

func createDir(tempDirPath string) {
	err := os.Mkdir(tempDirPath, 0755)
	if err != nil {
		fmt.Printf("%sError when creating a directory %s:%s\n", tags.Err, tempDirPath, err.Error())
	}
}

func deleteDir(tempDirPath string) {
	err := os.RemoveAll(tempDirPath)
	if err != nil {
		fmt.Printf("%sError when deleting a %s: %s\n", tags.Log, tempDirPath, err.Error())
	} else {
		fmt.Printf("%sFile %s successfully deleted\n", tags.Log, tempDirPath)
	}
}
