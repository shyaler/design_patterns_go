package main

func main() {
	builder1 := &EnumeratorBuilderObj{}
	enumerator1 := builder1.WithSourceTraverser(
		&FileTraverser{fileUrl: "https://fileUrl"}).WithTargetTraverser(
		&BlobTraverser{blobUrl: "https://blobUrl"}).WithTraversingChannel(
		&OrderedChannel{}).Build()

	enumerator1.Enumerate()

	builder2 := &EnumeratorBuilderObj{}
	enumerator2 := builder2.WithSourceTraverser(
		&BlobTraverser{blobUrl: "https://blobUrl"}).WithTargetTraverser(
		&BlobTraverser{blobUrl: "https://blobUrl"}).WithTraversingChannel(
		&SlidingWindowChannel{}).Build()

	enumerator2.Enumerate()
}
