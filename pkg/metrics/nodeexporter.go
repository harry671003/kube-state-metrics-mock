package metrics

import (
	"strconv"

	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	nodeCPUSecondsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "node_cpu_seconds_total",
		Help: "Seconds the CPUs spent in each mode.",
	}, []string{"kubernetes_node", "cpu", "mode"})
)

type NodeExporterMetrics struct{}

func (m *NodeExporterMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(nodeCPUSecondsTotal)
}

func (m *NodeExporterMetrics) Update(cluster *cluster.Cluster) {
	modes := []string{
		"idle",
		"iowait",
		"irq",
		"nice",
		"softirq",
		"steal",
		"system",
		"user",
	}
	for _, n := range cluster.Nodes {
		for i := 0; i < 7; i++ {
			for _, m := range modes {
				nodeCPUSecondsTotal.WithLabelValues(n, strconv.Itoa(i), m).Add(util.RandFloatBetween(5, 0))
			}
		}
	}
}
