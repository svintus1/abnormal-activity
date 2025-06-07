package main

import (
	"fmt"
	"os"
	"script/pkg/system"
	"script/pkg/tags"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1070.003 Clear Command History\n", tags.Info)
	addBashHistoryFile()
	removeBashHistory()
}

func addBashHistoryFile() {
	system.CopyFile("assets/.bash_history", "/root/.bash_history", 0600)
}

func removeBashHistory() {
	err := os.Remove("/root/.bash_history")
	if err != nil {
		fmt.Printf("%sОшибка при удалении ~/.bash_history: %s\n", tags.Err, err.Error())
	} else {
		fmt.Printf("%sУдален файл ~/.bash_history\n", tags.Info)
	}
}
