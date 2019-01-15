package main

import (
	"../ex_03.09-fractal_net/fractal"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", fractalHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func fractalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	strMap := r.URL.Query()
	if len(strMap) > 0 {
		if len(strMap.Get("frame")) > 0 {
			frame, _ := strconv.ParseFloat(strMap.Get("frame"), 64)
			if frame >= 0 && frame <= 4 {
				f1.Frame = frame
			}
		}
		if len(strMap.Get("long")) > 0 {
			long, _ := strconv.ParseFloat(strMap.Get("long"), 64)
			if long >= -2 && long <= 2 {
				f1.Long = long
			}
		}
		if len(strMap.Get("lati")) > 0 {
			lati, _ := strconv.ParseFloat(strMap.Get("lati"), 64)
			if lati >= -2 && lati <= 2 {
				f1.Lati = lati
			}
		}
		if len(strMap.Get("smpl")) > 0 {
			smpl, _ := strconv.Atoi(strMap.Get("smpl"))
			if smpl >= 0 && smpl <= 3 {
				f1.Smp_r = smpl
			}
		}
	}
	fmt.Fprintf(w, headHTML)
	fractal.Draw32(w, fractal.Mandelbrot32, f1)
	fmt.Fprintf(w, footHTML)
}

var (
	f1 = fractal.Config{
		Width:  600,
		Height: 600,
		Smp_r:  2,
		Frame:  5E-1,
		Long:   -125E-2,
		Lati:   25E-2,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}
)
