package main

import (
	"bufio"
	"bytes"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	count, err := counter(p, bufio.ScanWords)
	if err != nil {
		return len(p), err
	}
	*w += WordCounter(count)
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count, err := counter(p, bufio.ScanLines)
	if err != nil {
		return len(p), err
	}
	*l += LineCounter(count)
	return len(p), nil
}

func counter(p []byte, fn bufio.SplitFunc) (int, error) {
	in := bufio.NewScanner(bytes.NewReader(p))
	in.Split(fn)
	var count int
	for in.Scan() {
		count++
	}
	err := in.Err()
	return count, err
}

func main() {
}
