package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*

# HELP kube_namespace_created Unix creation timestamp
# Type kube_namespace_created gauge

# HELP kube_namespace_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_namespace_annotations gauge

# HELP kube_namespace_status_phase kubernetes namespace status phase.
# Type kube_namespace_status_phase gauge

# HELP kube_namespace_status_condition The condition of a namespace.
# Type kube_namespace_status_condition gauge

*/

var (
	namespaceLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_namespace_labels",
		Help: "Kubernetes labels converted to Prometheus labels.",
	}, []string{"namespace"})
)

type NamespaceMetrics struct{}

func (m *NamespaceMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(namespaceLabels)
}

func (m *NamespaceMetrics) Update(cluster *cluster.Cluster) {
	for _, ns := range cluster.Namespaces {
		namespaceLabels.WithLabelValues(ns).Set(1)
	}
}
