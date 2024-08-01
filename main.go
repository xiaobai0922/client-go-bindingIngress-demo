package main

import (
	"github.com/client-go-bindingIngress-demo/internal"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	// 1. config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/bailu/.kube/config")
	if err != nil {
		clusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("Error building kubeconfig: ", err)
		}
		config = clusterConfig
	}

	// 2. clientSet
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Error building clientset: ", err)
	}

	// 3. informer
	factory := informers.NewSharedInformerFactory(clientset, 0)
	serviceInformer := factory.Core().V1().Services()
	ingressInformer := factory.Networking().V1().Ingresses()

	controller := internal.NewController(clientset, serviceInformer, ingressInformer)
	stopChan := make(chan struct{})
	factory.Start(stopChan)
	factory.WaitForCacheSync(stopChan)

	controller.Run(stopChan)
}
