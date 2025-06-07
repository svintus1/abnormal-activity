package main

import (
	"flag"
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
	fmt.Printf("%sT1070.002 Clear Linux or Mac System Logs\n", tags.Info)
	forceFlag := flag.String("force", "", "Удалить все без уточнения")
	flag.Parse()
	force := *forceFlag == "yes"
	fmt.Println("Данный скрипт рекурсивно удалит директории /var/log и /run/log")
	fmt.Print("Если вы согласны с этим, что напишите yes: ")
	var userInput string
	fmt.Scan(&userInput)
	if !force {
		if userInput != "yes" {
			return
		}
	}
	err := os.RemoveAll("/var/log")
	if err != nil {
		fmt.Printf("%sошибка при удалении /var/log: %s\n", tags.Err, err.Error())
	}
	err = os.RemoveAll("/run/log")
	if err != nil {
		fmt.Printf("%sошибка при удалении /run/log: %s\n", tags.Err, err.Error())
	}
	fmt.Printf("%sСкрипт завершил свою работу\n", tags.Info)
}
