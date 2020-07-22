package main

import "testing"

func index(res []string, class string) int {
	for i := range res {
		if res[i] == class {
			return i + 1
		}
	}
	return 0
}

func TestToposortmap(t *testing.T) {
	fname := "toposortmap"
	res := toposortmap(prereqs)
	for class, items := range prereqs {
		i := index(res, class)
		for prereq, _ := range items {
			j := index(res, prereq)
			if j > i {
				t.Errorf(fname+
					": %q can not occur before %q",
					prereq, class)
			}
		}
	}
}
