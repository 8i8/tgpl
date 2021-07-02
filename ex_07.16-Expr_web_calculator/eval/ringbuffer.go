package eval

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

const defaltLen = 5

type List []string

func (l List) String() string {
	buf := strings.Builder{}
	for _, elm := range l {
		buf.WriteString(elm)
		buf.WriteByte('\n')
	}
	return buf.String()
}

type ringbuffer struct {
	buf []string
	i   int
}

// NewBuffer returns a new buffer of the size requested, if no value is
// given then the default value is used.
func NewBuffer(max ...int) *ringbuffer {
	var m int
	if len(max) == 0 {
		m = defaltLen
	} else {
		m = max[0]
	}
	return &ringbuffer{
		buf: make([]string, 0, m),
	}
}

// Load takes a string of comma separated expressions, splits them at
// the commas and then loads the buffer with the resulting strings.
func (r *ringbuffer) Load(value string) error {
	const fname = "ringbuffer.Load"
	byt, err := base64.RawStdEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("%s: %w", fname, err)
	}
	strs := strings.Split(string(byt), "&&")
	if len(strs) > cap(r.buf) {
		return errors.New("number of elements added excede the buffer length")
	}
	for i := range strs {
		r.Add(strs[i])
	}
	return nil
}

// Unload returns a comma seperated string that contains all of the
// expressions in the buffer.
func (r *ringbuffer) Unload() string {
	buf := bytes.Buffer{}
	strs := r.List()
	if len(strs) == 0 {
		return ""
	}
	buf.WriteString(strs[0])
	for i := range strs[1:] {
		buf.WriteString("&&")
		buf.WriteString(strs[i+1])
	}
	return base64.RawStdEncoding.EncodeToString(buf.Bytes())
}

// Reset clears the buffers memory.
func (r *ringbuffer) Reset() {
	r.buf = r.buf[:0]
	r.i = 0
}

// Add adds the value to the buffer, pushing out the oldest existing
// value if the buffer is already full.
func (r *ringbuffer) Add(c string) {
	if len(r.buf) < cap(r.buf) {
		r.buf = append(r.buf, c)
		r.i = (r.i + 1) % len(r.buf)
		return
	}
	r.i = (r.i + 1) % len(r.buf)
	r.buf[r.i] = c
}

// MoveUp moves the value to the top of the queue if it already
// exists in the queue and adds it to the queue if it does not.
func (r *ringbuffer) MoveUp(c string) {
	x, no := notInArray(r, c)
	if no {
		r.Add(c)
		return
	}
	switch {
	case r.i == x:
		// No need to do anything, c is already at the top of
		// the queue, replace just in case c has been updated.
		r.buf[r.i] = c
	case r.i > x:
		// Shunt values back to fill x's spot leaving the head
		// free, add c to the top of the queue.
		copy(r.buf[x:r.i], r.buf[x+1:r.i+1])
		r.buf[r.i] = c
	case r.i < x:
		// Shunt the values beneath x up to fill x's spot
		// leaving room to grow the head of the queue by one and
		// then fill it with c.
		r.i++
		copy(r.buf[r.i+1:x+1], r.buf[r.i:x])
		r.buf[r.i] = c
	}
}

// notInArray returns false with the index of the value if it exists in
// the array and true if it is not found.
func notInArray(r *ringbuffer, c string) (int, bool) {
	for i := 0; i < len(r.buf); i++ {
		if r.buf[i] == c {
			return i, false
		}
	}
	return 0, true
}

// List outputs the current buffers content as a list.
func (r *ringbuffer) List() List {
	out := make([]string, 0, len(r.buf))
	for i := 0; i < len(r.buf); i++ {
		out = append(out, r.buf[(r.i+1+i)%len(r.buf)])
	}
	return out
}
