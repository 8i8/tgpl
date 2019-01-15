package fractal

import (
	"io/ioutil"
	"testing"
)

func BenchmarkFractal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fractal(ioutil.Discard)
	}
}
