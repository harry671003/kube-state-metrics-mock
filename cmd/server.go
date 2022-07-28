package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/harry671003/KubeStateMetricsMock/pkg/cluster"
	"github.com/harry671003/KubeStateMetricsMock/pkg/config"
	"github.com/harry671003/KubeStateMetricsMock/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	cluster := cluster.NewCluster(&config.DefaultConfig)

	fmt.Printf("%v\n", cluster.Namespaces)
	fmt.Printf("%v\n", cluster.StatefulSets)
	update(cluster, reg)
	serve("", 9009, reg)
}

func serve(host string, port int, reg *prometheus.Registry) {
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving on: %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func update(cluster *cluster.Cluster, reg *prometheus.Registry) {
	metrics := []metrics.Metrics{
		&metrics.CertMetrics{},
		&metrics.StatefulSetMetrics{},
	}

	for _, m := range metrics {
		m.Register(reg)
	}

	go func() {
		for {
			for _, m := range metrics {
				m.Update(cluster)
			}

			time.Sleep(2 * time.Second)
		}
	}()
}
