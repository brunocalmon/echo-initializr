package logic

import (
	"fmt"
	"os"
	"strings"

	"github.com/brunocalmon/echo-initializr/data"
)

func CreateFiles(namespace, version, outputDir string, port int, files data.Files) {
	for _, file := range files {
		createFile(namespace, version, outputDir, port, file)
	}
}

func createFile(namespace, version, outputDir string, port int, file data.File) {
	f, err := os.Create(file.PathWithFile(outputDir))
	if err != nil {
		fmt.Println(err)
		return
	}
	if file.Name == "environment.go" {
		file.Content = fmt.Sprintf(file.Content, port)
	}
	if file.Name == "README.md" {
		file.Content = fmt.Sprintf(file.Content, namespace)
	}
	if file.Name == "main.go" {
		index := strings.LastIndex(namespace, "/")
		minimalistName := ""
		if index != -1 {
			minimalistName = namespace[index+1:]
		} else {
			minimalistName = namespace
		}

		file.Content = fmt.Sprintf(file.Content, namespace, namespace, minimalistName)
	}
	if file.Name == "go.mod" {
		file.Content = fmt.Sprintf(file.Content, namespace, version)
	}

	l, err := f.WriteString(file.Content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
