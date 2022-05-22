package telexporter

import (
	"io"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

// StdoutMetricOption sets the value of an option for a Config.
type StdoutMetricOption = stdoutmetric.Option

// WithStdoutMetricWriter sets the export stream destination.
func WithStdoutMetricWriter(w io.Writer) StdoutMetricOption {
	return stdoutmetric.WithWriter(w)
}

// WithStdoutMetricPrettyPrint sets the export stream format to use JSON.
func WithStdoutMetricPrettyPrint() StdoutMetricOption {
	return stdoutmetric.WithPrettyPrint()
}

// WithoutStdoutTimestamps sets the export stream to not include timestamps.
func WithoutStdoutTimestamps() StdoutMetricOption {
	return stdoutmetric.WithoutTimestamps()
}

// WithStdoutMetricAttributeEncoder sets the attribute encoder used in export.
func WithStdoutMetricAttributeEncoder(enc attribute.Encoder) StdoutMetricOption {
	return stdoutmetric.WithAttributeEncoder(enc)
}

// StdoutMetricExporter is an OpenTelemetry metric exporter that transmits telemetry to
// the local STDOUT.
type StdoutMetricExporter = stdoutmetric.Exporter

// NewStdoutMetric creates an Exporter with the passed options.
func NewStdoutMetric(options ...StdoutMetricOption) (*StdoutMetricExporter, error) {
	return stdoutmetric.New(options...)
}

// StdoutTraceOption sets the value of an option for a Config.
type StdoutTraceOption = stdouttrace.Option

// WithStdoutTraceWriter sets the export stream destination.
func WithStdoutTraceWriter(w io.Writer) StdoutTraceOption {
	return stdouttrace.WithWriter(w)
}

// WithStdoutTracePrettyPrint sets the export stream format to use JSON.
func WithStdoutTracePrettyPrint() StdoutTraceOption {
	return stdouttrace.WithPrettyPrint()
}

// WithoutStdoutTraceTimestamps sets the export stream to not include timestamps.
func WithoutStdoutTraceTimestamps() StdoutTraceOption {
	return stdouttrace.WithoutTimestamps()
}

// NewStdoutTrace creates an Exporter with the passed options.
func NewStdoutTrace(options ...StdoutTraceOption) (*StdoutTraceExporter, error) {
	return stdouttrace.New(options...)
}

// StdoutTraceExporter is an implementation of trace.SpanSyncer that writes spans to stdout.
type StdoutTraceExporter = stdouttrace.Exporter
