package rpm

import (
	"errors"
	"fmt"
	"os/exec"
	"script/pkg/tags"
	"strings"
)

type Dnf struct{}

func (*Dnf) Update() error {
	cmd := exec.Command("dnf", "makecache")
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при выполнении dnf makecache: " + err.Error())
	}
	fmt.Printf("%sОбновлен список пакетов (dnf makecache)\n", tags.Log)
	return nil
}

func (*Dnf) InstallPackage(pkg string) error {
	cmd := exec.Command("dnf", "install", "-y", pkg)
	err := cmd.Run()
	if err != nil {
		return errors.New("ошибка при устанвоке пакета " + pkg + ": " + err.Error())
	}
	fmt.Printf("%sПакет %s успешно установлен\n", tags.Log, pkg)
	return nil
}

func (*Yum) IsPackageInstalled(pkg string) bool {
	cmd := exec.Command("rpm", "-q", pkg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%sПакет %s не установлен: %s\n", tags.Log, pkg, strings.TrimSpace(string(output)))
		return false
	}
	fmt.Printf("%sПакет %s установлен: %s\n", tags.Log, pkg, strings.TrimSpace(string(output)))
	return true
}
