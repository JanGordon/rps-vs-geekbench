package main

import (
	"fmt"
	"html"
	"log"
	"math"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out := ""
		for i := 0; i < 100; i++ {
			out += fmt.Sprint(math.Sqrt(50 * 12566))
		}
		fmt.Fprintf(w, out, html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
