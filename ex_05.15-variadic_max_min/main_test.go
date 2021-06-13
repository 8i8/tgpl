package main

import "testing"

func TestMax(t *testing.T) {
	const fname = "TestMax"
	r, err := max(10, 23, 2, 50, 145, 49)
	if err != nil {
		t.Errorf("%s: did not expect and error: %s", fname, err)
	}
	if r != 145 {
		t.Errorf("%s: got %d want 145", fname, r)
	}
	_, err = max()
	if err != errEmpty {
		err := "(*errors.errorString:\"no value given\")"
		t.Errorf("%s: got (%+T:%q) want %s", fname, errEmpty, errEmpty, err)
	}
}

func TestMin(t *testing.T) {
	const fname = "TestMin"
	r, err := min(10, 23, 2, 50, 145, 49)
	if err != nil {
		t.Errorf("%s: did not expect and error: %s", fname, err)
	}
	if r != 2 {
		t.Errorf("%s: got %d want 145", fname, r)
	}
	_, err = min()
	if err != errEmpty {
		err := "(*errors.errorString:\"no value given\")"
		t.Errorf("%s: got (%+T:%q) want %s", fname, errEmpty, errEmpty, err)
	}
}
