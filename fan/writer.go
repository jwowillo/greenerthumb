package main

import (
	"fmt"
	"io"
)

// MultiWriter is like io.MultiWriter except it doesn't stop writing after an
// error.
type MultiWriter struct {
	writers []io.Writer
}

// NewMultiWriter which writes to all the io.Writers.
func NewMultiWriter(ws []io.Writer) *MultiWriter {
	return &MultiWriter{writers: ws}
}

func (w *MultiWriter) Write(bs []byte) (int, error) {
	var errMessages []string
	for _, subw := range w.writers {
		if _, err := subw.Write(bs); err != nil {
			errMessages = append(errMessages, err.Error())
		}
	}
	if errMessages != nil {
		return -1, fmt.Errorf("%v", errMessages)
	}
	return len(bs), nil
}
