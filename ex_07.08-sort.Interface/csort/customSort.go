package csort

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

// SortFnBuffer maintains a ring buffer of the last n sort functions that
// have been called.
type SortFnBuffer struct {
	Funcs  []SortFn
	list   []SortFn
	i      int
	l      int
	Max    int
	SortBy func(string) SortFn
}

// List returns a slice for itteration from the sort order ring buffer.
func (s *SortFnBuffer) List() []SortFn {
	s.list = s.list[:0]
	for i := 0; i < s.l; i++ {
		s.list = append(s.list, s.Funcs[((s.i+s.Max)-i-1)%s.Max])
	}
	return s.list
}

// SortFunction runs all of the conditional functions in the list of sort
// function returned from the SortFnBuffer List call.
func SortFunction(buf *SortFnBuffer) func(x, y interface{}) bool {
	return func(x, y interface{}) bool {
		for _, fn := range buf.List() {
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

// NewSortBuffer returns a new sort buffer, a ring buffer that retains
// the last n sort function calls in hitorical order.
func NewSortBuffer(s func(string) SortFn, v ...int) *SortFnBuffer {
	var l int
	if len(v) == 0 {
		l = 5
	} else {
		l = v[0]
	}
	return &SortFnBuffer{
		Funcs:  make([]SortFn, l),
		list:   make([]SortFn, 0, l),
		Max:    l,
		SortBy: s,
	}
}

// Add adds a sort function to the buffer.
func (s *SortFnBuffer) Add(cmd string) {
	s.i %= s.Max
	s.Funcs[s.i] = s.SortBy(cmd)
	s.i++
	if s.l < s.Max {
		s.l++
	}
}
