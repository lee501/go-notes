package main

import "os"

type BufferFileWriter struct {
	fs            *os.File
	cache         []byte
	cacheEndIndex int
}

func NewBufferFileWriter(fs *os.File, cacheSize int) *BufferFileWriter {
	return &BufferFileWriter{
		fs:            fs,
		cache:         make([]byte, cacheSize),
		cacheEndIndex: 0,
	}
}

func (w *BufferFileWriter) WriteByte(cont []byte) {
	if len(cont) > len(w.cache) {
		w.Flush()
		w.fs.Write(cont)
	} else {
		if len(cont)+w.cacheEndIndex > len(w.cache) {
			w.Flush()
		}
		copy(w.cache[w.cacheEndIndex:], cont)
		w.cacheEndIndex += len(cont)
	}
}

func (w *BufferFileWriter) Flush() {
	w.fs.Write(w.cache[0:w.cacheEndIndex])
	w.cacheEndIndex = 0
}

func (w *BufferFileWriter) WriteString(cont string) {
	w.WriteByte([]byte(cont))
}

func WriteWithBuffer(outFile string) {
	fs, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	writer := NewBufferFileWriter(fs, 2048)
	defer writer.Flush()
	for i := 0; i < 10000; i++ {
		writer.WriteString("this file use with cache")
	}
}
