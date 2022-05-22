package telsdk

import (
	"go.opentelemetry.io/otel/sdk/metric/aggregator"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/exponential/mapping"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/exponential/mapping/exponent"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/exponential/mapping/logarithm"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/lastvalue"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/sum"
	"go.opentelemetry.io/otel/sdk/metric/number"
	"go.opentelemetry.io/otel/sdk/metric/sdkapi"
)

// Aggregator implements a specific aggregation behavior, e.g., a
// behavior to track a sequence of updates to an instrument.  Counter
// instruments commonly use a simple Sum aggregator, but for the
// distribution instruments (Histogram, GaugeObserver) there are a
// number of possible aggregators with different cost and accuracy
// tradeoffs.
//
// Note that any Aggregator may be attached to any instrument--this is
// the result of the OpenTelemetry API/SDK separation.  It is possible
// to attach a Sum aggregator to a Histogram instrument.
type Aggregator = aggregator.Aggregator

// NewInconsistentAggregatorError formats an error describing an attempt to
// Checkpoint or Merge different-type aggregators.  The result can be unwrapped as
// an ErrInconsistentType.
func NewInconsistentAggregatorError(a1, a2 Aggregator) error {
	return aggregator.NewInconsistentAggregatorError(a1, a2)
}

// AggregatorRangeTest is a common routine for testing for valid input values.
// This rejects NaN values.  This rejects negative values when the
// metric instrument does not support negative values, including
// monotonic counter metrics and absolute Histogram metrics.
func AggregatorRangeTest(num number.Number, descriptor *sdkapi.Descriptor) error {
	return aggregator.RangeTest(num, descriptor)
}

// Mapping is the interface of an exponential histogram mapper.
type Mapping = mapping.Mapping

var (
	// ErrUnderflow is returned when computing the lower boundary
	// of an index that maps into a denormalized floating point value.
	ErrUnderflow = mapping.ErrUnderflow
	// ErrOverflow is returned when computing the lower boundary
	// of an index that maps into +Inf.
	ErrOverflow = mapping.ErrOverflow
)

const (
	// MinScale defines the point at which the exponential mapping
	// function becomes useless for float64.  With scale -10, ignoring
	// subnormal values, bucket indices range from -1 to 1.
	MinScale = exponent.MinScale

	// MaxScale is the largest scale supported in this code.  Use
	// ../logarithm for larger scales.
	MaxScale = exponent.MaxScale

	// SignificandWidth is the size of an IEEE 754 double-precision
	// floating-point significand.
	SignificandWidth = exponent.SignificandWidth
	// ExponentWidth is the size of an IEEE 754 double-precision
	// floating-point exponent.
	ExponentWidth = exponent.ExponentWidth

	// SignificandMask is the mask for the significand of an IEEE 754
	// double-precision floating-point value: 0xFFFFFFFFFFFFF.
	SignificandMask = exponent.SignificandMask

	// ExponentBias is the exponent bias specified for encoding
	// the IEEE 754 double-precision floating point exponent: 1023.
	ExponentBias = exponent.ExponentBias

	// ExponentMask are set to 1 for the bits of an IEEE 754
	// floating point exponent: 0x7FF0000000000000.
	ExponentMask = exponent.ExponentMask

	// SignMask selects the sign bit of an IEEE 754 floating point
	// number.
	SignMask = exponent.SignMask

	// MinNormalExponent is the minimum exponent of a normalized
	// floating point: -1022.
	MinNormalExponent = exponent.MinNormalExponent

	// MaxNormalExponent is the maximum exponent of a normalized
	// floating point: 1023.
	MaxNormalExponent = exponent.MaxNormalExponent

	// MinValue is the smallest normal number.
	MinValue = exponent.MinValue

	// MaxValue is the largest normal number.
	MaxValue = exponent.MaxValue
)

// NewExponentMapping constructs an exponential mapping function, used for scales <= 0.
func NewExponentMapping(scale int32) (mapping.Mapping, error) {
	return exponent.NewMapping(scale)
}

const (
	// MinScale ensures that the ../exponent mapper is used for
	// zero and negative scale values.  Do not use the logarithm
	// mapper for scales <= 0.
	LogarithmMinScale = logarithm.MinScale

	// MaxScale is selected as the largest scale that is possible
	// in current code, considering there are 10 bits of base-2
	// exponent combined with scale-bits of range.  At this scale,
	// the growth factor is 0.0000661%.
	//
	// Scales larger than 20 complicate the logic in cmd/prebuild,
	// because math/big overflows when exponent is math.MaxInt32
	// (== the index of math.MaxFloat64 at scale=21),
	//
	// At scale=20, index values are in the interval [-0x3fe00000,
	// 0x3fffffff], having 31 bits of information.  This is
	// sensible given that the OTLP exponential histogram data
	// point uses a signed 32 bit integer for indices.
	LogarithmMaxScale = logarithm.MaxScale

	// MaxValue is the largest normal number.
	LogarithmMaxValue = logarithm.MaxValue

	// MinValue is the smallest normal number.
	LogarithmMinValue = logarithm.MinValue
)

// NewLogarithmMapping constructs a logarithm mapping function, used for scales > 0.
func NewLogarithmMapping(scale int32) (mapping.Mapping, error) {
	return logarithm.NewMapping(scale)
}

type (
	// Aggregator observe events and counts them in pre-determined buckets.
	// It also calculates the sum and count of all events.
	HistogramAggregator = histogram.Aggregator

	// HistogramOption configures a histogram config.
	HistogramOption = histogram.Option
)

// HistogramWithExplicitBoundaries sets the ExplicitBoundaries configuration option of a config.
func HistogramWithExplicitBoundaries(explicitBoundaries []float64) HistogramOption {
	return histogram.WithExplicitBoundaries(explicitBoundaries)
}

// New returns a new aggregator for computing Histograms.
//
// A Histogram observe events and counts them in pre-defined buckets.
// And also provides the total sum and count of all observations.
//
// Note that this aggregator maintains each value using independent
// atomic operations, which introduces the possibility that
// checkpoints are inconsistent.
func NewHistogram(cnt int, desc *sdkapi.Descriptor, opts ...HistogramOption) []HistogramAggregator {
	return histogram.New(cnt, desc, opts...)
}

type (
	// Aggregator aggregates lastValue events.
	LastValueAggregator = lastvalue.Aggregator
)

// NewLastValue returns a new lastValue aggregator.  This aggregator retains the
// last value and timestamp that were recorded.
func NewLastValue(cnt int) []LastValueAggregator {
	return lastvalue.New(cnt)
}

// SumAggregator aggregates counter events.
type SumAggregator = sum.Aggregator

// NewAggregator returns a new counter aggregator implemented by atomic
// operations.  This aggregator implements the aggregation.Sum
// export interface.
func NewSumAggregator(cnt int) []SumAggregator {
	return sum.New(cnt)
}
