// ####### EXERCISE ####### //


























































package main

import (
	"fmt"
	"net/http"
	"strings"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You struck gold ! :D I am root")
}

func myName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", strings.TrimLeft(r.URL.String(), "/"))
}

func main() {
	fmt.Println("Listening on :8080")
	http.HandleFunc("/", root)
	http.HandleFunc("/poorva", myName)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
