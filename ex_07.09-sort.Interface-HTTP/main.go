// Exercise 7.9: Use the html/template package (4.6) to replace
// printTracks with a function that displays the tracks as an HTML
// table. Use the solution to the previous exercise to arrange that each
// click on a column head makes an HTTP request to sort the table.
package main

import (
	"fmt"
	"os"
	"os/signal"

	"sortInterface/data"
	"sortInterface/http"
)

func main() {
	go http.Serve(data.Data())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
