package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"testing"
)

func BenchmarkCountingSha1(b *testing.B) {
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	for i := 0; i < b.N; i++ {
		BitComp1(c1, c2)
	}
}

func BenchmarkCountingSha2(b *testing.B) {
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	for i := 0; i < b.N; i++ {
		BitComp2(c1, c2)
	}
}

func BenchmarkCountingSha3(b *testing.B) {
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	for i := 0; i < b.N; i++ {
		BitComp3(c1, c2)
	}
}

func BenchmarkCountingSha4(b *testing.B) {
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	for i := 0; i < b.N; i++ {
		BitComp4(c1, c2)
	}
}

func BenchmarkCountingSha5(b *testing.B) {

	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	d1 := data{}
	d2 := data{}
	r1 := bytes.NewReader(c1[:])
	r2 := bytes.NewReader(c2[:])

	if err := binary.Read(r1, binary.LittleEndian, &d1); err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	if err := binary.Read(r2, binary.LittleEndian, &d2); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	for i := 0; i < b.N; i++ {
		BitComp5(d1, d2)
	}
}

func BenchmarkCountingSha6(b *testing.B) {
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	for i := 0; i < b.N; i++ {
		BitsDifference(&c1, &c2)
	}
}
func BenchmarkCountingSha7(b *testing.B) {
	c1 := sha256.Sum256([]byte("This"))
	c2 := sha256.Sum256([]byte("That"))
	for i := 0; i < b.N; i++ {
		DifferentBits(c1, c2)
	}
}
