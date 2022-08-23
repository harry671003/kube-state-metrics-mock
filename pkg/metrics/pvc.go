package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*

# HELP kube_persistentvolumeclaim_labels Kubernetes labels converted to Prometheus labels.
# Type kube_persistentvolumeclaim_labels gauge

# HELP kube_persistentvolumeclaim_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_persistentvolumeclaim_annotations gauge

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

# HELP pod_pvc_allocation pod_pvc_allocation Bytes used by a PVC attached to a pod
# TYPE pod_pvc_allocation gauge
pod_pvc_allocation{namespace="cortex",persistentvolume="pvc-0040496c-ee7f-4a6f-995f-6abce1f009e4",persistentvolumeclaim="data-ingester-358",pod="unmounted-pvs"} 5.36870912e+11

*/

var (
	pvcInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_persistentvolumeclaim_info",
		Help: "pvc storage resource requests in bytes",
	}, []string{"namespace", "persistentvolumeclaim", "storageclass", "volumename"})

	pvcCapacityBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_persistentvolume_capacity_bytes",
		Help: "pv storage capacity in bytes",
	}, []string{"persistentvolume"})

	podPVCAllocation = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pod_pvc_allocation",
		Help: "Bytes used by a PVC attached to a pod",
	}, []string{"namespace", "persistentvolume", "persistentvolumeclaim", "pod"})

	pvcResourceRequestsStorageBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_persistentvolumeclaim_resource_requests_storage_bytes",
		Help: "The capacity of storage requested by the persistent volume claim.",
	}, []string{"namespace", "persistentvolumeclaim"})
)

type PVCMetrics struct{}

func (m *PVCMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(pvcInfo)
	r.MustRegister(pvcCapacityBytes)
	r.MustRegister(podPVCAllocation)
	r.MustRegister(pvcResourceRequestsStorageBytes)
}

func (m *PVCMetrics) Update(cluster *cluster.Cluster) {
	for ns, pvcs := range cluster.PVCs {
		for _, pvc := range pvcs {
			podForPVC := util.GetFixedAssignment(pvc.ClaimName, cluster.Pods[ns])
			pvcInfo.WithLabelValues(ns, pvc.ClaimName, "default", pvc.VolumeName).Set(1)
			pvcCapacityBytes.WithLabelValues(pvc.VolumeName).Set(100 * 1024 * 1024 * 1024)
			podPVCAllocation.WithLabelValues(ns, pvc.VolumeName, pvc.ClaimName, podForPVC).Set(util.RandFloatBetween(100*1024*1024*1024, 50*1024*1024*1024))
			pvcResourceRequestsStorageBytes.WithLabelValues(ns, pvc.ClaimName).Set(50 * 1024 * 1024 * 1024)
		}
	}
}
