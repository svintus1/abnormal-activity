package files

import (
	"errors"
	"os"
	"strings"
)

type File struct {
	Path string
}

func NewFile(path string) *File {
	file, _ := os.Create(path)
	defer file.Close()
	return &File{
		Path: path,
	}
}

func (file *File) ReadFileLines() ([]string, error) {
	content, err := os.ReadFile(file.Path)
	if err != nil {
		return nil, errors.New("ошибка при чтении файла " + file.Path + ": " + err.Error())
	}
	return strings.Split(string(content), "\n"), nil
}

func (file *File) WriteFileLines(content []string, perm os.FileMode) error {
	data := []byte(strings.Join(content, "\n") + "\n")
	tmpPath := file.Path + ".tmp"
	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return errors.New("ошибка при открытии файла " + tmpPath + ": " + err.Error())
	}
	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		return errors.New("ошибка при записи в файл " + tmpPath + ": " + err.Error())
	}
	tmpFile.Sync()
	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpPath)
		return errors.New("ошибка при закрытии файла " + tmpPath + ": " + err.Error())
	}
	return os.Rename(tmpPath, file.Path)
}
