package main

import (
	"context"
	"fmt"

	kubernetesapi "github.com/brewinski/go-forth-and-kube/src/kubernetes-api"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	items, err := kubernetesapi.GetNamespaces(clientset, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, item := range items {
			fmt.Printf("%+v\n", item)
		}
	}

	// podWatcher, err := kubernetesapi.WatchPods(clientset, ctx, "default")
	// if err != nil {
	// 	log.Default().Println(err)
	// 	log.Fatal(err)
	// }

	// deploymentWatcher, err := kubernetesapi.WatchDeployments(clientset, ctx, "default")
	// if err != nil {
	// 	log.Default().Println(err)
	// 	log.Fatal(err)
	// }

	// for {
	// 	select {
	// 	case podEvent := <-podWatcher:
	// 		fmt.Printf("Pod event: %+v\n", podEvent)
	// 	case deploymentEvent := <-deploymentWatcher:
	// 		fmt.Printf("Deployment event: %+v\n", deploymentEvent)
	// 	}
	// }
}
