package proxy

import (
	"fmt"
	"log"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type exporter struct {
	sync.Mutex
	service  string
	counters map[string]prometheus.Counter
}

func newExporter(service string) *exporter {
	return &exporter{
		service:  service,
		counters: make(map[string]prometheus.Counter),
	}
}

func (e *exporter) inc(route string, status int) {
	e.Lock()
	defer e.Unlock()

	key := fmt.Sprintf("%d:%s", status, route)

	if _, ok := e.counters[key]; !ok {
		e.counters[key] = prometheus.NewCounter(prometheus.CounterOpts{
			Name: "http_route_count",
			Help: "Number of requests for a given route and HTTP status code.",
			ConstLabels: prometheus.Labels{
				"service": e.service,
				"route":   route,
				"status":  fmt.Sprintf("%d", status),
			},
		})

		if err := prometheus.Register(e.counters[key]); err != nil {
			log.Fatalf("could not register route %s (%d): %s", route, status, err)
		}
	}

	if _, ok := e.counters[key]; ok { // Ensure registration not failed
		e.counters[key].Inc()
	}
}
