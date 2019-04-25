package msg

import (
	"io"
	"strconv"
	"unicode/utf8"
)

// ErrWriter is a wrapper around the io.Writer that encapsulates the error
// handling abstracting it from the function that is using the writer.
type ErrWriter struct {
	w   io.Writer
	Err error
}

// NewErrWriter creates a new ErrWriter.
func NewErrWriter(w io.Writer) *ErrWriter {
	ew := &ErrWriter{w: w}
	return ew
}

func (ew *ErrWriter) Write(buf []byte) {
	if ew.Err != nil {
		return
	}

	_, ew.Err = ew.w.Write(buf)
}

func (ew *ErrWriter) WriteRune(r rune) {
	if ew.Err != nil {
		return
	}
	a := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(a, r)
	_, ew.Err = ew.w.Write(a)
}

func (ew *ErrWriter) WriteInt(i int) {
	if ew.Err != nil {
		return
	}

	_, ew.Err = ew.w.Write([]byte(strconv.Itoa(i)))
}
