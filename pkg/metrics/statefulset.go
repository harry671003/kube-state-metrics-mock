package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ssCreated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_created",
		Help: "Unix creation timestamp.",
	}, []string{"namespace", "statefulset"})
	ssReplicasNum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_replicas",
		Help: "The number of replicas per StatefulSet.",
	}, []string{"namespace", "statefulset"})
	ssReplicasAvailable = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_replicas_available",
		Help: "The number of available replicas per StatefulSet.",
	}, []string{"namespace", "statefulset"})
	ssReplicasCurrent = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_replicas_current",
		Help: "The number of current replicas per StatefulSet.",
	}, []string{"namespace", "statefulset"})
	ssReplicasReady = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_replicas_ready",
		Help: "The number of ready replicas per StatefulSet.",
	}, []string{"namespace", "statefulset"})
	ssReplicasUpdated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_replicas_updated",
		Help: "The number of updated replicas per StatefulSet.",
	}, []string{"namespace", "statefulset"})
	ssReplicasGeneration = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_observed_generation",
		Help: "The generation observed by the StatefulSet controller.",
	})
	ssReplicasDesired = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_statefulset_replicas",
		Help: "Number of desired pods for a StatefulSet.",
	})
	ssMetaGeneration = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_statefulset_metadata_generation",
		Help: "Sequence number representing a specific generation of the desired state for the StatefulSet.",
	})
	ssAnnotations = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_annotations",
		Help: "Kubernetes annotations converted to Prometheus labels.",
	}, []string{"namespace", "statefulset"})
	ssLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_statefulset_labels",
		Help: "Kubernetes labels converted to Prometheus labels.",
	}, []string{"namespace", "statefulset"})
	ssCurRevision = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_current_revision",
		Help: "Indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).",
	})
	ssUpdateRevision = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_statefulset_status_update_revision",
		Help: "Indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)",
	})
)

type StatefulSetMetrics struct{}

func (m *StatefulSetMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(ssCreated)
	r.MustRegister(ssReplicasNum)
	r.MustRegister(ssReplicasAvailable)
	r.MustRegister(ssReplicasCurrent)
	r.MustRegister(ssReplicasReady)
	r.MustRegister(ssReplicasUpdated)
	r.MustRegister(ssReplicasGeneration)
	r.MustRegister(ssReplicasDesired)
	r.MustRegister(ssMetaGeneration)
	r.MustRegister(ssAnnotations)
	r.MustRegister(ssLabels)
	r.MustRegister(ssCurRevision)
	r.MustRegister(ssUpdateRevision)
}

func (m *StatefulSetMetrics) Update(cluster *cluster.Cluster) {
	m.updateCreated(cluster)
	m.updateReplicas(cluster)
}

func (m *StatefulSetMetrics) updateCreated(cluster *cluster.Cluster) {
	for ns, statefulsets := range cluster.StatefulSets {
		for _, s := range statefulsets {
			ssCreated.WithLabelValues(ns, s).SetToCurrentTime()
		}
	}
}

func (m *StatefulSetMetrics) updateReplicas(cluster *cluster.Cluster) {
	for ns, statefulsets := range cluster.StatefulSets {
		for _, s := range statefulsets {
			totalReplicas := float64(cluster.NumReplicasPerStatefulset)
			rand := util.RandBetween(totalReplicas, 0)
			ssReplicasNum.WithLabelValues(ns, s).Set(rand)
			ssReplicasAvailable.WithLabelValues(ns, s).Set(rand)
			ssReplicasCurrent.WithLabelValues(ns, s).Set(rand)
			ssReplicasReady.WithLabelValues(ns, s).Set(rand)
			ssReplicasUpdated.WithLabelValues(ns, s).Set(rand)
			ssLabels.WithLabelValues(ns, s).Set(1)
			ssAnnotations.WithLabelValues(ns, s).Set(1)
		}
	}
}
