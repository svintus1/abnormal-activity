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
	fileName         = "myapp"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1036.004 Masquerade Task or Service\n", tags.Info)
	os.Mkdir(newDir, 0600)
	defer func() {
		os.RemoveAll(newDir)
		fmt.Printf("%sУдалена директория %s\n", tags.Log, newDir)
	}()
	startFakeServices()
}

func startFakeServices() {
	for _, fakeService := range fakeServiceNames {
		filePath := filepath.Join(newDir, fakeService)
		system.CopyFile(filepath.Join("assets", fileName), filePath, 0700)
		err := exec.Command(filePath).Run()
		if err != nil {
			fmt.Printf("%sОшибка при запуске %s: %s\n", tags.Log, filePath, err.Error())
		} else {
			fmt.Printf("%sУспешно запущен %s\n", tags.Log, filePath)
		}
	}
}
