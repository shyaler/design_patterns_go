package main

type Enumerator interface {
	Enumerate()
}

type SyncEnumerator struct {
	sourceTraverser Traverser
	targetTraverser Traverser
	channel TraversingChannel
}

func (enumerator*SyncEnumerator) Enumerate() {
	enumerator.channel.setUpChan(10)
	enumerator.sourceTraverser.Traverse()
	enumerator.targetTraverser.Traverse()
}
