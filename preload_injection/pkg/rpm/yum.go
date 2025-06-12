package rpm

import (
	"errors"
	"fmt"
	"os/exec"
	"script/pkg/tags"
	"strings"
)

type Yum struct{}

func (*Yum) Update() error {
	cmd := exec.Command("yum", "check-update")
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при выполнении yum check-update: " + err.Error())
	}
	fmt.Printf("%sОбновлен список пакетов (yum check-update)\n", tags.Log)
	return nil
}

func (*Yum) InstallPackage(pkg string) error {
	cmd := exec.Command("yum", "install", "-y", pkg)
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при устанвоке пакета " + pkg + ": " + err.Error())
	}
	fmt.Printf("%sПакет %s успешно установлен\n", tags.Log, pkg)
	return nil
}

func (*Dnf) IsPackageInstalled(pkg string) bool {
	cmd := exec.Command("rpm", "-q", pkg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%sПакет %s не установлен: %s\n", tags.Log, pkg, strings.TrimSpace(string(output)))
		return false
	}
	fmt.Printf("%sПакет %s установлен: %s\n", tags.Log, pkg, strings.TrimSpace(string(output)))
	return true
}
