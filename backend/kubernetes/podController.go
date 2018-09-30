package kubernetes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type podData struct {
	Name      string      `json:"name"`
	StartTime int64       `json:"startTime"`
	PodIP     string      `json:"podIP"`
	Phase     v1.PodPhase `json:"phase"`
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

func getNameSpace(podID string) string {
	clientset := createClient()

	var podNameSpace string

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Panic(err.Error())
	}

	for i := range pods.Items {
		if pods.Items[i].Name == podID {
			podNameSpace = pods.Items[i].Namespace
		}
	}
	log.Printf("The pod %s is in the %s\n", podID, podNameSpace)

	return podNameSpace
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
		podData.StartTime = pods.Items[i].Status.StartTime.Unix()
		podData.PodIP = pods.Items[i].Status.PodIP
		podData.Phase = pods.Items[i].Status.Phase

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
			podData.StartTime = pods.Items[i].Status.StartTime.Unix()
			podData.PodIP = pods.Items[i].Status.PodIP
			podData.Phase = pods.Items[i].Status.Phase

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

// DeletePod deletes a specified pod within the cluster.
func DeletePod(w http.ResponseWriter, podID string) {
	clientset := createClient()
	podNameSpace := getNameSpace(podID)

	if len(podNameSpace) == 0 {
		log.Printf("Pod not found\n")
	} else {
		err := clientset.CoreV1().Pods(podNameSpace).Delete(podID, &metav1.DeleteOptions{})
		if err != nil {
			log.Panic(err.Error())
		}
		log.Printf("Pod %v was deleted\n", podID)
	}
}
