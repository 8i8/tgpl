package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const prefix = "http://"
const prefixs = "https://"

var ioIn = flag.Bool("I", false, "Accept input from stdio.")

func main() {

	flag.Parse()
	start := time.Now()
	ch := make(chan string)

	if *ioIn {
		count := 0
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() != "" {
				go fetch(input.Text(), ch) // start a goroutine
				count++
			}
		}
		for i := 0; i < count; i++ {
			fmt.Println(<-ch) // receive from channel ch
		}
	} else {
		for _, url := range os.Args[1:] {
			go fetch(url, ch) // start a goroutine
		}
		for range os.Args[1:] {
			fmt.Println(<-ch) // receive from channel ch
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	testcase := strings.ToLower(url)
	if strings.HasPrefix(testcase, prefix) == false &&
		strings.HasPrefix(testcase, prefixs) == false {
		url = prefix + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint("error: ", err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
