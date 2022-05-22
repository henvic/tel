package telexporter

import (
	"go.opentelemetry.io/otel/exporters/prometheus"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
)

// PrometheusExporter supports Prometheus pulls.  It does not implement the
// sdk/export/metric.Exporter interface--instead it creates a pull
// controller and reads the latest checkpointed data on-scrape.
type PrometheusExporter = prometheus.Exporter

// ErrUnsupportedAggregator is returned for unrepresentable aggregator
// types.
var ErrPrometheusUnsupportedAggregator = prometheus.ErrUnsupportedAggregator

// Config is a set of configs for the tally reporter.
type PrometheusConfig = prometheus.Config

// NewPrometheus returns a new Prometheus exporter using the configured metric
// controller.  See controller.New().
func NewPrometheus(config PrometheusConfig, ctrl *controller.Controller) (*PrometheusExporter, error) {
	return prometheus.New(config, ctrl)
}
