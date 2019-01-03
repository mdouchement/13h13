package proxy

import (
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	forwardpkg "github.com/vulcand/oxy/forward"
)

// ListenAndServe starts the proxy and the metrics servers.
// The proxy server listens on paddr and send requests to the forward address.
// The metrics server listens on maddr and use service to add label to the metrics.
func ListenAndServe(paddr, maddr, service string, forward *url.URL) error {
	exporter := newExporter(service)

	fwd, err := forwardpkg.New(forwardpkg.Logger(logrus.New()))
	if err != nil {
		return err
	}

	errCh := make(chan error)

	go func() {
		// Proxy
		server := http.NewServeMux()
		server.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			path := req.URL.Path
			sw := statusWriter{ResponseWriter: w}

			req.URL = forward
			fwd.ServeHTTP(&sw, req)

			exporter.inc(path, sw.status)
		}))
		errCh <- http.ListenAndServe(paddr, server)
	}()

	go func() {
		// Prometheus
		server := http.NewServeMux()
		server.Handle("/metrics", promhttp.Handler())
		errCh <- http.ListenAndServe(maddr, server)
	}()

	return <-errCh
}
