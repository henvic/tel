package tel

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// TracerConfig is a group of options for a Tracer.
type TracerConfig = trace.TracerConfig

// NewTracerConfig applies all the options to a returned TracerConfig.
func NewTracerConfig(options ...TracerOption) TracerConfig {
	return trace.NewTracerConfig(options...)
}

// TracerOption applies an option to a TracerConfig.
type TracerOption = trace.TracerOption

// SpanConfig is a group of options for a Span.
type SpanConfig = trace.SpanConfig

// NewSpanStartConfig applies all the options to a returned SpanConfig.
// No validation is performed on the returned SpanConfig (e.g. no uniqueness
// checking or bounding of data), it is left to the SDK to perform this
// action.
func NewSpanStartConfig(options ...SpanStartOption) SpanConfig {
	return trace.NewSpanStartConfig(options...)
}

// NewSpanEndConfig applies all the options to a returned SpanConfig.
// No validation is performed on the returned SpanConfig (e.g. no uniqueness
// checking or bounding of data), it is left to the SDK to perform this
// action.
func NewSpanEndConfig(options ...SpanEndOption) SpanConfig {
	return trace.NewSpanEndConfig(options...)
}

// SpanStartOption applies an option to a SpanConfig. These options are applicable
// only when the span is created.
type SpanStartOption = trace.SpanStartOption

// SpanEndOption applies an option to a SpanConfig. These options are
// applicable only when the span is ended.
type SpanEndOption = trace.SpanEndOption

// EventConfig is a group of options for an Event.
type EventConfig = trace.EventConfig

// NewEventConfig applies all the EventOptions to a returned EventConfig. If no
// timestamp option is passed, the returned EventConfig will have a Timestamp
// set to the call time, otherwise no validation is performed on the returned
// EventConfig.
func NewEventConfig(options ...EventOption) EventConfig {
	return trace.NewEventConfig(options...)
}

// EventOption applies span event options to an EventConfig.
type EventOption = trace.EventOption

// SpanOption are options that can be used at both the beginning and end of a span.
type SpanOption = trace.SpanOption

// SpanStartEventOption are options that can be used at the start of a span, or with an event.
type SpanStartEventOption = trace.SpanStartEventOption

// SpanEndEventOption are options that can be used at the end of a span, or with an event.
type SpanEndEventOption = trace.SpanEndEventOption

// WithAttributes adds the attributes related to a span life-cycle event.
// These attributes are used to describe the work a Span represents when this
// option is provided to a Span's start or end events. Otherwise, these
// attributes provide additional information about the event being recorded
// (e.g. error, state change, processing progress, system event).
//
// If multiple of these options are passed the attributes of each successive
// option will extend the attributes instead of overwriting. There is no
// guarantee of uniqueness in the resulting attributes.
func WithAttributes(attributes ...KeyValue) SpanStartEventOption {
	return trace.WithAttributes(attributes...)
}

// SpanEventOption are options that can be used with an event or a span.
type SpanEventOption = trace.SpanEventOption

// WithTimestamp sets the time of a Span or Event life-cycle moment (e.g.
// started, stopped, errored).
func WithTimestamp(t time.Time) SpanEventOption {
	return trace.WithTimestamp(t)
}

// WithStackTrace sets the flag to capture the error with stack trace (e.g. true, false).
func WithStackTrace(b bool) SpanEndEventOption {
	return trace.WithStackTrace(b)
}

// WithLinks adds links to a Span. The links are added to the existing Span
// links, i.e. this does not overwrite. Links with invalid span context are ignored.
func WithLinks(links ...Link) SpanStartOption {
	return trace.WithLinks(links...)
}

// WithNewRoot specifies that the Span should be treated as a root Span. Any
// existing parent span context will be ignored when defining the Span's trace
// identifiers.
func WithNewRoot() SpanStartOption {
	return trace.WithNewRoot()
}

// WithSpanKind sets the SpanKind of a Span.
func WithSpanKind(kind SpanKind) SpanStartOption {
	return trace.WithSpanKind(kind)
}

// TraceWithInstrumentationVersion sets the instrumentation version.
func TraceWithInstrumentationVersion(version string) TracerOption {
	return trace.WithInstrumentationVersion(version)
}

// TraceWithSchemaURL sets the schema URL for the Tracer.
func TraceWithSchemaURL(schemaURL string) TracerOption {
	return trace.WithSchemaURL(schemaURL)
}

// ContextWithSpan returns a copy of parent with span set as the current Span.
func ContextWithSpan(parent context.Context, span Span) context.Context {
	return trace.ContextWithSpan(parent, span)
}

// ContextWithSpanContext returns a copy of parent with sc as the current
// Span. The Span implementation that wraps sc is non-recording and performs
// no operations other than to return sc as the SpanContext from the
// SpanContext method.
func ContextWithSpanContext(parent context.Context, sc SpanContext) context.Context {
	return trace.ContextWithSpanContext(parent, sc)
}

// ContextWithRemoteSpanContext returns a copy of parent with rsc set explicly
// as a remote SpanContext and as the current Span. The Span implementation
// that wraps rsc is non-recording and performs no operations other than to
// return rsc as the SpanContext from the SpanContext method.
func ContextWithRemoteSpanContext(parent context.Context, rsc SpanContext) context.Context {
	return ContextWithSpanContext(parent, rsc.WithRemote(true))
}

// SpanFromContext returns the current Span from ctx.
//
// If no Span is currently set in ctx an implementation of a Span that
// performs no operations is returned.
func SpanFromContext(ctx context.Context) Span {
	return trace.SpanFromContext(ctx)
}

// SpanContextFromContext returns the current Span's SpanContext.
func SpanContextFromContext(ctx context.Context) SpanContext {
	return trace.SpanFromContext(ctx).SpanContext()
}

// NewNoopTracerProvider returns an implementation of TracerProvider that
// performs no operations. The Tracer and Spans created from the returned
// TracerProvider also perform no operations.
func NewNoopTracerProvider() TracerProvider {
	return trace.NewNoopTracerProvider()
}

// FlagsSampled is a bitmask with the sampled bit set. A SpanContext
// with the sampling bit set means the span is sampled.
const FlagsSampled = TraceFlags(0x01)

// TraceID is a unique identity of a trace.
// nolint:revive // revive complains about stutter of `trace.TraceID`.
type TraceID = trace.TraceID

// SpanID is a unique identity of a span in a trace.
type SpanID = trace.SpanID

// TraceIDFromHex returns a TraceID from a hex string if it is compliant with
// the W3C trace-context specification.  See more at
// https://www.w3.org/TR/trace-context/#trace-id
// nolint:revive // revive complains about stutter of `trace.TraceIDFromHex`.
func TraceIDFromHex(h string) (trace.TraceID, error) {
	return trace.TraceIDFromHex(h)
}

// SpanIDFromHex returns a SpanID from a hex string if it is compliant
// with the w3c trace-context specification.
// See more at https://www.w3.org/TR/trace-context/#parent-id
func SpanIDFromHex(h string) (SpanID, error) {
	return trace.SpanIDFromHex(h)
}

// TraceFlags contains flags that can be set on a SpanContext.
type TraceFlags byte

// SpanContextConfig contains mutable fields usable for constructing
// an immutable SpanContext.
type SpanContextConfig = trace.SpanContextConfig

// NewSpanContext constructs a SpanContext using values from the provided
// SpanContextConfig.
func NewSpanContext(config SpanContextConfig) SpanContext {
	return trace.NewSpanContext(config)
}

// SpanContext contains identifying trace information about a Span.
type SpanContext = trace.SpanContext

// Span is the individual component of a trace. It represents a single named
// and timed operation of a workflow that is traced. A Tracer is used to
// create a Span and it is then up to the operation the Span represents to
// properly end the Span when the operation itself ends.
//
// Warning: methods may be added to this interface in minor releases.
type Span = trace.Span

// Link is the relationship between two Spans. The relationship can be within
// the same Trace or across different Traces.
//
// For example, a Link is used in the following situations:
//
//   1. Batch Processing: A batch of operations may contain operations
//      associated with one or more traces/spans. Since there can only be one
//      parent SpanContext, a Link is used to keep reference to the
//      SpanContext of all operations in the batch.
//   2. Public Endpoint: A SpanContext for an in incoming client request on a
//      public endpoint should be considered untrusted. In such a case, a new
//      trace with its own identity and sampling decision needs to be created,
//      but this new trace needs to be related to the original trace in some
//      form. A Link is used to keep reference to the original SpanContext and
//      track the relationship.
type Link = trace.Link

// LinkFromContext returns a link encapsulating the SpanContext in the provided ctx.
func LinkFromContext(ctx context.Context, attrs ...KeyValue) Link {
	return trace.LinkFromContext(ctx, attrs...)
}

// SpanKind is the role a Span plays in a Trace.
type SpanKind = trace.SpanKind

// As a convenience, these match the proto definition, see
// https://github.com/open-telemetry/opentelemetry-proto/blob/30d237e1ff3ab7aa50e0922b5bebdd93505090af/opentelemetry/proto/trace/v1/trace.proto#L101-L129
//
// The unspecified value is not a valid `SpanKind`. Use `ValidateSpanKind()`
// to coerce a span kind to a valid value.
const (
	// SpanKindUnspecified is an unspecified SpanKind and is not a valid
	// SpanKind. SpanKindUnspecified should be replaced with SpanKindInternal
	// if it is received.
	SpanKindUnspecified = trace.SpanKindUnspecified
	// SpanKindInternal is a SpanKind for a Span that represents an internal
	// operation within an application.
	SpanKindInternal = trace.SpanKindInternal
	// SpanKindServer is a SpanKind for a Span that represents the operation
	// of handling a request from a client.
	SpanKindServer = trace.SpanKindServer
	// SpanKindClient is a SpanKind for a Span that represents the operation
	// of client making a request to a server.
	SpanKindClient = trace.SpanKindClient
	// SpanKindProducer is a SpanKind for a Span that represents the operation
	// of a producer sending a message to a message broker. Unlike
	// SpanKindClient and SpanKindServer, there is often no direct
	// relationship between this kind of Span and a SpanKindConsumer kind. A
	// SpanKindProducer Span will end once the message is accepted by the
	// message broker which might not overlap with the processing of that
	// message.
	SpanKindProducer = trace.SpanKindProducer
	// SpanKindConsumer is a SpanKind for a Span that represents the operation
	// of a consumer receiving a message from a message broker. Like
	// SpanKindProducer Spans, there is often no direct relationship between
	// this Span and the Span that produced the message.
	SpanKindConsumer = trace.SpanKindConsumer
)

// ValidateSpanKind returns a valid span kind value.  This will coerce
// invalid values into the default value, SpanKindInternal.
func ValidateSpanKind(spanKind SpanKind) SpanKind {
	return trace.ValidateSpanKind(spanKind)
}

// Tracer is the creator of Spans.
//
// Warning: methods may be added to this interface in minor releases.
type Tracer = trace.Tracer

// TracerProvider provides access to instrumentation Tracers.
//
// Warning: methods may be added to this interface in minor releases.
type TracerProvider = trace.TracerProvider

// TraceState provides additional vendor-specific trace identification
// information across different distributed tracing systems. It represents an
// immutable list consisting of key/value pairs, each pair is referred to as a
// list-member.
//
// TraceState conforms to the W3C Trace Context specification
// (https://www.w3.org/TR/trace-context-1). All operations that create or copy
// a TraceState do so by validating all input and will only produce TraceState
// that conform to the specification. Specifically, this means that all
// list-member's key/value pairs are valid, no duplicate list-members exist,
// and the maximum number of list-members (32) is not exceeded.
type TraceState = trace.TraceState

// ParseTraceState attempts to decode a TraceState from the passed
// string. It returns an error if the input is invalid according to the W3C
// Trace Context specification.
func ParseTraceState(tracestate string) (TraceState, error) {
	return trace.ParseTraceState(tracestate)
}
