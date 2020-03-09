package main

import (
	"fmt"
	"github.com/STRRL/sample-controller-rabbit/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)
import "k8s.io/client-go/rest"

func main() {
	fmt.Println("Hi, a lovely Rabbit here.")
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfig
		home, err := os.UserHomeDir()
		kubeconfig := filepath.Join(home, ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
			os.Exit(1)
		}
	}

	rabbitClientSet := versioned.NewForConfigOrDie(config)
	list, err := rabbitClientSet.StrrlV1alpha1().Rabbits("default").List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, item := range list.Items {
		fmt.Printf("%+v\n", item.Spec)
	}
}
