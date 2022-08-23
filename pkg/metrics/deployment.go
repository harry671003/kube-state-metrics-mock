package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	deployCreated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_created",
		Help: "Unix creation timestamp.",
	}, []string{"namespace", "deployment"})
	deployReplicasNum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_status_replicas",
		Help: "The number of replicas per deployment.",
	}, []string{"namespace", "deployment"})
	deployReplicasAvailable = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_status_replicas_available",
		Help: "The number of available replicas per deployment.",
	}, []string{"namespace", "deployment"})
	deployReplicasUnavailable = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_status_replicas_unavailable",
		Help: "The number of available replicas per deployment.",
	}, []string{"namespace", "deployment"})
	deployReplicasStatusUpdated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_status_replicas_updated",
		Help: "The number of updated replicas per deployment.",
	}, []string{"namespace", "deployment"})
	deployReplicasGeneration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_status_observed_generation",
		Help: "The generation observed by the deployment controller.",
	}, []string{"namespace", "deployment"})
	deployStatusCondition = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_status_condition",
		Help: "The current status conditions of a deployment.",
	}, []string{"namespace", "deployment"})
	deploySpecReplicasDesired = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_spec_replicas",
		Help: "Number of desired pods for a deployment.",
	}, []string{"namespace", "deployment"})
	deploySpecPaused = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_spec_paused",
		Help: "Whether the deployment is paused and will not be processed by the deployment controller.",
	}, []string{"namespace", "deployment"})
	deploySpecMaxUnavailable = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_spec_strategy_rollingupdate_max_unavailable",
		Help: "Maximum number of unavailable replicas during a rolling update of a deployment.",
	}, []string{"namespace", "deployment"})
	deploySpecSurge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_spec_strategy_rollingupdate_max_surge",
		Help: "Maximum number of replicas that can be scheduled above the desired number of replicas during a rolling update of a deployment.",
	}, []string{"namespace", "deployment"})
	deployMetaGeneration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_metadata_generation",
		Help: "Sequence number representing a specific generation of the desired state.",
	}, []string{"namespace", "deployment"})
	deployAnnotations = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_annotations",
		Help: "Kubernetes annotations converted to Prometheus labels.",
	}, []string{"namespace", "statefulset"})
	deployLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_deployment_labels",
		Help: "Kubernetes labels converted to Prometheus labels.",
	}, []string{"namespace", "statefulset"})
)

type DeploymentMetrics struct{}

func (m *DeploymentMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(deployCreated)
	r.MustRegister(deployReplicasNum)
	r.MustRegister(deployReplicasAvailable)
	r.MustRegister(deployReplicasUnavailable)
	r.MustRegister(deployReplicasStatusUpdated)
	r.MustRegister(deployReplicasGeneration)
	r.MustRegister(deployStatusCondition)
	r.MustRegister(deploySpecReplicasDesired)
	r.MustRegister(deploySpecPaused)
	r.MustRegister(deploySpecMaxUnavailable)
	r.MustRegister(deploySpecSurge)
	r.MustRegister(deployMetaGeneration)
	r.MustRegister(deployAnnotations)
	r.MustRegister(deployLabels)
}

func (m *DeploymentMetrics) Update(cluster *cluster.Cluster) {
	m.updateCreated(cluster)
	m.updateReplicas(cluster)
}

func (m *DeploymentMetrics) updateCreated(cluster *cluster.Cluster) {
	for ns, statefulsets := range cluster.StatefulSets {
		for _, s := range statefulsets {
			deployCreated.WithLabelValues(ns, s).SetToCurrentTime()
		}
	}
}

func (m *DeploymentMetrics) updateReplicas(cluster *cluster.Cluster) {
	for ns, deployments := range cluster.Deployments {
		for _, d := range deployments {
			totalReplicas := float64(cluster.NumReplicasPerDeployment)
			rand := util.RandBetween(0, totalReplicas)
			unavailable := totalReplicas - rand

			deployReplicasNum.WithLabelValues(ns, d).Set(rand)
			deployReplicasAvailable.WithLabelValues(ns, d).Set(rand)
			deployReplicasUnavailable.WithLabelValues(ns, d).Set(unavailable)
			deployReplicasStatusUpdated.WithLabelValues(ns, d).Set(rand)
			deploySpecPaused.WithLabelValues(ns, d).Set(0)
			deploySpecMaxUnavailable.WithLabelValues(ns, d).Set(rand)
			deploySpecSurge.WithLabelValues(ns, d).Set(rand)
			deployLabels.WithLabelValues(ns, d).Set(1)
			deployAnnotations.WithLabelValues(ns, d).Set(1)
		}
	}
}
