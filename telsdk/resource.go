package telsdk

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

var (
	// ErrPartialResource is returned by a detector when complete source
	// information for a Resource is unavailable or the source information
	// contains invalid values that are omitted from the returned Resource.
	ErrPartialResource = resource.ErrPartialResource
)

// ResourceDetector detects OpenTelemetry resource information.
type ResourceDetector = resource.Detector

// DetectResource calls all input detectors sequentially and merges each result with the previous one.
// It returns the merged error too.
func DetectResource(ctx context.Context, detectors ...ResourceDetector) (*Resource, error) {
	return resource.Detect(ctx, detectors...)
}

// ResourceOption is the interface that applies a configuration option.
type ResourceOption = resource.Option

// WithAttributes adds attributes to the configured Resource.
func WithAttributes(attributes ...attribute.KeyValue) ResourceOption {
	return resource.WithAttributes(attributes...)
}

// WithDetectors adds detectors to be evaluated for the configured resource.
func WithDetectors(detectors ...ResourceDetector) ResourceOption {
	return resource.WithDetectors(detectors...)
}

// WithFromEnv adds attributes from environment variables to the configured resource.
func WithFromEnv() ResourceOption {
	return resource.WithFromEnv()
}

// WithHost adds attributes from the host to the configured resource.
func WithHost() ResourceOption {
	return resource.WithHost()
}

// WithTelemetrySDK adds TelemetrySDK version info to the configured resource.
func WithTelemetrySDK() ResourceOption {
	return resource.WithTelemetrySDK()
}

// WithSchemaURL sets the schema URL for the configured resource.
func WithSchemaURL(schemaURL string) ResourceOption {
	return resource.WithSchemaURL(schemaURL)
}

// WithOS adds all the OS attributes to the configured Resource.
// See individual WithOS* functions to configure specific attributes.
func WithOS() ResourceOption {
	return resource.WithOS()
}

// WithOSType adds an attribute with the operating system type to the configured Resource.
func WithOSType() ResourceOption {
	return resource.WithOSType()
}

// WithOSDescription adds an attribute with the operating system description to the
// configured Resource. The formatted string is equivalent to the output of the
// `uname -snrvm` command.
func WithOSDescription() ResourceOption {
	return resource.WithOSDescription()
}

// WithProcess adds all the Process attributes to the configured Resource.
//
// Warning! This option will include process command line arguments. If these
// contain sensitive information it will be included in the exported resource.
//
// This option is equivalent to calling WithProcessPID,
// WithProcessExecutableName, WithProcessExecutablePath,
// WithProcessCommandArgs, WithProcessOwner, WithProcessRuntimeName,
// WithProcessRuntimeVersion, and WithProcessRuntimeDescription. See each
// option function for information about what resource attributes each
// includes.
func WithProcess() ResourceOption {
	return resource.WithProcess()
}

// WithProcessPID adds an attribute with the process identifier (PID) to the
// configured Resource.
func WithProcessPID() ResourceOption {
	return resource.WithProcessPID()
}

// WithProcessExecutableName adds an attribute with the name of the process
// executable to the configured Resource.
func WithProcessExecutableName() ResourceOption {
	return resource.WithProcessExecutableName()
}

// WithProcessExecutablePath adds an attribute with the full path to the process
// executable to the configured Resource.
func WithProcessExecutablePath() ResourceOption {
	return resource.WithProcessExecutablePath()
}

// WithProcessCommandArgs adds an attribute with all the command arguments (including
// the command/executable itself) as received by the process to the configured
// Resource.
//
// Warning! This option will include process command line arguments. If these
// contain sensitive information it will be included in the exported resource.
func WithProcessCommandArgs() ResourceOption {
	return resource.WithProcessCommandArgs()
}

// WithProcessOwner adds an attribute with the username of the user that owns the process
// to the configured Resource.
func WithProcessOwner() ResourceOption {
	return resource.WithProcessOwner()
}

// WithProcessRuntimeName adds an attribute with the name of the runtime of this
// process to the configured Resource.
func WithProcessRuntimeName() ResourceOption {
	return resource.WithProcessRuntimeName()
}

// WithProcessRuntimeVersion adds an attribute with the version of the runtime of
// this process to the configured Resource.
func WithProcessRuntimeVersion() ResourceOption {
	return resource.WithProcessRuntimeVersion()
}

// WithProcessRuntimeDescription adds an attribute with an additional description
// about the runtime of the process to the configured Resource.
func WithProcessRuntimeDescription() ResourceOption {
	return resource.WithProcessRuntimeDescription()
}

// WithContainer adds all the Container attributes to the configured Resource.
// See individual WithContainer* functions to configure specific attributes.
func WithContainer() ResourceOption {
	return resource.WithContainer()
}

// WithContainerID adds an attribute with the id of the container to the configured Resource.
func WithContainerID() ResourceOption {
	return resource.WithContainerID()
}

// Resource describes an entity about which identifying information
// and metadata is exposed.  Resource is an immutable object,
// equivalent to a map from key to unique value.
//
// Resources should be passed and stored as pointers
// (`*resource.Resource`).  The `nil` value is equivalent to an empty
// Resource.
type Resource = resource.Resource

// NewResource returns a Resource combined from the user-provided detectors.
func NewResource(ctx context.Context, opts ...ResourceOption) (*Resource, error) {
	return resource.New(ctx, opts...)
}

// NewWithAttributes creates a resource from attrs and associates the resource with a
// schema URL. If attrs contains duplicate keys, the last value will be used. If attrs
// contains any invalid items those items will be dropped. The attrs are assumed to be
// in a schema identified by schemaURL.
func NewWithAttributes(schemaURL string, attrs ...attribute.KeyValue) *Resource {
	return resource.NewWithAttributes(schemaURL, attrs...)
}

// NewSchemaless creates a resource from attrs. If attrs contains duplicate keys,
// the last value will be used. If attrs contains any invalid items those items will
// be dropped. The resource will not be associated with a schema URL. If the schema
// of the attrs is known use NewWithAttributes instead.
func NewSchemaless(attrs ...attribute.KeyValue) *Resource {
	return resource.NewSchemaless(attrs...)
}

// Merge creates a new resource by combining resource a and b.
//
// If there are common keys between resource a and b, then the value
// from resource b will overwrite the value from resource a, even
// if resource b's value is empty.
//
// The SchemaURL of the resources will be merged according to the spec rules:
// https://github.com/open-telemetry/opentelemetry-specification/blob/bad49c714a62da5493f2d1d9bafd7ebe8c8ce7eb/specification/resource/sdk.md#merge
// If the resources have different non-empty schemaURL an empty resource and an error
// will be returned.
func Merge(a, b *Resource) (*Resource, error) {
	return resource.Merge(a, b)
}

// Empty returns an instance of Resource with no attributes. It is
// equivalent to a `nil` Resource.
func Empty() *Resource {
	return resource.Empty()
}

// Default returns an instance of Resource with a default
// "service.name" and OpenTelemetrySDK attributes.
func Default() *Resource {
	return resource.Default()
}

// Environment returns an instance of Resource with attributes
// extracted from the OTEL_RESOURCE_ATTRIBUTES environment variable.
func Environment() *Resource {
	return resource.Environment()
}

// Defaults for BatchSpanProcessorOptions.
const (
	DefaultMaxQueueSize       = trace.DefaultMaxQueueSize
	DefaultScheduleDelay      = trace.DefaultScheduleDelay
	DefaultExportTimeout      = trace.DefaultExportTimeout
	DefaultMaxExportBatchSize = trace.DefaultMaxExportBatchSize
)

// BatchSpanProcessorOption configures a BatchSpanProcessor.
type BatchSpanProcessorOption = trace.BatchSpanProcessorOption

// BatchSpanProcessorOptions is configuration settings for a
// BatchSpanProcessor.
type BatchSpanProcessorOptions = trace.BatchSpanProcessorOptions

// NewBatchSpanProcessor creates a new SpanProcessor that will send completed
// span batches to the exporter with the supplied options.
//
// If the exporter is nil, the span processor will preform no action.
func NewBatchSpanProcessor(exporter SpanExporter, options ...BatchSpanProcessorOption) SpanProcessor {
	return trace.NewBatchSpanProcessor(exporter, options...)
}

// WithMaxQueueSize returns a BatchSpanProcessorOption that configures the
// maximum queue size allowed for a BatchSpanProcessor.
func WithMaxQueueSize(size int) BatchSpanProcessorOption {
	return func(o *BatchSpanProcessorOptions) {
		o.MaxQueueSize = size
	}
}

// WithMaxExportBatchSize returns a BatchSpanProcessorOption that configures
// the maximum export batch size allowed for a BatchSpanProcessor.
func WithMaxExportBatchSize(size int) BatchSpanProcessorOption {
	return trace.WithMaxExportBatchSize(size)
}

// WithBatchTimeout returns a BatchSpanProcessorOption that configures the
// maximum delay allowed for a BatchSpanProcessor before it will export any
// held span (whether the queue is full or not).
func WithBatchTimeout(delay time.Duration) BatchSpanProcessorOption {
	return trace.WithBatchTimeout(delay)
}

// WithExportTimeout returns a BatchSpanProcessorOption that configures the
// amount of time a BatchSpanProcessor waits for an exporter to export before
// abandoning the export.
func WithExportTimeout(timeout time.Duration) BatchSpanProcessorOption {
	return trace.WithExportTimeout(timeout)
}

// WithBlocking returns a BatchSpanProcessorOption that configures a
// BatchSpanProcessor to wait for enqueue operations to succeed instead of
// dropping data when the queue is full.
func WithBlocking() BatchSpanProcessorOption {
	return trace.WithBlocking()
}

// Event is a thing that happened during a Span's lifetime.
type Event = trace.Event

// IDGenerator allows custom generators for TraceID and SpanID.
type IDGenerator = trace.IDGenerator

// Link is the relationship between two Spans. The relationship can be within
// the same Trace or across different Traces.
type Link = trace.Link

// TracerProvider is an OpenTelemetry TracerProvider. It provides Tracers to
// instrumentation so it can trace operational flow through a system.
type TracerProvider = trace.TracerProvider

// The passed opts are used to override these default values and configure the
// returned TracerProvider appropriately.
func NewTracerProvider(opts ...TracerProviderOption) *TracerProvider {
	return trace.NewTracerProvider(opts...)
}

// TracerProviderOption configures a TracerProvider.
type TracerProviderOption = trace.TracerProviderOption

// WithSyncer registers the exporter with the TracerProvider using a
// SimpleSpanProcessor.
//
// This is not recommended for production use. The synchronous nature of the
// SimpleSpanProcessor that will wrap the exporter make it good for testing,
// debugging, or showing examples of other feature, but it will be slow and
// have a high computation resource usage overhead. The WithBatcher option is
// recommended for production use instead.
func WithSyncer(e SpanExporter) TracerProviderOption {
	return trace.WithSyncer(e)
}

// WithBatcher registers the exporter with the TracerProvider using a
// BatchSpanProcessor configured with the passed opts.
func WithBatcher(e SpanExporter, opts ...BatchSpanProcessorOption) TracerProviderOption {
	return trace.WithBatcher(e, opts...)
}

// WithSpanProcessor registers the SpanProcessor with a TracerProvider.
func WithSpanProcessor(sp SpanProcessor) TracerProviderOption {
	return trace.WithSpanProcessor(sp)
}

// WithResource returns a TracerProviderOption that will configure the
// Resource r as a TracerProvider's Resource. The configured Resource is
// referenced by all the Tracers the TracerProvider creates. It represents the
// entity producing telemetry.
//
// If this option is not used, the TracerProvider will use the
// resource.Default() Resource by default.
func WithResource(r *resource.Resource) TracerProviderOption {
	return trace.WithResource(r)
}

// WithIDGenerator returns a TracerProviderOption that will configure the
// IDGenerator g as a TracerProvider's IDGenerator. The configured IDGenerator
// is used by the Tracers the TracerProvider creates to generate new Span and
// Trace IDs.
//
// If this option is not used, the TracerProvider will use a random number
// IDGenerator by default.
func WithIDGenerator(g IDGenerator) TracerProviderOption {
	return trace.WithIDGenerator(g)
}

// WithSampler returns a TracerProviderOption that will configure the Sampler
// s as a TracerProvider's Sampler. The configured Sampler is used by the
// Tracers the TracerProvider creates to make their sampling decisions for the
// Spans they create.
//
// This option overrides the Sampler configured through the OTEL_TRACES_SAMPLER
// and OTEL_TRACES_SAMPLER_ARG environment variables. If this option is not used
// and the sampler is not configured through environment variables or the environment
// contains invalid/unsupported configuration, the TracerProvider will use a
// ParentBased(AlwaysSample) Sampler by default.
func WithSampler(s Sampler) TracerProviderOption {
	return trace.WithSampler(s)
}

// WithSpanLimits returns a TracerProviderOption that configures a
// TracerProvider to use the SpanLimits sl. These SpanLimits bound any Span
// created by a Tracer from the TracerProvider.
//
// If any field of sl is zero or negative it will be replaced with the default
// value for that field.
//
// If this or WithRawSpanLimits are not provided, the TracerProvider will use
// the limits defined by environment variables, or the defaults if unset.
// Refer to the NewSpanLimits documentation for information about this
// relationship.
//
// Deprecated: Use WithRawSpanLimits instead which allows setting unlimited
// and zero limits. This option will be kept until the next major version
// incremented release.
func WithSpanLimits(sl SpanLimits) TracerProviderOption {
	return trace.WithSpanLimits(sl)
}

// WithRawSpanLimits returns a TracerProviderOption that configures a
// TracerProvider to use these limits. These limits bound any Span created by
// a Tracer from the TracerProvider.
//
// The limits will be used as-is. Zero or negative values will not be changed
// to the default value like WithSpanLimits does. Setting a limit to zero will
// effectively disable the related resource it limits and setting to a
// negative value will mean that resource is unlimited. Consequentially, this
// means that the zero-value SpanLimits will disable all span resources.
// Because of this, limits should be constructed using NewSpanLimits and
// updated accordingly.
//
// If this or WithSpanLimits are not provided, the TracerProvider will use the
// limits defined by environment variables, or the defaults if unset. Refer to
// the NewSpanLimits documentation for information about this relationship.
func WithRawSpanLimits(limits SpanLimits) TracerProviderOption {
	return trace.WithRawSpanLimits(limits)
}

// Sampler decides whether a trace should be sampled and exported.
type Sampler = trace.Sampler

// SamplingParameters contains the values passed to a Sampler.
type SamplingParameters = trace.SamplingParameters

// SamplingDecision indicates whether a span is dropped, recorded and/or sampled.
type SamplingDecision = trace.SamplingDecision

// Valid sampling decisions.
const (
	// Drop will not record the span and all attributes/events will be dropped.
	Drop SamplingDecision = trace.Drop

	// Record indicates the span's `IsRecording() == true`, but `Sampled` flag
	// *must not* be set.
	RecordOnly = trace.RecordOnly

	// RecordAndSample has span's `IsRecording() == true` and `Sampled` flag
	// *must* be set.
	RecordAndSample = trace.RecordAndSample
)

// SamplingResult conveys a SamplingDecision, set of Attributes and a Tracestate.
type SamplingResult = trace.SamplingResult

// TraceIDRatioBased samples a given fraction of traces. Fractions >= 1 will
// always sample. Fractions < 0 are treated as zero. To respect the
// parent trace's `SampledFlag`, the `TraceIDRatioBased` sampler should be used
// as a delegate of a `Parent` sampler.
//nolint:revive // revive complains about stutter of `trace.TraceIDRatioBased`
func TraceIDRatioBased(fraction float64) Sampler {
	return trace.TraceIDRatioBased(fraction)
}

// AlwaysSample returns a Sampler that samples every trace.
// Be careful about using this sampler in a production application with
// significant traffic: a new trace will be started and exported for every
// request.
func AlwaysSample() Sampler {
	return trace.AlwaysSample()
}

// NeverSample returns a Sampler that samples no traces.
func NeverSample() Sampler {
	return trace.NeverSample()
}

// ParentBased returns a composite sampler which behaves differently,
// based on the parent of the span. If the span has no parent,
// the root(Sampler) is used to make sampling decision. If the span has
// a parent, depending on whether the parent is remote and whether it
// is sampled, one of the following samplers will apply:
// - remoteParentSampled(Sampler) (default: AlwaysOn)
// - remoteParentNotSampled(Sampler) (default: AlwaysOff)
// - localParentSampled(Sampler) (default: AlwaysOn)
// - localParentNotSampled(Sampler) (default: AlwaysOff)
func ParentBased(root Sampler, samplers ...ParentBasedSamplerOption) Sampler {
	return trace.ParentBased(root, samplers...)
}

// ParentBasedSamplerOption configures the sampler for a particular sampling case.
type ParentBasedSamplerOption = trace.ParentBasedSamplerOption

// WithRemoteParentSampled sets the sampler for the case of sampled remote parent.
func WithRemoteParentSampled(s Sampler) ParentBasedSamplerOption {
	return trace.WithRemoteParentSampled(s)
}

// WithRemoteParentNotSampled sets the sampler for the case of remote parent
// which is not sampled.
func WithRemoteParentNotSampled(s Sampler) ParentBasedSamplerOption {
	return trace.WithRemoteParentNotSampled(s)
}

// WithLocalParentSampled sets the sampler for the case of sampled local parent.
func WithLocalParentSampled(s Sampler) ParentBasedSamplerOption {
	return trace.WithLocalParentSampled(s)
}

// WithLocalParentNotSampled sets the sampler for the case of local parent
// which is not sampled.
func WithLocalParentNotSampled(s Sampler) ParentBasedSamplerOption {
	return trace.WithLocalParentNotSampled(s)
}

// NewSimpleSpanProcessor returns a new SpanProcessor that will synchronously
// send completed spans to the exporter immediately.
//
// This SpanProcessor is not recommended for production use. The synchronous
// nature of this SpanProcessor make it good for testing, debugging, or
// showing examples of other feature, but it will be slow and have a high
// computation resource usage overhead. The BatchSpanProcessor is recommended
// for production use instead.
func NewSimpleSpanProcessor(exporter SpanExporter) SpanProcessor {
	return trace.NewSimpleSpanProcessor(exporter)
}

// ReadOnlySpan allows reading information from the data structure underlying a
// trace.Span. It is used in places where reading information from a span is
// necessary but changing the span isn't necessary or allowed.
//
// Warning: methods may be added to this interface in minor releases.
type ReadOnlySpan = trace.ReadOnlySpan

// ReadWriteSpan exposes the same methods as trace.Span and in addition allows
// reading information from the underlying data structure.
// This interface exposes the union of the methods of trace.Span (which is a
// "write-only" span) and ReadOnlySpan. New methods for writing or reading span
// information should be added under trace.Span or ReadOnlySpan, respectively.
//
// Warning: methods may be added to this interface in minor releases.
type ReadWriteSpan = trace.ReadWriteSpan

// TraceStatus is the classified state of a Span.
type TraceStatus = trace.Status

// SpanExporter handles the delivery of spans to external receivers. This is
// the final component in the trace export pipeline.
type SpanExporter = trace.SpanExporter

const (
	// DefaultAttributeValueLengthLimit is the default maximum allowed
	// attribute value length, unlimited.
	DefaultAttributeValueLengthLimit = trace.DefaultAttributeValueLengthLimit

	// DefaultAttributeCountLimit is the default maximum number of attributes
	// a span can have.
	DefaultAttributeCountLimit = trace.DefaultAttributeCountLimit

	// DefaultEventCountLimit is the default maximum number of events a span
	// can have.
	DefaultEventCountLimit = trace.DefaultEventCountLimit

	// DefaultLinkCountLimit is the default maximum number of links a span can
	// have.
	DefaultLinkCountLimit = trace.DefaultLinkCountLimit

	// DefaultAttributePerEventCountLimit is the default maximum number of
	// attributes a span event can have.
	DefaultAttributePerEventCountLimit = trace.DefaultAttributePerEventCountLimit

	// DefaultAttributePerLinkCountLimit is the default maximum number of
	// attributes a span link can have.
	DefaultAttributePerLinkCountLimit = trace.DefaultAttributePerLinkCountLimit
)

// SpanLimits represents the limits of a span.
type SpanLimits = trace.SpanLimits

// NewSpanLimits returns a SpanLimits with all limits set to the value their
// corresponding environment variable holds, or the default if unset.
//
// • AttributeValueLengthLimit: OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT
// (default: unlimited)
//
// • AttributeCountLimit: OTEL_SPAN_ATTRIBUTE_COUNT_LIMIT (default: 128)
//
// • EventCountLimit: OTEL_SPAN_EVENT_COUNT_LIMIT (default: 128)
//
// • AttributePerEventCountLimit: OTEL_EVENT_ATTRIBUTE_COUNT_LIMIT (default:
// 128)
//
// • LinkCountLimit: OTEL_SPAN_LINK_COUNT_LIMIT (default: 128)
//
// • AttributePerLinkCountLimit: OTEL_LINK_ATTRIBUTE_COUNT_LIMIT (default: 128)
func NewSpanLimits() SpanLimits {
	return trace.NewSpanLimits()
}

// SpanProcessor is a processing pipeline for spans in the trace signal.
// SpanProcessors registered with a TracerProvider and are called at the start
// and end of a Span's lifecycle, and are called in the order they are
// registered.
type SpanProcessor = trace.SpanProcessor
