package kubernetesapi

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

func GetDeployments(clientset *kubernetes.Clientset, ctx context.Context,
	namespace string) ([]v1.Deployment, error) {

	list, err := clientset.AppsV1().Deployments(namespace).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func GetNamespaces(clientset *kubernetes.Clientset, ctx context.Context) ([]string, error) {
	list, err := clientset.CoreV1().Namespaces().
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var namespaces []string
	for _, item := range list.Items {
		namespaces = append(namespaces, item.Name)
	}
	return namespaces, nil
}

func WatchPods(clientset *kubernetes.Clientset, ctx context.Context,
	namespace string) (<-chan watch.Event, error) {

	watcher, err := clientset.CoreV1().Pods(namespace).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return watcher.ResultChan(), nil
}

func WatchDeployments(clientset *kubernetes.Clientset, ctx context.Context,
	namespace string) (<-chan watch.Event, error) {

	watcher, err := clientset.AppsV1().Deployments(namespace).Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return watcher.ResultChan(), nil
}
