package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	bogopsGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ze_bogops_at_octo",
			Help: "Les bogops chez OCTO.",
		},
		[]string{"level"},
	)
)

func initbogopsGauge() {
	bogopsGauge.WithLabelValues("Consultant").Set(11)
	bogopsGauge.WithLabelValues("Confirmé").Set(16)
	bogopsGauge.WithLabelValues("Senior").Set(4)
	bogopsGauge.WithLabelValues("Manager Ref").Set(1)
	bogopsGauge.WithLabelValues("Confirmed Mgr Ref").Set(1)
}

func main() {
	initbogopsGauge()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
