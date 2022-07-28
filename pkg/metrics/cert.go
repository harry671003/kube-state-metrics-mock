package metrics

import (
	"github.com/harry671003/KubeStateMetricsMock/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	csrAnnotations = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_certificatesigningrequest_annotations",
		Help: "Kubernetes labels converted to Prometheus labels.",
	})
	csrCreated = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_certificatesigningrequest_created",
		Help: "Unix creation timestamp.",
	})
	csrCondition = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_certificatesigningrequest_condition",
		Help: "The number of each certificatesigningrequest condition.",
	})
	csrCertLength = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kube_certificatesigningrequest_cert_length",
		Help: "Length of the issued cert.",
	})
)

type CertMetrics struct{}

func (c *CertMetrics) Register(r *prometheus.Registry) {
	r.MustRegister(csrAnnotations)
	r.MustRegister(csrCreated)
	r.MustRegister(csrCondition)
	r.MustRegister(csrCertLength)
}

func (c *CertMetrics) Update(config *config.Config) {
	csrAnnotations.Inc()
	csrCreated.Inc()
	csrCertLength.Inc()
	csrCondition.Inc()
}
