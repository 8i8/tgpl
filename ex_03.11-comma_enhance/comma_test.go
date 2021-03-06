package main

import (
	"testing"
)

func TestCommaBuf(t *testing.T) {
	testComma(t, commaBuf)
	testComma(t, comma)
}

func testComma(t *testing.T, test func(string) string) {

	test1(t, "1", "1", test)
	test1(t, "10", "10", test)
	test1(t, "100", "100", test)
	test1(t, "1000", "1,000", test)
	test1(t, "10000", "10,000", test)
	test1(t, "100000", "100,000", test)
	test1(t, "1000000", "1,000,000", test)
	test1(t, "10000000", "10,000,000", test)
	test1(t, "100000000", "100,000,000", test)
	test1(t, "1000000000", "1,000,000,000", test)
	test1(t, "10000000000", "10,000,000,000", test)
	test1(t, "100000000000", "100,000,000,000", test)
	test1(t, "1000000000000", "1,000,000,000,000", test)

	test1(t, "-1", "-1", test)
	test1(t, "-10", "-10", test)
	test1(t, "-100", "-100", test)
	test1(t, "-1000", "-1,000", test)
	test1(t, "-10000", "-10,000", test)
	test1(t, "-100000", "-100,000", test)
	test1(t, "-1000000", "-1,000,000", test)
	test1(t, "-10000000", "-10,000,000", test)
	test1(t, "-100000000", "-100,000,000", test)
	test1(t, "-1000000000", "-1,000,000,000", test)
	test1(t, "-10000000000", "-10,000,000,000", test)
	test1(t, "-100000000000", "-100,000,000,000", test)
	test1(t, "-1000000000000", "-1,000,000,000,000", test)

	test1(t, "1.0", "1.0", test)
	test1(t, "10.0", "10.0", test)
	test1(t, "100.0", "100.0", test)
	test1(t, "1000.0", "1,000.0", test)
	test1(t, "10000.0", "10,000.0", test)
	test1(t, "100000.0", "100,000.0", test)
	test1(t, "1000000.0", "1,000,000.0", test)
	test1(t, "10000000.0", "10,000,000.0", test)
	test1(t, "100000000.0", "100,000,000.0", test)
	test1(t, "1000000000.0", "1,000,000,000.0", test)
	test1(t, "10000000000.0", "10,000,000,000.0", test)
	test1(t, "100000000000.0", "100,000,000,000.0", test)
	test1(t, "1000000000000.0", "1,000,000,000,000.0", test)

	test1(t, "-1.0", "-1.0", test)
	test1(t, "-10.0", "-10.0", test)
	test1(t, "-100.0", "-100.0", test)
	test1(t, "-1000.0", "-1,000.0", test)
	test1(t, "-10000.0", "-10,000.0", test)
	test1(t, "-100000.0", "-100,000.0", test)
	test1(t, "-1000000.0", "-1,000,000.0", test)
	test1(t, "-10000000.0", "-10,000,000.0", test)
	test1(t, "-100000000.0", "-100,000,000.0", test)
	test1(t, "-1000000000.0", "-1,000,000,000.0", test)
	test1(t, "-10000000000.0", "-10,000,000,000.0", test)
	test1(t, "-100000000000.0", "-100,000,000,000.0", test)
	test1(t, "-1000000000000.0", "-1,000,000,000,000.0", test)

	test1(t, "1.00", "1.00", test)
	test1(t, "10.00", "10.00", test)
	test1(t, "100.00", "100.00", test)
	test1(t, "1000.00", "1,000.00", test)
	test1(t, "10000.00", "10,000.00", test)
	test1(t, "100000.00", "100,000.00", test)
	test1(t, "1000000.00", "1,000,000.00", test)
	test1(t, "10000000.00", "10,000,000.00", test)
	test1(t, "100000000.00", "100,000,000.00", test)
	test1(t, "1000000000.00", "1,000,000,000.00", test)
	test1(t, "10000000000.00", "10,000,000,000.00", test)
	test1(t, "100000000000.00", "100,000,000,000.00", test)
	test1(t, "1000000000000.00", "1,000,000,000,000.00", test)

	test1(t, "-1.00", "-1.00", test)
	test1(t, "-10.00", "-10.00", test)
	test1(t, "-100.00", "-100.00", test)
	test1(t, "-1000.00", "-1,000.00", test)
	test1(t, "-10000.00", "-10,000.00", test)
	test1(t, "-100000.00", "-100,000.00", test)
	test1(t, "-1000000.00", "-1,000,000.00", test)
	test1(t, "-10000000.00", "-10,000,000.00", test)
	test1(t, "-100000000.00", "-100,000,000.00", test)
	test1(t, "-1000000000.00", "-1,000,000,000.00", test)
	test1(t, "-10000000000.00", "-10,000,000,000.00", test)
	test1(t, "-100000000000.00", "-100,000,000,000.00", test)
	test1(t, "-1000000000000.00", "-1,000,000,000,000.00", test)

	test1(t, "1.000", "1.000", test)
	test1(t, "10.000", "10.000", test)
	test1(t, "100.000", "100.000", test)
	test1(t, "1000.000", "1,000.000", test)
	test1(t, "10000.000", "10,000.000", test)
	test1(t, "100000.000", "100,000.000", test)
	test1(t, "1000000.000", "1,000,000.000", test)
	test1(t, "10000000.000", "10,000,000.000", test)
	test1(t, "100000000.000", "100,000,000.000", test)
	test1(t, "1000000000.000", "1,000,000,000.000", test)
	test1(t, "10000000000.000", "10,000,000,000.000", test)
	test1(t, "100000000000.000", "100,000,000,000.000", test)
	test1(t, "1000000000000.000", "1,000,000,000,000.000", test)

	test1(t, "-1.000", "-1.000", test)
	test1(t, "-10.000", "-10.000", test)
	test1(t, "-100.000", "-100.000", test)
	test1(t, "-1000.000", "-1,000.000", test)
	test1(t, "-10000.000", "-10,000.000", test)
	test1(t, "-100000.000", "-100,000.000", test)
	test1(t, "-1000000.000", "-1,000,000.000", test)
	test1(t, "-10000000.000", "-10,000,000.000", test)
	test1(t, "-100000000.000", "-100,000,000.000", test)
	test1(t, "-1000000000.000", "-1,000,000,000.000", test)
	test1(t, "-10000000000.000", "-10,000,000,000.000", test)
	test1(t, "-100000000000.000", "-100,000,000,000.000", test)
	test1(t, "-1000000000000.000", "-1,000,000,000,000.000", test)
}

func test1(t *testing.T, control, want string, test func(string) string) {
	str := test(control)
	if str != want {
		t.Errorf(`returned %v wanted %v`, str, want)
	}
}

func BenchmarkCommaRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commaRec("10000000000")
	}
}

func BenchmarkCommaBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commaBuf("10000000000")
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma("10000000000")
	}
}
