package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

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
	teaPot = promauto.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "ze_status_code",
			Help: "Un status code.",
		},
		getStatusCode,
	)
)

func initbogopsGauge() {
	bogopsGauge.WithLabelValues("Consultant").Set(42)
}

func getStatusCode() float64 {

	min := 200
	max := 208

	httpBinUrl := fmt.Sprintf("https://httpbin.org/status/%d", min+rand.Intn(max-min+1))
	fmt.Println(httpBinUrl)

	response, err := http.Get(httpBinUrl)
	if err != nil {
		fmt.Println(err.Error())
		return math.NaN()
	}
	return float64(response.StatusCode)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	initbogopsGauge()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
