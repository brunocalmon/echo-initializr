package logic

import (
	"fmt"
	"os"
	"strings"

	"github.com/brunocalmon/echo-initializr/data"
)

//CreateFiles creates files on folders
func CreateFiles(structureData map[string]string, files data.Files) {
	for _, file := range files {
		createFile(structureData, file)
	}
}

func createFile(structureData map[string]string, file data.File) {
	path := file.Path(structureData["outputDir"])
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Could not create the directory %v\n", err)
		panic("Could not create the directory: " + path)
	}
	fmt.Println("Folder " + path + " created.")

	createdFile, err := os.Create(file.PathWithFile(structureData["outputDir"]))
	if err != nil {
		fmt.Println(err)
		return
	}

	file = contentFormater(structureData, file)

	l, err := createdFile.WriteString(file.Content)
	if err != nil {
		fmt.Println(err)
		createdFile.Close()
		return
	}

	fmt.Println(l, "written successfully: "+structureData["outputDir"])
	err = createdFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func contentFormater(structureData map[string]string, file data.File) data.File {
	structureData["application_context"] = extractApplicationName(structureData)
	structureData["application_name"] = extractApplicationName(structureData)

	for _, keyword := range file.Keys {
		file.Content = strings.Replace(file.Content, "{{"+keyword+"}}", structureData[keyword], -1)
	}

	return file
}

func extractApplicationName(structureData map[string]string) (applicationName string) {
	index := strings.LastIndex(structureData["namespace"], "/")
	if index != -1 {
		applicationName = structureData["namespace"][index+1:]
	} else {
		applicationName = structureData["namespace"]
	}
	return applicationName
}
