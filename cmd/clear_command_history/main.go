package main

import (
	"fmt"
	"os"
	"os/exec"
	"script/pkg/system"
	"script/pkg/tags"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1070.003 Clear Command History\n", tags.Info)
	clearCommandHistory()
	removeBashHistroy()
}

func clearCommandHistory() {
	err := exec.Command("history", "-c").Run()
	if err != nil {
		fmt.Printf("%sОшибка при очистке истории команд: %s\n", tags.Err, err.Error())
	} else {
		fmt.Printf("%sОчищена история команд\n", tags.Info)
	}
	err = exec.Command("history", "-w").Run()
	if err != nil {
		fmt.Printf("%sОшибка при перезаписи .bash_history: %s\n", tags.Err, err.Error())
	} else {
		fmt.Printf("%sПерезаписыан файл .bash_history\n", tags.Info)
	}
}

func removeBashHistroy() {
	err := os.Remove("~/.bash_history")
	if err != nil {
		fmt.Printf("%sОшибка при удалении ~/.bash_history: %s\n", tags.Err, err.Error())
	} else {
		fmt.Printf("%sУдален файл ~/.bash_history\n", tags.Info)
	}
}
