package promhttp

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetPrefixForMetrics(name string) {
	prometheus.DefaultRegisterer.Unregister(collectors.NewGoCollector())
	prometheus.DefaultRegisterer.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	prometheus.DefaultRegisterer = prometheus.WrapRegistererWithPrefix(fmt.Sprintf("%s_", name), prometheus.DefaultRegisterer)
	prometheus.DefaultRegisterer.MustRegister(collectors.NewGoCollector())
	prometheus.DefaultRegisterer.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
}

func AddHandler() http.Handler {
	return promhttp.Handler()
}
