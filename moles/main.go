package main

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Alderman77/kubernetes-whack-a-mole/moles/logging"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", handle)

	fmt.Println("About to listen on 8080. Go to http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", logging.Logger(r))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm a mole")
}
