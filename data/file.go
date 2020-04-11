package data

type FileContent interface {
	Path(outputDir string)
	PathWithFile(outputDir string)
}

type File struct {
	Namespace string
	Folder    string
	Name      string
	Content   string
}

func (f File) Path(outputDir string) string {
	return outputDir + "/" + f.Namespace + f.Folder
}

func (f File) PathWithFile(outputDir string) string {
	return outputDir + "/" + f.Namespace + f.Folder + "/" + f.Name
}

type Files []File
