package msg

import (
	"bufio"
	"io"
)

type ErrReader struct {
	r   *bufio.Reader
	Err error
}

// NewErrReader creates a new ErrReader.
func NewErrReader(r io.Reader) *ErrReader {
	br := bufio.NewReader(r)
	er := &ErrReader{r: br}
	return er
}

func (er *ErrReader) ReadRune() (rune, int, error) {
	var r rune
	var i int
	if er.Err != nil {
		return r, i, er.Err
	}
	r, i, er.Err = er.r.ReadRune()
	return r, i, er.Err
}
