// not complete
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":80", nil)
	fmt.Println("Listening on port 80")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "code/index.html")
}
