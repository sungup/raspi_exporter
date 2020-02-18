package main

import (
	"fmt"
	"net/http"
	"raspi_exporter/common"
)

type MetricHandler struct {
	http.Handler

	thermal *common.MetricCollector
}

func newMetricHandler(opts *common.RaspiExpOpts) *MetricHandler {
	return &MetricHandler{}
}

func (h *MetricHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`Hello World!!`))
}

func main() {
	// 1. Argument parsing
	opts := common.ArgParse()

	// 2. Add Checking Prerequisite
	if err := common.CheckPrerequisite(opts); err != nil {
		panic(err)
	}

	// 3. Make and assign handler
	handler := newMetricHandler(opts)

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(`<h1>Raspi Exporter</h1><a href="/metrics">Metrics</a>`))
	})

	http.Handle("/metrics", handler)

	// 4. Start server
	fmt.Printf("Connect to %s\n", opts.ServerAddr())

	if err := http.ListenAndServe(opts.ListenAddr(), nil); err != nil {
		panic(err)
	}
}
