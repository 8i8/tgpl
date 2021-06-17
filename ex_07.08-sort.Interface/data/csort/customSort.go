package csort

import (
	"bytes"
	"strings"
)

const (
	Equal = iota
	Left
	Right
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  customSort
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

type Custom struct {
	t    []interface{}
	less func(x, y interface{}) bool
}

func New(less func(x, y interface{}) bool, t []interface{}) *Custom {
	c := new(Custom)
	c.t = t
	c.less = less
	return c
}

func (x Custom) Len() int           { return len(x.t) }
func (x Custom) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x Custom) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  SortBuffer
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// SortFn are the functions used by the customSort less() function call.
type SortFn func(x, y interface{}) int

// SortBuffer maintains a ring buffer of the last n sort functions that
// have been called.
type SortBuffer struct {
	Funcs  []SortFn
	list   []SortFn
	cmds   []string
	slist  bytes.Buffer
	i      int
	l      int
	Max    int
	SortBy func(string) SortFn
}

// NewSortBuffer returns a new sort buffer, a ring buffer that retains
// the last n sort function calls in hitorical order.
func NewSortBuffer(s func(string) SortFn, v ...int) *SortBuffer {
	var l int
	if len(v) == 0 {
		l = 5
	} else {
		l = v[0]
	}
	return &SortBuffer{
		Funcs:  make([]SortFn, l),
		list:   make([]SortFn, 0, l),
		cmds:   make([]string, l),
		Max:    l,
		SortBy: s,
	}
}

// Add adds a sort function to the buffer.
func (s *SortBuffer) Load(cmd ...string) {
	const fname = "Load"
	if len(cmd) == 0 {
		return
	}
	for i := range cmd {
		s.i %= s.Max
		s.Funcs[s.i] = s.SortBy(cmd[i])
		s.cmds[s.i] = cmd[i]
		s.i++
		if s.l < s.Max {
			s.l++
		}
	}
}

func reverse(str string) string {
	if val := strings.Split(str, "-"); len(val) > 1 {
		return val[0]
	}
	return str + "-rev"
}

// Add adds a sort function to the buffer.
func (s *SortBuffer) Add(cmd string) {
	const fname = "Add"
	if cmd == "" {
		return
	}
	prev := s.Peek()
	if prev == cmd || prev == cmd+"-rev" {
		cmd = reverse(prev)
		s.Funcs[s.i-1] = s.SortBy(cmd)
		s.cmds[s.i-1] = cmd
		return
	}
	s.i %= s.Max
	s.Funcs[s.i] = s.SortBy(cmd)
	s.cmds[s.i] = cmd
	s.i++
	if s.l < s.Max {
		s.l++
	}
}

// Peek returns the last added commands name.
func (s *SortBuffer) Peek() string {
	if s.l == 0 {
		return "NUL"
	}
	return s.cmds[(s.i+s.l-1)%s.l]
}

// List returns a slice for itteration from the sort order ring buffer.
func (s *SortBuffer) List() []SortFn {
	if s.l == 0 {
		return nil
	}
	s.list = s.list[:0]
	for i := 0; i < s.l; i++ {
		s.list = append(s.list, s.Funcs[((s.i-1+s.l)-i)%s.l])
	}
	return s.list
}

func (s *SortBuffer) String() string {
	if s.l == 0 {
		return "NUL"
	}
	s.slist.Reset()
	s.slist.WriteString(s.cmds[(s.i)%s.l])
	for i := 1; i < s.l; i++ {
		s.slist.WriteByte(',')
		s.slist.WriteString(s.cmds[((s.i)+i)%s.l])
	}
	return s.slist.String()
}

// LoadSortFn runs all of the conditional functions in the list of sort
// function returned from the SortBuffer List call.
func (s *SortBuffer) LoadSortFn() func(x, y interface{}) bool {
	return func(x, y interface{}) bool {
		for _, fn := range s.List() {
			if res := fn(x, y); res > Equal {
				if res == Left {
					return true
				}
				return false
			}
		}
		return false
	}
}
