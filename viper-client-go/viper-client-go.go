package main

import (
	"fmt"

	"github.com/spf13/viper"
	apiv1 "k8s.io/api/core/v1"
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

	// Show namespaces
	coreV1Client := clientSet.CoreV1()
	nameSpaces, err := coreV1Client.Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("List of namespaces in cluster:", viper.Get("current-context"))
	for _, n := range nameSpaces.Items {
		fmt.Println(n.Name)
	}

	// Show all deployments in all namespaces
	v1Beta1Client := clientSet.AppsV1beta1()
	deployments, err := v1Beta1Client.Deployments(apiv1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("List of deployments in cluster:", viper.Get("current-context"))
	for _, d := range deployments.Items {
		fmt.Println(d.Name)
	}
}
