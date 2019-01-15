package main

import (
	"testing"
)

func TestReadyString(t *testing.T) {
	str := readyString("This is a test")
	want := "thisisatest"
	if str != want {
		t.Errorf("error: received %v wanted %v.", str, want)
	}
}

func TestAnogram(t *testing.T) {
	s1 := "This is a test?"
	s2 := "Is a test this?"
	res := anagram(s1, s2)
	if !res {
		t.Errorf("error: received %v wanted %v.", res, true)
	}

	s1 = "This is not an anagram."
	s2 = "Is a test this?"
	res = anagram(s1, s2)
	if res {
		t.Errorf("error: received %v wanted %v.", res, false)
	}

	s1 = `the cédille Ç, the accent aigu é, the accent circonflexe â, ê,
		î, ô, û, the accent grave à, è, ù and the accent tréma ë, ï, ü.`
	s2 = `the cédille Ç, the accent aigu é, the accent circonflexe â, ê,
		î, ô, û, the accent grave à, è, ù and the accent tréma ë, ï, ü.`
	res = anagram(s1, s2)
	if !res {
		t.Errorf("error: received %v wanted %v.", res, true)
	}

	s1 = "This is a test �"
	s2 = "This is a test x"
	res = anagram(s1, s2)
	if res {
		t.Errorf("error: received %v wanted %v.", res, true)
	}

	s1 = "This is a test x"
	s2 = "This is a test �"
	res = anagram(s1, s2)
	if res {
		t.Errorf("error: received %v wanted %v.", res, true)
	}
}
