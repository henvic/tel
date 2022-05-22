package telexporter

import (
	"log"
	"net/http"

	zkmodel "github.com/openzipkin/zipkin-go/model"
	"go.opentelemetry.io/otel/exporters/zipkin"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ZkipKinSpanModels converts OpenTelemetry spans into Zipkin model spans.
// This is used for exporting to Zipkin compatible tracing services.
func ZkipKinSpanModels(batch []tracesdk.ReadOnlySpan) []zkmodel.SpanModel {
	return zipkin.SpanModels(batch)
}

// ZipkinExporter exports spans to the zipkin collector.
type ZipkinExporter = zipkin.Exporter

// ZipkinOption defines a function that configures the exporter.
type ZipkinOption = zipkin.Option

// WithZkipKinLogger configures the exporter to use the passed logger.
func WithZkipKinLogger(logger *log.Logger) ZipkinOption {
	return zipkin.WithLogger(logger)
}

// WithClient configures the exporter to use the passed HTTP client.
func WithClient(client *http.Client) ZipkinOption {
	return zipkin.WithClient(client)
}

// New creates a new Zipkin exporter.
func New(collectorURL string, opts ...ZipkinOption) (*ZipkinExporter, error) {
	return zipkin.New(collectorURL, opts...)
}
