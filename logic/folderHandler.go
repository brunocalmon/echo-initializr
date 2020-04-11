package logic

import (
	"fmt"
	"os"

	"github.com/brunocalmon/echo-initializr/data"
)

func CreateFolders(outputDir string, files data.Files) {
	for _, file := range files {
		createFolder(outputDir, file)
	}
}

func createFolder(outputDir string, file data.File) error {
	path := file.Path(outputDir)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create the directory %v\n", err)
		panic("Could not create the directory: " + path)
	}

	fmt.Println("Folder " + path + " created.")

	return nil
}
