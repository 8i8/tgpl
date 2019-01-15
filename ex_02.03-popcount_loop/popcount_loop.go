package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tgpl/ex_02.03-popcount_loop/popcount"
)

func main() {
	if len(os.Args) > 1 {
		for _, in := range os.Args[1:] {
			num, _ := strconv.ParseFloat(in, 64)
			run(uint64(num))
		}
	} else {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			num, _ := strconv.ParseFloat(in.Text(), 64)
			run(uint64(num))
		}
	}
}

func run(v uint64) {
	fmt.Println(popcount.PopCount(v))
}
