package opencensusbridge

import (
	"go.opencensus.io/metric/metricexport"
	octrace "go.opencensus.io/trace"
	"go.opentelemetry.io/otel/bridge/opencensus"
	"go.opentelemetry.io/otel/sdk/metric/export"
	"go.opentelemetry.io/otel/trace"
)

// NewOpenCensusTracer returns an implementation of the OpenCensus Tracer interface which
// uses OpenTelemetry APIs.  Using this implementation of Tracer "upgrades"
// libraries that use OpenCensus to OpenTelemetry to facilitate a migration.
func NewOpenCensusTracer(tracer trace.Tracer) octrace.Tracer {
	return opencensus.NewTracer(tracer)
}

// OpenTelemetrySpanContextToOpenCensus converts from an OpenTelemetry SpanContext to an
// OpenCensus SpanContext, and handles any incompatibilities with the global
// error handler.
func OpenTelemetrySpanContextToOpenCensus(sc trace.SpanContext) octrace.SpanContext {
	return opencensus.OTelSpanContextToOC(sc)
}

// OpenCensusSpanContextToOpenTelemetry converts from an OpenCensus SpanContext to an
// OpenTelemetry SpanContext.
func OpenCensusSpanContextToOpenTelemetry(sc octrace.SpanContext) trace.SpanContext {
	return opencensus.OCSpanContextToOTel(sc)
}

// NewOpenCensusMetricExporter returns an OpenCensus exporter that exports to an
// OpenTelemetry exporter.
func NewOpenCensusMetricExporter(base export.Exporter) metricexport.Exporter {
	return opencensus.NewMetricExporter(base)
}
