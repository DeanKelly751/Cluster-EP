package main

import (
	"fmt"
	"log"
	"os"


	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var config *rest.Config
	var err error

	// Check if running inside a Kubernetes cluster
	if _, exists := os.LookupEnv("KUBERNETES_SERVICE_HOST"); exists {
		// Use in-cluster configuration
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Failed to load in-cluster configuration: %v", err)
		}
		fmt.Println("Using in-cluster configuration")
	} else {
		// Fall back to kubeconfig file for local development
		kubeconfigPath := clientcmd.RecommendedHomeFile // Default kubeconfig path (~/.kube/config)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			log.Fatalf("Failed to load kubeconfig from file: %v", err)
		}
		fmt.Println("Using local kubeconfig file")
	}

	// Print the cluster endpoint
	fmt.Printf("Cluster Endpoint: %s\n", config.Host)


}
