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
	fileName = "myapp"
	newDir   = "/tmp/bin"
	fakeCmd  = "su"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1574.007 Path Interception by PATH Environment Variable\n", tags.Info)
	err := os.Mkdir(newDir, 0755)
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
	}
	fakeCmdPath := filepath.Join(newDir, fakeCmd)
	defer func() {
		err := os.RemoveAll(newDir)
		if err != nil {
			fmt.Printf("%sОшибка при удалении %s: %s\n", tags.Log, fakeCmdPath, err.Error())
		} else {
			fmt.Printf("%sФайл %s успешно удален", tags.Log, fakeCmdPath)
		}
	}()
	system.UpdatePath([]string{newDir})
	err = system.CopyFile(filepath.Join("assets", fileName), fakeCmdPath, 0755)
	if err != nil {
		fmt.Printf("%sОшибка при копировании файла %s: %s\n", tags.Err, fileName, err.Error())
		return
	}
	fmt.Printf("%sФайл %s успешно скопирован в %s", tags.Info, fileName, fakeCmdPath)
	out, err := exec.Command("su").Output()
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
		return
	}
	fmt.Print(string(out))
}
