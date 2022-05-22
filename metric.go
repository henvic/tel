package tel

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/asyncfloat64"
	"go.opentelemetry.io/otel/metric/instrument/asyncint64"
	"go.opentelemetry.io/otel/metric/instrument/syncfloat64"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/unit"
)

// MeterConfig contains options for Meters.
type MeterConfig = metric.MeterConfig

// MeterOption is an interface for applying Meter options.
type MeterOption = metric.MeterOption

// NewMeterConfig creates a new MeterConfig and applies
// all the given options.
func NewMeterConfig(opts ...MeterOption) MeterConfig {
	return metric.NewMeterConfig(opts...)
}

// WithInstrumentationVersion sets the instrumentation version.
func WithInstrumentationVersion(version string) MeterOption {
	return metric.WithInstrumentationVersion(version)
}

// WithSchemaURL sets the schema URL.
func WithSchemaURL(schemaURL string) MeterOption {
	return metric.WithSchemaURL(schemaURL)
}

// MeterProvider provides access to named Meter instances, for instrumenting
// an application or library.
type MeterProvider = metric.MeterProvider

// Meter provides access to instrument instances for recording metrics.
type Meter = metric.Meter

type MetricUnit = unit.Unit

// Units defined by OpenTelemetry.
const (
	Dimensionless = unit.Dimensionless
	Bytes         = unit.Bytes
	Milliseconds  = unit.Milliseconds
)

// InstrumentConfig contains options for metric instrument descriptors.
type InstrumentConfig = instrument.Config

// InstrumentOption is an interface for applying metric instrument options.
type InstrumentOption = instrument.Option

// NewInstrumentConfig creates a new Config and applies all the given options.
func NewInstrumentConfig(opts ...InstrumentOption) InstrumentConfig {
	return instrument.NewConfig(opts...)
}

// WithInstrumentDescription applies provided description.
func WithInstrumentDescription(desc string) InstrumentOption {
	return instrument.WithDescription(desc)
}

// WithInstrumentUnit applies provided unit.
func WithInstrumentUnit(unit MetricUnit) InstrumentOption {
	return instrument.WithUnit(unit)
}

// Asynchronous instruments are instruments that are updated within a Callback.
// If an instrument is observed outside of it's callback it should be an error.
//
// This interface is used as a grouping mechanism.
type Asynchronous = instrument.Asynchronous

// Synchronous instruments are updated in line with application code.
//
// This interface is used as a grouping mechanism.
type Synchronous = instrument.Synchronous

// InstrumentProvider provides access to individual instruments.
type AsyncFloat64InstrumentProvider = asyncfloat64.InstrumentProvider

// AsyncFloat64Counter is an instrument that records increasing values.
type AsyncFloat64Counter = asyncfloat64.Counter

// AsyncFloat64UpDownCounter is an instrument that records increasing or decresing values.
type AsyncFloat64UpDownCounter = asyncfloat64.UpDownCounter

// AsyncFloat64Gauge is an instrument that records independent readings.
type AsyncFloat64Gauge = asyncfloat64.Gauge

// AsyncInt64InstrumentProvider provides access to individual instruments.
type AsyncInt64InstrumentProvider = asyncint64.InstrumentProvider

// AsyncInt64Counter is an instrument that records increasing values.
type AsyncInt64Counter = asyncint64.Counter

// AsyncInt64UpDownCounter is an instrument that records increasing or decresing values.
type AsyncInt64UpDownCounter = asyncint64.UpDownCounter

// Gauge is an instrument that records independent readings.
type Gauge = asyncint64.Gauge

// InstrumentProvider provides access to individual instruments.
type InstrumentProvider = syncfloat64.InstrumentProvider

// SyncFloat64Counter is an instrument that records increasing values.
type SyncFloat64Counter = syncfloat64.Counter

// SyncFloat64UpDownCounter is an instrument that records increasing or decresing values.
type SyncFloat64UpDownCounter = syncfloat64.UpDownCounter

// SyncFloat64Histogram is an instrument that records a distribution of values.
type SyncFloat64Histogram = syncfloat64.Histogram

// SyncInt64InstrumentProvider provides access to individual instruments.
type SyncInt64InstrumentProvider = syncint64.InstrumentProvider

// SyncInt64Counter is an instrument that records increasing values.
type SyncInt64Counter = syncint64.Counter

// SyncInt64UpDownCounter is an instrument that records increasing or decresing values.
type SyncInt64UpDownCounter = syncint64.UpDownCounter

// SyncInt64Histogram is an instrument that records a distribution of values.
type SyncInt64Histogram = syncint64.Histogram

// Meter returns a Meter from the global MeterProvider. The
// instrumentationName must be the name of the library providing
// instrumentation. This name may be the same as the instrumented code only if
// that code provides built-in instrumentation. If the instrumentationName is
// empty, then a implementation defined default name will be used instead.
//
// This is short for MeterProvider().Meter(name).
func GlobalMeter(instrumentationName string, opts ...MeterOption) Meter {
	return global.Meter(instrumentationName, opts...)
}

// MeterProvider returns the registered global trace provider.
// If none is registered then a No-op MeterProvider is returned.
func GlobalMeterProvider() MeterProvider {
	return global.MeterProvider()
}

// SetGlobalMeterProvider registers `mp` as the global meter provider.
func SetGlobalMeterProvider(mp MeterProvider) {
	global.SetMeterProvider(mp)
}
