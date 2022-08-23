package metrics

import (
	"strings"

	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*

# HELP kube_pod_info Information about pod.
# Type kube_pod_info gauge

# HELP kube_pod_init_container_info Information about an init container in a pod.
# Type kube_pod_init_container_info gauge

# HELP kube_pod_info Information about pod.
# Type kube_pod_info gauge

# HELP kube_pod_init_container_info Information about an init container in a pod.
# Type kube_pod_init_container_info gauge

# HELP kube_pod_init_container_resource_limits The number of requested limit resource by an init container.
# Type kube_pod_init_container_resource_limits gauge

# HELP kube_pod_init_container_resource_requests The number of requested request resource by an init container.
# Type kube_pod_init_container_resource_requests gauge

# HELP kube_pod_init_container_status_last_terminated_reason Describes the last reason the init container was in terminated state.
# Type kube_pod_init_container_status_last_terminated_reason gauge

# HELP kube_pod_init_container_status_ready Describes whether the init containers readiness check succeeded.
# Type kube_pod_init_container_status_ready gauge

# HELP kube_pod_init_container_status_restarts_total The number of restarts for the init container.
# Type kube_pod_init_container_status_restarts_total counter

# HELP kube_pod_init_container_status_running Describes whether the init container is currently in running state.
# Type kube_pod_init_container_status_running gauge

# HELP kube_pod_init_container_status_terminated Describes whether the init container is currently in terminated state.
# Type kube_pod_init_container_status_terminated gauge

# HELP kube_pod_init_container_status_terminated_reason Describes the reason the init container is currently in terminated state.
# Type kube_pod_init_container_status_terminated_reason gauge

# HELP kube_pod_init_container_status_waiting Describes whether the init container is currently in waiting state.
# Type kube_pod_init_container_status_waiting gauge

# HELP kube_pod_init_container_status_waiting_reason Describes the reason the init container is currently in waiting state.
# Type kube_pod_init_container_status_waiting_reason gauge

# HELP kube_pod_overhead_cpu_cores The pod overhead in regards to cpu cores associated with running a pod.
# Type kube_pod_overhead_cpu_cores gauge

# HELP kube_pod_overhead_memory_bytes The pod overhead in regards to memory associated with running a pod.
# Type kube_pod_overhead_memory_bytes gauge



# HELP kube_pod_restart_policy Describes the restart policy in use by this pod.
# Type kube_pod_restart_policy gauge

# HELP kube_pod_runtimeclass_name_info The runtimeclass associated with the pod.
# Type kube_pod_runtimeclass_name_info gauge

# HELP kube_pod_spec_volumes_persistentvolumeclaims_info Information about persistentvolumeclaim volumes in a pod.
# Type kube_pod_spec_volumes_persistentvolumeclaims_info gauge

# HELP kube_pod_spec_volumes_persistentvolumeclaims_readonly Describes whether a persistentvolumeclaim is mounted read only.
# Type kube_pod_spec_volumes_persistentvolumeclaims_readonly gauge

# HELP kube_pod_start_time Start time in unix timestamp for a pod.
# Type kube_pod_start_time gauge

# HELP kube_pod_status_phase The pods current phase.
# Type kube_pod_status_phase gauge

# HELP kube_pod_status_ready Describes whether the pod is ready to serve requests.
# Type kube_pod_status_ready gauge

# HELP kube_pod_status_reason The pod status reasons
# Type kube_pod_status_reason gauge

# HELP kube_pod_status_scheduled Describes the status of the scheduling process for the pod.
# Type kube_pod_status_scheduled gauge

# HELP kube_pod_status_scheduled_time Unix timestamp when pod moved into scheduled status
# Type kube_pod_status_scheduled_time gauge

# HELP kube_pod_status_unschedulable Describes the unschedulable status for the pod.
# Type kube_pod_status_unschedulable gauge

*/
var (
	podCompletionTime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_completion_time",
		Help: "Completion time in unix timestamp for a pod.",
	}, []string{"namespace", "pod"})
	podContainerInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_info",
		Help: "Information about a container in a pod.",
	}, []string{"namespace", "pod", "container"})
	podContainerResourceLimits = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_resource_limits",
		Help: "The number of requested limit resource by a container.",
	}, []string{"namespace", "pod", "container", "resource", "unit", "node"})
	podContainerResourceRequests = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_resource_requests",
		Help: "The number of requested request resource by a container.",
	}, []string{"namespace", "pod", "container", "resource", "unit", "node"})
	podContainerStateStarted = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_state_started",
		Help: "Start time in unix timestamp for a pod container.",
	}, []string{"namespace", "pod", "container"})
	podContainerLastTerminatedReason = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_last_terminated_reason",
		Help: "Describes the last reason the container was in terminated state.",
	}, []string{"namespace", "pod", "container", "reason"})
	podContainerStatusReady = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_ready",
		Help: "Describes whether the containers readiness check succeeded.",
	}, []string{"namespace", "pod", "container"})
	podContainerStatusRestartsTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_restarts_total",
		Help: "The number of container restarts per container.",
	}, []string{"namespace", "pod", "container", "reason"})
	podContainerStatusRunning = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_running",
		Help: "Describes whether the container is currently in running state.",
	}, []string{"namespace", "pod", "container", "uid"})
	podContainerStatusTerminated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_terminated",
		Help: "Describes whether the container is currently in terminated state.",
	}, []string{"namespace", "pod", "container"})
	podContainerStatusTerminatedReason = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_terminated_reason",
		Help: "Describes the reason the container is currently in terminated state.",
	}, []string{"namespace", "pod", "container", "reason"})
	podContainerStatusWaiting = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_waiting",
		Help: "Describes whether the container is currently in waiting state.",
	}, []string{"namespace", "pod", "container"})
	podContainerStatusWaitingReason = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_container_status_waiting_reason",
		Help: "Describes the reason the container is currently in waiting state.",
	}, []string{"namespace", "pod", "container", "reason"})
	podCreated = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_created",
		Help: "Unix creation timestamp",
	}, []string{"namespace", "pod"})
	podDeletionTimestamp = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_deletion_timestamp",
		Help: "Unix deletion timestamp",
	}, []string{"namespace", "pod"})
	podLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_labels",
		Help: "Kubernetes labels converted to Prometheus labels.",
	}, []string{"namespace", "pod"})
	podAnnotations = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_annotations",
		Help: "Kubernetes annotations converted to Prometheus labels.",
	}, []string{"namespace", "pod"})
	podOwner = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kube_pod_owner",
		Help: "Information about the Pod's owner.",
	}, []string{"namespace", "owner_is_controller", "owner_name", "pod", "owner_kind"})
)

type PodMetrics struct{}

func (m *PodMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(podCompletionTime)
	r.MustRegister(podContainerInfo)
	r.MustRegister(podContainerLastTerminatedReason)
	r.MustRegister(podContainerResourceLimits)
	r.MustRegister(podContainerResourceRequests)
	r.MustRegister(podContainerStateStarted)
	r.MustRegister(podContainerStatusRestartsTotal)
	r.MustRegister(podContainerStatusRunning)
	r.MustRegister(podContainerStatusReady)
	r.MustRegister(podContainerStatusTerminated)
	r.MustRegister(podContainerStatusTerminatedReason)
	r.MustRegister(podContainerStatusWaiting)
	r.MustRegister(podContainerStatusWaitingReason)
	r.MustRegister(podCreated)
	r.MustRegister(podDeletionTimestamp)
	r.MustRegister(podLabels)
	r.MustRegister(podAnnotations)
	r.MustRegister(podOwner)
}

func (m *PodMetrics) Update(cluster *cluster.Cluster) {
	m.updatePodMetrics(cluster)
	m.updateContainerMetrics(cluster)
}

func (m *PodMetrics) updatePodMetrics(cluster *cluster.Cluster) {
	for ns, pods := range cluster.Pods {
		for _, p := range pods {
			podCreated.WithLabelValues(ns, p).SetToCurrentTime()
			podDeletionTimestamp.WithLabelValues(ns, p).SetToCurrentTime()
			podCompletionTime.WithLabelValues(ns, p).SetToCurrentTime()
		}
	}
}

func (m *PodMetrics) updateContainerMetrics(cluster *cluster.Cluster) {
	oom := "OOMKilled"
	for ns, pods := range cluster.Containers {
		for p, containers := range pods {
			for _, c := range containers {
				uid := util.SHA1(c)
				node := util.GetFixedAssignment(p, cluster.Nodes)
				podContainerInfo.WithLabelValues(ns, p, c).Set(1)
				podContainerLastTerminatedReason.WithLabelValues(ns, p, c, oom).Set(1)
				podContainerResourceLimits.WithLabelValues(ns, p, c, "cpu", "core", node).Set(1)
				podContainerResourceLimits.WithLabelValues(ns, p, c, "memory", "byte", node).Set(1)
				podContainerResourceLimits.WithLabelValues(ns, p, c, "nvidia_com_gpu", "core", node).Set(1)
				podContainerResourceRequests.WithLabelValues(ns, p, c, "cpu", "core", node).Set(1)
				podContainerResourceRequests.WithLabelValues(ns, p, c, "nvidia_com_gpu", "core", node).Set(1)
				podContainerResourceRequests.WithLabelValues(ns, p, c, "memory", "byte", node).Set(1)
				podContainerStateStarted.WithLabelValues(ns, p, c).SetToCurrentTime()
				podContainerStatusRestartsTotal.WithLabelValues(ns, p, c, oom).Set(util.RandBetween(2, 0))
				podContainerStatusRunning.WithLabelValues(ns, p, c, uid).Set(1)
				podContainerStatusTerminated.WithLabelValues(ns, p, c).Set(1)
				podContainerStatusTerminatedReason.WithLabelValues(ns, p, c, oom).Set(1)
				podContainerStatusWaiting.WithLabelValues(ns, p, c).Set(1)
				podContainerStatusWaitingReason.WithLabelValues(ns, p, c, oom).Set(1)
				if strings.Contains(p, "statefulset") {
					podOwner.WithLabelValues(ns, "true", p, p, "DaemonSet")
				} else {
					podOwner.WithLabelValues(ns, "true", p, p, "ReplicaSet")
				}
				podAnnotations.WithLabelValues(ns, p).Set(1)
				podLabels.WithLabelValues(ns, p).Set(1)
			}
		}
	}
}
