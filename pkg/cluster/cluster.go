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
	NumReplicas  int
	Namespaces   []string
	StatefulSets map[string][]string
	Deployments  map[string][]string
	Pods         map[string][]string
	Containers   map[string](map[string][]string)
}

func (c *Cluster) Init(cfg *config.Config) {
	c.NumReplicas = cfg.ClusterConfig.NumReplicas

	c.Namespaces = []string{}
	c.StatefulSets = map[string][]string{}
	c.Deployments = map[string][]string{}
	c.Pods = map[string][]string{}
	c.Containers = map[string]map[string][]string{}

	for i := 0; i < cfg.ClusterConfig.NumNamespaces; i++ {
		ns := c.createString("namespace", i)
		c.Namespaces = append(c.Namespaces, ns)

		c.initStatefulsets(ns, cfg.ClusterConfig.NumStatefulsets)
		c.initDeployments(ns, cfg.ClusterConfig.NumDeployments)
		c.initPods(ns, c.StatefulSets[ns], cfg.ClusterConfig.NumReplicas)
		c.initPods(ns, c.Deployments[ns], cfg.ClusterConfig.NumReplicas)
		c.initContainers(ns, c.Pods[ns], cfg.ClusterConfig.NumContainers)
	}
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

func (c *Cluster) initPods(ns string, deployments []string, numReplicas int) {
	for _, d := range deployments {
		for i := 0; i < numReplicas; i++ {
			c.Pods[ns] = append(c.Pods[ns], fmt.Sprintf("%s-%s-%d", d, "pod", i))
		}
	}
}

func (c *Cluster) initContainers(ns string, pods []string, n int) {
	c.Containers[ns] = map[string][]string{}
	for _, p := range pods {
		for i := 0; i < n; i++ {
			c.Containers[ns][p] = append(c.Containers[ns][p], fmt.Sprintf("%s-%s-%d", p, "container", i))
		}
	}
}

func (c *Cluster) createString(resource string, num int) string {
	return fmt.Sprintf("%s-%d", resource, num)
}
