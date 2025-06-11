package main

import (
	"fmt"
	"os"
	"script/pkg/system"
	"script/pkg/tags"
)

var (
	newBashHistoryFile  = "assets/.bash_history"
	rootBashHistoryFile = "/root/.bash_history"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1070.003 Clear Command History\n", tags.Info)

	err := addBashHistoryFile()
	if err != nil {
		fmt.Printf("%sError when coping %s: %s\n", tags.Err, newBashHistoryFile, err.Error())
		return
	}

	err = deleteBashHistory()
	if err != nil {
		fmt.Printf("%sError when deleting %s: %s\n", tags.Err, rootBashHistoryFile, err.Error())
		return
	}

	fmt.Printf("%sScript executed successfully\n", tags.Info)
}

func addBashHistoryFile() error {
	err := system.CopyFile(newBashHistoryFile, rootBashHistoryFile, 0600)
	if err != nil {
		return err
	}
	fmt.Printf("%sCopied the file %s to %s\n", tags.Log, newBashHistoryFile, rootBashHistoryFile)
	return nil
}

func deleteBashHistory() error {
	err := os.Remove(rootBashHistoryFile)
	if err != nil {
		return err
	}
	fmt.Printf("%sDeleted the file %s\n", tags.Log, rootBashHistoryFile)
	return nil
}
