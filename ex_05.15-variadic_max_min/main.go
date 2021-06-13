// Exercise 5.15: Write variadic functions max and min, analogous to
// sum.  What should these functions do when called with no arguments?
// Write variants that require at least one argument.
package main

import (
	"errors"
	"math"
)

var errEmpty = errors.New("no value given")

// max returns the maximum value of all those intergers parsed,
// returning an error if it recieved no values.
func max(vals ...int) (max int, err error) {
	if len(vals) == 0 {
		err = errEmpty
	}
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

// max returns the minimum value of all those intergers parsed,
// returning an error if it recieves no values.
func min(vals ...int) (min int, err error) {
	if len(vals) == 0 {
		err = errEmpty
	}
	min = math.MaxInt64
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return
}

func main() {
}
