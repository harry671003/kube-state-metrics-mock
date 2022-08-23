package metrics

import (
	"fmt"

	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*



# HELP kube_node_labels kube_node_labels all labels for each node prefixed with label_
# TYPE kube_node_labels gauge

# HELP kube_node_status_allocatable kube_node_status_allocatable node allocatable
# TYPE kube_node_status_allocatable gauge

# HELP kube_node_status_allocatable_cpu_cores kube_node_status_allocatable_cpu_cores node allocatable cpu cores
# TYPE kube_node_status_allocatable_cpu_cores gauge

# HELP kube_node_status_allocatable_memory_bytes kube_node_status_allocatable_memory_bytes node allocatable memory in bytes
# TYPE kube_node_status_allocatable_memory_bytes gauge

# HELP kube_node_status_capacity kube_node_status_capacity node capacity
# TYPE kube_node_status_capacity gauge

# HELP kube_node_status_capacity_cpu_cores kube_node_status_capacity_cpu_cores Node Capacity CPU Cores
# TYPE kube_node_status_capacity_cpu_cores gauge

# HELP kube_node_status_capacity_memory_bytes kube_node_status_capacity_memory_bytes Node Capacity Memory Bytes
# TYPE kube_node_status_capacity_memory_bytes gauge

# HELP kube_node_status_condition kube_node_status_condition condition status for nodes
# TYPE kube_node_status_condition gauge

# HELP kube_persistentvolume_status_phase kube_persistentvolume_status_phase pv status phase
# TYPE kube_persistentvolume_status_phase gauge

# HELP kube_persistentvolumeclaim_resource_requests_storage_bytes kube_persistentvolumeclaim_resource_requests_storage_bytes pvc storage resource requests in bytes
# TYPE kube_persistentvolumeclaim_resource_requests_storage_bytes gauge

# HELP kube_pod_container_resource_limits kube_pod_container_resource_limits pods container resource limits
# TYPE kube_pod_container_resource_limits gauge

# HELP kube_pod_container_resource_limits_cpu_cores kube_pod_container_resource_limits_cpu_cores pods container cpu cores resource limits
# TYPE kube_pod_container_resource_limits_cpu_cores gauge

# HELP kube_pod_container_resource_limits_memory_bytes kube_pod_container_resource_limits_memory_bytes pods container memory bytes resource limits
# TYPE kube_pod_container_resource_limits_memory_bytes gauge

# HELP kube_pod_container_status_restarts_total kube_pod_container_status_restarts_total total container restarts
# TYPE kube_pod_container_status_restarts_total counter

# HELP kube_pod_labels kube_pod_labels all labels for each pod prefixed with label_
# TYPE kube_pod_labels gauge

# HELP kube_pod_owner kube_pod_owner Information about the Pod's owner
# TYPE kube_pod_owner gauge

# HELP kube_pod_status_phase kube_pod_container_status_terminated_reason Describes the reason the container is currently in terminated state.
# TYPE kube_pod_status_phase gauge

# HELP kubecost_cluster_info kubecost_cluster_info ClusterInfo
# TYPE kubecost_cluster_info gauge

#1",productanalytics="true",provider="AWS",provisioner="EKS",region="us-west-2",remotereadenabled="false",thanosenabled="false",valuesreporting="true",version="1.21+"} 1
# HELP kubecost_cluster_management_cost kubecost_cluster_management_cost Hourly cost paid as a cluster management fee.
# TYPE kubecost_cluster_management_cost gauge

# HELP kubecost_load_balancer_cost kubecost_load_balancer_cost Hourly cost of load balancer
# TYPE kubecost_load_balancer_cost gauge

# HELP kubecost_network_internet_egress_cost kubecost_network_internet_egress_cost Total cost per GB of internet egress.
# TYPE kubecost_network_internet_egress_cost gauge

# HELP kubecost_network_region_egress_cost kubecost_network_region_egress_cost Total cost per GB egress across regions
# TYPE kubecost_network_region_egress_cost gauge

# HELP kubecost_network_zone_egress_cost kubecost_network_zone_egress_cost Total cost per GB egress across zones
# TYPE kubecost_network_zone_egress_cost gauge

# HELP pod_pvc_allocation pod_pvc_allocation Bytes used by a PVC attached to a pod
# TYPE pod_pvc_allocation gauge

# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter

# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge

# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge

# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge

# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge

# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge

# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge

# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge

# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter

# HELP service_selector_labels service_selector_labels Service Selector Labels
# TYPE service_selector_labels gauge

*/

var (
	deployMatchLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "deployment_match_labels",
		Help: "Deployment Match Labels",
	}, []string{"namespace", "deployment", "label_name"})

	ssMatchLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "statefulSet_match_labels",
		Help: "StatefulSet Match Labels",
	}, []string{"namespace", "statefulset", "label_name"})

	containerCPUAlloc = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "container_cpu_allocation",
		Help: "Percent of a single CPU used in a minute",
	}, []string{"namespace", "node", "pod", "container"})

	containerGPUAlloc = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "container_gpu_allocation",
		Help: "Percent of a single GPU used in a minute",
	}, []string{"namespace", "node", "pod", "container"})

	containerMemAllocBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "container_memory_allocation_bytes",
		Help: "Bytes of RAM used",
	}, []string{"namespace", "node", "pod", "container"})

	nodeCPUHourlyCost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "node_cpu_hourly_cost",
		Help: "hourly cost for each cpu on this node",
	}, []string{"instance", "node", "instance_type", "provider_id", "region"})

	nodeGPUCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "node_gpu_count",
		Help: "count of gpu on this node",
	}, []string{"instance", "node", "instance_type", "provider_id", "region"})

	nodeGPUHourlyCost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "node_gpu_hourly_cost",
		Help: "hourly cost for each gpu on this node",
	}, []string{"instance", "node", "instance_type", "provider_id", "region"})

	nodeRAMHourlyCost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "node_ram_hourly_cost",
		Help: "hourly cost for each gb of ram on this node",
	}, []string{"instance", "node", "instance_type", "provider_id", "region"})

	nodeTotalHourlyCost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "node_total_hourly_cost",
		Help: "Total node cost per hour",
	}, []string{"instance", "node", "instance_type", "provider_id", "region"})

	pvHourlyCost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pv_hourly_cost",
		Help: "Cost per GB per hour on a persistent disk",
	}, []string{"persistentvolume", "provider_id", "volumenname"})

	nodeIsSpot = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kubecost_node_is_spot",
		Help: "Cloud provider info about node preemptibility",
	}, []string{"instance", "node", "instance_type", "provider_id", "region"})

	serviceSelectorLabels = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "service_selector_labels",
		Help: "Service Selector Labels",
	}, []string{"namespace", "label_name", "service"})

	lbCost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "kubecost_load_balancer_cost",
		Help: "Hourly cost of load balancer",
	}, []string{"namespace", "service_name"})
)

type CostModelMetrics struct{}

func (m *CostModelMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(deployMatchLabels)
	r.MustRegister(ssMatchLabels)
	r.MustRegister(containerCPUAlloc)
	r.MustRegister(containerGPUAlloc)
	r.MustRegister(containerMemAllocBytes)
	r.MustRegister(nodeCPUHourlyCost)
	r.MustRegister(nodeGPUCount)
	r.MustRegister(nodeGPUHourlyCost)
	r.MustRegister(nodeRAMHourlyCost)
	r.MustRegister(nodeTotalHourlyCost)
	r.MustRegister(pvHourlyCost)
	r.MustRegister(nodeIsSpot)
	r.MustRegister(serviceSelectorLabels)
	r.MustRegister(lbCost)
}

func (m *CostModelMetrics) Update(cluster *cluster.Cluster) {
	m.UpdateMatchLabels(cluster)
	m.UpdateContainer(cluster)
	m.UpdateNode(cluster)
}

func (m *CostModelMetrics) UpdateMatchLabels(cluster *cluster.Cluster) {
	for ns, deploys := range cluster.Deployments {
		for _, d := range deploys {
			deployMatchLabels.WithLabelValues(ns, d, d).Set(1)
		}
	}

	for ns, ss := range cluster.StatefulSets {
		for _, s := range ss {
			deployMatchLabels.WithLabelValues(ns, s, s).Set(1)
		}
	}

	for ns, services := range cluster.Services {
		for _, s := range services {
			serviceSelectorLabels.WithLabelValues(ns, s, s).Set(1)
			lbCost.WithLabelValues(ns, s).Set(util.RandFloatBetween(0.2, 0))
		}
	}
}

func (m *CostModelMetrics) UpdateContainer(cluster *cluster.Cluster) {
	for ns, pods := range cluster.Containers {
		for p, containers := range pods {
			for _, c := range containers {
				node := util.GetFixedAssignment(c, cluster.Nodes)
				containerCPUAlloc.WithLabelValues(ns, node, p, c).Set(util.RandFloatBetween(1, 0))
				containerGPUAlloc.WithLabelValues(ns, node, p, c).Set(util.RandFloatBetween(1, 0))
				containerMemAllocBytes.WithLabelValues(ns, node, p, c).Set(util.RandFloatBetween(17179869184, 0))
			}
		}
	}
}

func (m *CostModelMetrics) UpdateNode(cluster *cluster.Cluster) {
	for _, node := range cluster.Nodes {
		instanceType := util.GetFixedAssignment(node, cluster.InstanceTypes)
		providerId := fmt.Sprintf("aws:///us-west-2a/i-%s", node)
		region := "us-west-2"
		nodeCPUHourlyCost.WithLabelValues(node, node, instanceType, providerId, region).Set(util.RandFloatBetween(1, 0))
		nodeGPUCount.WithLabelValues(node, node, instanceType, providerId, region).Set(1)
		nodeGPUHourlyCost.WithLabelValues(node, node, instanceType, providerId, region).Set(util.RandFloatBetween(1, 0))
		nodeRAMHourlyCost.WithLabelValues(node, node, instanceType, providerId, region).Set(util.RandFloatBetween(1, 0))
		nodeTotalHourlyCost.WithLabelValues(node, node, instanceType, providerId, region).Set(util.RandFloatBetween(1, 0))
		nodeIsSpot.WithLabelValues(node, node, instanceType, providerId, region).Set(0)
	}
}

func (m *CostModelMetrics) UpdatePV(cluster *cluster.Cluster) {
	for _, pvcs := range cluster.PVCs {
		for _, pvc := range pvcs {
			providerId := fmt.Sprintf("aws:///us-west-2a/pvc-%s", pvc)
			pvHourlyCost.WithLabelValues(pvc.VolumeName, providerId, pvc.VolumeName).Set(util.RandFloatBetween(1, 0))
		}
	}
}
