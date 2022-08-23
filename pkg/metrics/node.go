package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*

# HELP kube_node_created Unix creation timestamp
# Type kube_node_created gauge

# HELP kube_node_info Information about a cluster node.
# Type kube_node_info gauge

# HELP kube_node_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_node_annotations gauge

# HELP kube_node_labels Kubernetes labels converted to Prometheus labels.
# Type kube_node_labels gauge

# HELP kube_node_role The role of a cluster node.
# Type kube_node_role gauge

# HELP kube_node_spec_taint The taint of a cluster node.
# Type kube_node_spec_taint gauge

# HELP kube_node_spec_unschedulable Whether a node can schedule new pods.
# Type kube_node_spec_unschedulable gauge

# HELP kube_node_status_allocatable The allocatable for different resources of a node that are available for scheduling.
# Type kube_node_status_allocatable gauge

# HELP kube_node_status_capacity The capacity for different resources of a node.
# Type kube_node_status_capacity gauge

# HELP kube_node_status_condition The condition of a cluster node.
# Type kube_node_status_condition gauge
*/

var (
	nodeLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_node_labels",
		Help: "Kubernetes labels converted to Prometheus labels.",
	}, []string{"node", "label_beta_kubernetes_io_instance_type", "label_eks_amazonaws_com_capacityType"})
)

type NodeMetrics struct{}

func (m *NodeMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(nodeLabels)
}

func (m *NodeMetrics) Update(cluster *cluster.Cluster) {
	for _, n := range cluster.Nodes {
		nodeLabels.WithLabelValues(n, "r5.2xlarge", "ON_DEMAND").Set(1)
	}
}
