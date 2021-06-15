package main

import (
	"sort"
	"testing"
)

func BenchmarkList(b *testing.B) {
	buf := NewSortFnBuffer()
	buf.Add("year")
	buf.Add("title")
	buf.Add("artist")
	for i := 0; i < b.N; i++ {
		sort.Sort(CustomSort{tracks, sortOrder(buf)})
	}
}
