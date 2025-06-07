package deb

import (
	"errors"
	"fmt"
	"os/exec"
	"script/pkg/tags"
	"strings"
)

type Apt struct{}

func (*Apt) Update() error {
	cmd := exec.Command("apt", "update")
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при выполнении apt update: " + err.Error())
	}
	fmt.Printf("%sОбновлен список пакетов (apt update)\n", tags.Log)
	return nil
}

func (*Apt) InstallPackage(pkg string) error {
	cmd := exec.Command("apt", "install", "-y", pkg)
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при устанвоке пакета " + pkg + ": " + err.Error())
	}
	fmt.Printf("%sПакет %s успешно установлен\n", tags.Log, pkg)
	return nil
}

func (*Apt) IsPackageInstalled(pkg string) bool {
	cmd := exec.Command("dpkg", "-s", pkg)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("%sНе обнаружен пакет(ошибка при поиске) %s:%s\n", tags.Log, pkg, err.Error())
		return false
	}
	if strings.Contains(string(output), "Status: install ok installed") {
		fmt.Printf("%sОбнаружен пакет %s\n", tags.Log, pkg)
		return true
	}
	fmt.Printf("%sНе обнаружен пакет %s\n", tags.Log, pkg)
	return false
}
