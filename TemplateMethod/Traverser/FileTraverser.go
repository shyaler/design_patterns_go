package traverser

import (
	"fmt"
	"net/url"
)

type FileTraverser struct {
	TraverserBase
}

func (tr FileTraverser) ReadDirectory(url *url.URL) {
	fmt.Printf("File Traverser reading directory %s\n", url.String())
}

func (tr FileTraverser) ReadFile(url *url.URL) {
	fmt.Printf("File Traverser reading file %s\n", url.String())
}
