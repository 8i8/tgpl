// Exercise 7.12: Change the handler for /list to print its ouptut as an
// HTML table, not test. You may find the html/template package (4.6)
// useful.
package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var templ = template.Must(template.ParseFiles("./assets/templ.gohtml"))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(db.onlineList))
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/read", http.HandlerFunc(db.read))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

type Entity struct {
	Item  string
	Price dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) onlineList(w http.ResponseWriter, req *http.Request) {
	data := make([]Entity, 0, len(db))
	for item, price := range db {
		data = append(data, Entity{item, price})
	}
	w.Header().Set("content-type", "text/html")
	templ.ExecuteTemplate(w, "main", data)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	item := values.Get("item")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}
	price, err := getDollars(w, values.Get("price"))
	if err != nil {
		return
	}
	db[item] = price
	db[item] = price
	fmt.Fprintf(w, "created: %q: %q\n", item, price)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	item := values.Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	price, err := getDollars(w, values.Get("price"))
	if err != nil {
		return
	}
	db[item] = price
	fmt.Fprintf(w, "updated: %q: %q\n", item, price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "deleted: %q\n", item)
}

var errUser = errors.New("user error")
var errParse = errors.New("there is something wrong with the value given")
var errValue = errors.New("value not permitted")

func getDollars(w http.ResponseWriter, v string) (dollars, error) {
	if v == "" {
		w.WriteHeader(http.StatusUnprocessableEntity) // 404
		fmt.Fprint(w, "please specify a price\n")
		return 0, errUser
	}
	price, err := strconv.ParseFloat(v, 32)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		fmt.Fprintf(w, "parse error: %q: %s\n", v, errParse)
		fmt.Fprintf(os.Stderr, "parse failed: %q: %s\n", v, err)
		return 0, errParse
	}
	if price <= 0 {
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		fmt.Fprintf(w, "price: %q: %s\n", v, err)
		return 0, errValue
	}
	return dollars(price), nil
}
