package main

import (
	"fmt"

	"github.com/spf13/viper"
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

}
