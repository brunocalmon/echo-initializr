package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var files Files

// Initializr reads all folders and files from the given resource and achitecture storing their informations in memory.
func Initializr(namespace, architecture string) Files {
	resourceChoosen := "./resources/" + architecture
	resources := readResources(resourceChoosen, len(resourceChoosen))
	for _, file := range resources {
		createFile(File{Namespace: namespace, Folder: file["dir"], Name: file["fileName"], Content: file["content"], Keys: stringToArray(file["keys"])})
	}

	return files
}

func stringToArray(keys string) []string {
	if keys == "" {
		return []string{}
	}

	var array []string
	dec := json.NewDecoder(strings.NewReader(keys))
	err := dec.Decode(&array)
	if err != nil {
		fmt.Println("keys mapped into the files couldn't be parse to arrays.")
		panic("keys mapped into the files couldn't be parse to arrays.")
	}
	return array
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
			fileInfos, err := readFile(dirName, resource.Name(), resourceSubstringIndex)
			if err != nil {
				panic(err)
			}
			files = append(files, fileInfos)
		}
	}
	return files
}

func createFile(f File) {
	files = append(files, f)
}

func readFile(dir, fileName string, resourceSubstringIndex int) (response map[string]string, err error) {
	fmt.Printf("reading sample: %s\n", fileName)
	if strings.Contains(fileName, ".sample") {
		file, err := os.Open(dir + "/" + fileName)
		defer file.Close()

		if err != nil {
			return nil, err
		}

		response = make(map[string]string)
		response["dir"] = dir[resourceSubstringIndex:]
		response["fileName"] = fileName[:strings.LastIndex(fileName, ".sample")]

		// Start reading from the file with a reader.
		reader := bufio.NewReader(file)

		var line string
		lineNumber := 0
		containsHeader := false
		for {
			line, err = reader.ReadString('\n')

			if err != nil {
				if err == io.EOF {
					response["content"] += line
				}
				break
			}

			if lineNumber == 0 {
				containsHeader = strings.Contains(line, "header:")
			}

			if !containsHeader && lineNumber < 4 || lineNumber >= 4 {
				response["content"] += line
				goto End
			}

			if lineNumber < 4 && strings.Contains(line, "keys") {
				response["keys"] = line[strings.Index(line, "[") : strings.LastIndex(line, "]")+1]
			}

		End:
			lineNumber++
		}

		if err != io.EOF {
			fmt.Printf(" > Failed!: %v\n", err)
			return nil, err
		}
		return response, nil
	}
	panic("No files found!")
}
