package data

import "github.com/brunocalmon/echo-initializr/global"

type FileContent interface {
	Path()
	PathWithFile()
}

type File struct {
	Namespace string
	Folder    string
	Name      string
	Content   string
}

func (f File) Path() string {
	var outputDir string = global.Global["outputDir"].(string)
	return outputDir + "/" + f.Namespace + f.Folder
}

func (f File) PathWithFile() string {
	var outputDir string = global.Global["outputDir"].(string)
	return outputDir + "/" + f.Namespace + f.Folder + "/" + f.Name
}

type Files []File
