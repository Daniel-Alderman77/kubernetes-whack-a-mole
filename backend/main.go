package main

import (
	"fmt"
	"net/http"

	"github.com/Daniel-Alderman77/kubernetes-whack-a-mole/backend/kubernetes"
	"github.com/Daniel-Alderman77/kubernetes-whack-a-mole/backend/logging"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handle)
	r.HandleFunc("/getPods", handlePodData)
	r.HandleFunc("/getMolePods", handleMolePodData)

	fmt.Println("About to listen on 8080. Go to http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", logging.Logger(r))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	fmt.Fprintln(w, "Kubernetes Whack-A-Mole")
}

func handlePodData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, kubernetes.ListPods(w))
}

func handleMolePodData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, kubernetes.ListMolePods(w))
}
