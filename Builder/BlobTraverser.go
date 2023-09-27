package main

import "fmt"

type BlobTraverser struct {
	blobUrl string
}

func (blob *BlobTraverser) Traverse() {
	fmt.Printf("Start traversing az blob %s\n", blob.blobUrl)
}
