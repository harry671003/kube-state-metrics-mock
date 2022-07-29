package cluster

import (
	"fmt"

	"github.com/harry671003/KubeStateMetricsMock/pkg/config"
)

func NewCluster(cfg *config.Config) *Cluster {
	c := &Cluster{}
	c.Init(cfg)

	return c
}

// Defines a K8s cluster
type Cluster struct {
	Namespaces   []string
	StatefulSets map[string][]string
	Deployments  map[string][]string
	Pods         []string
}

func (c *Cluster) Init(cfg *config.Config) {
	c.Namespaces = []string{}
	c.StatefulSets = map[string][]string{}
	c.Deployments = map[string][]string{}
	for i := 0; i < cfg.ClusterConfig.NumNamespaces; i++ {
		ns := c.createString("namespace", i)
		c.Namespaces = append(c.Namespaces, ns)

		c.initStatefulsets(ns, cfg.ClusterConfig.NumStatefulsets)
		c.initDeployments(ns, cfg.ClusterConfig.NumDeployments)
	}

	c.initPods(cfg.ClusterConfig.NumPods)
}

func (c *Cluster) initStatefulsets(namespace string, n int) {
	for i := 0; i < n; i++ {
		c.StatefulSets[namespace] = append(c.StatefulSets[namespace], c.createString("statefulset", i))
	}
}

func (c *Cluster) initDeployments(namespace string, n int) {
	for i := 0; i < n; i++ {
		c.Deployments[namespace] = append(c.Deployments[namespace], c.createString("deploy", i))
	}
}

func (c *Cluster) initPods(numReplicas int) {
	for i := 0; i < numReplicas; i++ {
		c.Pods = append(c.Pods, c.createString("pod", i))
	}
}

func (c *Cluster) createString(resource string, num int) string {
	return fmt.Sprintf("%s-%d", resource, num)
}
