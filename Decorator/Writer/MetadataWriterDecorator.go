package writer

import "fmt"

type MetadataWriter struct {
	WriterDecoratorBase
	metadata map[string]string
}

func NewMetadataWriter(writer Writer, metadata map[string]string) *MetadataWriter {
	return &MetadataWriter{WriterDecoratorBase{writer: writer}, metadata}
}

func (me *MetadataWriter) Write(str string) {
	me.Decorate(str, fmt.Sprintf("Metadata Writer write metadata\n%s", me.getStringMetadata()))
}

func (me *MetadataWriter) getStringMetadata() (result string) {
	result = ""
	for key, value := range me.metadata {
		result += fmt.Sprintf("Metadata Key: %s, Value: %s\n", key, value)
	}

	return
}
