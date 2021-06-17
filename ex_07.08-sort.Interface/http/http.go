package http

import (
	"fmt"
	"log"
	"net/http"
	"sortInterface/data"
	"sortInterface/data/csort"
	"strings"
	"text/template"
)

var templ = template.Must(template.ParseFiles("assets/templ.gohtml"))

type links struct {
	Href, Tag, Next string
}

func reverse(str string) string {
	if val := strings.Split(str, "-"); len(val) > 1 {
		return val[0]
	}
	return str + "-rev"
}

func table(buf *csort.SortBuffer, tracks []*data.Track) http.HandlerFunc {
	const fname = "http.HandlerFunc: table"
	return func(res http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("search")
		if err != nil && err != http.ErrNoCookie {
			log.Printf("%s: %s", fname, err)
		} else if err == http.ErrNoCookie {
			str := req.URL.RawQuery
			if str != "" {
				buf.Add(str)
			} else {
				buf.Add("title")
			}
		} else if err == nil {
			vars := strings.Split(cookie.Value, ",")
			prev := vars[0]
			next := req.URL.RawQuery
			if next == prev {
				next = reverse(next)
			}
			buf.Add(vars...)
			buf.Add(next)
		}
		cookie = &http.Cookie{
			Name:     "search",
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Value:    buf.String(),
		}
		d := struct {
			Tracks []*data.Track
			Link   links
		}{
			data.Sort(buf, tracks),
			links{
				"\"/stable\"",
				"stable sort",
				"/",
			},
		}
		fmt.Println(cookie.Value)
		http.SetCookie(res, cookie)
		templ.ExecuteTemplate(res, "main", d)
	}
}

func stable(tracks []*data.Track) http.HandlerFunc {
	const fname = "http.HandlerFunc: stable"
	var prev string
	return func(res http.ResponseWriter, req *http.Request) {

		next := req.URL.RawQuery
		if next == "" {
			next = "title"
		}
		if prev == next {
			next = reverse(next)
		}
		d := struct {
			Tracks []*data.Track
			Link   links
		}{
			data.StableSort(tracks, next),
			links{
				"\"/\"",
				"ring buffer sort",
				"/stable",
			},
		}
		templ.ExecuteTemplate(res, "main", d)
		prev = next
	}
}

func Serve(buf *csort.SortBuffer, tracks []*data.Track) {
	http.HandleFunc("/", table(buf, tracks))
	http.HandleFunc("/stable", stable(tracks))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
