package writer

import "fmt"

type WriterDecorator interface {
	Decorate(str, decorate string)
}

type WriterDecoratorBase struct {
	writer Writer
}

func (base *WriterDecoratorBase) Decorate(str, decorator string) {
	base.writer.Write(str)
	fmt.Println(decorator)
}
