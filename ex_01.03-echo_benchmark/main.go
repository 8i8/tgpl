package main

import (
	"os"

	"tgpl/ex_01.03-echo_benchmark/echo"
)

func main() {
	println(echo.echo1(os.Args))
	println(echo.echo2(os.Args))
	println(echo.echo3(os.Args))
}
