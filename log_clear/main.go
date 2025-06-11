package main

import (
	"flag"
	"fmt"
	"os"
	"script/pkg/system"
	"script/pkg/tags"
)

var (
	logDirs = []string{
		"/var/log",
		"/run/log",
	}
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1070.002 Clear Linux or Mac System Logs\n", tags.Info)

	force := parseForceFlag()
	if !confirmAction(force) {
		return
	}

	deleteLogDirs()

	fmt.Printf("%sThe script has completed its work\n", tags.Info)
}

func parseForceFlag() bool {
	forceFlag := flag.String("force", "", "Force delete")
	flag.Parse()
	return *forceFlag == "yes"
}

func confirmAction(force bool) bool {
	if force {
		return true
	}

	fmt.Println("This script will recursively delete the directories /var/log and /run/log")
	fmt.Print("If you agree with this you will write \"yes\": ")

	var userInput string
	fmt.Scan(&userInput)

	return userInput == "yes"
}

func deleteLogDirs() {
	for _, dir := range logDirs {
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Printf("%sError when deleting %s: %s\n", tags.Log, dir, err.Error())
		} else {
			fmt.Printf("%sDeleted directory %s\n", tags.Log, dir)
		}
	}
}
