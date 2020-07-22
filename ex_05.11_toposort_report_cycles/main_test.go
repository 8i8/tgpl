package main

import (
	"errors"
	"testing"
)

var test graph

func init() {
	test = make(graph)
	test["algorithms"] = buildMap("data structures")
	test["calculus"] = buildMap("linear algebra")
	test["compiler"] = buildMap(
		"data structures",
		"formal languages",
		"computer organization",
	)
	test["data structures"] = buildMap("discrete math")
	test["databases"] = buildMap("data structures")
	test["discrete math"] = buildMap("intro to programming")
	test["formal languages"] = buildMap("discrete math")
	test["networks"] = buildMap("operating systems")
	test["operating systems"] = buildMap(
		"data structures",
		"computer organization")
	test["operating systems"] = buildMap(
		"data structures",
		"computer organization")
	test["programming languages"] = buildMap(
		"data structures",
		"computer organization")
	test["linear algebra"] = buildMap("calculus")
}

func TestToposortmap(t *testing.T) {
	fname := "toposortmap"
	_, err := toposortmap(test)
	if !errors.Is(err, errCycle) {
		t.Errorf(fname+": expected %q received %q", errCycle, err)
	}
}

func TestIsCycle(t *testing.T) {
	fname := "isCycle"
	res := prereqs.isCycle("algorithms", "linear algebra")
	if res {
		t.Errorf(fname+": expected false received %t", res)
	}
	res = prereqs.isCycle("calculus", "linear algebra")
	if !res {
		t.Errorf(fname+": expected true received %t", res)
	}
}
