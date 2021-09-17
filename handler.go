package galog

import (
	"io"
)

//Handler writes logs to somewhere
type Handler interface {
	Write(p []byte) (n int, err error)
	Close() error
}

//StreamHandler writes logs to a specified io Writer, maybe stdout, stderr, etc...
type StreamHandler struct {
	w io.Writer
}

// NewStreamHandler return StreamHandler
func NewStreamHandler(w io.Writer) (*StreamHandler, error) {
	h := new(StreamHandler)

	h.w = w

	return h, nil
}

func (h *StreamHandler) Write(b []byte) (n int, err error) {
	return h.w.Write(b)
}

// Close do nothing
func (h *StreamHandler) Close() error {
	return nil
}

//NullHandler does nothing, it discards anything.
type NullHandler struct {
}

// NewNullHandler return NullHandler
func NewNullHandler() (*NullHandler, error) {
	return new(NullHandler), nil
}

// Write do nothing
func (h *NullHandler) Write(b []byte) (n int, err error) {
	return len(b), nil
}

// Close do nothing
func (h *NullHandler) Close() error {
	return nil
}
