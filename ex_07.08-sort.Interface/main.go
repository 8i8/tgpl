package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"sortInterface/data"
	"sortInterface/http"
)

func main() {

	log.SetFlags(log.Llongfile)
	go http.Serve(data.Data())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
