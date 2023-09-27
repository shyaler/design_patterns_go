package writer

import "fmt"

type Writer interface {
	Write(str string)
}

type FileWriter struct {
	FilePath string
}

func (file *FileWriter) Write(str string) {
	fmt.Printf("File writer writing %s to path %s\n", str, file.FilePath)
}

type FolderWriter struct {
	FolderPath string
}

func (folder *FolderWriter) Write(str string) {
	fmt.Printf("Folder writer writing %s to path %s\n", str, folder.FolderPath)
}
