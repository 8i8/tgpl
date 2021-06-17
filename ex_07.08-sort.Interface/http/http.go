package http

import (
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

func table(buf *csort.SortBuffer, tracks []*data.Track) http.HandlerFunc {
	const fname = "http.HandlerFunc: table"
	return func(res http.ResponseWriter, req *http.Request) {

		cookie, err := req.Cookie("search")
		if err != nil && err != http.ErrNoCookie {
			log.Printf("%s: %s", fname, err)
		} else if err == http.ErrNoCookie {
			buf.Add(req.URL.RawQuery)
		} else if err == nil {
			buf.Load(strings.Split(cookie.Value, ",")...)
			buf.Add(req.URL.RawQuery)
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
		http.SetCookie(res, cookie)
		templ.ExecuteTemplate(res, "main", d)
	}
}

func reverse(str string) string {
	if val := strings.Split(str, "-"); len(val) > 1 {
		return val[0]
	}
	return str + "-rev"
}

func stable(tracks []*data.Track) http.HandlerFunc {
	const fname = "http.HandlerFunc: stable"
	var prev string
	return func(res http.ResponseWriter, req *http.Request) {

		next := req.URL.RawQuery
		if prev == next {
			next = reverse(prev)
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

func destroyCookie() http.HandlerFunc {
	const fname = "destroyCookie"
	return func(res http.ResponseWriter, req *http.Request) {
		cookie := &http.Cookie{
			Name:     "search",
			MaxAge:   -1,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Value:    "",
		}
		http.SetCookie(res, cookie)
		templ.ExecuteTemplate(res, "main", nil)
	}
}

// Serve startes the server.
func Serve(buf *csort.SortBuffer, tracks []*data.Track) {
	http.HandleFunc("/", table(buf, tracks))
	http.HandleFunc("/stable", stable(tracks))
	http.HandleFunc("/cookie", destroyCookie())
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
