package main

import "fmt"

type FileTraverser struct {
	fileUrl string
}

func (file *FileTraverser) Traverse() {
	fmt.Printf("Start traversing az file %s\n", file.fileUrl)
}
