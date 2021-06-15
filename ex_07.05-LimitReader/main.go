// Exercise 7.5: The LimitReader function in the io package accepts and
// io.Reader and a number of bytes n, and returns another Reader that
// reads from r but reports an end-of-file condition after n bytes.
// Impliement it.
//	func LimitReader(r io.Reader, n int64) io.Reader
package main

import (
	"io"
)

type Limiter struct {
	r io.Reader // Reader
	n int       // Bytes remaining
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &Limiter{r, int(n)}
}

func (l Limiter) Read(b []byte) (n int, err error) {
	if len(b) > l.n {
		b = b[:l.n]
		err = io.EOF
	}
	n, errRead := l.r.Read(b)
	if errRead != nil && err != io.EOF {
		err = errRead
	}
	return
}

func main() {
}
