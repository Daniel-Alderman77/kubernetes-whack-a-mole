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
	r.HandleFunc("/deletePod/{podID}", handleDeletePod)

	fmt.Println("About to listen on 8080. Go to http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", logging.Logger(r))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Kubernetes Whack-A-Mole")
}

func handlePodData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, kubernetes.ListPods(w))
}

func handleMolePodData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, kubernetes.ListMolePods(w))
}

func handleDeletePod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	podID := vars["podID"]

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	kubernetes.DeletePod(w, podID)
}
