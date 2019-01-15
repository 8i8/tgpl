package main

import "fmt"

// With Floating-point literals
const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
)

// With integer literal, using KB as the multiplier
const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1000, KB * KB, MB * KB, GB * KB, TB * GB, PB * KB, EB * KB, ZB * KB
)

// With integer literal, using an extra x const as the multiplier
const (
	x, KB, MB, GB, TB, PB, EB, ZB, YB = 1000, x, x * x, MB * x, GB * x, TB * GB, PB * x, EB * x, ZB * x
)

// With rune literal
const (
	x, KB, MB, GB, TB, PB, EB, ZB, YB = 'Ϩ', x, x * x, MB * x, GB * x, TB * GB, PB * x, EB * x, ZB * x
)

func main() {
	fmt.Println("vim-go")
}
