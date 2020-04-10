package logic

import (
	"fmt"
	"os"
	"strings"

	"github.com/brunocalmon/echo-initializr/data"
	"github.com/brunocalmon/echo-initializr/global"
)

func CreateFiles() {

	files, ok := global.Global["files"].(data.Files)

	if ok {
		for _, file := range files {
			createFile(file)
		}
	} else {
		fmt.Println("Could not access or convert sample files interface!")
		panic("stopped")
	}
}

func createFile(file data.File) {
	f, err := os.Create(file.PathWithFile())
	if err != nil {
		fmt.Println(err)
		return
	}

	namespace, okn := global.Global["namespace"].(string)
	version, okv := global.Global["version"].(string)
	port, okp := global.Global["port"].(int)
	if okn && okv && okp {
		if file.Name == "environment.go" {
			file.Content = fmt.Sprintf(file.Content, port)
		}
		if file.Name == "README.md" {
			file.Content = fmt.Sprintf(file.Content, namespace)
		}
		if file.Name == "main.go" {
			file.Content = fmt.Sprintf(file.Content, namespace, namespace)
		}
		if file.Name == "go.mod" {
			index := strings.LastIndex(namespace, "/")
			moduleName := ""
			if index != -1 {
				moduleName = namespace[index+1:]
			} else {
				moduleName = namespace
			}

			file.Content = fmt.Sprintf(file.Content, moduleName, version)
		}
	} else {
		fmt.Println("Could not access or convert sample files interface!")
		panic("stopped")
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
