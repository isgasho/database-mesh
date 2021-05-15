package kubernetes

import (
	"context"
	"log"
	"sync"

	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Client is built upon the real Kubernetes client-go
type Client struct {
	Config    *rest.Config
	ClientSet kubernetes.Interface
}

var (
	c    *Client
	once sync.Once
	err  error
)

// InclusterClientset return a kubernetes client
func InclusterClientset() (*Client, error) {
	once.Do(func() {
		cfg, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln(err)
		}

		clientset, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			log.Fatalln(err)
		}
		c = &Client{
			Config:    cfg,
			ClientSet: clientset,
		}
	})
	return c, nil
}

// GetPod will retrieve a pod from the cluster
func (c *Client) GetPod(namespace, name string) (*v1.Pod, error) {
	return c.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, meta_v1.GetOptions{})
}
