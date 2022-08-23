package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*

# HELP kube_persistentvolumeclaim_labels Kubernetes labels converted to Prometheus labels.
# Type kube_persistentvolumeclaim_labels gauge

# HELP kube_persistentvolumeclaim_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_persistentvolumeclaim_annotations gauge

# HELP kube_persistentvolumeclaim_info Information about persistent volume claim.
# Type kube_persistentvolumeclaim_info gauge

# HELP kube_persistentvolumeclaim_status_phase The phase the persistent volume claim is currently in.
# Type kube_persistentvolumeclaim_status_phase gauge

# HELP kube_persistentvolumeclaim_resource_requests_storage_bytes The capacity of storage requested by the persistent volume claim.
# Type kube_persistentvolumeclaim_resource_requests_storage_bytes gauge

# HELP kube_persistentvolumeclaim_access_mode The access mode(s) specified by the persistent volume claim.
# Type kube_persistentvolumeclaim_access_mode gauge

# HELP kube_persistentvolumeclaim_status_condition Information about status of different conditions of persistent volume claim.
# Type kube_persistentvolumeclaim_status_condition gauge

# HELP kube_persistentvolume_claim_ref Information about the Persistent Volume Claim Reference.
# Type kube_persistentvolume_claim_ref gauge

*/

var (
	pvcInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_persistentvolumeclaim_info",
		Help: "pvc storage resource requests in bytes",
	}, []string{"namespace", "persistentvolumeclaim", "storageclass", "volumename"})
)

type PVCMetrics struct{}

func (m *PVCMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(pvcInfo)
}

func (m *PVCMetrics) Update(cluster *cluster.Cluster) {
	for ns, pvcs := range cluster.PVCs {
		for _, pvc := range pvcs {
			pvcInfo.WithLabelValues(ns, pvc, "default", pvc).Set(1)
		}
	}
}
