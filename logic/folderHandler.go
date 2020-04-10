package logic

import (
	"fmt"
	"os"

	"github.com/brunocalmon/echo-initializr/data"
	"github.com/brunocalmon/echo-initializr/global"
)

func CreateFolders() {
	files, ok := global.Global["files"].(data.Files)
	if ok {
		for _, file := range files {
			createFolder(file)
		}
	} else {
		fmt.Println("Could not access or convert sample files interface!")
		panic("stopped")
	}
}

func createFolder(file data.File) error {
	path := file.Path()
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create the directory %v\n", err)
		panic("Could not create the directory: " + path)
	}

	fmt.Println("Folder " + path + " created.")

	return nil
}
