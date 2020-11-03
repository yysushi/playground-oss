package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ping from %s\n", r.RemoteAddr)
		fmt.Fprintf(w, "pong\n")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
