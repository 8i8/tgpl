package main

import (
	"testing"
)

func TestCommaBuf(t *testing.T) {
	testComma(t, commaRec, "commaRec")
	testComma(t, comma, "comma")
	testComma(t, commaBf1, "commaBf1")
	testComma(t, commaBf2, "commaBf2")
	testComma(t, commaBf3, "cammoBf3")
}

func testComma(t *testing.T, test func(string) string, name string) {

	test1(t, test, name, "1", "1")
	test1(t, test, name, "10", "10")
	test1(t, test, name, "100", "100")
	test1(t, test, name, "1000", "1,000")
	test1(t, test, name, "10000", "10,000")
	test1(t, test, name, "100000", "100,000")
	test1(t, test, name, "1000000", "1,000,000")
	test1(t, test, name, "10000000", "10,000,000")
	test1(t, test, name, "100000000", "100,000,000")
	test1(t, test, name, "1000000000", "1,000,000,000")
	test1(t, test, name, "10000000000", "10,000,000,000")
	test1(t, test, name, "100000000000", "100,000,000,000")
	test1(t, test, name, "1000000000000", "1,000,000,000,000")
}

func test1(t *testing.T, test func(string) string, name, sample, want string) {
	str := test(sample)
	if str != want {
		t.Errorf(`%s returned %v wanted %v`, name, str, want)
	}
}

func BenchmarkCommaRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchComma(commaRec)
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchComma(comma)
	}
}

func BenchmarkCommaBf1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchComma(commaBf1)
	}
}

func BenchmarkCommaBf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchComma(commaBf2)
	}
}

func BenchmarkCommaBf3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchComma(commaBf3)
	}
}

func benchComma(benchFunc func(string) string) {

	bench1("1", benchFunc)
	bench1("10", benchFunc)
	bench1("100", benchFunc)
	bench1("1000", benchFunc)
	bench1("10000", benchFunc)
	bench1("100000", benchFunc)
	bench1("1000000", benchFunc)
	bench1("10000000", benchFunc)
	bench1("100000000", benchFunc)
	bench1("1000000000", benchFunc)
	bench1("10000000000", benchFunc)
	bench1("100000000000", benchFunc)
	bench1("1000000000000", benchFunc)
}

func bench1(control string, benchFunc func(string) string) {
	benchFunc(control)
}
