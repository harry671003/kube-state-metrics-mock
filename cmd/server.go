package main

import (
	"flag"
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
	configFile := flag.String("config_file", "", "Config file")
	flag.Parse()

	cfg, err := config.NewConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	reg := prometheus.NewRegistry()

	cluster := cluster.NewCluster(cfg)

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
		&metrics.DeploymentMetrics{},
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
