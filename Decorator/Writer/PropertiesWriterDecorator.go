package writer

import (
	"fmt"
	"time"
)

type PropertiesWriter struct {
	WriterDecoratorBase
	creationTime time.Time
	modifyTime   time.Time
}

func NewPropertiesWriter(writer Writer, creationTime, modeifyTime time.Time) *PropertiesWriter {
	return &PropertiesWriter{
		WriterDecoratorBase: WriterDecoratorBase{writer: writer},
		creationTime:        creationTime,
		modifyTime:          modeifyTime,
	}
}

func (prop *PropertiesWriter) Write(str string) {
	prop.Decorate(str, fmt.Sprintf("Properties Writer write properties\n%s", prop.getStringProperties()))
}

func (prop *PropertiesWriter) getStringProperties() (result string) {
	result = fmt.Sprintf("Creation Time: %s\n", prop.creationTime.Format("YYYY/MM/DD HH:MM:SS"))
	result += fmt.Sprintf("Modify Time: %s\n", prop.modifyTime.Format("YYYY/MM/DD HH:MM:SS"))
	return
}
