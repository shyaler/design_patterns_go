package traverser

import (
	"fmt"
	"net/url"
)

type BlobTraverser struct {
	TraverserBase
}

func (tr BlobTraverser) ReadDirectory(url *url.URL) {
	fmt.Printf("Blob Traverser reading directory %s\n", url.String())
}

func (tr BlobTraverser) ReadFile(url *url.URL) {
	fmt.Printf("Blob Traverser reading file %s\n", url.String())
}
