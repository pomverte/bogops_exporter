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
	teaPot = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "ze_status_code",
			Help: "Un status code.",
		},
	)
)

func initbogopsGauge() {
	bogopsGauge.WithLabelValues("Consultant").Set(11)
	bogopsGauge.WithLabelValues("Confirmé").Set(16)
	bogopsGauge.WithLabelValues("Senior").Set(4)
	bogopsGauge.WithLabelValues("Manager Ref").Set(1)
	bogopsGauge.WithLabelValues("Confirmed Mgr Ref").Set(1)
}

func intiStatusCode() {

	rand.Seed(time.Now().UnixNano())
	min := 200
	max := 208

	go func() { // go routine (asynchrone)
		for ; true; <-time.Tick(5 * time.Second) {
			httpBinUrl := fmt.Sprintf("https://httpbin.org/status/%d", min+rand.Intn(max-min+1))
			fmt.Println(httpBinUrl)
			response, err := http.Get(httpBinUrl)
			if err != nil {
				fmt.Println(err.Error())
				teaPot.Set(math.NaN())
			} else {
				teaPot.Set(float64(response.StatusCode))
			}
		}
	}()
}

func main() {
	initbogopsGauge()

	intiStatusCode()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
