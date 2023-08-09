package main

import (
	"context"

	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	namespace := "telepresence"

	// items, err := kubernetesapi.GetPods(clientset, ctx, namespace)
}
