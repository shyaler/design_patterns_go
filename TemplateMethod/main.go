package main

import (
	traverser "main/Traverser"

	"net/url"
)

func main() {
	blobTraverser := &traverser.TraverserBase{ITraverser: traverser.BlobTraverser{}}
	blobTraverser.Traverse(&url.URL{Host: "127.0.0.1", Path: "blob/path"})

	fileTraverser := &traverser.TraverserBase{ITraverser: traverser.FileTraverser{}}
	fileTraverser.Traverse(&url.URL{Host: "127.0.0.1", Path: "file/path"})
}
