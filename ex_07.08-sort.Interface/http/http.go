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
		http.SetCookie(res, cookie)
		templ.ExecuteTemplate(res, "main", data.Sort(buf, tracks))
	}
}

func Serve(buf *csort.SortBuffer, tracks []*data.Track) {
	http.HandleFunc("/", table(buf, tracks))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
