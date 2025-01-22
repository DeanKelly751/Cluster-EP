package main

import (
	"fmt"
	"log"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Load the kubeconfig
	kubeconfigPath := clientcmd.RecommendedHomeFile // Path to the kubeconfig file (usually ~/.kube/config)
	config, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	// Get the current context
	currentContext := config.CurrentContext
	if currentContext == "" {
		log.Fatalf("No current context is set in the kubeconfig")
	}
	fmt.Printf("Current Context: %s\n", currentContext)

	// Get cluster information for the current context
	context := config.Contexts[currentContext]
	if context == nil {
		log.Fatalf("Context %s not found in kubeconfig", currentContext)
	}

	cluster := config.Clusters[context.Cluster]
	if cluster == nil {
		log.Fatalf("Cluster %s not found in kubeconfig", context.Cluster)
	}

	// Print the cluster endpoint
	fmt.Printf("Cluster Endpoint: %s\n", cluster.Server)
}
