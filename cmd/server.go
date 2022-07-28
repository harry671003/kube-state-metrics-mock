package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/harry671003/KubeStateMetricsMock/pkg/config"
	"github.com/harry671003/KubeStateMetricsMock/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	collect(&config.DefaultConfig, reg)
	serve("", 9009, reg)
}

func serve(host string, port int, reg *prometheus.Registry) {
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving on: %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func collect(config *config.Config, reg *prometheus.Registry) {
	metrics := []metrics.Metrics{
		&metrics.CertMetrics{},
	}

	for _, m := range metrics {
		m.Register(reg)
	}

	go func() {
		for {
			for _, m := range metrics {
				m.Update(config)
			}

			time.Sleep(2 * time.Second)
		}
	}()
}
