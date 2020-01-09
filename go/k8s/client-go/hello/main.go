package main

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// TODO: fix
const configPath = "/home/vagrant/.kube/config"

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Printf("failed to BuildConfigFromFlags: %s\n", err.Error())
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("failed to NewForConfig: %s\n", err.Error())
		os.Exit(1)
	}

	p, err := clientset.CoreV1().Pods("kube-system").Get("kube-scheduler-lab-control-plane", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("failed to get Pod: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%+v\n", p)
}
