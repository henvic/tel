package telsdk

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/unit"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	"go.opentelemetry.io/otel/sdk/metric/controller/basic"
	ctime "go.opentelemetry.io/otel/sdk/metric/controller/time"
	"go.opentelemetry.io/otel/sdk/metric/export"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	"go.opentelemetry.io/otel/sdk/metric/number"
	basicProcessor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/processor/reducer"
	"go.opentelemetry.io/otel/sdk/metric/registry"
	"go.opentelemetry.io/otel/sdk/metric/sdkapi"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
)

// BasicControllerOption is the interface that applies the value to a configuration option.
type BasicControllerOption = basic.Option

// WithBasicControllerResource sets the Resource configuration option of a Config by merging it
// with the Resource configuration in the environment.
func WithBasicControllerResource(r *resource.Resource) BasicControllerOption {
	return basic.WithResource(r)
}

// WithCollectPeriod sets the CollectPeriod configuration option of a Config.
func WithBasicControllerCollectPeriod(period time.Duration) BasicControllerOption {
	return basic.WithCollectPeriod(period)
}

// WithBasicControllerCollectTimeout sets the CollectTimeout configuration option of a Config.
func WithBasicControllerCollectTimeout(timeout time.Duration) BasicControllerOption {
	return basic.WithCollectTimeout(timeout)
}

// WithBasicControllerExporter sets the exporter configuration option of a Config.
func WithBasicControllerExporter(exporter export.Exporter) BasicControllerOption {
	return basic.WithExporter(exporter)
}

// WithBasicControllerPushTimeout sets the PushTimeout configuration option of a Config.
func WithBasicControllerPushTimeout(timeout time.Duration) BasicControllerOption {
	return basic.WithPushTimeout(timeout)
}

// DefaultPeriod is used for:
//
// - the minimum time between calls to Collect()
// - the timeout for Export()
// - the timeout for Collect().
const BasicControllerDefaultPeriod = basic.DefaultPeriod

// ErrBasicControllerStarted indicates that a controller was started more
// than once.
var ErrBasicControllerStarted = basic.ErrControllerStarted

// Controller organizes and synchronizes collection of metric data in
// both "pull" and "push" configurations.  This supports two distinct
// modes:
//
// - Push and Pull: Start() must be called to begin calling the exporter;
//   Collect() is called periodically by a background thread after starting
//   the controller.
// - Pull-Only: Start() is optional in this case, to call Collect periodically.
//   If Start() is not called, Collect() can be called manually to initiate
//   collection
//
// The controller supports mixing push and pull access to metric data
// using the export.Reader RWLock interface.  Collection will
// be blocked by a pull request in the basic controller.
type BasicController = basic.Controller

// NewBasicController constructs a Controller using the provided checkpointer factory
// and options (including optional exporter) to configure a metric
// export pipeline.
func NewBasicController(checkpointerFactory export.CheckpointerFactory, opts ...BasicControllerOption) *BasicController {
	return basic.New(checkpointerFactory, opts...)
}

// Several types below are created to match "github.com/benbjohnson/clock"
// so that it remains a test-only dependency.

// Clock keeps track of time for a metric SDK.
type Clock = ctime.Clock

// Ticker signals time intervals.
type Ticker = ctime.Ticker

// RealClock wraps the time package and uses the system time to tell time.
type RealClock = ctime.RealClock

// RealTicker wraps the time package and uses system time to tick time
// intervals.
type RealTicker = ctime.RealTicker

// MetricExportProcessor is responsible for deciding which kind of aggregation to
// use (via AggregatorSelector), gathering exported results from the
// SDK during collection, and deciding over which dimensions to group
// the exported data.
//
// The SDK supports binding only one of these interfaces, as it has
// the sole responsibility of determining which Aggregator to use for
// each record.
//
// The embedded AggregatorSelector interface is called (concurrently)
// in instrumentation context to select the appropriate Aggregator for
// an instrument.
//
// The `Process` method is called during collection in a
// single-threaded context from the SDK, after the aggregator is
// checkpointed, allowing the processor to build the set of metrics
// currently being exported.
type MetricExportProcessor = export.Processor

// AggregatorSelector supports selecting the kind of Aggregator to
// use at runtime for a specific metric instrument.
type AggregatorSelector = export.AggregatorSelector

// Checkpointer is the interface used by a Controller to coordinate
// the Processor with Accumulator(s) and Exporter(s).  The
// StartCollection() and FinishCollection() methods start and finish a
// collection interval.  Controllers call the Accumulator(s) during
// collection to process Accumulations.
type Checkpointer = export.Checkpointer

// CheckpointerFactory is an interface for producing configured
// Checkpointer instances.
type CheckpointerFactory = export.CheckpointerFactory

// Exporter handles presentation of the checkpoint of aggregate
// metrics.  This is the final stage of a metrics export pipeline,
// where metric data are formatted for a specific system.
type Exporter = export.Exporter

// InstrumentationLibraryReader is an interface for exporters to iterate
// over one instrumentation library of metric data at a time.
type InstrumentationLibraryReader = export.InstrumentationLibraryReader

// Reader allows a controller to access a complete checkpoint of
// aggregated metrics from the Processor for a single library of
// metric data.  This is passed to the Exporter which may then use
// ForEach to iterate over the collection of aggregated metrics.
type Reader = export.Reader

// Metadata contains the common elements for exported metric data that
// are shared by the Accumulator->Processor and Processor->Exporter
// steps.
type Metadata = export.Metadata

// Accumulation contains the exported data for a single metric instrument
// and attribute set, as prepared by an Accumulator for the Processor.
type Accumulation = export.Accumulation

// Record contains the exported data for a single metric instrument
// and attribute set, as prepared by the Processor for the Exporter.
// This includes the effective start and end time for the aggregation.
type Record = export.Record

// NewMetricExportAccumulation allows Accumulator implementations to construct new
// Accumulations to send to Processors. The Descriptor, attributes, and
// Aggregator represent aggregate metric events received over a single
// collection period.
func NewMetricExportAccumulation(descriptor *sdkapi.Descriptor, attrs *attribute.Set, agg Aggregator) Accumulation {
	return export.NewAccumulation(descriptor, attrs, agg)
}

// NewRecord allows Processor implementations to construct export records.
// The Descriptor, attributes, and Aggregator represent aggregate metric
// events received over a single collection period.
func NewRecord(descriptor *sdkapi.Descriptor, attrs *attribute.Set, agg aggregation.Aggregation, start, end time.Time) Record {
	return export.NewRecord(descriptor, attrs, agg, start, end)
}

type (
	// Aggregation is an interface returned by the Aggregator
	// containing an interval of metric data.
	MetricExportAggregation = aggregation.Aggregation

	// Sum returns an aggregated sum.
	Sum = aggregation.Sum

	// Count returns the number of values that were aggregated.
	Count = aggregation.Count

	// LastValue returns the latest value that was aggregated.
	LastValue = aggregation.LastValue

	// Buckets represents histogram buckets boundaries and counts.
	//
	// For a Histogram with N defined boundaries, e.g, [x, y, z].
	// There are N+1 counts: [-inf, x), [x, y), [y, z), [z, +inf].
	Buckets = aggregation.Buckets

	// Histogram returns the count of events in pre-determined buckets.
	Histogram = aggregation.Histogram

	// AggregatorKind is a short name for the Aggregator that produces an
	// Aggregation, used for descriptive purpose only.  Kind is a
	// string to allow user-defined Aggregators.
	//
	// When deciding how to handle an Aggregation, Exporters are
	// encouraged to decide based on conversion to the above
	// interfaces based on strength, not on Kind value, when
	// deciding how to expose metric data.  This enables
	// user-supplied Aggregators to replace builtin Aggregators.
	//
	// For example, test for a Histogram before testing for a
	// Sum, and so on.
	AggregatorKind = aggregation.Kind
)

// Kind description constants.
const (
	SumKind       = aggregation.SumKind
	HistogramKind = aggregation.HistogramKind
	LastValueKind = aggregation.LastValueKind
)

// Sentinel errors for Aggregation interface.
var (
	ErrNegativeInput    = aggregation.ErrNegativeInput
	ErrNaNInput         = aggregation.ErrNaNInput
	ErrInconsistentType = aggregation.ErrInconsistentType

	// ErrNoCumulativeToDelta is returned when requesting delta
	// export kind for a precomputed sum instrument.
	ErrNoCumulativeToDelta = aggregation.ErrNoCumulativeToDelta

	// ErrNoData is returned when (due to a race with collection)
	// the Aggregator is check-pointed before the first value is set.
	// The aggregator should simply be skipped in this case.
	ErrNoData = aggregation.ErrNoData
)

// MetricNumber describes the data type of the Number.
type MetricNumber = number.Kind

const (
	// Int64Kind means that the Number stores int64.
	Int64Kind = number.Int64Kind
	// Float64Kind means that the Number stores float64.
	Float64Kind = number.Float64Kind
)

// Number represents either an integral or a floating point value. It
// needs to be accompanied with a source of Kind that describes
// the actual type of the value stored within Number.
type Number = number.Number

// NewNumberFromRaw creates a new Number from a raw value.
func NewNumberFromRaw(r uint64) Number {
	return number.NewNumberFromRaw(r)
}

// NewInt64Number creates an integral Number.
func NewInt64Number(i int64) Number {
	return number.NewInt64Number(i)
}

// NewFloat64Number creates a floating point Number.
func NewFloat64Number(f float64) Number {
	return number.NewFloat64Number(f)
}

// NewNumberSignChange returns a number with the same magnitude and
// the opposite sign.  `kind` must describe the kind of number in `nn`.
func NewNumberSignChange(kind MetricNumber, nn Number) Number {
	return number.NewNumberSignChange(kind, nn)
}

type (
	// Processor is a basic metric processor.
	Processor = basicProcessor.Processor
)

// ErrInconsistentState is returned when the sequence of collection's starts and finishes are incorrectly balanced.
var ErrInconsistentState = basicProcessor.ErrInconsistentState

// ErrInvalidTemporality is returned for unknown metric.Temporality.
var ErrInvalidTemporality = basicProcessor.ErrInvalidTemporality

// New returns a basic Processor that is also a Checkpointer using the provided
// AggregatorSelector to select Aggregators.  The TemporalitySelector
// is consulted to determine the kind(s) of exporter that will consume
// data, so that this Processor can prepare to compute Cumulative Aggregations
// as needed.
func New(aselector export.AggregatorSelector, tselector aggregation.TemporalitySelector, opts ...BasicProcessorOption) *Processor {
	return basicProcessor.New(aselector, tselector, opts...)
}

// NewFactory returns a new basic CheckpointerFactory.
func NewFactory(aselector export.AggregatorSelector, tselector aggregation.TemporalitySelector, opts ...BasicProcessorOption) export.CheckpointerFactory {
	return basicProcessor.NewFactory(aselector, tselector, opts...)
}

// BasicProcessorOption configures a basic processor configuration.
type BasicProcessorOption = basicProcessor.Option

// WithMemory sets the memory behavior of a Processor. If this is true, the
// processor will report metric instruments and attribute sets that were
// previously reported but not updated in the most recent interval.
func WithMemory(memory bool) BasicProcessorOption {
	return basicProcessor.WithMemory(memory)
}

type (
	// ReducerProcessor implements "dimensionality reduction" by
	// filtering keys from export attribute sets.
	ReducerProcessor = reducer.Processor

	// ReducerAttributeFilterSelector selects an attribute filter based on the
	// instrument described by the descriptor.
	ReducerAttributeFilterSelector = reducer.AttributeFilterSelector
)

// NewReducer returns a dimensionality-reducing Processor that passes data to the
// next stage in an export pipeline.
func NewReducer(filterSelector ReducerAttributeFilterSelector, ckpter export.Checkpointer) *ReducerProcessor {
	return reducer.New(filterSelector, ckpter)
}

// UniqueInstrumentMeterImpl implements the metric.MeterImpl interface, adding
// uniqueness checking for instrument descriptors.
type UniqueInstrumentMeterImpl = registry.UniqueInstrumentMeterImpl

// NewUniqueInstrumentMeterImpl returns a wrapped metric.MeterImpl
// with the addition of instrument name uniqueness checking.
func NewUniqueInstrumentMeterImpl(impl sdkapi.MeterImpl) *UniqueInstrumentMeterImpl {
	return registry.NewUniqueInstrumentMeterImpl(impl)
}

// NewMetricKindMismatchError formats an error that describes a
// mismatched metric instrument definition.
func NewMetricKindMismatchError(desc sdkapi.Descriptor) error {
	return registry.NewMetricKindMismatchError(desc)
}

// Compatible determines whether two sdkapi.Descriptors are considered
// the same for the purpose of uniqueness checking.
func Compatible(candidate, existing sdkapi.Descriptor) bool {
	return registry.Compatible(candidate, existing)
}

// APIDescriptor contains all the settings that describe an instrument,
// including its name, metric kind, number kind, and the configurable
// options.
type APIDescriptor = sdkapi.Descriptor

// NewAPIDescriptor returns a Descriptor with the given contents.
func NewAPIDescriptor(name string, ikind InstrumentKind, nkind MetricNumber, description string, u unit.Unit) APIDescriptor {
	return sdkapi.NewDescriptor(name, ikind, nkind, description, u)
}

// InstrumentKind describes the kind of instrument.
type InstrumentKind = sdkapi.InstrumentKind

const (
	// HistogramInstrumentKind indicates a Histogram instrument.
	HistogramInstrumentKind = sdkapi.HistogramInstrumentKind
	// GaugeObserverInstrumentKind indicates an GaugeObserver instrument.
	GaugeObserverInstrumentKind = sdkapi.GaugeObserverInstrumentKind

	// CounterInstrumentKind indicates a Counter instrument.
	CounterInstrumentKind = sdkapi.CounterInstrumentKind
	// UpDownCounterInstrumentKind indicates a UpDownCounter instrument.
	UpDownCounterInstrumentKind = sdkapi.UpDownCounterInstrumentKind

	// CounterObserverInstrumentKind indicates a CounterObserver instrument.
	CounterObserverInstrumentKind = sdkapi.CounterObserverInstrumentKind
	// UpDownCounterObserverInstrumentKind indicates a UpDownCounterObserver
	// instrument.
	UpDownCounterObserverInstrumentKind = sdkapi.UpDownCounterObserverInstrumentKind
)

// NewNoopSyncInstrument returns a No-op implementation of the
// synchronous instrument interface.
func NewNoopSyncInstrument() SyncImpl {
	return sdkapi.NewNoopSyncInstrument()
}

// NewNoopAsyncInstrument returns a No-op implementation of the
// asynchronous instrument interface.
func NewNoopAsyncInstrument() AsyncImpl {
	return sdkapi.NewNoopAsyncInstrument()
}

// MeterImpl is the interface an SDK must implement to supply a Meter
// implementation.
type MeterImpl = sdkapi.MeterImpl

// InstrumentImpl is a common interface for synchronous and
// asynchronous instruments.
type InstrumentImpl = sdkapi.InstrumentImpl

// SyncImpl is the implementation-level interface to a generic
// synchronous instrument (e.g., Histogram and Counter instruments).
type SyncImpl = sdkapi.SyncImpl

// AsyncImpl is an implementation-level interface to an
// asynchronous instrument (e.g., Observer instruments).
type AsyncImpl = sdkapi.AsyncImpl

// AsyncRunner is expected to convert into an AsyncSingleRunner or an
// AsyncBatchRunner.  SDKs will encounter an error if the AsyncRunner
// does not satisfy one of these interfaces.
type AsyncRunner = sdkapi.AsyncRunner

// AsyncSingleRunner is an interface implemented by single-observer
// callbacks.
type AsyncSingleRunner = sdkapi.AsyncSingleRunner

// AsyncBatchRunner is an interface implemented by batch-observer
// callbacks.
type AsyncBatchRunner = sdkapi.AsyncBatchRunner

// NewAPIMeasurement constructs a single observation, a binding between
// an asynchronous instrument and a number.
func NewAPIMeasurement(inst SyncImpl, n Number) {
	sdkapi.NewMeasurement(inst, n)
}

// Measurement is a low-level type used with synchronous instruments
// as a direct interface to the SDK via `RecordBatch`.
type APIMeasurement = sdkapi.Measurement

// NewObservation constructs a single observation, a binding between
// an asynchronous instrument and a number.
func NewObservation(inst AsyncImpl, n Number) Observation {
	return sdkapi.NewObservation(inst, n)
}

// Observation is a low-level type used with asynchronous instruments
// as a direct interface to the SDK via `BatchObserver`.
type Observation = sdkapi.Observation

// WrapMeterImpl wraps impl to be a full implementation of a Meter.
func WrapMeterImpl(impl MeterImpl) metric.Meter {
	return sdkapi.WrapMeterImpl(impl)
}

// UnwrapMeterImpl unwraps the Meter to its bare MeterImpl.
func UnwrapMeterImpl(m metric.Meter) MeterImpl {
	return sdkapi.UnwrapMeterImpl(m)
}

// NewWithInexpensiveDistribution returns a simple aggregator selector
// that uses minmaxsumcount aggregators for `Histogram`
// instruments.  This selector is faster and uses less memory than the
// others in this package because minmaxsumcount aggregators maintain
// the least information about the distribution among these choices.
func NewWithInexpensiveDistribution() export.AggregatorSelector {
	return simple.NewWithInexpensiveDistribution()
}

// NewWithHistogramDistribution returns a simple aggregator selector
// that uses histogram aggregators for `Histogram` instruments.
// This selector is a good default choice for most metric exporters.
func NewWithHistogramDistribution(options ...histogram.Option) export.AggregatorSelector {
	return simple.NewWithHistogramDistribution(options...)
}
