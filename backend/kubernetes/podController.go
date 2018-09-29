package kubernetes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type podData struct {
	Name string `json:"name"`
}

func createClient() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

// ListPods returns a list of pods running within the cluster.
func ListPods(w http.ResponseWriter) string {
	clientset := createClient()

	var podList []podData

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Panic(err.Error())
	}
	log.Printf("There are %d pods in the cluster\n", len(pods.Items))

	var podData podData
	for i := range pods.Items {
		podData.Name = pods.Items[i].Name

		podList = append(podList, podData)
	}

	json, err := json.Marshal(podList)
	if err != nil {
		log.Panic(err.Error())
	}

	result := string(json)
	fmt.Println(result)

	return result
}

// ListMolePods returns a list of test mole pods running within the cluster.
func ListMolePods(w http.ResponseWriter) string {
	clientset := createClient()

	var podList []podData

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Panic(err.Error())
	}
	log.Printf("There are %d pods in the cluster\n", len(pods.Items))

	var podData podData
	for i := range pods.Items {
		if strings.Contains(pods.Items[i].Name, "test-moles") {
			podData.Name = pods.Items[i].Name

			podList = append(podList, podData)
		}
	}

	json, err := json.Marshal(podList)
	if err != nil {
		log.Panic(err.Error())
	}

	result := string(json)
	fmt.Println(result)

	return result
}
