package main

import (
	"../ex_02.1-tempconv_kelvin/tempconv"
	"fmt"
)

func main() {

	var c tempconv.Celsius = 100

	f := tempconv.CToF(c)
	k := tempconv.FToK(f)
	c = tempconv.KToC(k)
	k = tempconv.CToK(c)
	f = tempconv.KToF(k)
	c = tempconv.FToC(f)

	fmt.Printf("should be 100 -> %v\n", c)
}
