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
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1564.001 Hidden Files and Directories\n", tags.Info)
	err := os.Mkdir(hiddenDir, 0755)
	if err != nil {
		fmt.Printf("%sОшибка при создании директории %s: %s\n", tags.Err, hiddenDir, err.Error())
	} else {
		defer removeDirWithLog(hiddenDir)
	}
	filePath := filepath.Join(hiddenDir, "."+fileName)
	fmt.Printf("%sУспешно создана директория%s\n", tags.Info, hiddenDir)
	err = system.CopyFile(filepath.Join("assets", fileName), filePath, 0755)
	if err != nil {
		fmt.Printf("%sОшибка при копировании файла %s: %s\n", tags.Err, fileName, err.Error())
		return
	}
	fmt.Printf("%sУспешно сокопирован файл%s\n", tags.Info, fileName)
	out, err := exec.Command(filePath).Output()
	if err != nil {
		fmt.Printf("%sОшибка при выполнении файла %s: %s\n", tags.Err, filePath, err.Error())
		return
	}
	fmt.Print(string(out))
	fmt.Printf("%sФайл %s успешно выполнен\n", tags.Info, filePath)
}

func removeDirWithLog(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("%sОшибка при удалении директории %s: %s\n", tags.Log, path, err.Error())
	} else {
		fmt.Printf("%sДиректория %s рекурсивно уалена\n", tags.Log, path)
	}
}
