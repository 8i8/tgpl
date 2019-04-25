package msg

import (
	"strings"
	"testing"
)

func TestErrReadRune(t *testing.T) {

	re := NewErrReader(strings.NewReader("This is a test"))
	r1, _, _ := re.ReadRune()
	r2, _, _ := re.ReadRune()

	if r1 != 'T' {
		t.Errorf("test: TestErrReadRune failed, recieved `%c` expected `%c`", r1, 'T')
	}

	if r2 != 'h' {
		t.Errorf("test: TestErrReadRune failed, recieved `%c` expected `%c`", r2, 'h')
	}
}
