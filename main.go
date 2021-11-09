package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	var port string = ":8087"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello Rafael!!! ARGOCD</h1>"))
	})
	fmt.Printf("Web Service is Running. Port: %s", strings.Replace(port, ":", "", 1))
	http.ListenAndServe(port, nil)

}
