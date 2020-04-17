package data

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var files Files

// Initializr reads all folders and files from the given resource and achitecture storing their informations in memory.
func Initializr(namespace, architecture string) Files {
	resourceChoosen := "./resources/" + architecture
	resources := readResources(resourceChoosen, len(resourceChoosen))
	for _, file := range resources {
		createFile(File{Namespace: namespace, Folder: file["dir"], Name: file["fileName"], Content: file["content"]})
	}

	return files
}

func readResources(dirName string, resourceSubstringIndex int) []map[string]string {
	resources, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("Directory reading error", err)
		panic("Directory reading error")
	}

	files := make([]map[string]string, 0)
	for _, resource := range resources {
		if resource.IsDir() {
			fileInfos := readResources(fmt.Sprintf("%s/%s", dirName, resource.Name()), resourceSubstringIndex)
			files = append(files, fileInfos...)
		} else {
			fileInfos := readFile(dirName, resource.Name(), resourceSubstringIndex)
			files = append(files, fileInfos)
		}
	}
	return files
}

func readFile(dir, fileName string, resourceSubstringIndex int) map[string]string {
	if strings.Contains(fileName, ".sample") {
		data, err := ioutil.ReadFile(dir + "/" + fileName)
		if err != nil {
			fmt.Println("File reading error", err)
			panic("File reading error")
		}

		relativeDir := dir[resourceSubstringIndex:]
		realFileName := fileName[:strings.LastIndex(fileName, ".sample")]
		content := string(data)

		return map[string]string{"dir": relativeDir, "fileName": realFileName, "content": content}
	}
	panic("No files found!")
}

func createFile(f File) {
	files = append(files, f)
}
