// @Author huzejun 2024/1/2 13:05:00
package logs

import "io"

var writerAdapter = make(map[string]InitLogWriterFunc, 0)

type InitLogWriterFunc func() LogWriter

type LogWriter interface {
	Flush()
	io.Writer
}

func RegisterInitWriterFunc(adapterName string, writerFunc InitLogWriterFunc) {
	writerAdapter[adapterName] = writerFunc
}
