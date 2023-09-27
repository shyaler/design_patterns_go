package traverser

import (
	"fmt"
	"net/url"
)

type ITraverser interface {
	ReadDirectory(url *url.URL)
	ReadFile(url *url.URL)
}

type TraverserBase struct {
	ITraverser
}

func (tr *TraverserBase) Traverse(url *url.URL) {
	fmt.Println("Start traversing")

	directoryURL := url.JoinPath("Directory")
	tr.ReadDirectory(directoryURL)
	fileURL := url.JoinPath("File")
	tr.ReadFile(fileURL)

	fmt.Println("Finish traversing")
}
