package telsdk

import (
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/export"
)

// MetricAccumulator implements the OpenTelemetry Meter API.  The
// Accumulator is bound to a single export.Processor in
// `NewAccumulator()`.
//
// The Accumulator supports a Collect() API to gather and export
// current data.  Collect() should be arranged according to
// the processor model.  Push-based processors will setup a
// timer to call Collect() periodically.  Pull-based processors
// will call Collect() when a pull request arrives.
type MetricAccumulator = metric.Accumulator

// ErrUninitializedInstrument is returned when an instrument is used when uninitialized.
var ErrUninitializedInstrument = metric.ErrUninitializedInstrument

// ErrBadInstrument is returned when an instrument from another SDK is
// attempted to be registered with this SDK.
var ErrBadInstrument = metric.ErrBadInstrument

// NewMetricAccumulator constructs a new Accumulator for the given
// processor.  This Accumulator supports only a single processor.
//
// The Accumulator does not start any background process to collect itself
// periodically, this responsibility lies with the processor, typically,
// depending on the type of export.  For example, a pull-based
// processor will call Collect() when it receives a request to scrape
// current metric values.  A push-based processor should configure its
// own periodic collection.
func NewMetricAccumulator(processor export.Processor) *MetricAccumulator {
	return metric.NewAccumulator(processor)
}
