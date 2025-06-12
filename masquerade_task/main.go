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
	fakeServiceNames = []string{"sshd", "journald", "systemd", "dockerd"}
	newDir           = "/tmp/.hidden"
	filePath         = "assets/myapp"
)

func main() {

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1036.004 Masquerade Task or Service\n", tags.Info)

	err := createDir(newDir)
	if err == nil {
		defer deleteDir(newDir)
	}

	startFakeServices(fakeServiceNames)
}

func createDir(newDir string) error {
	err := os.Mkdir(newDir, 0600)
	if err != nil {
		fmt.Printf("%sError when creating a directory %s: %s\n", tags.Log, newDir, err.Error())
		return err
	} else {
		fmt.Printf("%sDirectory %s created\n", tags.Log, newDir)
	}
	return nil
}

func deleteDir(newDir string) {
	err := os.RemoveAll(newDir)
	if err != nil {
		fmt.Printf("%sError when deleting the %s directory: %s\n", tags.Log, newDir, err.Error())
	} else {
		fmt.Printf("%sDirectory %s deleted\n", tags.Log, newDir)
	}
}

func startFakeServices(fakeServiceNames []string) {
	for _, fakeService := range fakeServiceNames {
		newFilePath := filepath.Join(newDir, fakeService)
		system.CopyFile(filePath, newFilePath, 0700)
		startService(newFilePath)
	}
}

func startService(newFilePath string) {
	err := exec.Command(newFilePath).Run()
	if err != nil {
		fmt.Printf("%sError when starting %s: %s\n", tags.Log, newFilePath, err.Error())
	} else {
		fmt.Printf("%sSuccessfully started %s\n", tags.Log, newFilePath)
	}
}
