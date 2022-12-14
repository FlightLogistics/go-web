package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<H1>Good morning this is the goweb Service</H1>")
		fmt.Fprintf(w, "<p>It runs on www.heliflyer.org.uk</p>")
		fmt.Fprintf(w, "<p>I added another line</p>")
		fmt.Fprintf(w, "<p>I added another line again</p>")

		fmt.Fprintf(w, "<p>I added another line again now</p>")
		fmt.Fprintf(w, "<p>incredible</p>")
		fmt.Fprintf(w, "<p>Can this go on?</p>")
		fmt.Fprintf(w, "<p>Not sure</p>")
		fmt.Fprintf(w, "<p>probably?</p>")
		fmt.Fprintf(w, "<p>probably the latest?</p>")
	})

	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/greet/"):]
		fmt.Fprintf(w, "Hello %s\n", name)
	})

	http.ListenAndServe(":9990", nil)
}
