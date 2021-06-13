package main

import "testing"

func TestJoin(t *testing.T) {
	const fname = "TestJoin"
	str := join(":")
	if str != "" {
		t.Errorf("%s: want \"\" got %q", fname, str)
	}
	str = join("", "hello")
	if str != "hello" {
		t.Errorf("%s: want \"hello\" got %q", fname, str)
	}
	str = join(" ", "hello", "world")
	if str != "hello world" {
		t.Errorf("%s: want \"hello world\" got %q", fname, str)
	}
}
