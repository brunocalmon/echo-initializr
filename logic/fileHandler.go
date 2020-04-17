package logic

import (
	"fmt"
	"os"
	"strings"

	"github.com/brunocalmon/echo-initializr/data"
)

//CreateFiles creates files on folders
func CreateFiles(namespace, version, outputDir string, port int, files data.Files) {
	for _, file := range files {
		createFile(namespace, version, outputDir, port, file)
	}
}

func createFile(namespace, version, outputDir string, port int, file data.File) {
	path := file.Path(outputDir)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create the directory %v\n", err)
		panic("Could not create the directory: " + path)
	}
	fmt.Println("Folder " + path + " created.")

	createdFile, err := os.Create(file.PathWithFile(outputDir))
	if err != nil {
		fmt.Println(err)
		return
	}

	file = contentPicker(namespace, version, port, file)

	l, err := createdFile.WriteString(file.Content)
	if err != nil {
		fmt.Println(err)
		createdFile.Close()
		return
	}

	fmt.Println(l, "written successfully: "+outputDir)
	err = createdFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func contentPicker(namespace, version string, port int, file data.File) data.File {
	if file.Name == "environment.go" {
		index := strings.LastIndex(namespace, "/")
		minimalistName := ""
		if index != -1 {
			minimalistName = namespace[index+1:]
		} else {
			minimalistName = namespace
		}

		file.Content = fmt.Sprintf(file.Content, port, minimalistName, minimalistName)
	}
	if file.Name == "README.md" {
		file.Content = fmt.Sprintf(file.Content, namespace)
	}
	if file.Name == "main.go" {
		file.Content = fmt.Sprintf(file.Content, namespace, namespace)
	}
	if file.Name == "go.mod" {
		file.Content = fmt.Sprintf(file.Content, namespace, version)
	}
	return file
}
