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
	NumReplicasPerStatefulset int
	NumReplicasPerDeployment  int
	Nodes                     []string
	Namespaces                []string
	StatefulSets              map[string][]string
	Deployments               map[string][]string
	PVCs                      map[string][]string
	Pods                      map[string][]string
	Containers                map[string](map[string][]string)
	InstanceTypes             []string
}

func (c *Cluster) Init(cfg *config.Config) {
	c.NumReplicasPerStatefulset = cfg.ClusterConfig.NumReplicasPerDeployment

	c.Namespaces = []string{}
	c.StatefulSets = map[string][]string{}
	c.Deployments = map[string][]string{}
	c.PVCs = map[string][]string{}
	c.Pods = map[string][]string{}
	c.Containers = map[string]map[string][]string{}

	c.initNodes(cfg.ClusterConfig.NumNodes)
	c.initInstanceTypes()

	for i := 0; i < cfg.ClusterConfig.NumNamespaces; i++ {
		ns := c.createString("namespace", i)
		c.Namespaces = append(c.Namespaces, ns)

		c.initStatefulsets(ns, cfg.ClusterConfig.NumStatefulsetsPerNamespace)
		c.initDeployments(ns, cfg.ClusterConfig.NumDeploymentsPerNamespace)
		c.initPVCs(ns, cfg.ClusterConfig.NumPVCsPerNamespace)
		c.initPods(ns, c.StatefulSets[ns], cfg.ClusterConfig.NumReplicasPerStatefulset)
		c.initPods(ns, c.Deployments[ns], cfg.ClusterConfig.NumReplicasPerDeployment)
		c.initContainers(ns, c.Pods[ns], cfg.ClusterConfig.NumContainersPerPod)

	}
}

func (c *Cluster) initNodes(n int) {
	for i := 0; i < n; i++ {
		c.Nodes = append(c.Nodes, c.createString("node", i))
	}
}

func (c *Cluster) initInstanceTypes() {
	c.InstanceTypes = []string{
		"m5.2xlarge",
		"m5.xlarge",
		"m5.4xlarge",
		"m5.8xlarge",
		"r5.xlarge",
		"r5.2xlarge",
		"r5.4xlarge",
		"r5.8xlarge",
		"c5.xlarge",
		"c5.2xlarge",
		"c5.4xlarge",
		"c5.8xlarge",
	}
}

func (c *Cluster) initStatefulsets(namespace string, n int) {
	for i := 0; i < n; i++ {
		c.StatefulSets[namespace] = append(c.StatefulSets[namespace], c.createString("statefulset", i))
	}
}

func (c *Cluster) initPVCs(namespace string, n int) {
	for i := 0; i < n; i++ {
		c.PVCs[namespace] = append(c.PVCs[namespace], c.createString("pvc", i))
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
