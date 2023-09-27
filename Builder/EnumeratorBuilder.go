package main

type EnumeratorBuilder interface {
	Build() *SyncEnumerator
}

type EnumeratorBuilderObj struct {
	enumerator      *SyncEnumerator
	sourceTraverser Traverser
	targetTraverser Traverser
	channel         TraversingChannel
}

func (builder *EnumeratorBuilderObj) WithSourceTraverser(traverser Traverser) *EnumeratorBuilderObj {
	builder.sourceTraverser = traverser
	return builder
}

func (builder *EnumeratorBuilderObj) WithTargetTraverser(traverser Traverser) *EnumeratorBuilderObj {
	builder.targetTraverser = traverser
	return builder
}

func (builder *EnumeratorBuilderObj) WithTraversingChannel(channel TraversingChannel) *EnumeratorBuilderObj {
	builder.channel = channel
	return builder
}

func (builder *EnumeratorBuilderObj) Build() *SyncEnumerator {
	builder.enumerator = &SyncEnumerator{
		sourceTraverser: builder.sourceTraverser,
		targetTraverser: builder.targetTraverser,
		channel:         builder.channel,
	}

	return builder.enumerator
}
