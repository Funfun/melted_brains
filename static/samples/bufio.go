package bufio2

import "io"

const (
	defaultBufSize  = 4096
	defaultBufCount = 4
)

type AsyncWriter struct {
	err       error
	buf       []byte
	n         int
	size      int
	countSent int
	countRcvd int

	wr io.Writer

	dataChan   chan []byte
	resultChan chan error
}

func NewAsyncWriterSize(wr io.Writer, size int, count int) *AsyncWriter {
	b, ok := wr.(*AsyncWriter)
	if ok && len(b.buf) >= size {
		return b
	}
	if size <= 0 {
		size = defaultBufSize
	}
	b = new(AsyncWriter)
	b.size = size
	b.buf = make([]byte, size)
	b.wr = wr
	if count <= 0 {
		count = defaultBufCount
	}
	b.dataChan = make(chan []byte, count)
	b.resultChan = make(chan error, count+1)
	go func() {
		b.writeThread()
	}()
	return b
}

func NewAsyncWriter(wr io.Writer) *AsyncWriter {
	return NewAsyncWriterSize(wr, defaultBufSize, defaultBufCount)
}

func (b *AsyncWriter) writeThread() {
	var err error
	for buf := range b.dataChan {
		if err == nil {
			var n int
			n, err = b.wr.Write(buf)
			if n < len(buf) && err == nil {
				err = io.ErrShortWrite
			}
		}
		b.resultChan <- err
	}
}
