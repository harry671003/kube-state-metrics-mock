package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	Register(r *prometheus.Registry)
	Update(c *cluster.Cluster)
}
