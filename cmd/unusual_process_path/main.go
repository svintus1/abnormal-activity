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
	path     = "/var/log"
	fileName = "myapp"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sProcesses in Unusual Paths\n", tags.Info)
	err := system.CopyFile(filepath.Join("assets", fileName), path, 0755)
	if err != nil {
		fmt.Printf("%sError when copying a binary file\n", tags.Err)
		return
	}
	absPath := filepath.Join(path, fileName)
	defer func() {
		os.Remove(absPath)
		fmt.Printf("%sFile %s removed\n", tags.Log, absPath)
	}()
	cmd := exec.Command(absPath)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%sError when executing a binary file\n", tags.Err)
		return
	}
	fmt.Printf("%sFile %s successfully launched in %s\n", tags.Info, fileName, absPath)
}
