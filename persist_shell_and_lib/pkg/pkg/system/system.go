package system

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"script/pkg/files"
	"script/pkg/tags"
	"script/pkg/utils"
	"strings"
	"syscall"
)

type PkgManagerInterface interface {
	Update() error
	InstallPackage(string) error
	IsPackageInstalled(string) bool
}

func IsRoot() bool {
	return syscall.Getuid() == 0
}

func UpdatePath() {
	currentPath := os.Getenv("PATH")
	newPath := currentPath + ":/sbin:/usr/sbin"
	os.Setenv("PATH", newPath)
}

func ServiceRestart(service string) error {
	cmd := exec.Command("systemctl", "restart", service)
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при перезапуске сервиса " + service + ": " + err.Error())
	}
	fmt.Printf("%sПерезапущен сервис %s\n", tags.Log, service)
	return nil
}

func GetUidMin() (string, error) {
	loginDefs := files.NewFile("/etc/login.defs")
	data, err := loginDefs.ReadFileLines()
	if err != nil {
		return "", err
	}
	for _, line := range data {
		if strings.Contains(line, "UID_MIN") {
			parts := strings.Fields(line)
			return parts[1], nil
		}
	}
	return "", errors.New("не получилось узнать uid_min")
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "1.1.1.1:53")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func EditFileHosts() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	outboundIP := GetOutboundIP()
	fileHosts := files.NewFile("/etc/hosts")
	lines, err := fileHosts.ReadFileLines()
	if err != nil {
		return err
	}
	var newLines []string
	correctLine := outboundIP + "\t" + hostname
	for _, line := range lines {
		if strings.Contains(line, hostname) && !strings.HasPrefix(line, "#") && line != correctLine {
			line = "# " + line
		}
		newLines = append(newLines, line)
	}
	if !utils.ContainsInSlice(newLines, correctLine) {
		newLines = append(newLines, correctLine)
		fmt.Printf("%sВ /etc/hosts добавлена строка %s\n", tags.Log, correctLine)
	}
	err = fileHosts.WriteFileLines(newLines, 0640)
	if err != nil {
		return err
	}
	return nil
}

func IsCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func DetectPackageManager() (string, error) {
	if IsCommandAvailable("apt") {
		return "apt", nil
	} else if IsCommandAvailable("yum") {
		return "yum", nil
	} else if IsCommandAvailable("dnf") {
		return "dnf", nil
	}
	return "", errors.New("не удлалось определить пакетный менеджер")
}

func AddLineJournaldConf(path string, line string) error {
	journaldConf := files.NewFile(path)
	content, err := journaldConf.ReadFileLines()
	if err != nil {
		return errors.New("ошибка при чтении файла " + path + ": " + err.Error())
	}
	if !utils.ContainsInSlice(content, line) {
		content = append(content, line)
	}
	err = journaldConf.WriteFileLines(content, 0640)
	if err != nil {
		return errors.New("ошибка при записи файла " + path + ": " + err.Error())
	}
	return nil
}

func EnableServiceMaualRestart(serviceDir string, confName string) error {
	err := exec.Command("mkdir", "-p", serviceDir).Run()
	if err != nil {
		return err
	}
	os.WriteFile(filepath.Join(serviceDir, confName), []byte(`
[Unit]
RefuseManualStop=no
RefuseManualStart=no
`), 0644)

	err = exec.Command("systemctl", "daemon-reexec").Run()
	if err != nil {
		return err
	}
	err = exec.Command("systemctl", "daemon-reload").Run()
	if err != nil {
		return err
	}
	return nil
}

func IsServiceActive(service string) bool {
	cmd := exec.Command("systemctl", "is-active", service)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	status := strings.TrimSpace(string(output))
	return status == "active"
}

func InstallIfNotInstalled(pkg string, pkgManager PkgManagerInterface) error {
	if !pkgManager.IsPackageInstalled(pkg) {
		err := pkgManager.InstallPackage(pkg)
		if err != nil {
			return fmt.Errorf("ошибка при уставновке пакета %s: %s", pkg, err.Error())
		}
	}
	return nil
}

func CopyFile(path string, newPath string, perm os.FileMode) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла %s: %s", path, err.Error())
	}
	err = os.WriteFile(newPath, content, perm)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл %s: %s", path, err.Error())
	}
	return nil
}
