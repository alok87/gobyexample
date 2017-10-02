package main

import (
	"fmt"

	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.kube")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading  %s", err))
	}
	// Get the kubeclient version stored in .kube/config and show it
	fmt.Println("Kubeconfig version:", viper.Get("apiVersion"))

	// Make the clinet restConfig from the kubeconfig file
	clientConfig, err := clientcmd.BuildConfigFromFlags("", viper.ConfigFileUsed())
	if err != nil {
		panic(err)
	}

	// Create a clientset
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err)
	}

	// Create a client for namespace
	coreV1Client := clientSet.CoreV1()
	nameSpaces, err := coreV1Client.Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// Show all namespaces
	fmt.Println("List of namespaces in cluster:", viper.Get("current-context"))
	for _, n := range nameSpaces.Items {
		fmt.Println(n.Name)
	}

}
